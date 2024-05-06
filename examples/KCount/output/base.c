
//go:build ignore

#include "vmlinux.h"
#include "bpf_tracing.h"
#include "bpf_helpers.h"
#include "bpf_core_read.h" 


char __license[] SEC("license") = "Dual MIT/GPL";



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
SEC("kprobe/count_packets")
int kprobe_sys_read(struct pt_regs *ctx) {
    // 增加计数器
    __u64 zero = 0;
    bpf_map_increment_elem(&function_count, &zero);
    return 0;
}
    





