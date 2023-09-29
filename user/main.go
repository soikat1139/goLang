package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello There")
	


	reader := bufio.NewReader(os.Stdin)

	fmt.Println("PleasE RATE our Service Here ");

	//comma ok syntax in go 

	input,_ :=reader.ReadString('\n');

	fmt.Println("Thanks you ",input)

	fmt.Printf("The type of input is %T",input)

	numRating,err := strconv.ParseFloat(strings.TrimSpace((input)),64)

	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println("There You The Increased Rating",numRating+1)
	}

}