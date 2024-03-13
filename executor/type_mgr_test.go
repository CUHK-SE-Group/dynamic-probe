package executor

import (
	"fmt"
	"testing"

	"github.com/cilium/ebpf/btf"
)

func TestLoad(t *testing.T) {
	s, err := btf.LoadSpec("./victim")
	if err != nil {
		panic(err)
	}
	f := ParseFuncs(s)
	_ = f
	fmt.Println(f)
}
