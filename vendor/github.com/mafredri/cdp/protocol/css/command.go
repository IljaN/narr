// Code generated by cdpgen. DO NOT EDIT.

package css

import (
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/page"
)

// AddRuleArgs represents the arguments for AddRule in the CSS domain.
type AddRuleArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // The css style sheet identifier where a new rule should be inserted.
	RuleText     string       `json:"ruleText"`     // The text of a new rule.
	Location     SourceRange  `json:"location"`     // Text position of a new rule in the target style sheet.
	// NodeForPropertySyntaxValidation NodeId for the DOM node in whose
	// context custom property declarations for registered properties
	// should be validated. If omitted, declarations in the new rule text
	// can only be validated statically, which may produce incorrect
	// results if the declaration contains a var() for example.
	//
	// Note: This property is experimental.
	NodeForPropertySyntaxValidation *dom.NodeID `json:"nodeForPropertySyntaxValidation,omitempty"`
}

// NewAddRuleArgs initializes AddRuleArgs with the required arguments.
func NewAddRuleArgs(styleSheetID StyleSheetID, ruleText string, location SourceRange) *AddRuleArgs {
	args := new(AddRuleArgs)
	args.StyleSheetID = styleSheetID
	args.RuleText = ruleText
	args.Location = location
	return args
}

// SetNodeForPropertySyntaxValidation sets the NodeForPropertySyntaxValidation optional argument.
// NodeId for the DOM node in whose context custom property
// declarations for registered properties should be validated. If
// omitted, declarations in the new rule text can only be validated
// statically, which may produce incorrect results if the declaration
// contains a var() for example.
//
// Note: This property is experimental.
func (a *AddRuleArgs) SetNodeForPropertySyntaxValidation(nodeForPropertySyntaxValidation dom.NodeID) *AddRuleArgs {
	a.NodeForPropertySyntaxValidation = &nodeForPropertySyntaxValidation
	return a
}

// AddRuleReply represents the return values for AddRule in the CSS domain.
type AddRuleReply struct {
	Rule Rule `json:"rule"` // The newly created rule.
}

// CollectClassNamesArgs represents the arguments for CollectClassNames in the CSS domain.
type CollectClassNamesArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
}

// NewCollectClassNamesArgs initializes CollectClassNamesArgs with the required arguments.
func NewCollectClassNamesArgs(styleSheetID StyleSheetID) *CollectClassNamesArgs {
	args := new(CollectClassNamesArgs)
	args.StyleSheetID = styleSheetID
	return args
}

// CollectClassNamesReply represents the return values for CollectClassNames in the CSS domain.
type CollectClassNamesReply struct {
	ClassNames []string `json:"classNames"` // Class name list.
}

// CreateStyleSheetArgs represents the arguments for CreateStyleSheet in the CSS domain.
type CreateStyleSheetArgs struct {
	FrameID page.FrameID `json:"frameId"` // Identifier of the frame where "via-inspector" stylesheet should be created.
}

// NewCreateStyleSheetArgs initializes CreateStyleSheetArgs with the required arguments.
func NewCreateStyleSheetArgs(frameID page.FrameID) *CreateStyleSheetArgs {
	args := new(CreateStyleSheetArgs)
	args.FrameID = frameID
	return args
}

// CreateStyleSheetReply represents the return values for CreateStyleSheet in the CSS domain.
type CreateStyleSheetReply struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // Identifier of the created "via-inspector" stylesheet.
}

// ForcePseudoStateArgs represents the arguments for ForcePseudoState in the CSS domain.
type ForcePseudoStateArgs struct {
	NodeID              dom.NodeID `json:"nodeId"`              // The element id for which to force the pseudo state.
	ForcedPseudoClasses []string   `json:"forcedPseudoClasses"` // Element pseudo classes to force when computing the element's style.
}

// NewForcePseudoStateArgs initializes ForcePseudoStateArgs with the required arguments.
func NewForcePseudoStateArgs(nodeID dom.NodeID, forcedPseudoClasses []string) *ForcePseudoStateArgs {
	args := new(ForcePseudoStateArgs)
	args.NodeID = nodeID
	args.ForcedPseudoClasses = forcedPseudoClasses
	return args
}

