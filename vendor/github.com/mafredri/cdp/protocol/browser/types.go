// Code generated by cdpgen. DO NOT EDIT.

package browser

import (
	"github.com/mafredri/cdp/protocol/internal"
)

// ContextID
//
// Note: This type is experimental.
type ContextID = internal.BrowserContextID

// WindowID
//
// Note: This type is experimental.
type WindowID int

// WindowState The state of the browser window.
//
// Note: This type is experimental.
type WindowState string

// WindowState as enums.
const (
	WindowStateNotSet     WindowState = ""
	WindowStateNormal     WindowState = "normal"
	WindowStateMinimized  WindowState = "minimized"
	WindowStateMaximized  WindowState = "maximized"
	WindowStateFullscreen WindowState = "fullscreen"
)

func (e WindowState) Valid() bool {
	switch e {
	case "normal", "minimized", "maximized", "fullscreen":
		return true
	default:
		return false
	}
}

func (e WindowState) String() string {
	return string(e)
}

// Bounds Browser window bounds information
//
// Note: This type is experimental.
type Bounds struct {
	Left        *int        `json:"left,omitempty"`        // The offset from the left edge of the screen to the window in pixels.
	Top         *int        `json:"top,omitempty"`         // The offset from the top edge of the screen to the window in pixels.
	Width       *int        `json:"width,omitempty"`       // The window width in pixels.
	Height      *int        `json:"height,omitempty"`      // The window height in pixels.
	WindowState WindowState `json:"windowState,omitempty"` // The window state. Default to normal.
}

// PermissionType
//
// Note: This type is experimental.
type PermissionType string

// PermissionType as enums.
const (
	PermissionTypeNotSet                   PermissionType = ""
	PermissionTypeAccessibilityEvents      PermissionType = "accessibilityEvents"
	PermissionTypeAudioCapture             PermissionType = "audioCapture"
	PermissionTypeBackgroundSync           PermissionType = "backgroundSync"
	PermissionTypeBackgroundFetch          PermissionType = "backgroundFetch"
	PermissionTypeCapturedSurfaceControl   PermissionType = "capturedSurfaceControl"
	PermissionTypeClipboardReadWrite       PermissionType = "clipboardReadWrite"
	PermissionTypeClipboardSanitizedWrite  PermissionType = "clipboardSanitizedWrite"
	PermissionTypeDisplayCapture           PermissionType = "displayCapture"
	PermissionTypeDurableStorage           PermissionType = "durableStorage"
	PermissionTypeFlash                    PermissionType = "flash"
	PermissionTypeGeolocation              PermissionType = "geolocation"
	PermissionTypeIdleDetection            PermissionType = "idleDetection"
	PermissionTypeLocalFonts               PermissionType = "localFonts"
	PermissionTypeMidi                     PermissionType = "midi"
	PermissionTypeMidiSysex                PermissionType = "midiSysex"
	PermissionTypeNFC                      PermissionType = "nfc"
	PermissionTypeNotifications            PermissionType = "notifications"
	PermissionTypePaymentHandler           PermissionType = "paymentHandler"
	PermissionTypePeriodicBackgroundSync   PermissionType = "periodicBackgroundSync"
	PermissionTypeProtectedMediaIdentifier PermissionType = "protectedMediaIdentifier"
	PermissionTypeSensors                  PermissionType = "sensors"
	PermissionTypeStorageAccess            PermissionType = "storageAccess"
	PermissionTypeSpeakerSelection         PermissionType = "speakerSelection"
	PermissionTypeTopLevelStorageAccess    PermissionType = "topLevelStorageAccess"
	PermissionTypeVideoCapture             PermissionType = "videoCapture"
	PermissionTypeVideoCapturePanTiltZoom  PermissionType = "videoCapturePanTiltZoom"
	PermissionTypeWakeLockScreen           PermissionType = "wakeLockScreen"
	PermissionTypeWakeLockSystem           PermissionType = "wakeLockSystem"
	PermissionTypeWebAppInstallation       PermissionType = "webAppInstallation"
	PermissionTypeWindowManagement         PermissionType = "windowManagement"
)

