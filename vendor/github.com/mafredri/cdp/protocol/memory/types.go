// Code generated by cdpgen. DO NOT EDIT.

package memory

// PressureLevel Memory pressure level.
type PressureLevel string

// PressureLevel as enums.
const (
	PressureLevelNotSet   PressureLevel = ""
	PressureLevelModerate PressureLevel = "moderate"
	PressureLevelCritical PressureLevel = "critical"
)

func (e PressureLevel) Valid() bool {
	switch e {
	case "moderate", "critical":
		return true
	default:
		return false
	}
}

func (e PressureLevel) String() string {
	return string(e)
}

// SamplingProfileNode Heap profile sample.
type SamplingProfileNode struct {
	Size  float64  `json:"size"`  // Size of the sampled allocation.
	Total float64  `json:"total"` // Total bytes attributed to this sample.
	Stack []string `json:"stack"` // Execution stack at the point of allocation.
}

// SamplingProfile Array of heap profile samples.
type SamplingProfile struct {
	Samples []SamplingProfileNode `json:"samples"` // No description.
	Modules []Module              `json:"modules"` // No description.
}

// Module Executable module information
type Module struct {
	Name        string  `json:"name"`        // Name of the module.
	UUID        string  `json:"uuid"`        // UUID of the module.
	BaseAddress string  `json:"baseAddress"` // Base address where the module is loaded into memory. Encoded as a decimal or hexadecimal (0x prefixed) string.
	Size        float64 `json:"size"`        // Size of the module in bytes.
}

// DOMCounter DOM object counter data.
type DOMCounter struct {
	Name  string `json:"name"`  // Object name. Note: object names should be presumed volatile and clients should not expect the returned names to be consistent across runs.
	Count int    `json:"count"` // Object count.
}