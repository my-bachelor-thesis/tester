package structs

type OutgoingJson struct {
	ExitCode        int     `json:"exit_code"`
	Output          string  `json:"output"`
	CompilationTime float32 `json:"compilation_time"` // in s
	RealTime        float32 `json:"real_time"`        // in s
	KernelTime      float32 `json:"kernel_time"`      // in s
	UserTime        float32 `json:"user_time"`        // in s
	MaxRamUsage     float32 `json:"max_ram_usage"`    // in MB
	BinarySize      float32 `json:"binary_size"`      // in MB
}