// GetBackgroundColorsArgs represents the arguments for GetBackgroundColors in the CSS domain.
type GetBackgroundColorsArgs struct {
	NodeID dom.NodeID `json:"nodeId"` // Id of the node to get background colors for.
}

// NewGetBackgroundColorsArgs initializes GetBackgroundColorsArgs with the required arguments.
func NewGetBackgroundColorsArgs(nodeID dom.NodeID) *GetBackgroundColorsArgs {
	args := new(GetBackgroundColorsArgs)
	args.NodeID = nodeID
	return args
}

// GetBackgroundColorsReply represents the return values for GetBackgroundColors in the CSS domain.
type GetBackgroundColorsReply struct {
	BackgroundColors   []string `json:"backgroundColors,omitempty"`   // The range of background colors behind this element, if it contains any visible text. If no visible text is present, this will be undefined. In the case of a flat background color, this will consist of simply that color. In the case of a gradient, this will consist of each of the color stops. For anything more complicated, this will be an empty array. Images will be ignored (as if the image had failed to load).
	ComputedFontSize   *string  `json:"computedFontSize,omitempty"`   // The computed font size for this node, as a CSS computed value string (e.g. '12px').
	ComputedFontWeight *string  `json:"computedFontWeight,omitempty"` // The computed font weight for this node, as a CSS computed value string (e.g. 'normal' or '100').
}

// GetComputedStyleForNodeArgs represents the arguments for GetComputedStyleForNode in the CSS domain.
type GetComputedStyleForNodeArgs struct {
	NodeID dom.NodeID `json:"nodeId"` // No description.
}

// NewGetComputedStyleForNodeArgs initializes GetComputedStyleForNodeArgs with the required arguments.
func NewGetComputedStyleForNodeArgs(nodeID dom.NodeID) *GetComputedStyleForNodeArgs {
	args := new(GetComputedStyleForNodeArgs)
	args.NodeID = nodeID
	return args
}

// GetComputedStyleForNodeReply represents the return values for GetComputedStyleForNode in the CSS domain.
type GetComputedStyleForNodeReply struct {
	ComputedStyle []ComputedStyleProperty `json:"computedStyle"` // Computed style for the specified DOM node.
}

// GetInlineStylesForNodeArgs represents the arguments for GetInlineStylesForNode in the CSS domain.
type GetInlineStylesForNodeArgs struct {
	NodeID dom.NodeID `json:"nodeId"` // No description.
}

// NewGetInlineStylesForNodeArgs initializes GetInlineStylesForNodeArgs with the required arguments.
func NewGetInlineStylesForNodeArgs(nodeID dom.NodeID) *GetInlineStylesForNodeArgs {
	args := new(GetInlineStylesForNodeArgs)
	args.NodeID = nodeID
	return args
}

// GetInlineStylesForNodeReply represents the return values for GetInlineStylesForNode in the CSS domain.
type GetInlineStylesForNodeReply struct {
	InlineStyle     *Style `json:"inlineStyle,omitempty"`     // Inline style for the specified DOM node.
	AttributesStyle *Style `json:"attributesStyle,omitempty"` // Attribute-defined element style (e.g. resulting from "width=20 height=100%").
}

// GetMatchedStylesForNodeArgs represents the arguments for GetMatchedStylesForNode in the CSS domain.
type GetMatchedStylesForNodeArgs struct {
	NodeID dom.NodeID `json:"nodeId"` // No description.
}

// NewGetMatchedStylesForNodeArgs initializes GetMatchedStylesForNodeArgs with the required arguments.
func NewGetMatchedStylesForNodeArgs(nodeID dom.NodeID) *GetMatchedStylesForNodeArgs {
	args := new(GetMatchedStylesForNodeArgs)
	args.NodeID = nodeID
	return args
}

