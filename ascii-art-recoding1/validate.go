package main

import (
	"fmt"
	"errors"
)

func ValidateInput(s string) (rune, error) {
	if len(s) == 0 {
		return 0, nil
	}

	for _, ch := range s {
		if ch < 32 || ch > 126 {
			err := fmt.Sprintf("unsupported characters: %c", rune(ch))
			return rune(ch), errors.New(err)
		}
	}
	return 0, nil
}




// func ValidateInput(s string) (rune, error) {

// 	for i, ch := range s {
// 		if (ch < 32 || ch > 126) && ch != '\n' {
// 			return ch, fmt.Errorf("got %q at position %d", ch, i)
// 			//return rune(ch), errors.New(err)
// 		}
// 	}
// 	return 0, nil
// }


