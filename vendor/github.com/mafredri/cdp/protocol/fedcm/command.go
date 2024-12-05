// Code generated by cdpgen. DO NOT EDIT.

package fedcm

// EnableArgs represents the arguments for Enable in the FedCM domain.
type EnableArgs struct {
	DisableRejectionDelay *bool `json:"disableRejectionDelay,omitempty"` // Allows callers to disable the promise rejection delay that would normally happen, if this is unimportant to what's being tested. (step 4 of https://fedidcg.github.io/FedCM/#browser-api-rp-sign-in)
}

// NewEnableArgs initializes EnableArgs with the required arguments.
func NewEnableArgs() *EnableArgs {
	args := new(EnableArgs)

	return args
}

// SetDisableRejectionDelay sets the DisableRejectionDelay optional argument.
// Allows callers to disable the promise rejection delay that would
// normally happen, if this is unimportant to what's being tested.
// (step 4 of https://fedidcg.github.io/FedCM/#browser-api-rp-sign-in)
func (a *EnableArgs) SetDisableRejectionDelay(disableRejectionDelay bool) *EnableArgs {
	a.DisableRejectionDelay = &disableRejectionDelay
	return a
}

// SelectAccountArgs represents the arguments for SelectAccount in the FedCM domain.
type SelectAccountArgs struct {
	DialogID     string `json:"dialogId"`     // No description.
	AccountIndex int    `json:"accountIndex"` // No description.
}

// NewSelectAccountArgs initializes SelectAccountArgs with the required arguments.
func NewSelectAccountArgs(dialogID string, accountIndex int) *SelectAccountArgs {
	args := new(SelectAccountArgs)
	args.DialogID = dialogID
	args.AccountIndex = accountIndex
	return args
}

// ClickDialogButtonArgs represents the arguments for ClickDialogButton in the FedCM domain.
type ClickDialogButtonArgs struct {
	DialogID     string       `json:"dialogId"`     // No description.
	DialogButton DialogButton `json:"dialogButton"` // No description.
}

// NewClickDialogButtonArgs initializes ClickDialogButtonArgs with the required arguments.
func NewClickDialogButtonArgs(dialogID string, dialogButton DialogButton) *ClickDialogButtonArgs {
	args := new(ClickDialogButtonArgs)
	args.DialogID = dialogID
	args.DialogButton = dialogButton
	return args
}

// OpenURLArgs represents the arguments for OpenURL in the FedCM domain.
type OpenURLArgs struct {
	DialogID       string         `json:"dialogId"`       // No description.
	AccountIndex   int            `json:"accountIndex"`   // No description.
	AccountURLType AccountURLType `json:"accountUrlType"` // No description.
}

// NewOpenURLArgs initializes OpenURLArgs with the required arguments.
func NewOpenURLArgs(dialogID string, accountIndex int, accountURLType AccountURLType) *OpenURLArgs {
	args := new(OpenURLArgs)
	args.DialogID = dialogID
	args.AccountIndex = accountIndex
	args.AccountURLType = accountURLType
	return args
}

// DismissDialogArgs represents the arguments for DismissDialog in the FedCM domain.
type DismissDialogArgs struct {
	DialogID        string `json:"dialogId"`                  // No description.
	TriggerCooldown *bool  `json:"triggerCooldown,omitempty"` // No description.
}

// NewDismissDialogArgs initializes DismissDialogArgs with the required arguments.
func NewDismissDialogArgs(dialogID string) *DismissDialogArgs {
	args := new(DismissDialogArgs)
	args.DialogID = dialogID
	return args
}

// SetTriggerCooldown sets the TriggerCooldown optional argument.
func (a *DismissDialogArgs) SetTriggerCooldown(triggerCooldown bool) *DismissDialogArgs {
	a.TriggerCooldown = &triggerCooldown
	return a
}