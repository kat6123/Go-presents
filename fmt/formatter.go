package main

import (
	"fmt"
	"strings"
)

// START OMIT

func main() {
	hdr := Header("GO WALKTHROUGH")
	fmt.Printf("%2.3d\n", hdr)
	fmt.Printf("%1.5s\n", hdr)
}

// Header represents formattable header text.
type Header string

func (hdr Header) String() string{
	return "string"
}

// Format decorates the header with pounds and snowmen.
func (hdr Header) Format(f fmt.State, c rune) {
	wid, _ := f.Width()
	prec, _ := f.Precision()

	f.Write([]byte(strings.Repeat("#", wid)))
	f.Write([]byte(hdr))
	f.Write([]byte(strings.Repeat("☃", prec)))
}

// END OMIT
