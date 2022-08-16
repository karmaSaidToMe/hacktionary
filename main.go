package main

import (
	"fmt"
//	"os"
//	"strconv"
)

var lenghtPwd int

func main() {
	generatePwds()

	/*args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Hacktionary \n   -l [integer] password length option.")		
		os.Exit(1)
	}

	switch args[0] {
		case "-l":
			lenghtPwd, _ = strconv.Atoi(args[1])
			fmt.Println(lenghtPwd)
			fmt.Println()

			generatePwds()

		case "-h":
			fmt.Println("Hacktionary \n   -l [integer] password length option.")
	}*/
}

func generatePwds() {
	letters   := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	//integers  := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	
	for a := 0; a < len(letters); a++ {
		for b := 0; b < len(letters); b++ {
			for c := 0; c < len(letters); c++ {
				for d := 0; d < len(letters); d++ {
					for e := 0; e < len(letters); e++ {
						for f := 0; f < len(letters); f++ {
							for g := 0; g < len(letters); g++ {
								for h := 0; h < len(letters); h++ {
									string := fmt.Sprintf("%s%s%s%s%s%s%s%s\n", letters[a], letters[b], letters[c], letters[d], letters[e], letters[f], letters[g], letters[h])
									fmt.Print(string)
								}
							}
						}
					}
				}
			}
		}
	}
}

// func fileWrite(){
// 	file, err := os.Create("hacktionary.txt")

// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		os.Exit(1)
// 	}
// 	defer file.Close()
// 	file.WriteString(string)

// 	fmt.Println("Yup")
// }