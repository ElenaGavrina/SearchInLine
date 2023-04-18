package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	m, _ := reader.ReadString('\n')
	fmt.Println(strings.Contains(m, "b"))
	fmt.Println(strings.Count(m, "b"))
}