package helpers

import (
	"math/rand"
)

type Genre struct {
	Name   string
	AvgBPM int
}

func GetRandomBPM() int {
	// List of genres with their average BPMs
	genres := []Genre{
		{"Pop", 120},
		{"Rock", 130},
		{"Hip-Hop", 90},
		{"Electronic", 128},
		{"Jazz", 110},
		{"Classical", 60},
		{"Reggae", 80},
		{"Metal", 140},
		{"Punk", 160},
		{"Thrash", 180},
		{"chill out dude", 200},
	}

	randomGenre := genres[rand.Intn(len(genres))]
	randomBPM := rand.Intn(21) + (randomGenre.AvgBPM - 10)

	return randomBPM
}
