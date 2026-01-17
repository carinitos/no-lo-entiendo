package main

import (
	"fmt"
	"no_lo_entiendo/helpers"
)

func main() {
	flag := true
	for flag {
		var title string
		fmt.Print("Enter song title: ")
		fmt.Scan(&title)

		fmt.Println("reading from input file")
		words, err := helpers.ListAndReadTextFile()
		if err != nil {
			fmt.Printf("error: %s", err.Error())
			return
		}

		fmt.Println("generating lyrics...")
		err = helpers.GenerateCutUpSentences(words, title)
		if err != nil {
			fmt.Printf("error: %s", err.Error())
			return
		}

		fmt.Println("generating key and chords")
		helpers.GenerateChords()

		fmt.Println("checking capo usage")
		helpers.CapoCheck()

		fmt.Println("generating bpm")
		bpm := helpers.GetRandomBPM()
		fmt.Printf("bpm: %d \n", bpm)

		var choice int
		fmt.Print("Press 1 to generate again. Press 2 to exit application: ")
		_, err = fmt.Scan(&choice)

		if err != nil || choice == 2 {
			fmt.Println("exiting...")
			flag = false
			continue
		}

		fmt.Println("continuing...")
	}
}
