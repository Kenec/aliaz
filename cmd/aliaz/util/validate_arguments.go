package util

import (
	"fmt"
	"os"
)

func ValidateArguments(arguments ...string) {
	for i, str := range arguments {
		if str == "" {
			fmt.Printf("argument %d is empty or unset\n", i+1)
			os.Exit(1)
		}
	}
}
