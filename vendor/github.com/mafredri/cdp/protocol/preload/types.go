// Code generated by cdpgen. DO NOT EDIT.

package preload

import (
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/network"
)

// RuleSetID Unique id
type RuleSetID string

// RuleSet Corresponds to SpeculationRuleSet
type RuleSet struct {
	ID            RuleSetID          `json:"id"`                      // No description.
	LoaderID      network.LoaderID   `json:"loaderId"`                // Identifies a document which the rule set is associated with.
	SourceText    string             `json:"sourceText"`              // Source text of JSON representing the rule set. If it comes from `<script>` tag, it is the textContent of the node. Note that it is a JSON for valid case. See also: - https://wicg.github.io/nav-speculation/speculation-rules.html - https://github.com/WICG/nav-speculation/blob/main/triggers.md
	BackendNodeID *dom.BackendNodeID `json:"backendNodeId,omitempty"` // A speculation rule set is either added through an inline `<script>` tag or through an external resource via the 'Speculation-Rules' HTTP header. For the first case, we include the BackendNodeId of the relevant `<script>` tag. For the second case, we include the external URL where the rule set was loaded from, and also RequestId if Network domain is enabled. See also: - https://wicg.github.io/nav-speculation/speculation-rules.html#speculation-rules-script - https://wicg.github.io/nav-speculation/speculation-rules.html#speculation-rules-header
	URL           *string            `json:"url,omitempty"`           // No description.
	RequestID     *network.RequestID `json:"requestId,omitempty"`     // No description.
	ErrorType     RuleSetErrorType   `json:"errorType,omitempty"`     // Error information `errorMessage` is null iff `errorType` is null.
	// ErrorMessage is deprecated.
	//
	// Deprecated: TODO(https://crbug.com/1425354): Replace this property
	// with structured error.
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// RuleSetErrorType
type RuleSetErrorType string

// RuleSetErrorType as enums.
const (
	RuleSetErrorTypeNotSet                RuleSetErrorType = ""
	RuleSetErrorTypeSourceIsNotJSONObject RuleSetErrorType = "SourceIsNotJsonObject"
	RuleSetErrorTypeInvalidRulesSkipped   RuleSetErrorType = "InvalidRulesSkipped"
)

func (e RuleSetErrorType) Valid() bool {
	switch e {
	case "SourceIsNotJsonObject", "InvalidRulesSkipped":
		return true
	default:
		return false
	}
}

func (e RuleSetErrorType) String() string {
	return string(e)
}

// SpeculationAction The type of preloading attempted. It corresponds to
// mojom::SpeculationAction (although PrefetchWithSubresources is omitted as it
// isn't being used by clients).
type SpeculationAction string

// SpeculationAction as enums.
const (
	SpeculationActionNotSet    SpeculationAction = ""
	SpeculationActionPrefetch  SpeculationAction = "Prefetch"
	SpeculationActionPrerender SpeculationAction = "Prerender"
)

func (e SpeculationAction) Valid() bool {
	switch e {
	case "Prefetch", "Prerender":
		return true
	default:
		return false
	}
}

func (e SpeculationAction) String() string {
	return string(e)
}

// SpeculationTargetHint Corresponds to mojom::SpeculationTargetHint. See
// https://github.com/WICG/nav-speculation/blob/main/triggers.md#window-name-targeting-hints
type SpeculationTargetHint string

// SpeculationTargetHint as enums.
const (
	SpeculationTargetHintNotSet SpeculationTargetHint = ""
	SpeculationTargetHintBlank  SpeculationTargetHint = "Blank"
	SpeculationTargetHintSelf   SpeculationTargetHint = "Self"
)

func (e SpeculationTargetHint) Valid() bool {
	switch e {
	case "Blank", "Self":
		return true
	default:
		return false
	}
}

func (e SpeculationTargetHint) String() string {
	return string(e)
}

// AttemptKey A key that identifies a preloading attempt.
//
// The url used is the url specified by the trigger (i.e. the initial URL),
// and not the final url that is navigated to. For example, prerendering allows
// same-origin main frame navigations during the attempt, but the attempt is
// still keyed with the initial URL.
type AttemptKey struct {
	LoaderID   network.LoaderID      `json:"loaderId"`             // No description.
	Action     SpeculationAction     `json:"action"`               // No description.
	URL        string                `json:"url"`                  // No description.
	TargetHint SpeculationTargetHint `json:"targetHint,omitempty"` // No description.
}

// AttemptSource Lists sources for a preloading attempt, specifically the ids
// of rule sets that had a speculation rule that triggered the attempt, and the
// BackendNodeIds of <a href> or <area href> elements that triggered the
// attempt (in the case of attempts triggered by a document rule). It is
// possible for multiple rule sets and links to trigger a single attempt.
type AttemptSource struct {
	Key        AttemptKey          `json:"key"`        // No description.
	RuleSetIDs []RuleSetID         `json:"ruleSetIds"` // No description.
	NodeIDs    []dom.BackendNodeID `json:"nodeIds"`    // No description.
}

// PrerenderFinalStatus List of FinalStatus reasons for Prerender2.
type PrerenderFinalStatus string

// PrerenderFinalStatus as enums.
const (
	PrerenderFinalStatusNotSet                                                     PrerenderFinalStatus = ""
	PrerenderFinalStatusActivated                                                  PrerenderFinalStatus = "Activated"
	PrerenderFinalStatusDestroyed                                                  PrerenderFinalStatus = "Destroyed"
	PrerenderFinalStatusLowEndDevice                                               PrerenderFinalStatus = "LowEndDevice"
	PrerenderFinalStatusInvalidSchemeRedirect                                      PrerenderFinalStatus = "InvalidSchemeRedirect"
	PrerenderFinalStatusInvalidSchemeNavigation                                    PrerenderFinalStatus = "InvalidSchemeNavigation"
	PrerenderFinalStatusNavigationRequestBlockedByCSP                              PrerenderFinalStatus = "NavigationRequestBlockedByCsp"
	PrerenderFinalStatusMainFrameNavigation                                        PrerenderFinalStatus = "MainFrameNavigation"
	PrerenderFinalStatusMojoBinderPolicy                                           PrerenderFinalStatus = "MojoBinderPolicy"
	PrerenderFinalStatusRendererProcessCrashed                                     PrerenderFinalStatus = "RendererProcessCrashed"
	PrerenderFinalStatusRendererProcessKilled                                      PrerenderFinalStatus = "RendererProcessKilled"
	PrerenderFinalStatusDownload                                                   PrerenderFinalStatus = "Download"
	PrerenderFinalStatusTriggerDestroyed                                           PrerenderFinalStatus = "TriggerDestroyed"
	PrerenderFinalStatusNavigationNotCommitted                                     PrerenderFinalStatus = "NavigationNotCommitted"
	PrerenderFinalStatusNavigationBadHTTPStatus                                    PrerenderFinalStatus = "NavigationBadHttpStatus"
	PrerenderFinalStatusClientCertRequested                                        PrerenderFinalStatus = "ClientCertRequested"
	PrerenderFinalStatusNavigationRequestNetworkError                              PrerenderFinalStatus = "NavigationRequestNetworkError"
	PrerenderFinalStatusCancelAllHostsForTesting                                   PrerenderFinalStatus = "CancelAllHostsForTesting"
	PrerenderFinalStatusDidFailLoad                                                PrerenderFinalStatus = "DidFailLoad"
	PrerenderFinalStatusStop                                                       PrerenderFinalStatus = "Stop"
	PrerenderFinalStatusSSLCertificateError                                        PrerenderFinalStatus = "SslCertificateError"
	PrerenderFinalStatusLoginAuthRequested                                         PrerenderFinalStatus = "LoginAuthRequested"
	PrerenderFinalStatusUAChangeRequiresReload                                     PrerenderFinalStatus = "UaChangeRequiresReload"
	PrerenderFinalStatusBlockedByClient                                            PrerenderFinalStatus = "BlockedByClient"
	PrerenderFinalStatusAudioOutputDeviceRequested                                 PrerenderFinalStatus = "AudioOutputDeviceRequested"
	PrerenderFinalStatusMixedContent                                               PrerenderFinalStatus = "MixedContent"
	PrerenderFinalStatusTriggerBackgrounded                                        PrerenderFinalStatus = "TriggerBackgrounded"
	PrerenderFinalStatusMemoryLimitExceeded                                        PrerenderFinalStatus = "MemoryLimitExceeded"
	PrerenderFinalStatusDataSaverEnabled                                           PrerenderFinalStatus = "DataSaverEnabled"
	PrerenderFinalStatusTriggerURLHasEffectiveURL                                  PrerenderFinalStatus = "TriggerUrlHasEffectiveUrl"
	PrerenderFinalStatusActivatedBeforeStarted                                     PrerenderFinalStatus = "ActivatedBeforeStarted"
	PrerenderFinalStatusInactivePageRestriction                                    PrerenderFinalStatus = "InactivePageRestriction"
	PrerenderFinalStatusStartFailed                                                PrerenderFinalStatus = "StartFailed"
	PrerenderFinalStatusTimeoutBackgrounded                                        PrerenderFinalStatus = "TimeoutBackgrounded"
	PrerenderFinalStatusCrossSiteRedirectInInitialNavigation                       PrerenderFinalStatus = "CrossSiteRedirectInInitialNavigation"
	PrerenderFinalStatusCrossSiteNavigationInInitialNavigation                     PrerenderFinalStatus = "CrossSiteNavigationInInitialNavigation"
	PrerenderFinalStatusSameSiteCrossOriginRedirectNotOptInInInitialNavigation     PrerenderFinalStatus = "SameSiteCrossOriginRedirectNotOptInInInitialNavigation"
	PrerenderFinalStatusSameSiteCrossOriginNavigationNotOptInInInitialNavigation   PrerenderFinalStatus = "SameSiteCrossOriginNavigationNotOptInInInitialNavigation"
	PrerenderFinalStatusActivationNavigationParameterMismatch                      PrerenderFinalStatus = "ActivationNavigationParameterMismatch"
	PrerenderFinalStatusActivatedInBackground                                      PrerenderFinalStatus = "ActivatedInBackground"
	PrerenderFinalStatusEmbedderHostDisallowed                                     PrerenderFinalStatus = "EmbedderHostDisallowed"
	PrerenderFinalStatusActivationNavigationDestroyedBeforeSuccess                 PrerenderFinalStatus = "ActivationNavigationDestroyedBeforeSuccess"
	PrerenderFinalStatusTabClosedByUserGesture                                     PrerenderFinalStatus = "TabClosedByUserGesture"
	PrerenderFinalStatusTabClosedWithoutUserGesture                                PrerenderFinalStatus = "TabClosedWithoutUserGesture"
	PrerenderFinalStatusPrimaryMainFrameRendererProcessCrashed                     PrerenderFinalStatus = "PrimaryMainFrameRendererProcessCrashed"
	PrerenderFinalStatusPrimaryMainFrameRendererProcessKilled                      PrerenderFinalStatus = "PrimaryMainFrameRendererProcessKilled"
	PrerenderFinalStatusActivationFramePolicyNotCompatible                         PrerenderFinalStatus = "ActivationFramePolicyNotCompatible"
	PrerenderFinalStatusPreloadingDisabled                                         PrerenderFinalStatus = "PreloadingDisabled"
	PrerenderFinalStatusBatterySaverEnabled                                        PrerenderFinalStatus = "BatterySaverEnabled"
	PrerenderFinalStatusActivatedDuringMainFrameNavigation                         PrerenderFinalStatus = "ActivatedDuringMainFrameNavigation"
	PrerenderFinalStatusPreloadingUnsupportedByWebContents                         PrerenderFinalStatus = "PreloadingUnsupportedByWebContents"
	PrerenderFinalStatusCrossSiteRedirectInMainFrameNavigation                     PrerenderFinalStatus = "CrossSiteRedirectInMainFrameNavigation"
	PrerenderFinalStatusCrossSiteNavigationInMainFrameNavigation                   PrerenderFinalStatus = "CrossSiteNavigationInMainFrameNavigation"
	PrerenderFinalStatusSameSiteCrossOriginRedirectNotOptInInMainFrameNavigation   PrerenderFinalStatus = "SameSiteCrossOriginRedirectNotOptInInMainFrameNavigation"
	PrerenderFinalStatusSameSiteCrossOriginNavigationNotOptInInMainFrameNavigation PrerenderFinalStatus = "SameSiteCrossOriginNavigationNotOptInInMainFrameNavigation"
	PrerenderFinalStatusMemoryPressureOnTrigger                                    PrerenderFinalStatus = "MemoryPressureOnTrigger"
	PrerenderFinalStatusMemoryPressureAfterTriggered                               PrerenderFinalStatus = "MemoryPressureAfterTriggered"
	PrerenderFinalStatusPrerenderingDisabledByDevTools                             PrerenderFinalStatus = "PrerenderingDisabledByDevTools"
	PrerenderFinalStatusSpeculationRuleRemoved                                     PrerenderFinalStatus = "SpeculationRuleRemoved"
	PrerenderFinalStatusActivatedWithAuxiliaryBrowsingContexts                     PrerenderFinalStatus = "ActivatedWithAuxiliaryBrowsingContexts"
	PrerenderFinalStatusMaxNumOfRunningEagerPrerendersExceeded                     PrerenderFinalStatus = "MaxNumOfRunningEagerPrerendersExceeded"
	PrerenderFinalStatusMaxNumOfRunningNonEagerPrerendersExceeded                  PrerenderFinalStatus = "MaxNumOfRunningNonEagerPrerendersExceeded"
	PrerenderFinalStatusMaxNumOfRunningEmbedderPrerendersExceeded                  PrerenderFinalStatus = "MaxNumOfRunningEmbedderPrerendersExceeded"
	PrerenderFinalStatusPrerenderingURLHasEffectiveURL                             PrerenderFinalStatus = "PrerenderingUrlHasEffectiveUrl"
	PrerenderFinalStatusRedirectedPrerenderingURLHasEffectiveURL                   PrerenderFinalStatus = "RedirectedPrerenderingUrlHasEffectiveUrl"
	PrerenderFinalStatusActivationURLHasEffectiveURL                               PrerenderFinalStatus = "ActivationUrlHasEffectiveUrl"
	PrerenderFinalStatusJavaScriptInterfaceAdded                                   PrerenderFinalStatus = "JavaScriptInterfaceAdded"
	PrerenderFinalStatusJavaScriptInterfaceRemoved                                 PrerenderFinalStatus = "JavaScriptInterfaceRemoved"
	PrerenderFinalStatusAllPrerenderingCanceled                                    PrerenderFinalStatus = "AllPrerenderingCanceled"
	PrerenderFinalStatusWindowClosed                                               PrerenderFinalStatus = "WindowClosed"
	PrerenderFinalStatusSlowNetwork                                                PrerenderFinalStatus = "SlowNetwork"
	PrerenderFinalStatusOtherPrerenderedPageActivated                              PrerenderFinalStatus = "OtherPrerenderedPageActivated"
)

func (e PrerenderFinalStatus) Valid() bool {
	switch e {
	case "Activated", "Destroyed", "LowEndDevice", "InvalidSchemeRedirect", "InvalidSchemeNavigation", "NavigationRequestBlockedByCsp", "MainFrameNavigation", "MojoBinderPolicy", "RendererProcessCrashed", "RendererProcessKilled", "Download", "TriggerDestroyed", "NavigationNotCommitted", "NavigationBadHttpStatus", "ClientCertRequested", "NavigationRequestNetworkError", "CancelAllHostsForTesting", "DidFailLoad", "Stop", "SslCertificateError", "LoginAuthRequested", "UaChangeRequiresReload", "BlockedByClient", "AudioOutputDeviceRequested", "MixedContent", "TriggerBackgrounded", "MemoryLimitExceeded", "DataSaverEnabled", "TriggerUrlHasEffectiveUrl", "ActivatedBeforeStarted", "InactivePageRestriction", "StartFailed", "TimeoutBackgrounded", "CrossSiteRedirectInInitialNavigation", "CrossSiteNavigationInInitialNavigation", "SameSiteCrossOriginRedirectNotOptInInInitialNavigation", "SameSiteCrossOriginNavigationNotOptInInInitialNavigation", "ActivationNavigationParameterMismatch", "ActivatedInBackground", "EmbedderHostDisallowed", "ActivationNavigationDestroyedBeforeSuccess", "TabClosedByUserGesture", "TabClosedWithoutUserGesture", "PrimaryMainFrameRendererProcessCrashed", "PrimaryMainFrameRendererProcessKilled", "ActivationFramePolicyNotCompatible", "PreloadingDisabled", "BatterySaverEnabled", "ActivatedDuringMainFrameNavigation", "PreloadingUnsupportedByWebContents", "CrossSiteRedirectInMainFrameNavigation", "CrossSiteNavigationInMainFrameNavigation", "SameSiteCrossOriginRedirectNotOptInInMainFrameNavigation", "SameSiteCrossOriginNavigationNotOptInInMainFrameNavigation", "MemoryPressureOnTrigger", "MemoryPressureAfterTriggered", "PrerenderingDisabledByDevTools", "SpeculationRuleRemoved", "ActivatedWithAuxiliaryBrowsingContexts", "MaxNumOfRunningEagerPrerendersExceeded", "MaxNumOfRunningNonEagerPrerendersExceeded", "MaxNumOfRunningEmbedderPrerendersExceeded", "PrerenderingUrlHasEffectiveUrl", "RedirectedPrerenderingUrlHasEffectiveUrl", "ActivationUrlHasEffectiveUrl", "JavaScriptInterfaceAdded", "JavaScriptInterfaceRemoved", "AllPrerenderingCanceled", "WindowClosed", "SlowNetwork", "OtherPrerenderedPageActivated":
		return true
	default:
		return false
	}
}

func (e PrerenderFinalStatus) String() string {
	return string(e)
}

// Status Preloading status values, see also PreloadingTriggeringOutcome. This
// status is shared by prefetchStatusUpdated and prerenderStatusUpdated.
type Status string

// Status as enums.
const (
	StatusNotSet       Status = ""
	StatusPending      Status = "Pending"
	StatusRunning      Status = "Running"
	StatusReady        Status = "Ready"
	StatusSuccess      Status = "Success"
	StatusFailure      Status = "Failure"
	StatusNotSupported Status = "NotSupported"
)

func (e Status) Valid() bool {
	switch e {
	case "Pending", "Running", "Ready", "Success", "Failure", "NotSupported":
		return true
	default:
		return false
	}
}

func (e Status) String() string {
	return string(e)
}

// PrefetchStatus TODO(https://crbug.com/1384419): revisit the list of
// PrefetchStatus and filter out the ones that aren't necessary to the
// developers.
type PrefetchStatus string

// PrefetchStatus as enums.
const (
	PrefetchStatusNotSet                                                      PrefetchStatus = ""
	PrefetchStatusPrefetchAllowed                                             PrefetchStatus = "PrefetchAllowed"
	PrefetchStatusPrefetchFailedIneligibleRedirect                            PrefetchStatus = "PrefetchFailedIneligibleRedirect"
	PrefetchStatusPrefetchFailedInvalidRedirect                               PrefetchStatus = "PrefetchFailedInvalidRedirect"
	PrefetchStatusPrefetchFailedMIMENotSupported                              PrefetchStatus = "PrefetchFailedMIMENotSupported"
	PrefetchStatusPrefetchFailedNetError                                      PrefetchStatus = "PrefetchFailedNetError"
	PrefetchStatusPrefetchFailedNon2XX                                        PrefetchStatus = "PrefetchFailedNon2XX"
	PrefetchStatusPrefetchEvictedAfterCandidateRemoved                        PrefetchStatus = "PrefetchEvictedAfterCandidateRemoved"
	PrefetchStatusPrefetchEvictedForNewerPrefetch                             PrefetchStatus = "PrefetchEvictedForNewerPrefetch"
	PrefetchStatusPrefetchHeldback                                            PrefetchStatus = "PrefetchHeldback"
	PrefetchStatusPrefetchIneligibleRetryAfter                                PrefetchStatus = "PrefetchIneligibleRetryAfter"
	PrefetchStatusPrefetchIsPrivacyDecoy                                      PrefetchStatus = "PrefetchIsPrivacyDecoy"
	PrefetchStatusPrefetchIsStale                                             PrefetchStatus = "PrefetchIsStale"
	PrefetchStatusPrefetchNotEligibleBrowserContextOffTheRecord               PrefetchStatus = "PrefetchNotEligibleBrowserContextOffTheRecord"
	PrefetchStatusPrefetchNotEligibleDataSaverEnabled                         PrefetchStatus = "PrefetchNotEligibleDataSaverEnabled"
	PrefetchStatusPrefetchNotEligibleExistingProxy                            PrefetchStatus = "PrefetchNotEligibleExistingProxy"
	PrefetchStatusPrefetchNotEligibleHostIsNonUnique                          PrefetchStatus = "PrefetchNotEligibleHostIsNonUnique"
	PrefetchStatusPrefetchNotEligibleNonDefaultStoragePartition               PrefetchStatus = "PrefetchNotEligibleNonDefaultStoragePartition"
	PrefetchStatusPrefetchNotEligibleSameSiteCrossOriginPrefetchRequiredProxy PrefetchStatus = "PrefetchNotEligibleSameSiteCrossOriginPrefetchRequiredProxy"
	PrefetchStatusPrefetchNotEligibleSchemeIsNotHTTPS                         PrefetchStatus = "PrefetchNotEligibleSchemeIsNotHttps"
	PrefetchStatusPrefetchNotEligibleUserHasCookies                           PrefetchStatus = "PrefetchNotEligibleUserHasCookies"
	PrefetchStatusPrefetchNotEligibleUserHasServiceWorker                     PrefetchStatus = "PrefetchNotEligibleUserHasServiceWorker"
	PrefetchStatusPrefetchNotEligibleBatterySaverEnabled                      PrefetchStatus = "PrefetchNotEligibleBatterySaverEnabled"
	PrefetchStatusPrefetchNotEligiblePreloadingDisabled                       PrefetchStatus = "PrefetchNotEligiblePreloadingDisabled"
	PrefetchStatusPrefetchNotFinishedInTime                                   PrefetchStatus = "PrefetchNotFinishedInTime"
	PrefetchStatusPrefetchNotStarted                                          PrefetchStatus = "PrefetchNotStarted"
	PrefetchStatusPrefetchNotUsedCookiesChanged                               PrefetchStatus = "PrefetchNotUsedCookiesChanged"
	PrefetchStatusPrefetchProxyNotAvailable                                   PrefetchStatus = "PrefetchProxyNotAvailable"
	PrefetchStatusPrefetchResponseUsed                                        PrefetchStatus = "PrefetchResponseUsed"
	PrefetchStatusPrefetchSuccessfulButNotUsed                                PrefetchStatus = "PrefetchSuccessfulButNotUsed"
	PrefetchStatusPrefetchNotUsedProbeFailed                                  PrefetchStatus = "PrefetchNotUsedProbeFailed"
)

func (e PrefetchStatus) Valid() bool {
	switch e {
	case "PrefetchAllowed", "PrefetchFailedIneligibleRedirect", "PrefetchFailedInvalidRedirect", "PrefetchFailedMIMENotSupported", "PrefetchFailedNetError", "PrefetchFailedNon2XX", "PrefetchEvictedAfterCandidateRemoved", "PrefetchEvictedForNewerPrefetch", "PrefetchHeldback", "PrefetchIneligibleRetryAfter", "PrefetchIsPrivacyDecoy", "PrefetchIsStale", "PrefetchNotEligibleBrowserContextOffTheRecord", "PrefetchNotEligibleDataSaverEnabled", "PrefetchNotEligibleExistingProxy", "PrefetchNotEligibleHostIsNonUnique", "PrefetchNotEligibleNonDefaultStoragePartition", "PrefetchNotEligibleSameSiteCrossOriginPrefetchRequiredProxy", "PrefetchNotEligibleSchemeIsNotHttps", "PrefetchNotEligibleUserHasCookies", "PrefetchNotEligibleUserHasServiceWorker", "PrefetchNotEligibleBatterySaverEnabled", "PrefetchNotEligiblePreloadingDisabled", "PrefetchNotFinishedInTime", "PrefetchNotStarted", "PrefetchNotUsedCookiesChanged", "PrefetchProxyNotAvailable", "PrefetchResponseUsed", "PrefetchSuccessfulButNotUsed", "PrefetchNotUsedProbeFailed":
		return true
	default:
		return false
	}
}

func (e PrefetchStatus) String() string {
	return string(e)
}

// PrerenderMismatchedHeaders Information of headers to be displayed when the
// header mismatch occurred.
type PrerenderMismatchedHeaders struct {
	HeaderName      string  `json:"headerName"`                // No description.
	InitialValue    *string `json:"initialValue,omitempty"`    // No description.
	ActivationValue *string `json:"activationValue,omitempty"` // No description.
}