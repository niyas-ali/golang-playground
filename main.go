package main

import (
	p "goplayground/playground"
)

func main() {
	server := p.NewAPIServer()
	server.Start()
}
