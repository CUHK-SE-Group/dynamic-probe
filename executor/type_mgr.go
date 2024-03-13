package executor

import (
	"fmt"
	"strings"

	"github.com/cilium/ebpf/btf"
)

type FuncDefs struct {
	Funcs map[string]*btf.FuncProto
}

func ParseFuncs(spec *btf.Spec) *FuncDefs {
	f := &FuncDefs{
		Funcs: make(map[string]*btf.FuncProto),
	}
	iter := spec.Iterate()
	for iter.Next() {
		if ff, ok := iter.Type.(*btf.Func); ok {
			f.Funcs[ff.Name] = ff.Type.(*btf.FuncProto)
		}
	}
	return f
}

func (f *FuncDefs) GenerateGoStructDef() {

}

func (f *FuncDefs) GenerateCStructDef() {

}

func GenerateGo(spec *btf.Spec, structName string, packageName string) *strings.Builder {
	sb := &strings.Builder{}
	sb.WriteString(fmt.Sprintf("package %s\n", packageName))
	sb.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	// for _, field := range spec.Fields {
	// 	sb.WriteString(fmt.Sprintf("    %s %s\n", field.Name, field.Type))
	// }

	iter := spec.Iterate()
	for iter.Next() {
		switch v := iter.Type.(type) {
		case *btf.Func:
			fmt.Printf("Func, name=%v\n", v.Name)
			proto := v.Type.(*btf.FuncProto) // 断言Func类型为FuncProto
			fmt.Printf("Return type: %s\n", typeName(proto.Return))
			fmt.Println("Parameters:")
			for _, param := range proto.Params {
				fmt.Printf("  %s: %s\n", param.Name, typeName(param.Type))
			}
		}
	}
	sb.WriteString("}\n")
	return sb
}

func typeName(t btf.Type) string {
	switch v := t.(type) {
	case *btf.Int:
		return v.Name
	case *btf.Pointer:
		// 递归处理指针指向的类型
		return "*" + typeName(v.Target)
	case *btf.Array:
		// 递归处理数组元素的类型
		return fmt.Sprintf("Array of %s", typeName(v.Type))
	case *btf.Struct:
		return "struct " + v.Name
	case *btf.Union:
		return "union " + v.Name
	case *btf.Enum:
		return "enum " + v.Name
	case *btf.Typedef:
		return "typedef " + v.Name
	case *btf.Const:
		return "const " + typeName(v.Type)
	default:
		return fmt.Sprintf("unknown(%T)", t)
	}
}
