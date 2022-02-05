package main

func main() {
	// Shorter names are preferred, better with one letter, but they must be descriptive
	// Generally, the shorter the name, the more scoped it is
	index := 1 // Bad
	i := 2     // Good
	_, _ = index, i

	// Acronyms should remain in all caps
	parseHttp := false // Bad
	parseHTTP := false // Good
	_, _ = parseHttp, parseHTTP

	// Variables should follow camelCase or TitleCase, but not snake_case
	this_is_not_ok := true // Bad
	thisIsOK := true       // Good
	ThisIsOKEither := true // Good
	_, _, _ = this_is_not_ok, thisIsOK, ThisIsOKEither

	// Exported variables must be in TitleCase and begin with a capital letter
	ExportThisPlease := true
	_ = ExportThisPlease

	// Getters/Setters
	// Getters should have the same name as the referred variable, but in TitleCase
	// Setters should begin with Set*
	// TODO: Add examples

	// Interfaces with just one method must have the same name as the method with
	// the -er suffix. Ex.: Parse => Parser
	// TODO: Add examples
}
