package main

import ("example.com/greetings"
        "fmt"
        "log"
        "rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
    log.SetFlags(0)

	//Exercise 1
	fmt.Println("Exercise 1 result: ", quote.Go())

	//Exercise 2
    var greeting string
	greeting, err := greetings.Hello("Michael")
	if err != nil {
        log.Fatal(err)
    }
	fmt.Println("Exercise 2 result: ", greeting)

	//Exercise 3
	var names []string
	names = []string{"Michael", "Charlie"}

	hellos, err := greetings.Hellos(names)
	if err != nil {
        log.Fatal(err)
    }
	fmt.Println("Exercise 2 result: ", hellos)	
}
