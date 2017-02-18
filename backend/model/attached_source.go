package model

import (
	"errors"
	"github.com/manyminds/api2go/jsonapi"
	"strconv"
)

type AttachedSource struct {
	ID               int       `json:"-"`
	SourceID         string    `json:"-"`
	Source           *Source   `json:"-"`
	RendererID       string    `json:"-"`
	Renderer         *Renderer `json:"-"`
	PlaybackStatus   string    `json:"playback-status"`
	PlaybackPosition int64     `json:"playback-position"`
	IsMuted          bool      `json:"is-muted"`
	Volume           float64   `json:"volume"`
}

func (as AttachedSource) GetName() string {
	return "attached-sources"
}

func (as AttachedSource) GetID() string {
	return strconv.Itoa(as.ID)
}

func (as *AttachedSource) SetID(ids string) error {
	if ids == "" {
		return nil
	}

	id, err := strconv.Atoi(ids)
	as.ID = id
	return err
}

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

func (as AttachedSource) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}

	result = append(result, as.Source)
	result = append(result, as.Renderer)

	return result
}

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
