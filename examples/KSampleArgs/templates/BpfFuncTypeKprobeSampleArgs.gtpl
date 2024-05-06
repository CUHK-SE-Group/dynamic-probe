{{define "secFunction"}}
{{range $index, $element := .Functions}}
    {{if eq $element.Aim "BpfFuncTypeKprobeSampleArgs"}}
// 定义 BPF 程序的入口点
SEC("kprobe/{{$element.Name}}")
int kprobe_sys_open(struct pt_regs *ctx) {
    // 获取第一个参数
    u64 arg1 = PT_REGS_PARM1(ctx);
    // 获取第二个参数
    u64 arg2 = PT_REGS_PARM2(ctx);
    // 获取第三个参数
    u64 arg3 = PT_REGS_PARM3(ctx);
    // 获取第四个参数
    u64 arg4 = PT_REGS_PARM4(ctx);
    
    // 输出参数
    bpf_printk("sys_open args: %lld, %lld, %lld, %lld\n", arg1, arg2, arg3, arg4);
    
    return 0;
}
    {{end}}
{{end}}
{{end}}