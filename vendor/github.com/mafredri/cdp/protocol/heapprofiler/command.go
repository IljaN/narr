// Code generated by cdpgen. DO NOT EDIT.

package heapprofiler

import (
	"github.com/mafredri/cdp/protocol/runtime"
)

// AddInspectedHeapObjectArgs represents the arguments for AddInspectedHeapObject in the HeapProfiler domain.
type AddInspectedHeapObjectArgs struct {
	HeapObjectID HeapSnapshotObjectID `json:"heapObjectId"` // Heap snapshot object id to be accessible by means of $x command line API.
}

// NewAddInspectedHeapObjectArgs initializes AddInspectedHeapObjectArgs with the required arguments.
func NewAddInspectedHeapObjectArgs(heapObjectID HeapSnapshotObjectID) *AddInspectedHeapObjectArgs {
	args := new(AddInspectedHeapObjectArgs)
	args.HeapObjectID = heapObjectID
	return args
}

// GetHeapObjectIDArgs represents the arguments for GetHeapObjectID in the HeapProfiler domain.
type GetHeapObjectIDArgs struct {
	ObjectID runtime.RemoteObjectID `json:"objectId"` // Identifier of the object to get heap object id for.
}

// NewGetHeapObjectIDArgs initializes GetHeapObjectIDArgs with the required arguments.
func NewGetHeapObjectIDArgs(objectID runtime.RemoteObjectID) *GetHeapObjectIDArgs {
	args := new(GetHeapObjectIDArgs)
	args.ObjectID = objectID
	return args
}

// GetHeapObjectIDReply represents the return values for GetHeapObjectID in the HeapProfiler domain.
type GetHeapObjectIDReply struct {
	HeapSnapshotObjectID HeapSnapshotObjectID `json:"heapSnapshotObjectId"` // Id of the heap snapshot object corresponding to the passed remote object id.
}

// GetObjectByHeapObjectIDArgs represents the arguments for GetObjectByHeapObjectID in the HeapProfiler domain.
type GetObjectByHeapObjectIDArgs struct {
	ObjectID    HeapSnapshotObjectID `json:"objectId"`              // No description.
	ObjectGroup *string              `json:"objectGroup,omitempty"` // Symbolic group name that can be used to release multiple objects.
}

// NewGetObjectByHeapObjectIDArgs initializes GetObjectByHeapObjectIDArgs with the required arguments.
func NewGetObjectByHeapObjectIDArgs(objectID HeapSnapshotObjectID) *GetObjectByHeapObjectIDArgs {
	args := new(GetObjectByHeapObjectIDArgs)
	args.ObjectID = objectID
	return args
}

// SetObjectGroup sets the ObjectGroup optional argument. Symbolic
// group name that can be used to release multiple objects.
func (a *GetObjectByHeapObjectIDArgs) SetObjectGroup(objectGroup string) *GetObjectByHeapObjectIDArgs {
	a.ObjectGroup = &objectGroup
	return a
}

// GetObjectByHeapObjectIDReply represents the return values for GetObjectByHeapObjectID in the HeapProfiler domain.
type GetObjectByHeapObjectIDReply struct {
	Result runtime.RemoteObject `json:"result"` // Evaluation result.
}

// GetSamplingProfileReply represents the return values for GetSamplingProfile in the HeapProfiler domain.
type GetSamplingProfileReply struct {
	Profile SamplingHeapProfile `json:"profile"` // Return the sampling profile being collected.
}

// StartSamplingArgs represents the arguments for StartSampling in the HeapProfiler domain.
type StartSamplingArgs struct {
	SamplingInterval                 *float64 `json:"samplingInterval,omitempty"`                 // Average sample interval in bytes. Poisson distribution is used for the intervals. The default value is 32768 bytes.
	IncludeObjectsCollectedByMajorGC *bool    `json:"includeObjectsCollectedByMajorGC,omitempty"` // By default, the sampling heap profiler reports only objects which are still alive when the profile is returned via getSamplingProfile or stopSampling, which is useful for determining what functions contribute the most to steady-state memory usage. This flag instructs the sampling heap profiler to also include information about objects discarded by major GC, which will show which functions cause large temporary memory usage or long GC pauses.
	IncludeObjectsCollectedByMinorGC *bool    `json:"includeObjectsCollectedByMinorGC,omitempty"` // By default, the sampling heap profiler reports only objects which are still alive when the profile is returned via getSamplingProfile or stopSampling, which is useful for determining what functions contribute the most to steady-state memory usage. This flag instructs the sampling heap profiler to also include information about objects discarded by minor GC, which is useful when tuning a latency-sensitive application for minimal GC activity.
}

