package main

import (
	"server"
)

func main() {
	s := server.NewServer()
	s.Run(":8080")
}
