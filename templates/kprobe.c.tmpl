//go:build ignore

#include "vmlinux.h"
#include "bpf_tracing.h"
#include "bpf_helpers.h"
#include "bpf_core_read.h"

char __license[] SEC("license") = "Dual MIT/GPL";

{{range $index, $element := .TargetFunction}}
struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, u32);
    __type(value, u64);
    __uint(max_entries, 1024);
} start_time{{$index}} SEC(".maps");
{{end}}

{{range $index, $element := .TargetFunction}}
SEC("kprobe/{{.}}")
int gen_{{.}}(struct pt_regs *ctx) {
    u32 pid = bpf_get_current_pid_tgid();
    u64 ts = bpf_ktime_get_ns();

    bpf_map_update_elem(&start_time{{$index}}, &pid, &ts, BPF_ANY);
    return 0;
}
{{end}}

{{range $index, $element := .TargetFunction}}
SEC("kretprobe/{{.}}")
int gen_ret_{{.}}(struct pt_regs *ctx) {
    u32 pid = bpf_get_current_pid_tgid();
    u64 *tsp, delta;

    tsp = bpf_map_lookup_elem(&start_time{{$index}}, &pid);
    if (tsp != 0) {
        delta = bpf_ktime_get_ns() - *tsp;
        bpf_printk("{{.}} executed in %llu ns\n", delta);
        bpf_map_delete_elem(&start_time{{$index}}, &pid);
    }
    return 0;
}
{{end}}