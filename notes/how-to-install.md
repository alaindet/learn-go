# How to install Go 1.17 on Ubuntu 20.04

1. Get link to latest binary tarbal from https://golang.org/dl/
2. Install Go
   1. Option 1
      ```
      cd ~
      # Change URL below as needed
      curl -O https://golang.org/dl/go1.17.1.linux-amd64.tar.gz
      sudo tar -xzvf go1.13.3.linux-amd64.tar.gz -C /usr/local
      sudo chown -R root:root /usr/local/go
      sudo apt-get update
      sudo apt-get install golang-go
      ```
   2. Option 2
      ```
      sudo add-apt-repository ppa:longsleep/golang-backports
      sudo apt-get update
      sudo apt-get install golang-go
      ```
3. `mkdir -p $HOME/go/{bin,src}`
4. Open `~/.profile` and add at end
   ```
   export GOPATH=$HOME/go
   export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin
   ```
5. `source ~/.profile`
6. `go version` (Should print installed version)

## Test if it works

1. `mkdir $GOPATH/hello_world`
2. `cd $GOPATH/hello_world`
3. `touch main.go`
4. Edit `main.go` and add this
   ```
   package main

   import "fmt"

   func main() {
       fmt.Println("Hello World")
   }
   ```
5. `go run main.go` (Should print "Hello World" into the console)
