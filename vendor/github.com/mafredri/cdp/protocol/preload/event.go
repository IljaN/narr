// Code generated by cdpgen. DO NOT EDIT.

package preload

import (
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/protocol/page"
	"github.com/mafredri/cdp/rpcc"
)

// RuleSetUpdatedClient is a client for RuleSetUpdated events. Upsert.
// Currently, it is only emitted when a rule set added.
type RuleSetUpdatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*RuleSetUpdatedReply, error)
	rpcc.Stream
}

// RuleSetUpdatedReply is the reply for RuleSetUpdated events.
type RuleSetUpdatedReply struct {
	RuleSet RuleSet `json:"ruleSet"` // No description.
}

// RuleSetRemovedClient is a client for RuleSetRemoved events.
type RuleSetRemovedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*RuleSetRemovedReply, error)
	rpcc.Stream
}

// RuleSetRemovedReply is the reply for RuleSetRemoved events.
type RuleSetRemovedReply struct {
	ID RuleSetID `json:"id"` // No description.
}

// EnabledStateUpdatedClient is a client for PreloadEnabledStateUpdated events.
// Fired when a preload enabled state is updated.
type EnabledStateUpdatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*EnabledStateUpdatedReply, error)
	rpcc.Stream
}

// EnabledStateUpdatedReply is the reply for PreloadEnabledStateUpdated events.
type EnabledStateUpdatedReply struct {
	DisabledByPreference                        bool `json:"disabledByPreference"`                        // No description.
	DisabledByDataSaver                         bool `json:"disabledByDataSaver"`                         // No description.
	DisabledByBatterySaver                      bool `json:"disabledByBatterySaver"`                      // No description.
	DisabledByHoldbackPrefetchSpeculationRules  bool `json:"disabledByHoldbackPrefetchSpeculationRules"`  // No description.
	DisabledByHoldbackPrerenderSpeculationRules bool `json:"disabledByHoldbackPrerenderSpeculationRules"` // No description.
}

// PrefetchStatusUpdatedClient is a client for PrefetchStatusUpdated events.
// Fired when a prefetch attempt is updated.
type PrefetchStatusUpdatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*PrefetchStatusUpdatedReply, error)
	rpcc.Stream
}

// PrefetchStatusUpdatedReply is the reply for PrefetchStatusUpdated events.
type PrefetchStatusUpdatedReply struct {
	Key               AttemptKey        `json:"key"`               // No description.
	InitiatingFrameID page.FrameID      `json:"initiatingFrameId"` // The frame id of the frame initiating prefetch.
	PrefetchURL       string            `json:"prefetchUrl"`       // No description.
	Status            Status            `json:"status"`            // No description.
	PrefetchStatus    PrefetchStatus    `json:"prefetchStatus"`    // No description.
	RequestID         network.RequestID `json:"requestId"`         // No description.
}

// PrerenderStatusUpdatedClient is a client for PrerenderStatusUpdated events.
// Fired when a prerender attempt is updated.
type PrerenderStatusUpdatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*PrerenderStatusUpdatedReply, error)
	rpcc.Stream
}

// PrerenderStatusUpdatedReply is the reply for PrerenderStatusUpdated events.
type PrerenderStatusUpdatedReply struct {
	Key                     AttemptKey                   `json:"key"`                               // No description.
	Status                  Status                       `json:"status"`                            // No description.
	PrerenderStatus         PrerenderFinalStatus         `json:"prerenderStatus,omitempty"`         // No description.
	DisallowedMojoInterface *string                      `json:"disallowedMojoInterface,omitempty"` // This is used to give users more information about the name of Mojo interface that is incompatible with prerender and has caused the cancellation of the attempt.
	MismatchedHeaders       []PrerenderMismatchedHeaders `json:"mismatchedHeaders,omitempty"`       // No description.
}

// AttemptSourcesUpdatedClient is a client for PreloadingAttemptSourcesUpdated events.
// Send a list of sources for all preloading attempts in a document.
type AttemptSourcesUpdatedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*AttemptSourcesUpdatedReply, error)
	rpcc.Stream
}

// AttemptSourcesUpdatedReply is the reply for PreloadingAttemptSourcesUpdated events.
type AttemptSourcesUpdatedReply struct {
	LoaderID                 network.LoaderID `json:"loaderId"`                 // No description.
	PreloadingAttemptSources []AttemptSource  `json:"preloadingAttemptSources"` // No description.
}
