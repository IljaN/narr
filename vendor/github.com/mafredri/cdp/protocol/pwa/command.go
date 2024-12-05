// Code generated by cdpgen. DO NOT EDIT.

package pwa

import (
	"github.com/mafredri/cdp/protocol/target"
)

// GetOSAppStateArgs represents the arguments for GetOSAppState in the PWA domain.
type GetOSAppStateArgs struct {
	ManifestID string `json:"manifestId"` // The id from the webapp's manifest file, commonly it's the url of the site installing the webapp. See https://web.dev/learn/pwa/web-app-manifest.
}

// NewGetOSAppStateArgs initializes GetOSAppStateArgs with the required arguments.
func NewGetOSAppStateArgs(manifestID string) *GetOSAppStateArgs {
	args := new(GetOSAppStateArgs)
	args.ManifestID = manifestID
	return args
}

// GetOSAppStateReply represents the return values for GetOSAppState in the PWA domain.
type GetOSAppStateReply struct {
	BadgeCount   int           `json:"badgeCount"`   // No description.
	FileHandlers []FileHandler `json:"fileHandlers"` // No description.
}

// InstallArgs represents the arguments for Install in the PWA domain.
type InstallArgs struct {
	ManifestID            string  `json:"manifestId"`                      // No description.
	InstallURLOrBundleURL *string `json:"installUrlOrBundleUrl,omitempty"` // The location of the app or bundle overriding the one derived from the manifestId.
}

// NewInstallArgs initializes InstallArgs with the required arguments.
func NewInstallArgs(manifestID string) *InstallArgs {
	args := new(InstallArgs)
	args.ManifestID = manifestID
	return args
}

// SetInstallURLOrBundleURL sets the InstallURLOrBundleURL optional argument.
// The location of the app or bundle overriding the one derived from
// the manifestId.
func (a *InstallArgs) SetInstallURLOrBundleURL(installURLOrBundleURL string) *InstallArgs {
	a.InstallURLOrBundleURL = &installURLOrBundleURL
	return a
}

// UninstallArgs represents the arguments for Uninstall in the PWA domain.
type UninstallArgs struct {
	ManifestID string `json:"manifestId"` // No description.
}

// NewUninstallArgs initializes UninstallArgs with the required arguments.
func NewUninstallArgs(manifestID string) *UninstallArgs {
	args := new(UninstallArgs)
	args.ManifestID = manifestID
	return args
}

// LaunchArgs represents the arguments for Launch in the PWA domain.
type LaunchArgs struct {
	ManifestID string  `json:"manifestId"`    // No description.
	URL        *string `json:"url,omitempty"` // No description.
}

// NewLaunchArgs initializes LaunchArgs with the required arguments.
func NewLaunchArgs(manifestID string) *LaunchArgs {
	args := new(LaunchArgs)
	args.ManifestID = manifestID
	return args
}

// SetURL sets the URL optional argument.
func (a *LaunchArgs) SetURL(url string) *LaunchArgs {
	a.URL = &url
	return a
}

// LaunchReply represents the return values for Launch in the PWA domain.
type LaunchReply struct {
	TargetID target.ID `json:"targetId"` // ID of the tab target created as a result.
}

// LaunchFilesInAppArgs represents the arguments for LaunchFilesInApp in the PWA domain.
type LaunchFilesInAppArgs struct {
	ManifestID string   `json:"manifestId"` // No description.
	Files      []string `json:"files"`      // No description.
}

// NewLaunchFilesInAppArgs initializes LaunchFilesInAppArgs with the required arguments.
func NewLaunchFilesInAppArgs(manifestID string, files []string) *LaunchFilesInAppArgs {
	args := new(LaunchFilesInAppArgs)
	args.ManifestID = manifestID
	args.Files = files
	return args
}

// LaunchFilesInAppReply represents the return values for LaunchFilesInApp in the PWA domain.
type LaunchFilesInAppReply struct {
	TargetIDs []target.ID `json:"targetIds"` // IDs of the tab targets created as the result.
}

// OpenCurrentPageInAppArgs represents the arguments for OpenCurrentPageInApp in the PWA domain.
type OpenCurrentPageInAppArgs struct {
	ManifestID string `json:"manifestId"` // No description.
}

// NewOpenCurrentPageInAppArgs initializes OpenCurrentPageInAppArgs with the required arguments.
func NewOpenCurrentPageInAppArgs(manifestID string) *OpenCurrentPageInAppArgs {
	args := new(OpenCurrentPageInAppArgs)
	args.ManifestID = manifestID
	return args
}

// ChangeAppUserSettingsArgs represents the arguments for ChangeAppUserSettings in the PWA domain.
type ChangeAppUserSettingsArgs struct {
	ManifestID    string      `json:"manifestId"`              // No description.
	LinkCapturing *bool       `json:"linkCapturing,omitempty"` // If user allows the links clicked on by the user in the app's scope, or extended scope if the manifest has scope extensions and the flags `DesktopPWAsLinkCapturingWithScopeExtensions` and `WebAppEnableScopeExtensions` are enabled. Note, the API does not support resetting the linkCapturing to the initial value, uninstalling and installing the web app again will reset it. TODO(crbug.com/339453269): Setting this value on ChromeOS is not supported yet.
	DisplayMode   DisplayMode `json:"displayMode,omitempty"`   // No description.
}

// NewChangeAppUserSettingsArgs initializes ChangeAppUserSettingsArgs with the required arguments.
func NewChangeAppUserSettingsArgs(manifestID string) *ChangeAppUserSettingsArgs {
	args := new(ChangeAppUserSettingsArgs)
	args.ManifestID = manifestID
	return args
}

// SetLinkCapturing sets the LinkCapturing optional argument. If user
// allows the links clicked on by the user in the app's scope, or
// extended scope if the manifest has scope extensions and the flags
// `DesktopPWAsLinkCapturingWithScopeExtensions` and
// `WebAppEnableScopeExtensions` are enabled.
//
// Note, the API does not support resetting the linkCapturing to the
// initial value, uninstalling and installing the web app again will
// reset it.
//
// TODO(crbug.com/339453269): Setting this value on ChromeOS is not
// supported yet.
func (a *ChangeAppUserSettingsArgs) SetLinkCapturing(linkCapturing bool) *ChangeAppUserSettingsArgs {
	a.LinkCapturing = &linkCapturing
	return a
}

// SetDisplayMode sets the DisplayMode optional argument.
func (a *ChangeAppUserSettingsArgs) SetDisplayMode(displayMode DisplayMode) *ChangeAppUserSettingsArgs {
	a.DisplayMode = displayMode
	return a
}
