{{define "secFunction"}}
{{range $index, $element := .Functions}}
SEC("uprobe/{{$element.Name}}")
int bpf_prog_{{$element.Name}}(struct pt_regs *ctx) {
    u32 pid = bpf_get_current_pid_tgid();
    u64 ts = bpf_ktime_get_ns();
    struct {{$element.Name}}_args args = {};


    bpf_map_update_elem(&start_time_map, &pid, &ts, BPF_ANY);
    bpf_map_update_elem(&args_map_{{$index}}, &pid, &args, BPF_ANY);
    return 0;
}
{{end}}
{{end}}

{{define "count"}}

{{end}}