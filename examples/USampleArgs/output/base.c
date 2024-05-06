
//go:build ignore
#include "vmlinux.h"
#include "bpf_tracing.h"
#include "bpf_helpers.h"
#include "bpf_core_read.h" 
char __license[] SEC("license") = "Dual MIT/GPL";
//eBPF Event Definitions
struct packet_data {
    u32 destination_ip;
    u32 packet_size;
    u32 source_ip;
    
};
const struct packet_data *unused_packet_data __attribute__((unused));

// eBPF maps
struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, u32);
    __type(value, u64);
    __uint(max_entries, 1);
} function_count SEC(".maps");

struct {
    __uint(type, BPF_MAP_TYPE_ARRAY_OF_MAPS);
    __type(key, u32);
    __type(value, u64);
    __uint(max_entries, 10);
} array_of_packet_maps SEC(".maps");



    
SEC("uprobe:Uarg")
int uprobe_function_sample_args(struct pt_regs *ctx) {
    // 定义参数变量
    u64 arg1, arg2, arg3;

    // 从用户空间内存中读取函数参数
    bpf_probe_read_user(&arg1, sizeof(arg1), (void *)(ctx->di));
    bpf_probe_read_user(&arg2, sizeof(arg2), (void *)(ctx->si));
    bpf_probe_read_user(&arg3, sizeof(arg3), (void *)(ctx->dx));

    // 输出函数参数到内核日志
    bpf_printk("Uarg: arg1=%lld, arg2=%lld, arg3=%lld\n", arg1, arg2, arg3);

    return 0;
}

    

    





