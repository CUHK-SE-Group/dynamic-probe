package controller

import "kprobe/executor"

type BpfRuntimeConfig struct {
	EBPFProgram  executor.EBPFProgram `toml:"eBPFProgram"`
	ObjectBinary string
}
