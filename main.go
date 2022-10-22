package main

import (
	"fmt"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func main() {

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.CtrlC {
			return true, nil // Stop listener by returning true on Ctrl+C
		}

		fmt.Println("\r" + key.String()) // Print every key press
		return false, nil                // Return false to continue listening
	})
}
