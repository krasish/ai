package utils

import (
	"fmt"
	"log"
)

func readUntilCorrectU(u *uint) {
	for _, err := fmt.Scanf("%d", u); err != nil; {
		log.Println(err)
		_, err = fmt.Scanf("%d", u)
	}
}

func readUntilCorrect(i *int) {
	for _, err := fmt.Scanf("%d", i); err != nil; {
		log.Println(err)
		_, err = fmt.Scanf("%d", i)
	}
}

func readUntilCorrectB(b *bool) {
	for _, err := fmt.Scanf("%t", b); err != nil; {
		log.Println(err)
		_, err = fmt.Scanf("%t", b)
	}
}

func readUntilCorrectYesNo(b *bool) {
	var answer string
	for _, err := fmt.Scanf("%s", &answer); err != nil || !(answer == "yes" || answer == "no"); {
		fmt.Println("The answer may be 'yes' or 'no'")
		_, err = fmt.Scanf("%s", &answer)
	}
	switch answer {
	case "yes":
		*b = true
	case "no":
		*b = false
	}
}

func ReadUnsigned(u *uint, prompt string) {
	fmt.Print(prompt)
	readUntilCorrectU(u)
}

func ReadInt(i *int, prompt string) {
	fmt.Print(prompt)
	readUntilCorrect(i)
}

func ReadBool(b *bool, prompt string) {
	fmt.Print(prompt)
	readUntilCorrectB(b)
}

func ReadYesNo(b *bool, prompt string) {
	fmt.Print(prompt)
	readUntilCorrectYesNo(b)
}

func printUnsigned(u uint, nl bool) {
	if nl {
		fmt.Println(u)
	} else {
		fmt.Print(u)
	}
}

func printInt(i int8, nl bool) {
	if nl {
		fmt.Println(i)
	} else {
		fmt.Print(i)
	}
}
