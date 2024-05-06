{{define "secFunction"}}
{{range $index, $element := .Functions}}
    {{if eq $element.Aim "BpfFuncTypeUprobeTime"}}
// 定义 BPF 程序的入口点
SEC("uprobe:{{$element.Name}}")
int uprobe_function_time(struct pt_regs *ctx) {
    u64 start_time = bpf_ktime_get_ns();
    u64 end_time = bpf_ktime_get_ns();
    u64 delta = end_time - start_time;

    // 输出执行时间到内核日志
    bpf_printk("{{$element.Name}} execution time: %lld ns\n", delta);

    return 0;
}
    {{end}}
{{end}}
{{end}}