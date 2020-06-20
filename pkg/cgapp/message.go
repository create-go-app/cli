package cgapp

import "fmt"

// SendMessage function for send message (colored or common)
func SendMessage(text, color string) {
	//
	if color == "" {
		// Send common text
		fmt.Printf("%v\n", text)
	} else {
		// Define variables
		var (
			red       string = "\033[0;31m"
			green     string = "\033[0;32m"
			cyan      string = "\033[0;36m"
			yellow    string = "\033[1;33m"
			noColor   string = "\033[0m"
			textColor string
		)

		// Switch color
		switch color {
		case "green":
			textColor = green
			break
		case "yellow":
			textColor = yellow
			break
		case "red":
			textColor = red
			break
		case "cyan":
			textColor = cyan
			break
		}

		// Send colored text
		fmt.Printf("%v%v%v\n", textColor, text, noColor)
	}
}
