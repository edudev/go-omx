package model

import "encoding/hex"
import "crypto/sha512"

type Source struct {
	Uri   string `json:"uri"`
}

func (c Source) GetID() string {
    hash := sha512.Sum512_256([]byte(c.Uri))
    return hex.EncodeToString(hash[:])
}
