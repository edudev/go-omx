package model

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"strconv"
)

// AttachedSource struct connects the source and the renderer and designates a running
// media source
type AttachedSource struct {
	ID               int       `json:"-"`
	SourceID         string    `json:"-"`
	Source           *Source   `json:"-"`
	RendererID       string    `json:"-"`
	Renderer         *Renderer `json:"-"`
	PlaybackStatus   string    `json:"playback-status"`
	PlaybackPosition float64   `json:"playback-position"`
	IsMuted          bool      `json:"is-muted"`
	Volume           float64   `json:"volume"`
}

// GetName returns this model's type to comply with JSON API
func (as AttachedSource) GetName() string {
	return "attached-sources"
}

// GetID returns this object's ID
func (as AttachedSource) GetID() string {
	return strconv.Itoa(as.ID)
}

// SetID is used when creating/updating AttachedSource instances
func (as *AttachedSource) SetID(ids string) error {
	if ids == "" {
		return nil
	}

	id, err := strconv.Atoi(ids)
	as.ID = id
	return err
}

// GetReferences is used for JSON API compliance to list all related models
func (as AttachedSource) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "source",
			Name: "source",
		},
		{
			Type: "renderer",
			Name: "renderer",
		},
	}
}

// GetReferencedIDs is used for JSON API compliance to list all related models' instances
func (as AttachedSource) GetReferencedIDs() []jsonapi.ReferenceID {
	result := []jsonapi.ReferenceID{}

	result = append(result, jsonapi.ReferenceID{
		ID:   as.SourceID,
		Type: "source",
		Name: "source",
	})

	result = append(result, jsonapi.ReferenceID{
		ID:   as.RendererID,
		Type: "renderer",
		Name: "renderer",
	})

	return result
}

// GetReferencedStructs is used for JSON API compliance to list all related models' instances
func (as AttachedSource) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	result = append(result, as.Source)
	result = append(result, as.Renderer)

	return result
}

// SetToOneReferenceID is used for JSON API compliance when creating/updating the object's related fields
func (as *AttachedSource) SetToOneReferenceID(name, ID string) error {
	if name == "source" {
		as.SourceID = ID
	} else if name == "renderer" {
		as.RendererID = ID
	} else {
		return errors.New("Invalid relation type")
	}

	return nil
}
