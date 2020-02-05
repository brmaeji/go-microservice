package data

import (
	"fmt"
	"microservice/models"
)

//enum of data adapter types

//AdapterType defines which data adapter will be used
type AdapterType int

const (
	//Undefined will be chosen if a wrong parameter is passed
	Undefined = iota
	//Memory Adapters will store data in memory maps
	Memory = iota
	//Postgres will use Postgres DB created with .env variables
	Postgres = iota
)

//Adapter defines the methods that the DBAdapter has to have
type Adapter interface {
	Find(name string) ([]models.BandMention, error)
	Increase(bm *models.BandMention) (*models.BandMention, error)
	Create(name string) error
}

//NewAdapter creates new a data adapter
func NewAdapter(t AdapterType) (Adapter, error) {
	switch t {
	case Memory:
		a, err := NewMemoryAdapter()
		if err != nil {
			return nil, err
		}
		return a, nil
		break
	case Postgres:
		a, err := NewPostgresAdapter()
		if err != nil {
			return nil, err
		}
		return a, nil
		break
	}

	return nil, fmt.Errorf("Unrecognized AdapterType with ID %v", t)
}
