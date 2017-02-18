package model

import "testing"

func TestAttachedSourceProperTypeName(t *testing.T) {
	as := AttachedSource{}
	if as.GetName() != "attached-sources" {
		t.Error("Attached Sources have a wrong JSON API name")
	}
}

func TestAttachedSourceHasProperID(t *testing.T) {
	as := AttachedSource{ID: 1024}
	if as.GetID() != "1024" {
		t.Error("Attached sources don't have a proper string ID")
	}
}

func TestAttachedSourceCanSetProperID(t *testing.T) {
	as := AttachedSource{ID: 1024}
	as.SetID("4096")
	if as.ID != 4096 {
		t.Error("Attached sources don't have a proper string ID setter")
	}
}
