package helpers

import (
	"fmt"
	"math/rand"
)

func CapoCheck() {
	if rand.Intn(100) >= 50 {
		fmt.Println("no capo")
	} else {
		capoPosition := rand.Intn(6) + 1 // 1 to 6
		fmt.Printf("capo on fret %d\n", capoPosition)
	}
}
