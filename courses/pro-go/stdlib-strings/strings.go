package main

import (
	"fmt"
	"strings"
	"unicode"
)

func stringsComparison() {
	fmt.Println(
		"Contains:",
		strings.Contains("Hello World", "Hell"),  // true
		strings.Contains("Hello World", "llo W"), // true
		strings.Contains("Hello World", "Word"),  // false
	)

	// True if any letter of substr is contained in string
	// Case-sensitive
	fmt.Println(
		"ContainsAny:",
		strings.ContainsAny("Hello World", "Hwz"), // true
		strings.ContainsAny("Hello World", "abc"), // false
	)

	fmt.Println(
		"ContainsRune:",
		strings.ContainsRune("Città", 'à'), // true
		strings.ContainsRune("Città", 'a'), // false
		strings.ContainsRune("Città", 'ç'), // false
	)

	// Checks if a string starts/ends with a given substring
	fmt.Println(
		"HasPrefix/HasSuffix:",
		strings.HasPrefix("Hello", "Hell"),        // true
		strings.HasPrefix("Hello", "Hello World"), // false
		strings.HasSuffix("Hello World", "rld"),   // true
	)
}

func stringsCasing1() {
	fmt.Println(
		strings.ToLower("NBA, NFL"), // nba, nfl
	)

	fmt.Println(
		strings.ToUpper("say it out loud"), // SAY IT OUT LOUD
	)

	fmt.Println(
		strings.Title("this is a sentence"), // This Is A Sentence
	)

	// strings.ToTitle is better when working with Unicode
	fmt.Println(
		strings.ToTitle("ǳ"), // ǲ
		strings.ToUpper("ǳ"), // Ǳ
	)
}

func stringsCasing2() {
	fmt.Println(
		unicode.IsLower('A'), // false
		unicode.IsLower('à'), // true
		unicode.IsUpper('A'), // true
		unicode.IsUpper('b'), // false
		unicode.IsUpper('À'), // true

		// TODO: This is strange
		unicode.IsTitle('A'), // false
		unicode.IsTitle('À'), // false
	)
}

func stringsInspect() {

	isLetterB := func(r rune) bool {
		return r == 'B' || r == 'b'
	}

	fmt.Println(
		strings.Count("aabbaaccddeeaabb", "aa"),          // 3
		strings.Index("aabbaaccddeeaabb", "aa"),          // 0
		strings.Index("aabbaaccddeeaabb", "zz"),          // -1
		strings.LastIndex("aabbaaccddeeaabb", "aa"),      // 12
		strings.LastIndex("aabbaaccddeeaabb", "zz"),      // -1
		strings.IndexAny("aabbaaccddeeaabb", "az"),       // 0
		strings.IndexAny("aabbaaccddeeaabb", "zy"),       // -1
		strings.LastIndexAny("aabbaaccddeeaabb", "az"),   // 13
		strings.LastIndexAny("aabbaaccddeeaabb", "ot"),   // -1
		strings.IndexFunc("aabbaaccddeeaabb", isLetterB), // 2
	)
}

func stringsSplit() {
	fmt.Println(
		strings.Fields("this is a sentence"), // [this is a sentence]
		strings.Fields("this   a  sentence"), // [this is a sentence]
		strings.FieldsFunc(
			"this_is-a+sentence_",
			func(r rune) bool {
				return r == '_' || r == '-' || r == '+'
			},
		), // [this is a sentece]
		strings.Split("aa bb cc", " "),        // [aa bb cc]
		strings.SplitN("aa_bb_cc_dd", "_", 3), // [aa bb cc_dd]
	)
}

func stringsTrim() {
	fmt.Println(
		// Trims on both sides, but *NOT* internally!
		strings.TrimSpace("  this is a  sentence  "),
		strings.Trim("aabbccdd", "ad"),               // bbcc
		strings.TrimLeft("aaFoo", "aa"),              // Foo
		strings.TrimRight("Barzz", "zz"),             // Bar
		strings.TrimPrefix("RemoveRemove", "Remove"), // Remove
		strings.TrimSuffix("RemoveRemove", "Remove"), // Remove
		// TrimFunc(s, func)
		// TrimLeftFunc(s, func)
		// TrimRightFunc(s, func)
	)
}

func stringsReplace() {
	mapper := func(r rune) rune {
		switch r {

		// Map 'b' to 'c'
		case 'b':
			return 'c'

		// Remove 'z'
		case 'z':
			return -1

		// Leave other letters there
		default:
			return r
		}
	}

	fmt.Println(
		strings.Replace("foobarfoo", "foo", "bar", 1),     // barbarfoo
		strings.Replace("foobarfoo", "foo", "bar", 2),     // barbarbar
		strings.Replace("foobarfoo", "foo", "bar", 10),    // barbarbar
		strings.ReplaceAll("azHelloazWorldaz", "az", "_"), // _Hello_World_
		strings.Map(mapper, "zI had a boat"),
	)

	// This accepts couples of strings interpreted as replaceThis => withThis
	// Ex.: boat => kayak, small => huge
	replacer := strings.NewReplacer(
		"boat", "kayak",
		"small", "huge",
	)
	replaced := replacer.Replace("It was a boat. A small boat.")
	fmt.Println("Replaced:", replaced)
}

func stringsBuild1() {
	fmt.Println(
		strings.Join(strings.Fields("I live in town"), "#"), // I#live#in#town
		strings.Repeat("-:", 3),                             // -:-:-:
	)
}

func stringsBuild2() {
	text := "It was a boat. A small boat."
	var builder strings.Builder

	builder.WriteString("This will be erased")
	builder.Reset()

	for _, word := range strings.Fields(text) {
		if word == "small" {
			builder.WriteString("very ")
		}
		builder.WriteString(word)
		builder.WriteRune(' ')
	}

	fmt.Println(
		"String:", builder.String(),
		"Cap:", builder.Cap(),
		"Len:", builder.Len(),
	)
}
