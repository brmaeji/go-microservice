package models

//BandMention is the document that associates the band with the number of times it was mentioned
type BandMention struct {
	Name     string `json:"name"`
	Mentions int    `json:"mentions"`
}
