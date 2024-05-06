
//go:build ignore

#include "vmlinux.h"
#include "bpf_tracing.h"
#include "bpf_helpers.h"
#include "bpf_core_read.h" 


char __license[] SEC("license") = "Dual MIT/GPL";



struct count_packets_args {
    
    u32 destination_ip;
    u32 packet_size;
    u32 source_ip;
};
const struct count_packets_args *unused_count_packets_args __attribute__((unused));
struct check_packets_args {
    
    u64 total_bytes;
    u64 total_packets;
};
const struct check_packets_args *unused_check_packets_args __attribute__((unused));


struct {
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, u32);
    __type(value, u64);
    __uint(max_entries, 1024);
} packet_count_map SEC(".maps");

struct {
    __uint(type, BPF_MAP_TYPE_ARRAY_OF_MAPS);
    __type(key, u32);
    __type(value, u64);
    __uint(max_entries, 10);
} array_of_packet_maps SEC(".maps");

struct {
    __uint(type, BPF_MAP_TYPE_ARRAY_OF_MAPS);
    __type(key, u32);
    __type(value, u64);
    __uint(max_entries, 10);
} args_map_0 SEC(".maps");

struct {
    __uint(type, BPF_MAP_TYPE_ARRAY_OF_MAPS);
    __type(key, u32);
    __type(value, u64);
    __uint(max_entries, 10);
} start_time_map SEC(".maps");

struct {
    __uint(type, BPF_MAP_TYPE_ARRAY_OF_MAPS);
    __type(key, u32);
    __type(value, u64);
    __uint(max_entries, 10);
} args_map_1 SEC(".maps");




SEC("uprobe/check_packets")
int bpf_prog_check_packets(struct pt_regs *ctx) {
    u32 pid = bpf_get_current_pid_tgid();
    u64 ts = bpf_ktime_get_ns();
    struct check_packets_args args = {};


    bpf_map_update_elem(&start_time_map, &pid, &ts, BPF_ANY);
    bpf_map_update_elem(&args_map_0, &pid, &args, BPF_ANY);
    return 0;
}

SEC("uprobe/count_packets")
int bpf_prog_count_packets(struct pt_regs *ctx) {
    u32 pid = bpf_get_current_pid_tgid();
    u64 ts = bpf_ktime_get_ns();
    struct count_packets_args args = {};


    bpf_map_update_elem(&start_time_map, &pid, &ts, BPF_ANY);
    bpf_map_update_elem(&args_map_1, &pid, &args, BPF_ANY);
    return 0;
}





