{{define "secFunction2"}}
{{range $index, $element := .Functions}}
    {{if eq $element.Aim "BpfFuncTypeKprobeTime"}}
// 定义 BPF 程序的入口点
SEC("kprobe/{{$element.Name}}")
int kprobe_{{$element.Name}}(struct pt_regs *ctx) {
    // 获取当前时间戳
    u64 start_time = bpf_ktime_get_ns();
    // 将时间戳保存到 BPF 上下文中
    bpf_context_store(&ctx->start_time, start_time);
    return 0;
}

// 定义 BPF 程序的出口点
SEC("kretprobe/{{$element.Name}}")
int kretprobe_{{$element.Name}}(struct pt_regs *ctx) {
    // 获取之前保存的时间戳
    u64 start_time;
    bpf_context_load(&ctx->start_time, &start_time);
    // 计算函数执行时间
    u64 end_time = bpf_ktime_get_ns();
    u64 delta = end_time - start_time;
    // 输出执行时间（可以替换为其他输出方式，如使用 perf_event 输出到用户空间）
    bpf_printk("{{$element.Name}} execution time: %lld ns\n", delta);
    return 0;
}

    {{end}}
{{end}}
{{end}}



