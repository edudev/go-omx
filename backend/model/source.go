package model

import "encoding/hex"
import "crypto/sha512"

// Source is a static media type, waiting to be played by a Renderer
type Source struct {
	URI string `json:"uri"`
}

// GetID is used to get a unique ID for each source
func (c Source) GetID() string {
	hash := sha512.Sum512_256([]byte(c.URI))
	return hex.EncodeToString(hash[:])
}