// GetMatchedStylesForNodeReply represents the return values for GetMatchedStylesForNode in the CSS domain.
type GetMatchedStylesForNodeReply struct {
	InlineStyle                 *Style                          `json:"inlineStyle,omitempty"`                 // Inline style for the specified DOM node.
	AttributesStyle             *Style                          `json:"attributesStyle,omitempty"`             // Attribute-defined element style (e.g. resulting from "width=20 height=100%").
	MatchedCSSRules             []RuleMatch                     `json:"matchedCSSRules,omitempty"`             // CSS rules matching this node, from all applicable stylesheets.
	PseudoElements              []PseudoElementMatches          `json:"pseudoElements,omitempty"`              // Pseudo style matches for this node.
	Inherited                   []InheritedStyleEntry           `json:"inherited,omitempty"`                   // A chain of inherited styles (from the immediate node parent up to the DOM tree root).
	InheritedPseudoElements     []InheritedPseudoElementMatches `json:"inheritedPseudoElements,omitempty"`     // A chain of inherited pseudo element styles (from the immediate node parent up to the DOM tree root).
	CSSKeyframesRules           []KeyframesRule                 `json:"cssKeyframesRules,omitempty"`           // A list of CSS keyframed animations matching this node.
	CSSPositionTryRules         []PositionTryRule               `json:"cssPositionTryRules,omitempty"`         // A list of CSS @position-try rules matching this node, based on the position-try-fallbacks property.
	ActivePositionFallbackIndex *int                            `json:"activePositionFallbackIndex,omitempty"` // Index of the active fallback in the applied position-try-fallback property, will not be set if there is no active position-try fallback.
	CSSPropertyRules            []PropertyRule                  `json:"cssPropertyRules,omitempty"`            // A list of CSS at-property rules matching this node.
	CSSPropertyRegistrations    []PropertyRegistration          `json:"cssPropertyRegistrations,omitempty"`    // A list of CSS property registrations matching this node.
	CSSFontPaletteValuesRule    *FontPaletteValuesRule          `json:"cssFontPaletteValuesRule,omitempty"`    // A font-palette-values rule matching this node.
	// ParentLayoutNodeID Id of the first parent element that does not
	// have display: contents.
	//
	// Note: This property is experimental.
	ParentLayoutNodeID *dom.NodeID `json:"parentLayoutNodeId,omitempty"`
}

// GetMediaQueriesReply represents the return values for GetMediaQueries in the CSS domain.
type GetMediaQueriesReply struct {
	Medias []Media `json:"medias"` // No description.
}

// GetPlatformFontsForNodeArgs represents the arguments for GetPlatformFontsForNode in the CSS domain.
type GetPlatformFontsForNodeArgs struct {
	NodeID dom.NodeID `json:"nodeId"` // No description.
}

// NewGetPlatformFontsForNodeArgs initializes GetPlatformFontsForNodeArgs with the required arguments.
func NewGetPlatformFontsForNodeArgs(nodeID dom.NodeID) *GetPlatformFontsForNodeArgs {
	args := new(GetPlatformFontsForNodeArgs)
	args.NodeID = nodeID
	return args
}

// GetPlatformFontsForNodeReply represents the return values for GetPlatformFontsForNode in the CSS domain.
type GetPlatformFontsForNodeReply struct {
	Fonts []PlatformFontUsage `json:"fonts"` // Usage statistics for every employed platform font.
}

// GetStyleSheetTextArgs represents the arguments for GetStyleSheetText in the CSS domain.
type GetStyleSheetTextArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
}

// NewGetStyleSheetTextArgs initializes GetStyleSheetTextArgs with the required arguments.
func NewGetStyleSheetTextArgs(styleSheetID StyleSheetID) *GetStyleSheetTextArgs {
	args := new(GetStyleSheetTextArgs)
	args.StyleSheetID = styleSheetID
	return args
}

// GetStyleSheetTextReply represents the return values for GetStyleSheetText in the CSS domain.
type GetStyleSheetTextReply struct {
	Text string `json:"text"` // The stylesheet text.
}

// GetLayersForNodeArgs represents the arguments for GetLayersForNode in the CSS domain.
type GetLayersForNodeArgs struct {
	NodeID dom.NodeID `json:"nodeId"` // No description.
}

// NewGetLayersForNodeArgs initializes GetLayersForNodeArgs with the required arguments.
func NewGetLayersForNodeArgs(nodeID dom.NodeID) *GetLayersForNodeArgs {
	args := new(GetLayersForNodeArgs)
	args.NodeID = nodeID
	return args
}

