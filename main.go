package main

import (
	"assignment1/identity"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// len(os.Args) = length of args
	if len(os.Args) <= 1 {
		fmt.Println("Please enter number absen")
		return
	}
	// os.Args[0] = tmp/go-build75578968/b001/exe/main
	// os.Args[1] = based on input terminal
	absen, err := strconv.Atoi(os.Args[1])
	_ = err
	if absen < 1 || absen > len(identity.Identity) {
		fmt.Println("Something went wrong. Please re-check your input")
		return
	}

	identity.Identity[absen-1].ShowData()
}
