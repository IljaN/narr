// Code generated by cdpgen. DO NOT EDIT.

package dom

import (
	"encoding/json"

	"github.com/mafredri/cdp/protocol/internal"
)

// NodeID Unique DOM node identifier.
type NodeID int

// BackendNodeID Unique DOM node identifier used to reference a node that may
// not have been pushed to the front-end.
type BackendNodeID int

// BackendNode Backend node with a friendly name.
type BackendNode struct {
	NodeType      int           `json:"nodeType"`      // `Node`'s nodeType.
	NodeName      string        `json:"nodeName"`      // `Node`'s nodeName.
	BackendNodeID BackendNodeID `json:"backendNodeId"` // No description.
}

// PseudoType Pseudo element type.
type PseudoType string

// PseudoType as enums.
const (
	PseudoTypeNotSet                  PseudoType = ""
	PseudoTypeFirstLine               PseudoType = "first-line"
	PseudoTypeFirstLetter             PseudoType = "first-letter"
	PseudoTypeBefore                  PseudoType = "before"
	PseudoTypeAfter                   PseudoType = "after"
	PseudoTypeMarker                  PseudoType = "marker"
	PseudoTypeBackdrop                PseudoType = "backdrop"
	PseudoTypeSelection               PseudoType = "selection"
	PseudoTypeSearchText              PseudoType = "search-text"
	PseudoTypeTargetText              PseudoType = "target-text"
	PseudoTypeSpellingError           PseudoType = "spelling-error"
	PseudoTypeGrammarError            PseudoType = "grammar-error"
	PseudoTypeHighlight               PseudoType = "highlight"
	PseudoTypeFirstLineInherited      PseudoType = "first-line-inherited"
	PseudoTypeScrollMarker            PseudoType = "scroll-marker"
	PseudoTypeScrollMarkerGroup       PseudoType = "scroll-marker-group"
	PseudoTypeScrollNextButton        PseudoType = "scroll-next-button"
	PseudoTypeScrollPrevButton        PseudoType = "scroll-prev-button"
	PseudoTypeScrollbar               PseudoType = "scrollbar"
	PseudoTypeScrollbarThumb          PseudoType = "scrollbar-thumb"
	PseudoTypeScrollbarButton         PseudoType = "scrollbar-button"
	PseudoTypeScrollbarTrack          PseudoType = "scrollbar-track"
	PseudoTypeScrollbarTrackPiece     PseudoType = "scrollbar-track-piece"
	PseudoTypeScrollbarCorner         PseudoType = "scrollbar-corner"
	PseudoTypeResizer                 PseudoType = "resizer"
	PseudoTypeInputListButton         PseudoType = "input-list-button"
	PseudoTypeViewTransition          PseudoType = "view-transition"
	PseudoTypeViewTransitionGroup     PseudoType = "view-transition-group"
	PseudoTypeViewTransitionImagePair PseudoType = "view-transition-image-pair"
	PseudoTypeViewTransitionOld       PseudoType = "view-transition-old"
	PseudoTypeViewTransitionNew       PseudoType = "view-transition-new"
)

func (e PseudoType) Valid() bool {
	switch e {
	case "first-line", "first-letter", "before", "after", "marker", "backdrop", "selection", "search-text", "target-text", "spelling-error", "grammar-error", "highlight", "first-line-inherited", "scroll-marker", "scroll-marker-group", "scroll-next-button", "scroll-prev-button", "scrollbar", "scrollbar-thumb", "scrollbar-button", "scrollbar-track", "scrollbar-track-piece", "scrollbar-corner", "resizer", "input-list-button", "view-transition", "view-transition-group", "view-transition-image-pair", "view-transition-old", "view-transition-new":
		return true
	default:
		return false
	}
}

func (e PseudoType) String() string {
	return string(e)
}

// ShadowRootType Shadow root type.
type ShadowRootType string

// ShadowRootType as enums.
const (
	ShadowRootTypeNotSet    ShadowRootType = ""
	ShadowRootTypeUserAgent ShadowRootType = "user-agent"
	ShadowRootTypeOpen      ShadowRootType = "open"
	ShadowRootTypeClosed    ShadowRootType = "closed"
)

func (e ShadowRootType) Valid() bool {
	switch e {
	case "user-agent", "open", "closed":
		return true
	default:
		return false
	}
}

