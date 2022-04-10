module myapp

go 1.18

// This does not exist yet
replace github.com/alaindet/gomitolo => ../gomitolo

require github.com/alaindet/gomitolo v0.0.0-00010101000000-000000000000

require github.com/joho/godotenv v1.4.0 // indirect