func (e PermissionType) Valid() bool {
	switch e {
	case "accessibilityEvents", "audioCapture", "backgroundSync", "backgroundFetch", "capturedSurfaceControl", "clipboardReadWrite", "clipboardSanitizedWrite", "displayCapture", "durableStorage", "flash", "geolocation", "idleDetection", "localFonts", "midi", "midiSysex", "nfc", "notifications", "paymentHandler", "periodicBackgroundSync", "protectedMediaIdentifier", "sensors", "storageAccess", "speakerSelection", "topLevelStorageAccess", "videoCapture", "videoCapturePanTiltZoom", "wakeLockScreen", "wakeLockSystem", "webAppInstallation", "windowManagement":
		return true
	default:
		return false
	}
}

func (e PermissionType) String() string {
	return string(e)
}

// PermissionSetting
//
// Note: This type is experimental.
type PermissionSetting string

// PermissionSetting as enums.
const (
	PermissionSettingNotSet  PermissionSetting = ""
	PermissionSettingGranted PermissionSetting = "granted"
	PermissionSettingDenied  PermissionSetting = "denied"
	PermissionSettingPrompt  PermissionSetting = "prompt"
)

func (e PermissionSetting) Valid() bool {
	switch e {
	case "granted", "denied", "prompt":
		return true
	default:
		return false
	}
}

func (e PermissionSetting) String() string {
	return string(e)
}

// PermissionDescriptor Definition of PermissionDescriptor defined in the
// Permissions API:
// https://w3c.github.io/permissions/#dom-permissiondescriptor.
//
// Note: This type is experimental.
type PermissionDescriptor struct {
	Name                     string `json:"name"`                               // Name of permission. See https://cs.chromium.org/chromium/src/third_party/blink/renderer/modules/permissions/permission_descriptor.idl for valid permission names.
	Sysex                    *bool  `json:"sysex,omitempty"`                    // For "midi" permission, may also specify sysex control.
	UserVisibleOnly          *bool  `json:"userVisibleOnly,omitempty"`          // For "push" permission, may specify userVisibleOnly. Note that userVisibleOnly = true is the only currently supported type.
	AllowWithoutSanitization *bool  `json:"allowWithoutSanitization,omitempty"` // For "clipboard" permission, may specify allowWithoutSanitization.
	AllowWithoutGesture      *bool  `json:"allowWithoutGesture,omitempty"`      // For "fullscreen" permission, must specify allowWithoutGesture:true.
	PanTiltZoom              *bool  `json:"panTiltZoom,omitempty"`              // For "camera" permission, may specify panTiltZoom.
}

// CommandID Browser command ids used by executeBrowserCommand.
//
// Note: This type is experimental.
type CommandID string

// CommandID as enums.
const (
	CommandIDNotSet         CommandID = ""
	CommandIDOpenTabSearch  CommandID = "openTabSearch"
	CommandIDCloseTabSearch CommandID = "closeTabSearch"
)

func (e CommandID) Valid() bool {
	switch e {
	case "openTabSearch", "closeTabSearch":
		return true
	default:
		return false
	}
}

func (e CommandID) String() string {
	return string(e)
}

// Bucket Chrome histogram bucket.
//
// Note: This type is experimental.
type Bucket struct {
	Low   int `json:"low"`   // Minimum value (inclusive).
	High  int `json:"high"`  // Maximum value (exclusive).
	Count int `json:"count"` // Number of samples.
}

// Histogram Chrome histogram.
//
// Note: This type is experimental.
type Histogram struct {
	Name    string   `json:"name"`    // Name.
	Sum     int      `json:"sum"`     // Sum of sample values.
	Count   int      `json:"count"`   // Total number of samples.
	Buckets []Bucket `json:"buckets"` // Buckets.
}