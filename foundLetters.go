package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter a string:")

	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')

	str = strings.TrimSpace(str)
	str = strings.ToLower(str)

	isValid := strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a")
	if isValid {
		fmt.Println("Found!")
		return
	}

	fmt.Println("Not Found!")
}
