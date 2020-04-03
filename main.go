package main

import (
	"fmt"

	"github.com/integrii/flaggy"
	"github.com/memcachier/bcrypt"
)

var (
	hashSubcommand *flaggy.Subcommand
	checkSubcommand *flaggy.Subcommand
	toHash string 
	hashedValue string
	plainValue string
)

func init() {
	flaggy.SetName("BCrypt tool")
	flaggy.SetDescription("A tool to hash values or to check hashed values against some plaintext.")

	hashSubcommand = flaggy.NewSubcommand("hash")
	hashSubcommand.AddPositionalValue(&toHash, "toHash", 1, true, "String to hash.")
	hashSubcommand.Description = "Hash a given value using bcrypt."

	checkSubcommand = flaggy.NewSubcommand("check")
	checkSubcommand.AddPositionalValue(&hashedValue, "hashed", 1, true, "Hashed value to check.")
	checkSubcommand.AddPositionalValue(&plainValue, "plain", 2, true, "Plain value to check against.")
	checkSubcommand.Description = "Check a bcrypt hash against some plaintext. You have to put the hash in '' so bash does not interpret the $ signs."

	flaggy.AttachSubcommand(hashSubcommand, 1)
	flaggy.AttachSubcommand(checkSubcommand, 1)

	flaggy.Parse()
}

func main() {
	if hashSubcommand.Used {
		hashValue()
	}

	if checkSubcommand.Used {
		checkValue()
	}
}

func hashValue() {
	salt, err := bcrypt.GenSalt(10)
	cipher, err := bcrypt.Crypt(toHash, salt)

	if err == nil {
		fmt.Println(cipher)
	} else {
		fmt.Println("An error occurred.")
	}
}

func checkValue() {
	match, err := bcrypt.Verify(plainValue, hashedValue)
	
	if err != nil {
		fmt.Println("An error occurred.")
		return
	}

	if match {
		fmt.Println("The values match.")
	} else {
		fmt.Println("The values do not match.")
	}
}