// GetLayersForNodeReply represents the return values for GetLayersForNode in the CSS domain.
type GetLayersForNodeReply struct {
	RootLayer LayerData `json:"rootLayer"` // No description.
}

// GetLocationForSelectorArgs represents the arguments for GetLocationForSelector in the CSS domain.
type GetLocationForSelectorArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
	SelectorText string       `json:"selectorText"` // No description.
}

// NewGetLocationForSelectorArgs initializes GetLocationForSelectorArgs with the required arguments.
func NewGetLocationForSelectorArgs(styleSheetID StyleSheetID, selectorText string) *GetLocationForSelectorArgs {
	args := new(GetLocationForSelectorArgs)
	args.StyleSheetID = styleSheetID
	args.SelectorText = selectorText
	return args
}

// GetLocationForSelectorReply represents the return values for GetLocationForSelector in the CSS domain.
type GetLocationForSelectorReply struct {
	Ranges []SourceRange `json:"ranges"` // No description.
}

// TrackComputedStyleUpdatesArgs represents the arguments for TrackComputedStyleUpdates in the CSS domain.
type TrackComputedStyleUpdatesArgs struct {
	PropertiesToTrack []ComputedStyleProperty `json:"propertiesToTrack"` // No description.
}

// NewTrackComputedStyleUpdatesArgs initializes TrackComputedStyleUpdatesArgs with the required arguments.
func NewTrackComputedStyleUpdatesArgs(propertiesToTrack []ComputedStyleProperty) *TrackComputedStyleUpdatesArgs {
	args := new(TrackComputedStyleUpdatesArgs)
	args.PropertiesToTrack = propertiesToTrack
	return args
}

// TakeComputedStyleUpdatesReply represents the return values for TakeComputedStyleUpdates in the CSS domain.
type TakeComputedStyleUpdatesReply struct {
	NodeIDs []dom.NodeID `json:"nodeIds"` // The list of node Ids that have their tracked computed styles updated.
}

// SetEffectivePropertyValueForNodeArgs represents the arguments for SetEffectivePropertyValueForNode in the CSS domain.
type SetEffectivePropertyValueForNodeArgs struct {
	NodeID       dom.NodeID `json:"nodeId"`       // The element id for which to set property.
	PropertyName string     `json:"propertyName"` // No description.
	Value        string     `json:"value"`        // No description.
}

// NewSetEffectivePropertyValueForNodeArgs initializes SetEffectivePropertyValueForNodeArgs with the required arguments.
func NewSetEffectivePropertyValueForNodeArgs(nodeID dom.NodeID, propertyName string, value string) *SetEffectivePropertyValueForNodeArgs {
	args := new(SetEffectivePropertyValueForNodeArgs)
	args.NodeID = nodeID
	args.PropertyName = propertyName
	args.Value = value
	return args
}

// SetPropertyRulePropertyNameArgs represents the arguments for SetPropertyRulePropertyName in the CSS domain.
type SetPropertyRulePropertyNameArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
	Range        SourceRange  `json:"range"`        // No description.
	PropertyName string       `json:"propertyName"` // No description.
}

// NewSetPropertyRulePropertyNameArgs initializes SetPropertyRulePropertyNameArgs with the required arguments.
func NewSetPropertyRulePropertyNameArgs(styleSheetID StyleSheetID, rang SourceRange, propertyName string) *SetPropertyRulePropertyNameArgs {
	args := new(SetPropertyRulePropertyNameArgs)
	args.StyleSheetID = styleSheetID
	args.Range = rang
	args.PropertyName = propertyName
	return args
}

// SetPropertyRulePropertyNameReply represents the return values for SetPropertyRulePropertyName in the CSS domain.
type SetPropertyRulePropertyNameReply struct {
	PropertyName Value `json:"propertyName"` // The resulting key text after modification.
}

// SetKeyframeKeyArgs represents the arguments for SetKeyframeKey in the CSS domain.
type SetKeyframeKeyArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
	Range        SourceRange  `json:"range"`        // No description.
	KeyText      string       `json:"keyText"`      // No description.
}

