package loops

import "fmt"

func Loops() {
	//looping
	arr := []int{10, 20, 30}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//infinite loop
	word := "Hi"
	for {
		if word == "Hello" {
			fmt.Println("Learn")
			break
		}
	}

	//while loop
	i := 0
	for i > 1 {
		i = i + 2
	}
	fmt.Println(i)

	//ranging
	for i, j := range arr { //i is index,j is value
		fmt.Println(i, j)
	}

	//ranging a string
	for index, char := range word {
		fmt.Println(index, char)
	}

	// ranging a map
	mapData := make(map[int]bool)
	for key, value := range mapData {
		fmt.Println(key, value)
	}

	// ranging a channel
	chnl := make(chan int)
	go func() {
		chnl <- 10
		chnl <- 100
		chnl <- 1000
	}()
	for i := range chnl {
		fmt.Println(i)
	}

}
