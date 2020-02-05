package helpers

import (
	"microservice/models"
	"testing"
)

var mapa map[string]*models.BandMention
var creationKeys []string

func init() {

	mapa = make(map[string]*models.BandMention)

	creationKeys = []string{
		"The Beatles",
		"The Cure",
		"Iron Maiden",
		"Metallica",
		"Queen",
	}

	for _, k := range creationKeys {
		mapa[k] = &models.BandMention{
			Name:     k,
			Mentions: 1,
		}
	}
}

func TestGetKeys(t *testing.T) {
	keys := GetKeys(mapa)
	if len(keys) != 5 {
		t.Errorf("couldn't get the correct keys number from map. Expected 5, got %v\n", len(keys))
	}

	for i, k := range keys {
		if k != keys[i] {
			t.Errorf("expected key %v, got %v", k, keys[i])
		}
	}
}
