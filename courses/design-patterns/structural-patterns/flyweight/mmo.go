package main

import "strings"

type User struct {
	FullName string
}

func NewUser(fullName string) *User {
	return &User{fullName}
}

var allNames []string

type UserCompact struct {
	names []uint8
}

func NewUserCompact(fullName string) *UserCompact {
	getOrAdd := func(s string) uint8 {
		for i := range allNames {
			if allNames[i] == s {
				return uint8(i)
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames) - 1)
	}

	user := UserCompact{}
	parts := strings.Split(fullName, " ") // <-- This is weak!
	for _, p := range parts {
		user.names = append(user.names, getOrAdd(p))
	}

	return &user
}

func (u *UserCompact) FullName() string {
	var parts []string
	for _, id := range u.names {
		parts = append(parts, allNames[id])
	}
	return strings.Join(parts, " ")
}

func mmoExample() {

	// This approach "wastes" memory
	// john := NewUser("John Doe")
	// jane := NewUser("Jane Doe")
	// anotherJane := NewUser("Jane Smith")

	// This approach saves memory since "Jane" is repeated
	// john2 := NewUserCompact("John Doe")
	// jane2 := NewUserCompact("Jane Doe")
	// anotherJane2 := NewUserCompact("Jane Smith")

	// TODO: Calculate weights and print them
}
