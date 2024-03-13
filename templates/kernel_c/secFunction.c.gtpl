{{define "secFunction"}}
{{range $index, $element := .TargetFunction}}
SEC("uprobe/{{$index}}")
int bpf_prog_{{$index}}(struct pt_regs *ctx) {
    u32 pid = bpf_get_current_pid_tgid();
    u64 ts = bpf_ktime_get_ns();
    struct {{$index}}_args args = {};

    bpf_printk("fuck 1");

    {{range $id, $val := $element}}
    args.arg{{$id}} = PT_REGS_PARM{{AddOne $id}}(ctx);
    {{end}}

    bpf_map_update_elem(&start_time_map, &pid, &ts, BPF_ANY);
    bpf_map_update_elem(&{{$index}}_args_map, &pid, &args, BPF_ANY);
    return 0;
}
{{end}}

{{end}}
