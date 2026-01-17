package helpers

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// use words and generate sentences
func GenerateCutUpSentences(words []string, filename string) error {
	if len(words) == 0 {
		return fmt.Errorf("no words available to create sentences")
	}

	sentences := []string{}
	for i := 0; i < 50; i++ { //this value is how many lines are generated, right now, it is 200
		sentenceLength := randomSentenceLength()
		var sentenceWords []string
		for j := 0; j < sentenceLength; j++ {
			word := words[rand.Intn(len(words))]
			sentenceWords = append(sentenceWords, word)
		}
		sentence := strings.Join(sentenceWords, " ") + "."
		sentences = append(sentences, sentence)
	}

	file, err := os.Create("output/" + filename + ".txt")
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range sentences {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}

	// Flush any buffered data to the file
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("failed to flush writer: %w", err)
	}

	fmt.Printf("file created: %s.txt \n", filename)
	return nil
}

func randomSentenceLength() int {
	return rand.Intn(20) + 1 //this value determines how many words per line, right now its up to 20
}
