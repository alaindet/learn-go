package main

type EnumLike interface {
	String() string
	Is(int) bool
	Allowed() bool
}