func (e ShadowRootType) String() string {
	return string(e)
}

// CompatibilityMode Document compatibility mode.
type CompatibilityMode string

// CompatibilityMode as enums.
const (
	CompatibilityModeNotSet            CompatibilityMode = ""
	CompatibilityModeQuirksMode        CompatibilityMode = "QuirksMode"
	CompatibilityModeLimitedQuirksMode CompatibilityMode = "LimitedQuirksMode"
	CompatibilityModeNoQuirksMode      CompatibilityMode = "NoQuirksMode"
)

func (e CompatibilityMode) Valid() bool {
	switch e {
	case "QuirksMode", "LimitedQuirksMode", "NoQuirksMode":
		return true
	default:
		return false
	}
}

func (e CompatibilityMode) String() string {
	return string(e)
}

// PhysicalAxes ContainerSelector physical axes
type PhysicalAxes string

// PhysicalAxes as enums.
const (
	PhysicalAxesNotSet     PhysicalAxes = ""
	PhysicalAxesHorizontal PhysicalAxes = "Horizontal"
	PhysicalAxesVertical   PhysicalAxes = "Vertical"
	PhysicalAxesBoth       PhysicalAxes = "Both"
)

func (e PhysicalAxes) Valid() bool {
	switch e {
	case "Horizontal", "Vertical", "Both":
		return true
	default:
		return false
	}
}

func (e PhysicalAxes) String() string {
	return string(e)
}

// LogicalAxes ContainerSelector logical axes
type LogicalAxes string

// LogicalAxes as enums.
const (
	LogicalAxesNotSet LogicalAxes = ""
	LogicalAxesInline LogicalAxes = "Inline"
	LogicalAxesBlock  LogicalAxes = "Block"
	LogicalAxesBoth   LogicalAxes = "Both"
)

func (e LogicalAxes) Valid() bool {
	switch e {
	case "Inline", "Block", "Both":
		return true
	default:
		return false
	}
}

func (e LogicalAxes) String() string {
	return string(e)
}

// ScrollOrientation Physical scroll orientation
type ScrollOrientation string

// ScrollOrientation as enums.
const (
	ScrollOrientationNotSet     ScrollOrientation = ""
	ScrollOrientationHorizontal ScrollOrientation = "horizontal"
	ScrollOrientationVertical   ScrollOrientation = "vertical"
)

func (e ScrollOrientation) Valid() bool {
	switch e {
	case "horizontal", "vertical":
		return true
	default:
		return false
	}
}

func (e ScrollOrientation) String() string {
	return string(e)
}

