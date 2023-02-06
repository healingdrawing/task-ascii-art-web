package ascii_art

import (
	"fmt"
	"os"
)

var unittest = false

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please enter some text")
		return
	}

	fmt.Print(AsciiToString(os.Args[1], "bannsdbfjhsdbf"))
}
