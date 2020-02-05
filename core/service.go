package service

import (
	"log"
	data "microservice/adapters/data"
)

//MentionService handles all functions related to the Band Mentions app
type MentionService struct {
	data.Adapter
}

//NewMentionsService creates a new instance of the MentionService
func NewMentionsService(t data.AdapterType) (MentionService, error) {

	dataAdapter, err := data.NewAdapter(t)
	if err != nil {
		return MentionService{}, err
	}

	service := MentionService{dataAdapter}

	log.Println("MentionService created...")
	return service, nil
}
