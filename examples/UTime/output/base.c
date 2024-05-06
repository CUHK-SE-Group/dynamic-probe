
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



    
// 定义 BPF 程序的入口点
SEC("uprobe:Utime")
int uprobe_function_time(struct pt_regs *ctx) {
    u64 start_time = bpf_ktime_get_ns();
    u64 end_time = bpf_ktime_get_ns();
    u64 delta = end_time - start_time;

    // 输出执行时间到内核日志
    bpf_printk("Utime execution time: %lld ns\n", delta);

    return 0;
}
    

    





