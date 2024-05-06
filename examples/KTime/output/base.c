
//go:build ignore

#include "vmlinux.h"
#include "bpf_tracing.h"
#include "bpf_helpers.h"
#include "bpf_core_read.h" 


char __license[] SEC("license") = "Dual MIT/GPL";


__u32 zero = 0;



struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, u32);
    __type(value, u64);
    __uint(max_entries, 1);
} timestamps SEC(".maps");

struct {
    __uint(type, BPF_MAP_TYPE_ARRAY_OF_MAPS);
    __type(key, u32);
    __type(value, u64);
    __uint(max_entries, 10);
} array_of_packet_maps SEC(".maps");




    
// 定义 BPF 程序的入口点
SEC("kprobe/check_packets")
int kprobe_check_packets(struct pt_regs *ctx) {
    // 记录当前时间
    u64 timestamp = bpf_ktime_get_ns();
    // 将时间戳保存到 BPF 映射中，以便用户空间程序访问
    bpf_map_update_elem(&timestamps, &zero, &timestamp, BPF_ANY);
    return 0;
}

// 定义 BPF 程序的出口点
SEC("kretprobe/check_packets")
int kretprobe_check_packets(struct pt_regs *ctx) {
    // 从 BPF 映射中获取之前保存的时间戳
    u64 *timestamp = bpf_map_lookup_elem(&timestamps, &zero);
    if (timestamp) {
        // 计算函数执行时间
        u64 delta = bpf_ktime_get_ns() - *timestamp;
        // 输出执行时间（可以替换为其他输出方式，如使用 perf_event 输出到用户空间）
        bpf_printk("check_packets execution time: %lld ns\n", delta);
        // 删除映射中的时间戳，准备下一次记录
        bpf_map_delete_elem(&timestamps, &zero);
    }
    return 0;
}
    

    
// 定义 BPF 程序的入口点
SEC("kprobe/ktimefunc")
int kprobe_ktimefunc(struct pt_regs *ctx) {
    // 记录当前时间
    u64 timestamp = bpf_ktime_get_ns();
    // 将时间戳保存到 BPF 映射中，以便用户空间程序访问
    bpf_map_update_elem(&timestamps, &zero, &timestamp, BPF_ANY);
    return 0;
}

// 定义 BPF 程序的出口点
SEC("kretprobe/ktimefunc")
int kretprobe_ktimefunc(struct pt_regs *ctx) {
    // 从 BPF 映射中获取之前保存的时间戳
    u64 *timestamp = bpf_map_lookup_elem(&timestamps, &zero);
    if (timestamp) {
        // 计算函数执行时间
        u64 delta = bpf_ktime_get_ns() - *timestamp;
        // 输出执行时间（可以替换为其他输出方式，如使用 perf_event 输出到用户空间）
        bpf_printk("ktimefunc execution time: %lld ns\n", delta);
        // 删除映射中的时间戳，准备下一次记录
        bpf_map_delete_elem(&timestamps, &zero);
    }
    return 0;
}
    



