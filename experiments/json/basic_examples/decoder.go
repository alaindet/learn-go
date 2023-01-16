package basic_examples

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type MyStruct struct {
	ID       uint64            `json:"id"`
	Name     string            `json:"name"`
	Hobbies  []string          `json:"hobbies"`
	Position [2]float64        `json:"position"`
	Extra    map[string]string `json:"extra"`
}

func DecodeJSON() {

	jsonData := []byte(`{
		"id": 123456789,
		"name": "John Doe",
		"hobbies": ["skiing", "chess"],
		"position": [46.460732, 8.213391],
		"extra": {
			"favoritePokemon": "Pikachu",
			"favoriteColor": "Green"
		}
	}`)

	var my MyStruct
	reader := bytes.NewReader(jsonData)
	_ = json.NewDecoder(reader).Decode(&my)

	fmt.Printf("%+v\n", my)
	/*
		{
			ID:123456789
			Name:John Doe
			Hobbies:[skiing chess]
			Position:[46.460732 8.213391]
			Extra:map[favoriteColor:Green favoritePokemon:Pikachu]
		}
	*/
}
