{{define "secFunction"}}
{{range $index, $element := .Functions}}
    {{if eq $element.Aim "BpfFuncTypeUprobeSampleReturn"}}
SEC("uprobe:{{$element.Name}}")
int uprobe_function_sample_return(struct pt_regs *ctx) {
    u64 retval = PT_REGS_RC(ctx); // 获取函数返回值

    // 输出函数返回值到内核日志
    bpf_printk("{{$element.Name}} returned: %lld\n", retval);

    return 0;
}
    {{end}}
{{end}}
{{end}}


