//go:build ignore

#include "vmlinux.h"
#include "bpf_tracing.h"
#include "bpf_helpers.h"
#include "bpf_core_read.h"

char __license[] SEC("license") = "Dual MIT/GPL";


{{range $key, $values := .TargetFunction }}
struct {{$key}}_func_event_data {
    u32 pid;
    u64 start_time;
    long ret;
    u64 duration;
    {{ range $index, $value := $values }}
    u64 arg{{$index}}; {{ end }}
};
const struct {{$key}}_func_event_data *unused_{{$key}} __attribute__((unused));
{{end}}

struct {
    __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
} events SEC(".maps");

{{range $key, $values := .TargetFunction }}
struct {{$key}}_args {
    {{ range $index, $value := $values }}
    u64 arg{{$index}}; {{ end }}
};
{{end}}

{{range $key, $values := .TargetFunction }}
struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, u32);
    __type(value, struct {{$key}}_args);
    __uint(max_entries, 1024);
} {{$key}}_args_map SEC(".maps");
{{end}}

struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, u32);
    __type(value, u64); // 只存储开始时间
    __uint(max_entries, 1024);
} start_time_map SEC(".maps");

{{range $index, $element := .TargetFunction}}
SEC("uprobe/{{$index}}")
int bpf_prog_{{$index}}(struct pt_regs *ctx) {
    u32 pid = bpf_get_current_pid_tgid();
    u64 ts = bpf_ktime_get_ns();
    struct {{$index}}_args args = {};

    bpf_printk("fuck 1");

    // 假设参数是long类型，根据需要调整
    {{range $id, $val := $element}}
    args.arg{{$id}} = PT_REGS_PARM{{AddOne $id}}(ctx);
    {{end}}

    bpf_map_update_elem(&start_time_map, &pid, &ts, BPF_ANY);
    bpf_map_update_elem(&{{$index}}_args_map, &pid, &args, BPF_ANY);
    return 0;
}
{{end}}

{{range $index, $element := .TargetFunction}}
SEC("uretprobe/{{$index}}")
int bpf_prog_ret_{{$index}}(struct pt_regs *ctx) {
    u32 pid = bpf_get_current_pid_tgid();
    u64 *tsp, delta;
    struct {{$index}}_args *argsp;
    struct {{$index}}_func_event_data data = {};

    tsp = bpf_map_lookup_elem(&start_time_map, &pid);
    argsp = bpf_map_lookup_elem(&{{$index}}_args_map, &pid);
    if (tsp != 0 && argsp != 0) {
        delta = bpf_ktime_get_ns() - *tsp;
        data.pid = pid;
        data.start_time = *tsp;
        data.duration = delta;
        {{range $id, $val := $element}}
        data.arg{{$id}} = argsp->arg{{$id}}; // 使用保存的参数值
        {{end}}
        data.ret = PT_REGS_RC(ctx); // 获取返回值

        bpf_printk("pointer %lld", data.arg0);

        // 将数据发送到用户空间
        bpf_perf_event_output(ctx, &events, BPF_F_CURRENT_CPU, &data, sizeof(data));

        bpf_map_delete_elem(&start_time_map, &pid);
        bpf_map_delete_elem(&{{$index}}_args_map, &pid);
    }
    return 0;
}
{{end}}
