{{define "secFunction"}}
{{range $index, $element := .Functions}}
    {{if eq $element.Aim "BpfFuncTypeUprobeCount"}}
// 定义 BPF 程序的入口点
SEC("uprobe:{{$element.Name}}")
int uprobe_function_count(struct pt_regs *ctx) {
    // 定义一个 BPF 映射用于统计调用次数
    BPF_MAP_DEF(count_map) = {
        .map_type = BPF_MAP_TYPE_HASH,
        .key_size = sizeof(u64),
        .value_size = sizeof(u64),
        .max_entries = 1,
    };

    // 将函数调用次数存储在 BPF 映射中
    u64 *count = bpf_map_lookup_elem(&count_map, &ctx->ip);
    if (count) {
        (*count)++;
    } else {
        u64 initial_count = 1;
        bpf_map_update_elem(&count_map, &ctx->ip, &initial_count, BPF_ANY);
    }

    return 0;
}
    {{end}}
{{end}}
{{end}}

