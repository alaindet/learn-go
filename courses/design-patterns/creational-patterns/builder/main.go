/*
The builder pattern provides an abstraction layer to create objects in steps, when
such objects require too many arguments and/or configuration to be created

- A builder is a separate piece of software used to build other things
- Builder can be fluent/chainable when functions return the receiver
- Builders can aggregate other sub-builders
*/
package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 4
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	trailingWhitespace := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", trailingWhitespace, e.name))

	if len(e.text) > 0 {
		nextTrailingWhitespace := strings.Repeat(" ", indentSize*(indent+1))
		sb.WriteString(fmt.Sprintf("%s%s\n", nextTrailingWhitespace, e.text))
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}

	sb.WriteString(fmt.Sprintf("%s</%s>\n", trailingWhitespace, e.name))
	return sb.String()
}

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName,
		HtmlElement{rootName, "", []HtmlElement{}},
	}
}

func (b *HtmlBuilder) String() string {
	return b.root.String()
}

func (b *HtmlBuilder) AddChild(childName, childText string) {
	childElement := HtmlElement{childName, childText, []HtmlElement{}}
	b.root.elements = append(b.root.elements, childElement)
}

func (b *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	b.AddChild(childName, childText)
	return b
}

func main() {
	// Built-in in Go
	sb := strings.Builder{}

	// Example 1
	sb.Reset()
	sb.WriteString("<ul>")
	sb.WriteString("<li>Hello</li>")
	sb.WriteString("<li>World</li>")
	sb.WriteString("</ul>")
	fmt.Println("Example 1: ", sb.String())

	// Example 2
	sb.Reset()
	words := []string{"Hello", "World"}
	sb.WriteString("<ul>")
	for _, w := range words {
		sb.WriteString(fmt.Sprintf("<li>%s</li>", w))
	}
	sb.WriteString("</ul>")
	fmt.Println("Example 2: ", sb.String())

	// Example 3
	ul := NewHtmlBuilder("ul")
	ul.AddChild("li", "Hello")
	ul.AddChild("li", "World")

	// Notice the weird syntax
	ul. // <-- This is weird
		AddChildFluent("li", "How are you?"). // <-- This is weird
		AddChildFluent("li", "I'm fine thanks")

	fmt.Println("Example 3")
	fmt.Println(ul.String())

	multipleBuildersExample()
	builderParameterExample()
	functionalBuilderExample()
}
