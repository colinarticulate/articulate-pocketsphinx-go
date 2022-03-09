package main

import (
	"fmt"
	"multimap"
	"sync"
	"time"
)

type inputsCapitalize struct {
	input []string
}

func makeDataTestCapitalize() [5][]string {
	var data [5][]string //make([]inputsCapitalize, 5, 5)

	data[0] = []string{"oneaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
	data[1] = []string{"twoaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
	data[2] = []string{"threeaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
	data[3] = []string{"fouraaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
	data[4] = []string{"fiveaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}

	return data
}

func test_Capitlaize_non_concurrent() {
	inputs := makeDataTestCapitalize()

	n := len(inputs)
	fmt.Printf("\n\nNon-Concurrent test: \n")
	start := time.Now()
	for i := 0; i < n; i++ {
		//starti := time.Now()
		var input = []string{inputs[i][0]}
		multimap.Capitalize(input)
		//elapsedi := time.Since(starti)

		//fmt.Printf("Capitalize: %s\tresult: %v\t\t%s\n", inputs[i], input, elapsedi)
	}
	fmt.Println("Total execution time:", time.Since(start))
}

func Capitalize_concurrent(i int, wg *sync.WaitGroup, inputs [5][]string) {

	defer wg.Done()

	//starti := time.Now()
	var input = []string{inputs[i][0]}
	multimap.Capitalize(input)
	//elapsedi := time.Since(starti)

	//fmt.Printf("Capitalize: %s\tresult: %v\t\t%s\n", inputs[i], input, elapsedi)

}

func test_Capitlaize_concurrent() {
	inputs := makeDataTestCapitalize()

	n := len(inputs)
	init := 0
	var wg sync.WaitGroup
	fmt.Printf("\n\nConcurrent test: \n")
	start := time.Now()
	for i := init; i < n; i++ {
		wg.Add(1)
		go Capitalize_concurrent(i, &wg, inputs)
	}
	wg.Wait()
	fmt.Println("Total execution time:", time.Since(start))
}

func main() {
	// Call our gcd() function
	x := 42
	y := 105
	g := multimap.Gcd(x, y)
	fmt.Println("The gcd of ", x, " and ", y, " is ", g)

	// Call the gcdmain() function
	args := []string{"gcdmain", "42", "105"}
	multimap.Gcdmain(args)

	// Call the count function
	fmt.Println(multimap.Count("Hello World", 'l'))

	// Call the capitalize function
	capitalizeMe := []string{"hello world"}
	multimap.Capitalize(capitalizeMe)
	fmt.Println(capitalizeMe[0])

	test_Capitlaize_concurrent()
	test_Capitlaize_non_concurrent()

}