// NewSetKeyframeKeyArgs initializes SetKeyframeKeyArgs with the required arguments.
func NewSetKeyframeKeyArgs(styleSheetID StyleSheetID, rang SourceRange, keyText string) *SetKeyframeKeyArgs {
	args := new(SetKeyframeKeyArgs)
	args.StyleSheetID = styleSheetID
	args.Range = rang
	args.KeyText = keyText
	return args
}

// SetKeyframeKeyReply represents the return values for SetKeyframeKey in the CSS domain.
type SetKeyframeKeyReply struct {
	KeyText Value `json:"keyText"` // The resulting key text after modification.
}

// SetMediaTextArgs represents the arguments for SetMediaText in the CSS domain.
type SetMediaTextArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
	Range        SourceRange  `json:"range"`        // No description.
	Text         string       `json:"text"`         // No description.
}

// NewSetMediaTextArgs initializes SetMediaTextArgs with the required arguments.
func NewSetMediaTextArgs(styleSheetID StyleSheetID, rang SourceRange, text string) *SetMediaTextArgs {
	args := new(SetMediaTextArgs)
	args.StyleSheetID = styleSheetID
	args.Range = rang
	args.Text = text
	return args
}

// SetMediaTextReply represents the return values for SetMediaText in the CSS domain.
type SetMediaTextReply struct {
	Media Media `json:"media"` // The resulting CSS media rule after modification.
}

// SetContainerQueryTextArgs represents the arguments for SetContainerQueryText in the CSS domain.
type SetContainerQueryTextArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
	Range        SourceRange  `json:"range"`        // No description.
	Text         string       `json:"text"`         // No description.
}

// NewSetContainerQueryTextArgs initializes SetContainerQueryTextArgs with the required arguments.
func NewSetContainerQueryTextArgs(styleSheetID StyleSheetID, rang SourceRange, text string) *SetContainerQueryTextArgs {
	args := new(SetContainerQueryTextArgs)
	args.StyleSheetID = styleSheetID
	args.Range = rang
	args.Text = text
	return args
}

// SetContainerQueryTextReply represents the return values for SetContainerQueryText in the CSS domain.
type SetContainerQueryTextReply struct {
	ContainerQuery ContainerQuery `json:"containerQuery"` // The resulting CSS container query rule after modification.
}

// SetSupportsTextArgs represents the arguments for SetSupportsText in the CSS domain.
type SetSupportsTextArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
	Range        SourceRange  `json:"range"`        // No description.
	Text         string       `json:"text"`         // No description.
}

// NewSetSupportsTextArgs initializes SetSupportsTextArgs with the required arguments.
func NewSetSupportsTextArgs(styleSheetID StyleSheetID, rang SourceRange, text string) *SetSupportsTextArgs {
	args := new(SetSupportsTextArgs)
	args.StyleSheetID = styleSheetID
	args.Range = rang
	args.Text = text
	return args
}

// SetSupportsTextReply represents the return values for SetSupportsText in the CSS domain.
type SetSupportsTextReply struct {
	Supports Supports `json:"supports"` // The resulting CSS Supports rule after modification.
}

// SetScopeTextArgs represents the arguments for SetScopeText in the CSS domain.
type SetScopeTextArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
	Range        SourceRange  `json:"range"`        // No description.
	Text         string       `json:"text"`         // No description.
}

// NewSetScopeTextArgs initializes SetScopeTextArgs with the required arguments.
func NewSetScopeTextArgs(styleSheetID StyleSheetID, rang SourceRange, text string) *SetScopeTextArgs {
	args := new(SetScopeTextArgs)
	args.StyleSheetID = styleSheetID
	args.Range = rang
	args.Text = text
	return args
}

// SetScopeTextReply represents the return values for SetScopeText in the CSS domain.
type SetScopeTextReply struct {
	Scope Scope `json:"scope"` // The resulting CSS Scope rule after modification.
}

// SetRuleSelectorArgs represents the arguments for SetRuleSelector in the CSS domain.
type SetRuleSelectorArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
	Range        SourceRange  `json:"range"`        // No description.
	Selector     string       `json:"selector"`     // No description.
}

