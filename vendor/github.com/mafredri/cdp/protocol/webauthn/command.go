// Code generated by cdpgen. DO NOT EDIT.

package webauthn

// EnableArgs represents the arguments for Enable in the WebAuthn domain.
type EnableArgs struct {
	EnableUI *bool `json:"enableUI,omitempty"` // Whether to enable the WebAuthn user interface. Enabling the UI is recommended for debugging and demo purposes, as it is closer to the real experience. Disabling the UI is recommended for automated testing. Supported at the embedder's discretion if UI is available. Defaults to false.
}

// NewEnableArgs initializes EnableArgs with the required arguments.
func NewEnableArgs() *EnableArgs {
	args := new(EnableArgs)

	return args
}

// SetEnableUI sets the EnableUI optional argument. Whether to enable
// the WebAuthn user interface. Enabling the UI is recommended for
// debugging and demo purposes, as it is closer to the real experience.
// Disabling the UI is recommended for automated testing. Supported at
// the embedder's discretion if UI is available. Defaults to false.
func (a *EnableArgs) SetEnableUI(enableUI bool) *EnableArgs {
	a.EnableUI = &enableUI
	return a
}

// AddVirtualAuthenticatorArgs represents the arguments for AddVirtualAuthenticator in the WebAuthn domain.
type AddVirtualAuthenticatorArgs struct {
	Options VirtualAuthenticatorOptions `json:"options"` // No description.
}

// NewAddVirtualAuthenticatorArgs initializes AddVirtualAuthenticatorArgs with the required arguments.
func NewAddVirtualAuthenticatorArgs(options VirtualAuthenticatorOptions) *AddVirtualAuthenticatorArgs {
	args := new(AddVirtualAuthenticatorArgs)
	args.Options = options
	return args
}

// AddVirtualAuthenticatorReply represents the return values for AddVirtualAuthenticator in the WebAuthn domain.
type AddVirtualAuthenticatorReply struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
}

// SetResponseOverrideBitsArgs represents the arguments for SetResponseOverrideBits in the WebAuthn domain.
type SetResponseOverrideBitsArgs struct {
	AuthenticatorID  AuthenticatorID `json:"authenticatorId"`            // No description.
	IsBogusSignature *bool           `json:"isBogusSignature,omitempty"` // If isBogusSignature is set, overrides the signature in the authenticator response to be zero. Defaults to false.
	IsBadUV          *bool           `json:"isBadUV,omitempty"`          // If isBadUV is set, overrides the UV bit in the flags in the authenticator response to be zero. Defaults to false.
	IsBadUP          *bool           `json:"isBadUP,omitempty"`          // If isBadUP is set, overrides the UP bit in the flags in the authenticator response to be zero. Defaults to false.
}

// NewSetResponseOverrideBitsArgs initializes SetResponseOverrideBitsArgs with the required arguments.
func NewSetResponseOverrideBitsArgs(authenticatorID AuthenticatorID) *SetResponseOverrideBitsArgs {
	args := new(SetResponseOverrideBitsArgs)
	args.AuthenticatorID = authenticatorID
	return args
}

// SetIsBogusSignature sets the IsBogusSignature optional argument. If
// isBogusSignature is set, overrides the signature in the
// authenticator response to be zero. Defaults to false.
func (a *SetResponseOverrideBitsArgs) SetIsBogusSignature(isBogusSignature bool) *SetResponseOverrideBitsArgs {
	a.IsBogusSignature = &isBogusSignature
	return a
}

// SetIsBadUV sets the IsBadUV optional argument. If isBadUV is set,
// overrides the UV bit in the flags in the authenticator response to
// be zero. Defaults to false.
func (a *SetResponseOverrideBitsArgs) SetIsBadUV(isBadUV bool) *SetResponseOverrideBitsArgs {
	a.IsBadUV = &isBadUV
	return a
}

// SetIsBadUP sets the IsBadUP optional argument. If isBadUP is set,
// overrides the UP bit in the flags in the authenticator response to
// be zero. Defaults to false.
func (a *SetResponseOverrideBitsArgs) SetIsBadUP(isBadUP bool) *SetResponseOverrideBitsArgs {
	a.IsBadUP = &isBadUP
	return a
}

// RemoveVirtualAuthenticatorArgs represents the arguments for RemoveVirtualAuthenticator in the WebAuthn domain.
type RemoveVirtualAuthenticatorArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
}

// NewRemoveVirtualAuthenticatorArgs initializes RemoveVirtualAuthenticatorArgs with the required arguments.
func NewRemoveVirtualAuthenticatorArgs(authenticatorID AuthenticatorID) *RemoveVirtualAuthenticatorArgs {
	args := new(RemoveVirtualAuthenticatorArgs)
	args.AuthenticatorID = authenticatorID
	return args
}

// AddCredentialArgs represents the arguments for AddCredential in the WebAuthn domain.
type AddCredentialArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
	Credential      Credential      `json:"credential"`      // No description.
}

// NewAddCredentialArgs initializes AddCredentialArgs with the required arguments.
func NewAddCredentialArgs(authenticatorID AuthenticatorID, credential Credential) *AddCredentialArgs {
	args := new(AddCredentialArgs)
	args.AuthenticatorID = authenticatorID
	args.Credential = credential
	return args
}

