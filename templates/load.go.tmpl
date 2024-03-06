package main

import (
	"log"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)


func main() {
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal(err)
	}

	objs := bpfObjects{}
	if err := loadBpfObjects(&objs, nil); err != nil {
		log.Fatalf("loading objects: %v", err)
	}
	defer objs.Close()

	{{ range $index, $element := .TargetFunction }}
	kp{{$index}}, err{{$index}} := link.Kprobe("{{$element}}", objs.Gen{{Transform $element}}, nil)
	if err{{$index}} != nil {
		log.Fatalf("opening kprobe: %s", err{{$index}})
	}
	defer kp{{$index}}.Close()
	{{end}}


	{{ range $index, $element := .TargetFunction }}
	retkp{{$index}}, retErr{{$index}} := link.Kprobe("{{$element}}", objs.GenRet{{Transform $element}}, nil)
	if retErr{{$index}} != nil {
		log.Fatalf("opening kprobe: %s", err{{$index}})
	}
	defer retkp{{$index}}.Close()
	{{end}}

	log.Println("Waiting for events..")
	for {

	}
}
