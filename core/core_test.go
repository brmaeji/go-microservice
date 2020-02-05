package service

import (
	"testing"
)

func TestNewMentionsService(t *testing.T) {

	_, err := NewMentionsService(1)
	if err != nil {
		t.Errorf("could not instantiate new MentionsService: %v\n", err)
	}
}