// Node DOM interaction is implemented in terms of mirror objects that
// represent the actual DOM nodes. DOMNode is a base node mirror type.
type Node struct {
	NodeID           NodeID                `json:"nodeId"`                     // Node identifier that is passed into the rest of the DOM messages as the `nodeId`. Backend will only push node with given `id` once. It is aware of all requested nodes and will only fire DOM events for nodes known to the client.
	ParentID         *NodeID               `json:"parentId,omitempty"`         // The id of the parent node if any.
	BackendNodeID    BackendNodeID         `json:"backendNodeId"`              // The BackendNodeId for this node.
	NodeType         int                   `json:"nodeType"`                   // `Node`'s nodeType.
	NodeName         string                `json:"nodeName"`                   // `Node`'s nodeName.
	LocalName        string                `json:"localName"`                  // `Node`'s localName.
	NodeValue        string                `json:"nodeValue"`                  // `Node`'s nodeValue.
	ChildNodeCount   *int                  `json:"childNodeCount,omitempty"`   // Child count for `Container` nodes.
	Children         []Node                `json:"children,omitempty"`         // Child nodes of this node when requested with children.
	Attributes       []string              `json:"attributes,omitempty"`       // Attributes of the `Element` node in the form of flat array `[name1, value1, name2, value2]`.
	DocumentURL      *string               `json:"documentURL,omitempty"`      // Document URL that `Document` or `FrameOwner` node points to.
	BaseURL          *string               `json:"baseURL,omitempty"`          // Base URL that `Document` or `FrameOwner` node uses for URL completion.
	PublicID         *string               `json:"publicId,omitempty"`         // `DocumentType`'s publicId.
	SystemID         *string               `json:"systemId,omitempty"`         // `DocumentType`'s systemId.
	InternalSubset   *string               `json:"internalSubset,omitempty"`   // `DocumentType`'s internalSubset.
	XMLVersion       *string               `json:"xmlVersion,omitempty"`       // `Document`'s XML version in case of XML documents.
	Name             *string               `json:"name,omitempty"`             // `Attr`'s name.
	Value            *string               `json:"value,omitempty"`            // `Attr`'s value.
	PseudoType       PseudoType            `json:"pseudoType,omitempty"`       // Pseudo element type for this node.
	PseudoIdentifier *string               `json:"pseudoIdentifier,omitempty"` // Pseudo element identifier for this node. Only present if there is a valid pseudoType.
	ShadowRootType   ShadowRootType        `json:"shadowRootType,omitempty"`   // Shadow root type.
	FrameID          *internal.PageFrameID `json:"frameId,omitempty"`          // Frame ID for frame owner elements.
	ContentDocument  *Node                 `json:"contentDocument,omitempty"`  // Content document for frame owner elements.
	ShadowRoots      []Node                `json:"shadowRoots,omitempty"`      // Shadow root list for given element host.
	TemplateContent  *Node                 `json:"templateContent,omitempty"`  // Content document fragment for template elements.
	PseudoElements   []Node                `json:"pseudoElements,omitempty"`   // Pseudo elements associated with this node.
	// ImportedDocument is deprecated.
	//
	// Deprecated: Deprecated, as the HTML Imports API has been removed
	// (crbug.com/937746). This property used to return the imported
	// document for the HTMLImport links. The property is always undefined
	// now.
	ImportedDocument  *Node             `json:"importedDocument,omitempty"`
	DistributedNodes  []BackendNode     `json:"distributedNodes,omitempty"`  // Distributed nodes for given insertion point.
	IsSVG             *bool             `json:"isSVG,omitempty"`             // Whether the node is SVG.
	CompatibilityMode CompatibilityMode `json:"compatibilityMode,omitempty"` // No description.
	AssignedSlot      *BackendNode      `json:"assignedSlot,omitempty"`      // No description.
}

// DetachedElementInfo A structure to hold the top-level node of a detached
// tree and an array of its retained descendants.
type DetachedElementInfo struct {
	TreeNode        Node     `json:"treeNode"`        // No description.
	RetainedNodeIDs []NodeID `json:"retainedNodeIds"` // No description.
}

// RGBA A structure holding an RGBA color.
type RGBA struct {
	R int      `json:"r"`           // The red component, in the [0-255] range.
	G int      `json:"g"`           // The green component, in the [0-255] range.
	B int      `json:"b"`           // The blue component, in the [0-255] range.
	A *float64 `json:"a,omitempty"` // The alpha component, in the [0-1] range (default: 1).
}

// Quad An array of quad vertices, x immediately followed by y for each point,
// points clock-wise.
type Quad []float64

// BoxModel Box model.
type BoxModel struct {
	Content      Quad              `json:"content"`                // Content box
	Padding      Quad              `json:"padding"`                // Padding box
	Border       Quad              `json:"border"`                 // Border box
	Margin       Quad              `json:"margin"`                 // Margin box
	Width        int               `json:"width"`                  // Node width
	Height       int               `json:"height"`                 // Node height
	ShapeOutside *ShapeOutsideInfo `json:"shapeOutside,omitempty"` // Shape outside coordinates
}

// ShapeOutsideInfo CSS Shape Outside details.
type ShapeOutsideInfo struct {
	Bounds      Quad              `json:"bounds"`      // Shape bounds
	Shape       []json.RawMessage `json:"shape"`       // Shape coordinate details
	MarginShape []json.RawMessage `json:"marginShape"` // Margin shape bounds
}

// Rect Rectangle.
type Rect struct {
	X      float64 `json:"x"`      // X coordinate
	Y      float64 `json:"y"`      // Y coordinate
	Width  float64 `json:"width"`  // Rectangle width
	Height float64 `json:"height"` // Rectangle height
}

// CSSComputedStyleProperty
type CSSComputedStyleProperty struct {
	Name  string `json:"name"`  // Computed style property name.
	Value string `json:"value"` // Computed style property value.
}