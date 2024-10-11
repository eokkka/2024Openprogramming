package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	i, _ := in.ReadString('\n')
	fmt.Println(i)

}
