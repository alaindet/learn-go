# JSON

- JSON Encoding is mainly performed via the standard `encoding/json` package
- `json.Marshal()` and `json.Encoder()` can be used to encode to JSON, with the former being preferred
- `json.Unmarshal()` and `json.Decoder()` can be used to decode from JSON, with the former being preferred
- The library is based on `struct tags`, which are a language feature to add annotations to struct fields
- Struct tags are short, simple and pretty idiomatic, but accessing them requires reflection which is broadly discouraged (WTF)
- Example usage for struct tags
  ```go
  func main() {
    type User struct {
      Name  string `firstTag:"The 1st tag of Name" secondTag:"2nd tag"`
      Email string `firstTag:"The 1st tag of Email"`
    }

    u := User{"Bob", "bob@mycompany.com"}
    t := reflect.TypeOf(u)

    for i := 0; i < t.NumField(); i++ {
      f := t.Field(i)
      fmt.Printf("\nField: User.%s\n", f.Name)
      fmt.Printf("\tTag value : %q\n", f.Tag)
      fmt.Printf("\tValue of 'firstTag': %q\n", f.Tag.Get("firstTag"))
      fmt.Printf("\tValue of 'secondTag': %q\n", f.Tag.Get("secondTag"))
    }

    /*
      Field: User.Name
        Tag value : "firstTag:\"The 1st tag of Name\" secondTag:\"2nd tag\""
        Value of 'firstTag': "The 1st tag of Name"
        Value of 'secondTag': "2nd tag"

      Field: User.Email
        Tag value : "firstTag:\"The 1st tag of Email\""
        Value of 'firstTag': "The 1st tag of Email"
        Value of 'secondTag': ""
    */
  }
  ```
- Other libraries exist, like
  - `https://github.com/francoispqt/gojay`
  - `https://github.com/json-iterator/go`
  - `https://github.com/goccy/go-json`
- I recommend `goccy/go-json` as of early 2023 since it's compatible with `encoding/json` and faster as well

## Go to JSON type conversions

| Go                                | JSON           |
|-----------------------------------|----------------|
| `bool`                            | boolean        |
| `string`                          | string         |
| `int*`, `uint*`, `float*`, `rune` | number         |
| `array`, `slice`                  | array          |
| `struct`, `map`                   | object         |
| `nil` pointers                    | null           |
| `chan`, `func`, `complex*`        | UNSUPPORTED    |
| `time.Time`                       | RFC3339        |
| `[]byte`                          | Base-64 string |

## JSON to Go type conversions

| JSON    | Go                                  |
|---------|-------------------------------------|
| boolean | `bool`                              |
| string  | `string`                            |
| number  | `int*`, `uint*`, `float*` or `rune` |
| array   | `array` or `slice`                  |
| object  | `struct` or `map`                   |

- Unsupported types **forces decoding to fail**
- Example RFC3339 date `"2023-01-06T10:17:08.8241416+01:00"`

## Encoding
- For encoding, `json.Marshal()` is preferred

```go
import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	fmt.Println(toJson(map[string]any{
		"bool":   true,
		"string": "Foo bar",
		"int64":  int64(1234),
		"int":    1234,
		"float":  12.34,
		"arr":    [2]string{"foo", "bar"},
		"slice":  []string{"baz", "qez"},
		"struct": struct{ Name string }{"Foo"},
		"map":    map[string]int{"foo": 123},
		"nil":    make([]string, 0, 3),
		"now":    time.Now(),
		"b64":    []byte("This is a string"),
	}))
  /*
  {
    "arr": [
      "foo",
      "bar"
    ],
    "b64": "VGhpcyBpcyBhIHN0cmluZw==",
    "bool": true,
    "float": 12.34,
    "int": 1234,
    "int64": 1234,
    "map": {
      "foo": 123
    },
    "nil": [],
    "now": "2009-11-10T23:00:00Z",
    "slice": [
      "baz",
      "qez"
    ],
    "string": "Foo bar",
    "struct": {
      "Name": "Foo"
    }
  }
  */

	fmt.Println(toJson(map[string]any{
		"fn": func() {},
		"ch": make(chan int),
	}))
  /*
  {}
  */
}

func toJson(data any) string {
	result, err := json.Marshal(data)
	if err != nil {
		return "{}"
	}
	return string(result)
}
```

## Decoding
- For decoding, `json.Decoder()` is preferred

```go
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

func main() {

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
			Extra:map[
        favoriteColor:Green
        favoritePokemon:Pikachu
      ]
		}
	*/
}
```