// NewStartSamplingArgs initializes StartSamplingArgs with the required arguments.
func NewStartSamplingArgs() *StartSamplingArgs {
	args := new(StartSamplingArgs)

	return args
}

// SetSamplingInterval sets the SamplingInterval optional argument.
// Average sample interval in bytes. Poisson distribution is used for
// the intervals. The default value is 32768 bytes.
func (a *StartSamplingArgs) SetSamplingInterval(samplingInterval float64) *StartSamplingArgs {
	a.SamplingInterval = &samplingInterval
	return a
}

// SetIncludeObjectsCollectedByMajorGC sets the IncludeObjectsCollectedByMajorGC optional argument.
// By default, the sampling heap profiler reports only objects which
// are still alive when the profile is returned via getSamplingProfile
// or stopSampling, which is useful for determining what functions
// contribute the most to steady-state memory usage. This flag
// instructs the sampling heap profiler to also include information
// about objects discarded by major GC, which will show which functions
// cause large temporary memory usage or long GC pauses.
func (a *StartSamplingArgs) SetIncludeObjectsCollectedByMajorGC(includeObjectsCollectedByMajorGC bool) *StartSamplingArgs {
	a.IncludeObjectsCollectedByMajorGC = &includeObjectsCollectedByMajorGC
	return a
}

// SetIncludeObjectsCollectedByMinorGC sets the IncludeObjectsCollectedByMinorGC optional argument.
// By default, the sampling heap profiler reports only objects which
// are still alive when the profile is returned via getSamplingProfile
// or stopSampling, which is useful for determining what functions
// contribute the most to steady-state memory usage. This flag
// instructs the sampling heap profiler to also include information
// about objects discarded by minor GC, which is useful when tuning a
// latency-sensitive application for minimal GC activity.
func (a *StartSamplingArgs) SetIncludeObjectsCollectedByMinorGC(includeObjectsCollectedByMinorGC bool) *StartSamplingArgs {
	a.IncludeObjectsCollectedByMinorGC = &includeObjectsCollectedByMinorGC
	return a
}

// StartTrackingHeapObjectsArgs represents the arguments for StartTrackingHeapObjects in the HeapProfiler domain.
type StartTrackingHeapObjectsArgs struct {
	TrackAllocations *bool `json:"trackAllocations,omitempty"` // No description.
}

// NewStartTrackingHeapObjectsArgs initializes StartTrackingHeapObjectsArgs with the required arguments.
func NewStartTrackingHeapObjectsArgs() *StartTrackingHeapObjectsArgs {
	args := new(StartTrackingHeapObjectsArgs)

	return args
}

// SetTrackAllocations sets the TrackAllocations optional argument.
func (a *StartTrackingHeapObjectsArgs) SetTrackAllocations(trackAllocations bool) *StartTrackingHeapObjectsArgs {
	a.TrackAllocations = &trackAllocations
	return a
}

// StopSamplingReply represents the return values for StopSampling in the HeapProfiler domain.
type StopSamplingReply struct {
	Profile SamplingHeapProfile `json:"profile"` // Recorded sampling heap profile.
}

// StopTrackingHeapObjectsArgs represents the arguments for StopTrackingHeapObjects in the HeapProfiler domain.
type StopTrackingHeapObjectsArgs struct {
	ReportProgress *bool `json:"reportProgress,omitempty"` // If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken when the tracking is stopped.
	// TreatGlobalObjectsAsRoots is deprecated.
	//
	// Deprecated: Deprecated in favor of `exposeInternals`.
	TreatGlobalObjectsAsRoots *bool `json:"treatGlobalObjectsAsRoots,omitempty"`
	CaptureNumericValue       *bool `json:"captureNumericValue,omitempty"` // If true, numerical values are included in the snapshot
	// ExposeInternals If true, exposes internals of the snapshot.
	//
	// Note: This property is experimental.
	ExposeInternals *bool `json:"exposeInternals,omitempty"`
}

// NewStopTrackingHeapObjectsArgs initializes StopTrackingHeapObjectsArgs with the required arguments.
func NewStopTrackingHeapObjectsArgs() *StopTrackingHeapObjectsArgs {
	args := new(StopTrackingHeapObjectsArgs)

	return args
}

