{{define "secInclude"}}
//go:build ignore
#include "vmlinux.h"
#include "bpf_tracing.h"
#include "bpf_helpers.h"
#include "bpf_core_read.h" {{end}}
