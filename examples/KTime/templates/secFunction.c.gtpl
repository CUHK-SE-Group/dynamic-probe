{{define "secFunction"}}
{{range $index, $element := .Functions}}
    {{if eq $element.Aim "BpfFuncTypeKprobeTime"}}
// 定义 BPF 程序的入口点
SEC("kprobe/{{$element.Name}}")
int kprobe_{{$element.Name}}(struct pt_regs *ctx) {
    // 记录当前时间
    u64 timestamp = bpf_ktime_get_ns();
    // 将时间戳保存到 BPF 映射中，以便用户空间程序访问
    bpf_map_update_elem(&timestamps, &zero, &timestamp, BPF_ANY);
    return 0;
}

// 定义 BPF 程序的出口点
SEC("kretprobe/{{$element.Name}}")
int kretprobe_{{$element.Name}}(struct pt_regs *ctx) {
    // 从 BPF 映射中获取之前保存的时间戳
    u64 *timestamp = bpf_map_lookup_elem(&timestamps, &zero);
    if (timestamp) {
        // 计算函数执行时间
        u64 delta = bpf_ktime_get_ns() - *timestamp;
        // 输出执行时间（可以替换为其他输出方式，如使用 perf_event 输出到用户空间）
        bpf_printk("{{$element.Name}} execution time: %lld ns\n", delta);
        // 删除映射中的时间戳，准备下一次记录
        bpf_map_delete_elem(&timestamps, &zero);
    }
    return 0;
}
    {{end}}
{{end}}
{{end}}