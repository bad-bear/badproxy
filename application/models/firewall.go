package models

type FireWall struct {
	Table       string
	Protocol    string
	Source      string
	Destination string
	Port        string
	Target      string
}
