/*
Flyweigtht design pattern

Optimizes space by externalizing repeating or potentially repeating data
associated with objects. Ex.: Objects in a collection have a favorite color ID
referring to an external color map or table; if two people have the same favorite
color, they have the same ID
*/
package main

func main() {
	badFormatterExample()
	goodFormatterExample()
}