// NewSetRuleSelectorArgs initializes SetRuleSelectorArgs with the required arguments.
func NewSetRuleSelectorArgs(styleSheetID StyleSheetID, rang SourceRange, selector string) *SetRuleSelectorArgs {
	args := new(SetRuleSelectorArgs)
	args.StyleSheetID = styleSheetID
	args.Range = rang
	args.Selector = selector
	return args
}

// SetRuleSelectorReply represents the return values for SetRuleSelector in the CSS domain.
type SetRuleSelectorReply struct {
	SelectorList SelectorList `json:"selectorList"` // The resulting selector list after modification.
}

// SetStyleSheetTextArgs represents the arguments for SetStyleSheetText in the CSS domain.
type SetStyleSheetTextArgs struct {
	StyleSheetID StyleSheetID `json:"styleSheetId"` // No description.
	Text         string       `json:"text"`         // No description.
}

// NewSetStyleSheetTextArgs initializes SetStyleSheetTextArgs with the required arguments.
func NewSetStyleSheetTextArgs(styleSheetID StyleSheetID, text string) *SetStyleSheetTextArgs {
	args := new(SetStyleSheetTextArgs)
	args.StyleSheetID = styleSheetID
	args.Text = text
	return args
}

// SetStyleSheetTextReply represents the return values for SetStyleSheetText in the CSS domain.
type SetStyleSheetTextReply struct {
	SourceMapURL *string `json:"sourceMapURL,omitempty"` // URL of source map associated with script (if any).
}

// SetStyleTextsArgs represents the arguments for SetStyleTexts in the CSS domain.
type SetStyleTextsArgs struct {
	Edits []StyleDeclarationEdit `json:"edits"` // No description.
	// NodeForPropertySyntaxValidation NodeId for the DOM node in whose
	// context custom property declarations for registered properties
	// should be validated. If omitted, declarations in the new rule text
	// can only be validated statically, which may produce incorrect
	// results if the declaration contains a var() for example.
	//
	// Note: This property is experimental.
	NodeForPropertySyntaxValidation *dom.NodeID `json:"nodeForPropertySyntaxValidation,omitempty"`
}

// NewSetStyleTextsArgs initializes SetStyleTextsArgs with the required arguments.
func NewSetStyleTextsArgs(edits []StyleDeclarationEdit) *SetStyleTextsArgs {
	args := new(SetStyleTextsArgs)
	args.Edits = edits
	return args
}

// SetNodeForPropertySyntaxValidation sets the NodeForPropertySyntaxValidation optional argument.
// NodeId for the DOM node in whose context custom property
// declarations for registered properties should be validated. If
// omitted, declarations in the new rule text can only be validated
// statically, which may produce incorrect results if the declaration
// contains a var() for example.
//
// Note: This property is experimental.
func (a *SetStyleTextsArgs) SetNodeForPropertySyntaxValidation(nodeForPropertySyntaxValidation dom.NodeID) *SetStyleTextsArgs {
	a.NodeForPropertySyntaxValidation = &nodeForPropertySyntaxValidation
	return a
}

// SetStyleTextsReply represents the return values for SetStyleTexts in the CSS domain.
type SetStyleTextsReply struct {
	Styles []Style `json:"styles"` // The resulting styles after modification.
}

// StopRuleUsageTrackingReply represents the return values for StopRuleUsageTracking in the CSS domain.
type StopRuleUsageTrackingReply struct {
	RuleUsage []RuleUsage `json:"ruleUsage"` // No description.
}

// TakeCoverageDeltaReply represents the return values for TakeCoverageDelta in the CSS domain.
type TakeCoverageDeltaReply struct {
	Coverage  []RuleUsage `json:"coverage"`  // No description.
	Timestamp float64     `json:"timestamp"` // Monotonically increasing time, in seconds.
}

// SetLocalFontsEnabledArgs represents the arguments for SetLocalFontsEnabled in the CSS domain.
type SetLocalFontsEnabledArgs struct {
	Enabled bool `json:"enabled"` // Whether rendering of local fonts is enabled.
}

// NewSetLocalFontsEnabledArgs initializes SetLocalFontsEnabledArgs with the required arguments.
func NewSetLocalFontsEnabledArgs(enabled bool) *SetLocalFontsEnabledArgs {
	args := new(SetLocalFontsEnabledArgs)
	args.Enabled = enabled
	return args
}
