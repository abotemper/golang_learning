package main

// has to have a package, on the convention we call it a main

//standard library
import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	fmt.Println("hello, world! println version")
	fmt.Print("hello, ")
	fmt.Print("world, print version\n")
	sayHelloWorld("hello, world, function version")

	var printVar1 string
	printVar1 = "hello, world, variable version 1"
	sayHelloWorld(printVar1)

	printVar2 := "hello, world, variable version 2" //:= is combining declaration and set value
	sayHelloWorld(printVar2)

	var printVar3 string = "hello, world, variable version 3"
	sayHelloWorld(printVar3)

	// var printVar4 string
	// printVar4 = doctor.Intro()
	var printVar4 string = doctor.Intro()
	fmt.Println(printVar4)

	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n') //\n means whenever user type an enter
	//when user type enter , it will store in userInput variable

	fmt.Println(userInput)

	//in go , you need noly one type loop

	// for {
	// 	userInput2, _ := reader.ReadString('\n')
	// 	fmt.Println((userInput2))
	// }

	for {
		fmt.Print("-> ")
		userInput2, _ := reader.ReadString('\n')

		userInput2 = strings.Replace(userInput2, "\r\n", "", -1)
		userInput2 = strings.Replace(userInput2, "\n", "", -1)

		if userInput2 == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput2))
		}

	}

}

func sayHelloWorld(whatToSay string) {
	fmt.Println(whatToSay)
}

//hello
