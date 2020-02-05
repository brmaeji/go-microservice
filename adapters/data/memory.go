package data

import (
	"fmt"
	"log"
	"microservice/helpers"
	"microservice/models"
	"strings"
)

//MemoryAdapter sets up an in-memory hashmap acting as a data repository
type MemoryAdapter struct {
	mentions map[string]*models.BandMention
}

//NewMemoryAdapter creates a new memory adapter
func NewMemoryAdapter() (*MemoryAdapter, error) {
	ms := make(map[string]*models.BandMention)

	log.Println("MemoryAdapter created...")
	return &MemoryAdapter{ms}, nil
}

//Find returns all of the bands that have a name that match given param
func (ma *MemoryAdapter) Find(name string) ([]models.BandMention, error) {
	if val, ok := ma.mentions[name]; ok {
		//we have an exact match
		return []models.BandMention{*val}, nil
	}

	//we'll filter all the name bands that contain the given key
	keys := helpers.GetKeys(ma.mentions)

	//do it case-insensitive
	upperCaseMatch := strings.ToUpper(name)
	results := []models.BandMention{}

	for _, k := range keys {
		upperCaseKey := strings.ToUpper(k)
		if strings.Contains(upperCaseKey, upperCaseMatch) {
			results = append(results, *ma.mentions[k])
		}
	}

	return results, nil
}

//Increase will add one to the mentions on a BandMention model
func (ma *MemoryAdapter) Increase(bm *models.BandMention) (*models.BandMention, error) {
	bm.Mentions = bm.Mentions + 1
	ma.mentions[bm.Name] = bm

	return bm, nil
}

//Create adds a new BandMention on the stored data with an initial value of 1 mention
func (ma *MemoryAdapter) Create(name string) error {
	if _, ok := ma.mentions[name]; ok {
		return fmt.Errorf("there is already a band created for name %v", name)
	}

	newBm := &models.BandMention{
		Mentions: 1,
		Name:     name,
	}

	ma.mentions[name] = newBm
	return nil
}
