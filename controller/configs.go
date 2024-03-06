package controller

type BpfRuntimeConfig struct {
	EBPFProgram EBPFProgram `toml:"eBPFProgram"`
}

// This struct will be sent to the template directly
type EBPFProgram struct {
	MaxEntries     int      `toml:"MaxEntries"`
	TargetFunction []string `toml:"TargetFunction"`
	FunctionName   []string `toml:"FunctionName"`
}
