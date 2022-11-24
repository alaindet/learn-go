/*
Singleton Design Pattern

A "singleton" is just a component with one single instance into the entire app

Reasons involve expensive creation or just business logic
Examples are specialized factories, database connections, loggers etc.
It is usually considered an anti-pattern due to misuse

To make singletons testable, they should implement an interface (ex.: Database)
*/

package main

import (
	"fmt"
	"sync"
)

// This is crucial to test singletons
type Database interface {
	GetPopulation(city string) int
}

type singletonDatabase struct {
	capitals map[string]int // [cityName]population
}

func (db *singletonDatabase) GetPopulation(city string) int {
	return db.capitals[city]
}

type dummyDatabase struct {
	capitals map[string]int
}

func (d *dummyDatabase) GetPopulation(city string) int {

	// Lazy-load predictable data
	if len(d.capitals) == 0 {
		d.capitals = map[string]int{
			"Alpha": 10,
			"Beta":  20,
			"Gamma": 30,
		}
	}

	return d.capitals[city]
}

func GetTotalPopulation(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {

	// This is lazy-loaded and thread safe!
	once.Do(func() {
		db := singletonDatabase{}
		capitals, err := readCapitalsData("capitals.txt")
		if err != nil {
			panic("cannot read the file")
		}
		db.capitals = capitals
		instance = &db
	})

	return instance
}

func main() {
	db := GetSingletonDatabase()
	population := db.GetPopulation("Sao Paulo")
	fmt.Println("Population of Sao Paulo", population)

	total := GetTotalPopulation(db, []string{"Tokyo", "Osaka"})
	fmt.Println("Total population", total)

	// Simulating a test with fake dependency
	mockDb := &dummyDatabase{}
	mockTotal := GetTotalPopulation(mockDb, []string{"Alpha", "Gamma"})
	fmt.Println("Mock total population", mockTotal)
}