// SetReportProgress sets the ReportProgress optional argument. If
// true 'reportHeapSnapshotProgress' events will be generated while
// snapshot is being taken when the tracking is stopped.
func (a *StopTrackingHeapObjectsArgs) SetReportProgress(reportProgress bool) *StopTrackingHeapObjectsArgs {
	a.ReportProgress = &reportProgress
	return a
}

// SetTreatGlobalObjectsAsRoots sets the TreatGlobalObjectsAsRoots optional argument.
//
// Deprecated:
// Deprecated in favor of `exposeInternals`.
func (a *StopTrackingHeapObjectsArgs) SetTreatGlobalObjectsAsRoots(treatGlobalObjectsAsRoots bool) *StopTrackingHeapObjectsArgs {
	a.TreatGlobalObjectsAsRoots = &treatGlobalObjectsAsRoots
	return a
}

// SetCaptureNumericValue sets the CaptureNumericValue optional argument.
// If true, numerical values are included in the snapshot
func (a *StopTrackingHeapObjectsArgs) SetCaptureNumericValue(captureNumericValue bool) *StopTrackingHeapObjectsArgs {
	a.CaptureNumericValue = &captureNumericValue
	return a
}

// SetExposeInternals sets the ExposeInternals optional argument. If
// true, exposes internals of the snapshot.
//
// Note: This property is experimental.
func (a *StopTrackingHeapObjectsArgs) SetExposeInternals(exposeInternals bool) *StopTrackingHeapObjectsArgs {
	a.ExposeInternals = &exposeInternals
	return a
}

// TakeHeapSnapshotArgs represents the arguments for TakeHeapSnapshot in the HeapProfiler domain.
type TakeHeapSnapshotArgs struct {
	ReportProgress *bool `json:"reportProgress,omitempty"` // If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
	// TreatGlobalObjectsAsRoots is deprecated.
	//
	// Deprecated: If true, a raw snapshot without artificial roots will
	// be generated. Deprecated in favor of `exposeInternals`.
	TreatGlobalObjectsAsRoots *bool `json:"treatGlobalObjectsAsRoots,omitempty"`
	CaptureNumericValue       *bool `json:"captureNumericValue,omitempty"` // If true, numerical values are included in the snapshot
	// ExposeInternals If true, exposes internals of the snapshot.
	//
	// Note: This property is experimental.
	ExposeInternals *bool `json:"exposeInternals,omitempty"`
}

// NewTakeHeapSnapshotArgs initializes TakeHeapSnapshotArgs with the required arguments.
func NewTakeHeapSnapshotArgs() *TakeHeapSnapshotArgs {
	args := new(TakeHeapSnapshotArgs)

	return args
}

// SetReportProgress sets the ReportProgress optional argument. If
// true 'reportHeapSnapshotProgress' events will be generated while
// snapshot is being taken.
func (a *TakeHeapSnapshotArgs) SetReportProgress(reportProgress bool) *TakeHeapSnapshotArgs {
	a.ReportProgress = &reportProgress
	return a
}

// SetTreatGlobalObjectsAsRoots sets the TreatGlobalObjectsAsRoots optional argument.
//
// Deprecated:
// If true, a raw snapshot without artificial roots will be generated.
// Deprecated in favor of `exposeInternals`.
func (a *TakeHeapSnapshotArgs) SetTreatGlobalObjectsAsRoots(treatGlobalObjectsAsRoots bool) *TakeHeapSnapshotArgs {
	a.TreatGlobalObjectsAsRoots = &treatGlobalObjectsAsRoots
	return a
}

// SetCaptureNumericValue sets the CaptureNumericValue optional argument.
// If true, numerical values are included in the snapshot
func (a *TakeHeapSnapshotArgs) SetCaptureNumericValue(captureNumericValue bool) *TakeHeapSnapshotArgs {
	a.CaptureNumericValue = &captureNumericValue
	return a
}

// SetExposeInternals sets the ExposeInternals optional argument. If
// true, exposes internals of the snapshot.
//
// Note: This property is experimental.
func (a *TakeHeapSnapshotArgs) SetExposeInternals(exposeInternals bool) *TakeHeapSnapshotArgs {
	a.ExposeInternals = &exposeInternals
	return a
}