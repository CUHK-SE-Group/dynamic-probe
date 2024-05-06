
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
SEC("kprobe/check_packets")
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
    

    





