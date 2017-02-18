package model

import "testing"

func TestSourceHasID(t *testing.T) {
	s := Source{URI: "test"}
	if s.GetID() != "" {
		return
	}
	t.Error("Unable to get Unique ID of Source")
}

func TestSourcesHaveUniqueID(t *testing.T) {
	s1 := Source{URI: "test"}
	s2 := Source{URI: "test"}
	s3 := Source{URI: "test2"}
	if s1.GetID() != s2.GetID() || s1.GetID() == s3.GetID() {
		t.Error("Sources don't have a unique ID based on URI")
	}
}
