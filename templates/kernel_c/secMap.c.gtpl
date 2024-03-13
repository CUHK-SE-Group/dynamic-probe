{{define "secMap"}}
{{range $key, $values := .Maps }}
struct {
    __uint(type, {{.Type}});
    __type(key, {{.Key}});
    __type(value, {{.Value}});
    __uint(max_entries, {{.MaxEntries}});
} {{.Name}} SEC(".maps");
{{end}}
{{end}}
