// Code generated by cdpgen. DO NOT EDIT.

package profiler

import (
	"github.com/mafredri/cdp/protocol/debugger"
	"github.com/mafredri/cdp/rpcc"
)

// ConsoleProfileFinishedClient is a client for ConsoleProfileFinished events.
type ConsoleProfileFinishedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ConsoleProfileFinishedReply, error)
	rpcc.Stream
}

// ConsoleProfileFinishedReply is the reply for ConsoleProfileFinished events.
type ConsoleProfileFinishedReply struct {
	ID       string            `json:"id"`              // No description.
	Location debugger.Location `json:"location"`        // Location of console.profileEnd().
	Profile  Profile           `json:"profile"`         // No description.
	Title    *string           `json:"title,omitempty"` // Profile title passed as an argument to console.profile().
}

// ConsoleProfileStartedClient is a client for ConsoleProfileStarted events.
// Sent when new profile recording is started using console.profile() call.
type ConsoleProfileStartedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ConsoleProfileStartedReply, error)
	rpcc.Stream
}

// ConsoleProfileStartedReply is the reply for ConsoleProfileStarted events.
type ConsoleProfileStartedReply struct {
	ID       string            `json:"id"`              // No description.
	Location debugger.Location `json:"location"`        // Location of console.profile().
	Title    *string           `json:"title,omitempty"` // Profile title passed as an argument to console.profile().
}

// PreciseCoverageDeltaUpdateClient is a client for PreciseCoverageDeltaUpdate events.
// Reports coverage delta since the last poll (either from an event like this,
// or from `takePreciseCoverage` for the current isolate. May only be sent if
// precise code coverage has been started. This event can be trigged by the
// embedder to, for example, trigger collection of coverage data immediately at
// a certain point in time.
type PreciseCoverageDeltaUpdateClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*PreciseCoverageDeltaUpdateReply, error)
	rpcc.Stream
}

// PreciseCoverageDeltaUpdateReply is the reply for PreciseCoverageDeltaUpdate events.
type PreciseCoverageDeltaUpdateReply struct {
	Timestamp float64          `json:"timestamp"` // Monotonically increasing time (in seconds) when the coverage update was taken in the backend.
	Occasion  string           `json:"occasion"`  // Identifier for distinguishing coverage events.
	Result    []ScriptCoverage `json:"result"`    // Coverage data for the current isolate.
}