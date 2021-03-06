// +build encrypt_passwords

package main

import (
	"fmt"
	"github.com/nightshaders/stpeter/auth"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {

	fmt.Print("Enter password: ")
	p, err := terminal.ReadPassword(0)
	fmt.Println()

	if err != nil {
		panic("Could not obtain password")
	}

	pp, err := auth.EncryptPassword(p, "salt")
	fmt.Println(pp)
}
