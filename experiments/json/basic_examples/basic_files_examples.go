package basic_examples

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func FromJsonFile() {
	dir, _ := os.Getwd()
	inputJson := filepath.Join(dir, "assets", "basic.input.json")
	data, err := os.ReadFile(inputJson)

	if err != nil {
		panic("File basic.input.json not found")
	}

	var parsedFromJson Person
	err = json.Unmarshal(data, &parsedFromJson)
	_ = err

	fmt.Println(parsedFromJson)
}

func ToJsonFile() {
	dir, _ := os.Getwd()
	outputJson := filepath.Join(dir, "assets", "basic.output.json")
	p := Person{"John Doe", 42}
	data, _ := json.Marshal(p)
	os.WriteFile(outputJson, data, 0664)

	fmt.Println(data)
}
