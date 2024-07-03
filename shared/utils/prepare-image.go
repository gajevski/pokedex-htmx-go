package utils

import "fmt"

func PreparePokemonImage(id int) string {
	if id < 10 {
		return fmt.Sprintf("00%d", id)
	}
	if id < 100 && id >= 10 {
		return fmt.Sprintf("0%d", id)
	}
	return fmt.Sprintf("%d", id)
}
