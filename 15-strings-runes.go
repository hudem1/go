package main

import (
	"fmt"
	"unicode/utf8"
)

// A bit complicated, so better take a look at the link below
// link: https://gobyexample.com/strings-and-runes
func strings_runes_main() {
	// A string is composed of runes each composed of bytes
	// Each rune can have different byte length (ex: a rune might be 2 bytes long while another is 3)
	const s = "สวัสดี"

	// Len counts the number of bytes
	fmt.Println("Len:", len(s))

	// Display each byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// Display each rune and the byte they start at (might be several bytes long)
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// Save as above!
	// Display each rune and the byte they start at (might be several bytes long)
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		// Just an additional function call
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
