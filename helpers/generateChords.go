package helpers

import (
	"fmt"
	"math/rand"
)

//generate key
//generate chord pattern

var chordWheel = map[string][]string{
	"C":  {"C", "Dm", "Em", "F", "G", "Am", "Bdim"},
	"G":  {"G", "Am", "Bm", "C", "D", "Em", "F#dim"},
	"D":  {"D", "Em", "F#m", "G", "A", "Bm", "C#dim"},
	"A":  {"A", "Bm", "C#m", "D", "E", "F#m", "G#dim"},
	"E":  {"E", "F#m", "G#m", "A", "B", "C#m", "D#dim"},
	"B":  {"B", "C#m", "D#m", "E", "F#", "G#m", "A#dim"},
	"F#": {"F#", "G#m", "A#m", "B", "C#", "D#m", "E#dim"},
	"Db": {"Db", "Ebm", "Fm", "Gb", "Ab", "Bbm", "Cdim"},
	"Ab": {"Ab", "Bbm", "Cm", "Db", "Eb", "Fm", "Gdim"},
	"Eb": {"Eb", "Fm", "Gm", "Ab", "Bb", "Cm", "Ddim"},
	"Bb": {"Bb", "Cm", "Dm", "Eb", "F", "Gm", "Adim"},
	"F":  {"F", "Gm", "Am", "Bb", "C", "Dm", "Edim"},
}

var chordProgressions = [][]string{
	{"I", "IV", "V", "IV"},
	{"I", "V", "vi", "IV"},
	{"ii", "V", "I"},
	{"I", "vi", "IV", "V"},
	{"I", "IV", "vi", "V"},
	{"vi", "IV", "I", "V"},
	{"i", "bVII", "III", "bVII"},
	{"I", "iii", "IV", "V"},
	{"I", "ii", "iii", "IV", "V"},
	{"IV", "I", "ii", "V"},
	{"i", "V", "i", "iv"},
	{"i", "bVII", "iv", "bVII"},
	{"I", "V", "vi", "iii"},
}

func pickRandomKey() string {
	keys := make([]string, 0, len(chordWheel))
	for key := range chordWheel {
		keys = append(keys, key)
	}
	return keys[rand.Intn(len(keys))]
}

func pickRandomProgression() []string {
	return chordProgressions[rand.Intn(len(chordProgressions))]
}

func mapProgressionToChords(key string, progression []string) []string {
	chords := chordWheel[key]

	degreeMap := map[string]int{
		"I": 0, "ii": 1, "iii": 2, "IV": 3, "V": 4, "vi": 5, "vii": 6,
		"i": 0, "bVII": 6, "III": 2, "iv": 3,
	}
	mappedChords := []string{}

	for _, degree := range progression {
		if index, exists := degreeMap[degree]; exists {
			mappedChords = append(mappedChords, chords[index])
		} else {
			mappedChords = append(mappedChords, "?")
		}
	}

	return mappedChords
}

func GenerateChords() {
	randKey := pickRandomKey()
	fmt.Printf("Random Key: %s\n", randKey)

	randProgression := pickRandomProgression()
	fmt.Printf("Chord Progression: %v\n", randProgression)

	mappedChords := mapProgressionToChords(randKey, randProgression)
	fmt.Printf("Mapped Chords: %v\n", mappedChords)
}
