//go:build ignore

#include "vmlinux.h"
#include "bpf_tracing.h"
#include "bpf_helpers.h"
#include "bpf_core_read.h"

char __license[] SEC("license") = "Dual MIT/GPL";



struct demonstrateDynamicMemory_func_event_data {
    u32 pid;
    u64 start_time;
    long ret;
    u64 duration;
    
    u64 arg0; 
};
const struct demonstrateDynamicMemory_func_event_data *unused_demonstrateDynamicMemory __attribute__((unused));

struct demonstratePointers_func_event_data {
    u32 pid;
    u64 start_time;
    long ret;
    u64 duration;
    
    u64 arg0; 
};
const struct demonstratePointers_func_event_data *unused_demonstratePointers __attribute__((unused));


struct {
    __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
} events SEC(".maps");


struct demonstrateDynamicMemory_args {
    
    u64 arg0; 
};

struct demonstratePointers_args {
    
    u64 arg0; 
};



struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, u32);
    __type(value, struct demonstrateDynamicMemory_args);
    __uint(max_entries, 1024);
} demonstrateDynamicMemory_args_map SEC(".maps");

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, u32);
    __type(value, struct demonstratePointers_args);
    __uint(max_entries, 1024);
} demonstratePointers_args_map SEC(".maps");


struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, u32);
    __type(value, u64); // 只存储开始时间
    __uint(max_entries, 1024);
} start_time_map SEC(".maps");


SEC("uprobe/demonstrateDynamicMemory")
int bpf_prog_demonstrateDynamicMemory(struct pt_regs *ctx) {
    u32 pid = bpf_get_current_pid_tgid();
    u64 ts = bpf_ktime_get_ns();
    struct demonstrateDynamicMemory_args args = {};

    bpf_printk("fuck 1");

    // 假设参数是long类型，根据需要调整
    
    args.arg0 = PT_REGS_PARM1(ctx);
    

    bpf_map_update_elem(&start_time_map, &pid, &ts, BPF_ANY);
    bpf_map_update_elem(&demonstrateDynamicMemory_args_map, &pid, &args, BPF_ANY);
    return 0;
}

SEC("uprobe/demonstratePointers")
int bpf_prog_demonstratePointers(struct pt_regs *ctx) {
    u32 pid = bpf_get_current_pid_tgid();
    u64 ts = bpf_ktime_get_ns();
    struct demonstratePointers_args args = {};

    bpf_printk("fuck 1");

    // 假设参数是long类型，根据需要调整
    
    args.arg0 = PT_REGS_PARM1(ctx);
    

    bpf_map_update_elem(&start_time_map, &pid, &ts, BPF_ANY);
    bpf_map_update_elem(&demonstratePointers_args_map, &pid, &args, BPF_ANY);
    return 0;
}



SEC("uretprobe/demonstrateDynamicMemory")
int bpf_prog_ret_demonstrateDynamicMemory(struct pt_regs *ctx) {
    u32 pid = bpf_get_current_pid_tgid();
    u64 *tsp, delta;
    struct demonstrateDynamicMemory_args *argsp;
    struct demonstrateDynamicMemory_func_event_data data = {};

    tsp = bpf_map_lookup_elem(&start_time_map, &pid);
    argsp = bpf_map_lookup_elem(&demonstrateDynamicMemory_args_map, &pid);
    if (tsp != 0 && argsp != 0) {
        delta = bpf_ktime_get_ns() - *tsp;
        data.pid = pid;
        data.start_time = *tsp;
        data.duration = delta;
        
        data.arg0 = argsp->arg0; // 使用保存的参数值
        
        data.ret = PT_REGS_RC(ctx); // 获取返回值

        bpf_printk("pointer %lld", data.arg0);

        // 将数据发送到用户空间
        bpf_perf_event_output(ctx, &events, BPF_F_CURRENT_CPU, &data, sizeof(data));

        bpf_map_delete_elem(&start_time_map, &pid);
        bpf_map_delete_elem(&demonstrateDynamicMemory_args_map, &pid);
    }
    return 0;
}

SEC("uretprobe/demonstratePointers")
int bpf_prog_ret_demonstratePointers(struct pt_regs *ctx) {
    u32 pid = bpf_get_current_pid_tgid();
    u64 *tsp, delta;
    struct demonstratePointers_args *argsp;
    struct demonstratePointers_func_event_data data = {};

    tsp = bpf_map_lookup_elem(&start_time_map, &pid);
    argsp = bpf_map_lookup_elem(&demonstratePointers_args_map, &pid);
    if (tsp != 0 && argsp != 0) {
        delta = bpf_ktime_get_ns() - *tsp;
        data.pid = pid;
        data.start_time = *tsp;
        data.duration = delta;
        
        data.arg0 = argsp->arg0; // 使用保存的参数值
        
        data.ret = PT_REGS_RC(ctx); // 获取返回值

        bpf_printk("pointer %lld", data.arg0);

        // 将数据发送到用户空间
        bpf_perf_event_output(ctx, &events, BPF_F_CURRENT_CPU, &data, sizeof(data));

        bpf_map_delete_elem(&start_time_map, &pid);
        bpf_map_delete_elem(&demonstratePointers_args_map, &pid);
    }
    return 0;
}

