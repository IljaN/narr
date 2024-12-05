// Code generated by cdpgen. DO NOT EDIT.

package cachestorage

import (
	"github.com/mafredri/cdp/protocol/storage"
)

// CacheID Unique identifier of the Cache object.
type CacheID string

// CachedResponseType type of HTTP response cached
type CachedResponseType string

// CachedResponseType as enums.
const (
	CachedResponseTypeNotSet         CachedResponseType = ""
	CachedResponseTypeBasic          CachedResponseType = "basic"
	CachedResponseTypeCORS           CachedResponseType = "cors"
	CachedResponseTypeDefault        CachedResponseType = "default"
	CachedResponseTypeError          CachedResponseType = "error"
	CachedResponseTypeOpaqueResponse CachedResponseType = "opaqueResponse"
	CachedResponseTypeOpaqueRedirect CachedResponseType = "opaqueRedirect"
)

func (e CachedResponseType) Valid() bool {
	switch e {
	case "basic", "cors", "default", "error", "opaqueResponse", "opaqueRedirect":
		return true
	default:
		return false
	}
}

func (e CachedResponseType) String() string {
	return string(e)
}

// DataEntry Data entry.
type DataEntry struct {
	RequestURL         string             `json:"requestURL"`         // Request URL.
	RequestMethod      string             `json:"requestMethod"`      // Request method.
	RequestHeaders     []Header           `json:"requestHeaders"`     // Request headers
	ResponseTime       float64            `json:"responseTime"`       // Number of seconds since epoch.
	ResponseStatus     int                `json:"responseStatus"`     // HTTP response status code.
	ResponseStatusText string             `json:"responseStatusText"` // HTTP response status text.
	ResponseType       CachedResponseType `json:"responseType"`       // HTTP response type
	ResponseHeaders    []Header           `json:"responseHeaders"`    // Response headers
}

// Cache Cache identifier.
type Cache struct {
	CacheID        CacheID         `json:"cacheId"`                 // An opaque unique id of the cache.
	SecurityOrigin string          `json:"securityOrigin"`          // Security origin of the cache.
	StorageKey     string          `json:"storageKey"`              // Storage key of the cache.
	StorageBucket  *storage.Bucket `json:"storageBucket,omitempty"` // Storage bucket of the cache.
	CacheName      string          `json:"cacheName"`               // The name of the cache.
}

// Header
type Header struct {
	Name  string `json:"name"`  // No description.
	Value string `json:"value"` // No description.
}

// CachedResponse Cached response
type CachedResponse struct {
	Body []byte `json:"body"` // Entry content, base64-encoded. (Encoded as a base64 string when passed over JSON)
}
