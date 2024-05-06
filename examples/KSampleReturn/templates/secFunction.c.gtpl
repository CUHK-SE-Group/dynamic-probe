{{define "secFunction"}}
{{range $index, $element := .Functions}}
{{if eq $element.Aim "BpfFuncTypeKprobeSampleReturn"}}
// 定义 BPF 程序的入口点
SEC("kprobe/{{$element.Name}}")
int kprobe_sys_open(struct pt_regs *ctx) {
    // 获取函数返回值
    int ret = PT_REGS_RC(ctx);
    
    // 输出函数返回值到内核日志
    bpf_printk("{{$element.Name}} returned: %d\n", ret);
    
    return 0;
}
    {{end}}
{{end}}
{{end}}