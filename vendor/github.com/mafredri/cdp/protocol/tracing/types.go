// Code generated by cdpgen. DO NOT EDIT.

package tracing

import (
	"encoding/json"
	"errors"
)

// MemoryDumpConfig Configuration for memory dump. Used only when
// "memory-infra" category is enabled.
//
// Note: This type is experimental.
type MemoryDumpConfig []byte

// MarshalJSON copies behavior of json.RawMessage.
func (m MemoryDumpConfig) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	return m, nil
}

// UnmarshalJSON copies behavior of json.RawMessage.
func (m *MemoryDumpConfig) UnmarshalJSON(data []byte) error {
	if m == nil {
		return errors.New("tracing.MemoryDumpConfig: UnmarshalJSON on nil pointer")
	}
	*m = append((*m)[0:0], data...)
	return nil
}

var _ json.Marshaler = (*MemoryDumpConfig)(nil)
var _ json.Unmarshaler = (*MemoryDumpConfig)(nil)

// TraceConfig
type TraceConfig struct {
	// RecordMode Controls how the trace buffer stores data.
	//
	// Values: "recordUntilFull", "recordContinuously", "recordAsMuchAsPossible", "echoToConsole".
	//
	// Note: This property is experimental.
	RecordMode *string `json:"recordMode,omitempty"`
	// TraceBufferSizeInKb Size of the trace buffer in kilobytes. If not
	// specified or zero is passed, a default value of 200 MB would be
	// used.
	//
	// Note: This property is experimental.
	TraceBufferSizeInKb *float64 `json:"traceBufferSizeInKb,omitempty"`
	// EnableSampling Turns on JavaScript stack sampling.
	//
	// Note: This property is experimental.
	EnableSampling *bool `json:"enableSampling,omitempty"`
	// EnableSystrace Turns on system tracing.
	//
	// Note: This property is experimental.
	EnableSystrace *bool `json:"enableSystrace,omitempty"`
	// EnableArgumentFilter Turns on argument filter.
	//
	// Note: This property is experimental.
	EnableArgumentFilter *bool    `json:"enableArgumentFilter,omitempty"`
	IncludedCategories   []string `json:"includedCategories,omitempty"` // Included category filters.
	ExcludedCategories   []string `json:"excludedCategories,omitempty"` // Excluded category filters.
	// SyntheticDelays Configuration to synthesize the delays in tracing.
	//
	// Note: This property is experimental.
	SyntheticDelays []string `json:"syntheticDelays,omitempty"`
	// MemoryDumpConfig Configuration for memory dump triggers. Used only
	// when "memory-infra" category is enabled.
	//
	// Note: This property is experimental.
	MemoryDumpConfig MemoryDumpConfig `json:"memoryDumpConfig,omitempty"`
}

// StreamFormat Data format of a trace. Can be either the legacy JSON format
// or the protocol buffer format. Note that the JSON format will be deprecated
// soon.
//
// Note: This type is experimental.
type StreamFormat string

// StreamFormat as enums.
const (
	StreamFormatNotSet StreamFormat = ""
	StreamFormatJSON   StreamFormat = "json"
	StreamFormatProto  StreamFormat = "proto"
)

func (e StreamFormat) Valid() bool {
	switch e {
	case "json", "proto":
		return true
	default:
		return false
	}
}

func (e StreamFormat) String() string {
	return string(e)
}

// StreamCompression Compression type to use for traces returned via streams.
//
// Note: This type is experimental.
type StreamCompression string

// StreamCompression as enums.
const (
	StreamCompressionNotSet StreamCompression = ""
	StreamCompressionNone   StreamCompression = "none"
	StreamCompressionGzip   StreamCompression = "gzip"
)

func (e StreamCompression) Valid() bool {
	switch e {
	case "none", "gzip":
		return true
	default:
		return false
	}
}

func (e StreamCompression) String() string {
	return string(e)
}

// MemoryDumpLevelOfDetail Details exposed when memory request explicitly
// declared. Keep consistent with memory_dump_request_args.h and
// memory_instrumentation.mojom
//
// Note: This type is experimental.
type MemoryDumpLevelOfDetail string

// MemoryDumpLevelOfDetail as enums.
const (
	MemoryDumpLevelOfDetailNotSet     MemoryDumpLevelOfDetail = ""
	MemoryDumpLevelOfDetailBackground MemoryDumpLevelOfDetail = "background"
	MemoryDumpLevelOfDetailLight      MemoryDumpLevelOfDetail = "light"
	MemoryDumpLevelOfDetailDetailed   MemoryDumpLevelOfDetail = "detailed"
)

func (e MemoryDumpLevelOfDetail) Valid() bool {
	switch e {
	case "background", "light", "detailed":
		return true
	default:
		return false
	}
}

func (e MemoryDumpLevelOfDetail) String() string {
	return string(e)
}

// Backend Backend type to use for tracing. `chrome` uses the
// Chrome-integrated tracing service and is supported on all platforms.
// `system` is only supported on Chrome OS and uses the Perfetto system tracing
// service. `auto` chooses `system` when the perfettoConfig provided to
// Tracing.start specifies at least one non-Chrome data source; otherwise uses
// `chrome`.
//
// Note: This type is experimental.
type Backend string

// Backend as enums.
const (
	BackendNotSet Backend = ""
	BackendAuto   Backend = "auto"
	BackendChrome Backend = "chrome"
	BackendSystem Backend = "system"
)

func (e Backend) Valid() bool {
	switch e {
	case "auto", "chrome", "system":
		return true
	default:
		return false
	}
}

func (e Backend) String() string {
	return string(e)
}