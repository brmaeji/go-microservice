package helpers

import (
	"microservice/models"
	"os"
)

//GetKeys will return a string slice of all the keys for that map
func GetKeys(mapa map[string]*models.BandMention) []string {

	keys := make([]string, len(mapa))

	i := 0
	for k := range mapa {
		keys[i] = k
		i++
	}

	return keys
}

//GetEnvVar reads env variable
func GetEnvVar(varName string) string {
	return os.Getenv(varName)
}
