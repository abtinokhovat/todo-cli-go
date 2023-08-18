package cmd

import (
	"bufio"
	"fmt"
)

func CommandBuilder(entity, op string) string {
	return fmt.Sprintf("%s-%s", entity, op)
}

func Scan(scanner *bufio.Scanner, message string) string {
	fmt.Println(message)
	scanner.Scan()
	return scanner.Text()
}
