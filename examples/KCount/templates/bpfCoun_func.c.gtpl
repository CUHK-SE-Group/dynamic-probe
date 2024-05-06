{{define "bpfCoun_func"}}
{{range $index, $element := .Functions}}
    {{if eq $element.Aim "BpfFuncTypeKprobeCount"}}
// 定义 BPF 程序的入口点
SEC("kprobe/{{$element.Name}}")
int kprobe_sys_read(struct pt_regs *ctx) {
    // 增加计数器
    __u64 zero = 0;
    bpf_map_increment_elem(&function_count, &zero);
    return 0;
}
    {{end}}
{{end}}
{{end}}