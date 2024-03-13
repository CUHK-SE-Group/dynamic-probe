package executor

import (
	"fmt"

	"github.com/cilium/ebpf/btf"
)

type MapType string
type Aim string

const (
	BpfMapTypeArray         MapType = "BPF_MAP_TYPE_ARRAY"
	BpfMapTypeHashOfMaps    MapType = "BPF_MAP_TYPE_HASH_OF_MAPS"
	BpfMapTypeArrayOfMaps   MapType = "BPF_MAP_TYPE_ARRAY_OF_MAPS"
	BpfMapTypeQueue         MapType = "BPF_MAP_TYPE_QUEUE"
	BpfMapTypeStack         MapType = "BPF_MAP_TYPE_STACK"
	BpfMapTypePercpuHash    MapType = "BPF_MAP_TYPE_PERCPU_HASH"
	BpfMapTypeLruHash       MapType = "BPF_MAP_TYPE_LRU_HASH"
	BpfMapTypeLruPercpuHash MapType = "BPF_MAP_TYPE_LRU_PERCPU_HASH"
)

const (
	BpfFuncTypeKprobeTime         Aim = "BpfFuncTypeKprobeTime"
	BpfFuncTypeKprobeCount        Aim = "BpfFuncTypeKprobeCount"
	BpfFuncTypeKprobeSampleArgs   Aim = "BpfFuncTypeKprobeSampleArgs"
	BpfFuncTypeKprobeSampleReturn Aim = "BpfFuncTypeKprobeSampleReturn"

	BpfFuncTypeUprobeTime         Aim = "BpfFuncTypeUprobeTime"
	BpfFuncTypeUprobeCount        Aim = "BpfFuncTypeUprobeCount"
	BpfFuncTypeUprobeSampleArgs   Aim = "BpfFuncTypeUprobeSampleArgs"
	BpfFuncTypeUprobeSampleReturn Aim = "BpfFuncTypeUprobeSampleReturn"
)

// This struct will be sent to the template directly
type EBPFProgram struct {
	Functions []FuncDef
	Maps      []EBPFMap
	Structs   []Struct
}

// the templates for maps
type EBPFMap struct {
	Type       MapType
	Key        string
	Value      string
	MaxEntries int

	Name string
}

// the template for structs
type Struct struct {
	Members map[string]string
	Name    string

	ForTransmission bool
}

type FuncDef struct {
	Btf  *btf.FuncProto `toml:"-"`
	Name string
	Aim  Aim
}

func (f *FuncDef) GetFuncName() string {
	return f.Name
}

func (f *FuncDef) GetParams() []string {
	results := []string{}
	for _, v := range f.Btf.Params {
		results = append(results, f.btf2C(v.Type))
	}
	return results
}

func (f *FuncDef) GetReturn() string {
	return f.btf2C(f.Btf.Return)
}

func (f *FuncDef) btf2C(tp btf.Type) string {
	fmt.Println("type: ", tp)
	return tp.TypeName()
}
