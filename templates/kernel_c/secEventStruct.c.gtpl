{{define "secEventStruct"}}
{{range $key, $values := .TargetFunction }}
struct {{$key}}_func_event_data {
    u32 pid;
    u64 start_time;
    long ret;
    u64 duration;
    {{ range $index, $value := $values }}
    u64 arg{{$index}}; {{ end }}
};
const struct {{$key}}_func_event_data *unused_{{$key}} __attribute__((unused));
{{end}}
{{end}}