// GetCredentialArgs represents the arguments for GetCredential in the WebAuthn domain.
type GetCredentialArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
	CredentialID    string          `json:"credentialId"`    // No description.
}

// NewGetCredentialArgs initializes GetCredentialArgs with the required arguments.
func NewGetCredentialArgs(authenticatorID AuthenticatorID, credentialID string) *GetCredentialArgs {
	args := new(GetCredentialArgs)
	args.AuthenticatorID = authenticatorID
	args.CredentialID = credentialID
	return args
}

// GetCredentialReply represents the return values for GetCredential in the WebAuthn domain.
type GetCredentialReply struct {
	Credential Credential `json:"credential"` // No description.
}

// GetCredentialsArgs represents the arguments for GetCredentials in the WebAuthn domain.
type GetCredentialsArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
}

// NewGetCredentialsArgs initializes GetCredentialsArgs with the required arguments.
func NewGetCredentialsArgs(authenticatorID AuthenticatorID) *GetCredentialsArgs {
	args := new(GetCredentialsArgs)
	args.AuthenticatorID = authenticatorID
	return args
}

// GetCredentialsReply represents the return values for GetCredentials in the WebAuthn domain.
type GetCredentialsReply struct {
	Credentials []Credential `json:"credentials"` // No description.
}

// RemoveCredentialArgs represents the arguments for RemoveCredential in the WebAuthn domain.
type RemoveCredentialArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
	CredentialID    string          `json:"credentialId"`    // No description.
}

// NewRemoveCredentialArgs initializes RemoveCredentialArgs with the required arguments.
func NewRemoveCredentialArgs(authenticatorID AuthenticatorID, credentialID string) *RemoveCredentialArgs {
	args := new(RemoveCredentialArgs)
	args.AuthenticatorID = authenticatorID
	args.CredentialID = credentialID
	return args
}

// ClearCredentialsArgs represents the arguments for ClearCredentials in the WebAuthn domain.
type ClearCredentialsArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
}

// NewClearCredentialsArgs initializes ClearCredentialsArgs with the required arguments.
func NewClearCredentialsArgs(authenticatorID AuthenticatorID) *ClearCredentialsArgs {
	args := new(ClearCredentialsArgs)
	args.AuthenticatorID = authenticatorID
	return args
}

// SetUserVerifiedArgs represents the arguments for SetUserVerified in the WebAuthn domain.
type SetUserVerifiedArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
	IsUserVerified  bool            `json:"isUserVerified"`  // No description.
}

// NewSetUserVerifiedArgs initializes SetUserVerifiedArgs with the required arguments.
func NewSetUserVerifiedArgs(authenticatorID AuthenticatorID, isUserVerified bool) *SetUserVerifiedArgs {
	args := new(SetUserVerifiedArgs)
	args.AuthenticatorID = authenticatorID
	args.IsUserVerified = isUserVerified
	return args
}

// SetAutomaticPresenceSimulationArgs represents the arguments for SetAutomaticPresenceSimulation in the WebAuthn domain.
type SetAutomaticPresenceSimulationArgs struct {
	AuthenticatorID AuthenticatorID `json:"authenticatorId"` // No description.
	Enabled         bool            `json:"enabled"`         // No description.
}

// NewSetAutomaticPresenceSimulationArgs initializes SetAutomaticPresenceSimulationArgs with the required arguments.
func NewSetAutomaticPresenceSimulationArgs(authenticatorID AuthenticatorID, enabled bool) *SetAutomaticPresenceSimulationArgs {
	args := new(SetAutomaticPresenceSimulationArgs)
	args.AuthenticatorID = authenticatorID
	args.Enabled = enabled
	return args
}

// SetCredentialPropertiesArgs represents the arguments for SetCredentialProperties in the WebAuthn domain.
type SetCredentialPropertiesArgs struct {
	AuthenticatorID   AuthenticatorID `json:"authenticatorId"`             // No description.
	CredentialID      string          `json:"credentialId"`                // No description.
	BackupEligibility *bool           `json:"backupEligibility,omitempty"` // No description.
	BackupState       *bool           `json:"backupState,omitempty"`       // No description.
}

// NewSetCredentialPropertiesArgs initializes SetCredentialPropertiesArgs with the required arguments.
func NewSetCredentialPropertiesArgs(authenticatorID AuthenticatorID, credentialID string) *SetCredentialPropertiesArgs {
	args := new(SetCredentialPropertiesArgs)
	args.AuthenticatorID = authenticatorID
	args.CredentialID = credentialID
	return args
}

// SetBackupEligibility sets the BackupEligibility optional argument.
func (a *SetCredentialPropertiesArgs) SetBackupEligibility(backupEligibility bool) *SetCredentialPropertiesArgs {
	a.BackupEligibility = &backupEligibility
	return a
}

// SetBackupState sets the BackupState optional argument.
func (a *SetCredentialPropertiesArgs) SetBackupState(backupState bool) *SetCredentialPropertiesArgs {
	a.BackupState = &backupState
	return a
}