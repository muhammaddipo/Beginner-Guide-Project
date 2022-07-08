package main

import (
	"fmt"

	"github.com/muhammaddipo/Beginner-Guide-Project.git/internal/network"
)

func main() {
	fmt.Println("Awesome CLI v0.0.2")

	network.Ping("127.0.0.1")
}
