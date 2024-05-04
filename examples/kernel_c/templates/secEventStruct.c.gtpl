{{define "secEventStruct"}}
{{range $key, $values := .Structs }}
struct {{.Name}} {
    {{range $key, $value := .Members }}
    {{$value}} {{$key}};{{end}}
};
{{if .Members}}const struct {{.Name}} *unused_{{.Name}} __attribute__((unused));{{end}}{{end}}{{end}}
