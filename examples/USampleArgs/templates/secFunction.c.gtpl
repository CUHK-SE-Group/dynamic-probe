{{define "secFunction"}}
{{range $index, $element := .Functions}}
    {{if eq $element.Aim "BpfFuncTypeUprobeSampleArgs"}}
SEC("uprobe:{{$element.Name}}")
int uprobe_function_sample_args(struct pt_regs *ctx) {
    // 定义参数变量
    u64 arg1, arg2, arg3;

    // 从用户空间内存中读取函数参数
    bpf_probe_read_user(&arg1, sizeof(arg1), (void *)(ctx->di));
    bpf_probe_read_user(&arg2, sizeof(arg2), (void *)(ctx->si));
    bpf_probe_read_user(&arg3, sizeof(arg3), (void *)(ctx->dx));

    // 输出函数参数到内核日志
    bpf_printk("{{$element.Name}}: arg1=%lld, arg2=%lld, arg3=%lld\n", arg1, arg2, arg3);

    return 0;
}

    {{end}}
{{end}}
{{end}}