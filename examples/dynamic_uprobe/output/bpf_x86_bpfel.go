// Code generated by bpf2go; DO NOT EDIT.
//go:build 386 || amd64

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfDemonstrateDynamicMemoryArgs struct{ Arg0 uint64 }

type bpfDemonstratePointersArgs struct{ Arg0 uint64 }

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	BpfProgDemonstrateDynamicMemory    *ebpf.ProgramSpec `ebpf:"bpf_prog_demonstrateDynamicMemory"`
	BpfProgDemonstratePointers         *ebpf.ProgramSpec `ebpf:"bpf_prog_demonstratePointers"`
	BpfProgRetDemonstrateDynamicMemory *ebpf.ProgramSpec `ebpf:"bpf_prog_ret_demonstrateDynamicMemory"`
	BpfProgRetDemonstratePointers      *ebpf.ProgramSpec `ebpf:"bpf_prog_ret_demonstratePointers"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	DemonstrateDynamicMemoryArgsMap *ebpf.MapSpec `ebpf:"demonstrateDynamicMemory_args_map"`
	DemonstratePointersArgsMap      *ebpf.MapSpec `ebpf:"demonstratePointers_args_map"`
	Events                          *ebpf.MapSpec `ebpf:"events"`
	StartTimeMap                    *ebpf.MapSpec `ebpf:"start_time_map"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	DemonstrateDynamicMemoryArgsMap *ebpf.Map `ebpf:"demonstrateDynamicMemory_args_map"`
	DemonstratePointersArgsMap      *ebpf.Map `ebpf:"demonstratePointers_args_map"`
	Events                          *ebpf.Map `ebpf:"events"`
	StartTimeMap                    *ebpf.Map `ebpf:"start_time_map"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.DemonstrateDynamicMemoryArgsMap,
		m.DemonstratePointersArgsMap,
		m.Events,
		m.StartTimeMap,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	BpfProgDemonstrateDynamicMemory    *ebpf.Program `ebpf:"bpf_prog_demonstrateDynamicMemory"`
	BpfProgDemonstratePointers         *ebpf.Program `ebpf:"bpf_prog_demonstratePointers"`
	BpfProgRetDemonstrateDynamicMemory *ebpf.Program `ebpf:"bpf_prog_ret_demonstrateDynamicMemory"`
	BpfProgRetDemonstratePointers      *ebpf.Program `ebpf:"bpf_prog_ret_demonstratePointers"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.BpfProgDemonstrateDynamicMemory,
		p.BpfProgDemonstratePointers,
		p.BpfProgRetDemonstrateDynamicMemory,
		p.BpfProgRetDemonstratePointers,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_x86_bpfel.o
var _BpfBytes []byte
