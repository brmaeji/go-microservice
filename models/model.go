package models

//BandMention is the document that associates the band with the number of times it was mentioned
type BandMention struct {
	Name     string `gorm:"primary_key;" json:"name"`
	Mentions int    `gorm:"not null;" json:"mentions"`
}
