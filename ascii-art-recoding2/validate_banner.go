package main

import (
	"fmt"
)

func ValidateBanner(banner map[rune][]string) error {

	for key, value := range banner {
		if key < 32 || key > 126 {
			return fmt.Errorf("unsupported characters: %c", key)
		}

		if len(value) != 8 {
			return fmt.Errorf("invalid lines: %c", key)
		}

	}
	if len(banner) != 95 {
		return fmt.Errorf("empty banner files")
	}
	return nil
}
