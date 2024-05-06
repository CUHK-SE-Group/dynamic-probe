{{define "secEventStruct"}}//eBPF Event Definitions{{range $key, $values := .Structs }}
struct {{.Name}} {
    {{range $key, $value := .Members }}{{$value}} {{$key}};
    {{end}}
};
{{if .ForTransmission}}const struct {{.Name}} *unused_{{.Name}} __attribute__((unused));
{{end}}{{end}}{{end}}