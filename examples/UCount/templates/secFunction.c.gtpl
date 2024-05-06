{{define "secFunction"}}
{{range $index, $element := .Functions}}
    {{if eq $element.Aim "BpfFuncTypeUprobeCount"}}
// 定义 BPF 程序的入口点
SEC("uprobe:{{$element.Name}}")
int uprobe_function_count(struct pt_regs *ctx) {
    // 将函数调用次数存储在 BPF 映射中
    u64 ip = PT_REGS_IP(ctx);
    u64 *count = bpf_map_lookup_elem(&function_count, &ip);
    if (count) {
        (*count)++;
    } else {
        u64 initial_count = 1;
        bpf_map_update_elem(&function_count, &ip, &initial_count, BPF_ANY);
    }

    return 0;
}
    {{end}}
{{end}}
{{end}}
