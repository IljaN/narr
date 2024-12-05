// Code generated by cdpgen. DO NOT EDIT.

package input

import (
	"encoding/json"
	"errors"
	"time"
)

// TouchPoint
type TouchPoint struct {
	X             float64  `json:"x"`                       // X coordinate of the event relative to the main frame's viewport in CSS pixels.
	Y             float64  `json:"y"`                       // Y coordinate of the event relative to the main frame's viewport in CSS pixels. 0 refers to the top of the viewport and Y increases as it proceeds towards the bottom of the viewport.
	RadiusX       *float64 `json:"radiusX,omitempty"`       // X radius of the touch area (default: 1.0).
	RadiusY       *float64 `json:"radiusY,omitempty"`       // Y radius of the touch area (default: 1.0).
	RotationAngle *float64 `json:"rotationAngle,omitempty"` // Rotation angle (default: 0.0).
	Force         *float64 `json:"force,omitempty"`         // Force (default: 1.0).
	// TangentialPressure The normalized tangential pressure, which has a
	// range of [-1,1] (default: 0).
	//
	// Note: This property is experimental.
	TangentialPressure *float64 `json:"tangentialPressure,omitempty"`
	TiltX              *float64 `json:"tiltX,omitempty"` // The plane angle between the Y-Z plane and the plane containing both the stylus axis and the Y axis, in degrees of the range [-90,90], a positive tiltX is to the right (default: 0)
	TiltY              *float64 `json:"tiltY,omitempty"` // The plane angle between the X-Z plane and the plane containing both the stylus axis and the X axis, in degrees of the range [-90,90], a positive tiltY is towards the user (default: 0).
	// Twist The clockwise rotation of a pen stylus around its own major
	// axis, in degrees in the range [0,359] (default: 0).
	//
	// Note: This property is experimental.
	Twist *int     `json:"twist,omitempty"`
	ID    *float64 `json:"id,omitempty"` // Identifier used to track touch sources between events, must be unique within an event.
}

// GestureSourceType
//
// Note: This type is experimental.
type GestureSourceType string

// GestureSourceType as enums.
const (
	GestureSourceTypeNotSet  GestureSourceType = ""
	GestureSourceTypeDefault GestureSourceType = "default"
	GestureSourceTypeTouch   GestureSourceType = "touch"
	GestureSourceTypeMouse   GestureSourceType = "mouse"
)

func (e GestureSourceType) Valid() bool {
	switch e {
	case "default", "touch", "mouse":
		return true
	default:
		return false
	}
}

func (e GestureSourceType) String() string {
	return string(e)
}

// MouseButton
type MouseButton string

// MouseButton as enums.
const (
	MouseButtonNotSet  MouseButton = ""
	MouseButtonNone    MouseButton = "none"
	MouseButtonLeft    MouseButton = "left"
	MouseButtonMiddle  MouseButton = "middle"
	MouseButtonRight   MouseButton = "right"
	MouseButtonBack    MouseButton = "back"
	MouseButtonForward MouseButton = "forward"
)

func (e MouseButton) Valid() bool {
	switch e {
	case "none", "left", "middle", "right", "back", "forward":
		return true
	default:
		return false
	}
}

func (e MouseButton) String() string {
	return string(e)
}

// TimeSinceEpoch UTC time in seconds, counted from January 1, 1970.
type TimeSinceEpoch float64

// String calls (time.Time).String().
func (t TimeSinceEpoch) String() string {
	return t.Time().String()
}

// Time parses the Unix time.
func (t TimeSinceEpoch) Time() time.Time {
	ts := float64(t) / 1
	secs := int64(ts)
	nsecs := int64((ts - float64(secs)) * 1000000000)
	return time.Unix(secs, nsecs)
}

// MarshalJSON implements json.Marshaler. Encodes to null if t is zero.
func (t TimeSinceEpoch) MarshalJSON() ([]byte, error) {
	if t == 0 {
		return []byte("null"), nil
	}
	f := float64(t)
	return json.Marshal(&f)
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *TimeSinceEpoch) UnmarshalJSON(data []byte) error {
	*t = 0
	if len(data) == 0 {
		return nil
	}
	var f float64
	if err := json.Unmarshal(data, &f); err != nil {
		return errors.New("input.TimeSinceEpoch: " + err.Error())
	}
	*t = TimeSinceEpoch(f)
	return nil
}

var _ json.Marshaler = (*TimeSinceEpoch)(nil)
var _ json.Unmarshaler = (*TimeSinceEpoch)(nil)

// DragDataItem
//
// Note: This type is experimental.
type DragDataItem struct {
	MimeType string  `json:"mimeType"`          // Mime type of the dragged data.
	Data     string  `json:"data"`              // Depending of the value of `mimeType`, it contains the dragged link, text, HTML markup or any other data.
	Title    *string `json:"title,omitempty"`   // Title associated with a link. Only valid when `mimeType` == "text/uri-list".
	BaseURL  *string `json:"baseURL,omitempty"` // Stores the base URL for the contained markup. Only valid when `mimeType` == "text/html".
}

// DragData
//
// Note: This type is experimental.
type DragData struct {
	Items              []DragDataItem `json:"items"`              // No description.
	Files              []string       `json:"files,omitempty"`    // List of filenames that should be included when dropping
	DragOperationsMask int            `json:"dragOperationsMask"` // Bit field representing allowed drag operations. Copy = 1, Link = 2, Move = 16
}
