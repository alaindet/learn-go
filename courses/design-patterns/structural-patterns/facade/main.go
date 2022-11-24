/*
Facade Design Pattern

The facade pattern provides a simple abstracted API hiding complexity of a
sub-system

Facades group stuff, but they should allow accessing the inner details if needed
*/
package main

type Buffer struct {
	widht, height int
	buffer        []rune
}

func NewBuffer(w, h int) *Buffer {
	return &Buffer{w, h, make([]rune, w*h)}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

type Viewport struct {
	buffer *Buffer
	offset int
}

func NewViewport(b *Buffer) *Viewport {
	return &Viewport{b, 0}
}

func (v *Viewport) GetCharAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

// This is a facade
type Console struct {
	buffer    []*Buffer
	viewports []*Viewport
	offset    int
}

func NewConsole() *Console {
	b := NewBuffer(200, 150)
	v := NewViewport(b)
	return &Console{[]*Buffer{b}, []*Viewport{v}, 0}
}

func (c *Console) GetCharAt(index int) rune {
	return c.viewports[0].GetCharAt(index)
}

func main() {
	c := NewConsole()
	c.GetCharAt(0)
}
