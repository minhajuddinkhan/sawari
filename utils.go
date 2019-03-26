package sawari

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
)

func TakeInput(msg string) string {

	blue := color.New(color.FgBlue)
	white := color.New(color.FgWhite)
	blue.Println(msg)
	reader := bufio.NewReader(os.Stdin)
	for {

		white.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		return text
	}
}
