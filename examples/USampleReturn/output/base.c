
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



    
SEC("uprobe:check_packets")
int uprobe_function_sample_return(struct pt_regs *ctx) {
    u64 retval = PT_REGS_RC(ctx); // 获取函数返回值

    // 输出函数返回值到内核日志
    bpf_printk("check_packets returned: %lld\n", retval);

    return 0;
}
    

    





