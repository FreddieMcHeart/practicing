package main

import "fmt"

func alphanumeric(str string) (result bool) {
	result = true
	if str == "" {
		result = false
	}

	for _, elem := range str {
		if string(elem) == " " {
			result = false
		} else if elem >= 48 && elem <= 57 {
			continue
		} else if elem >= 65 && elem <= 90 {
			continue
		} else if elem >= 97 && elem <= 122 {
			continue
		} else {
			result = false
		}
	}
	return result
}

func main() {
	var (
	//str = ""
	)
	fmt.Println(alphanumeric(".*?"))
	fmt.Println(alphanumeric("a"))
	fmt.Println(alphanumeric("Mazinkaiser"))
	fmt.Println(alphanumeric("hello world_"))
	fmt.Println(alphanumeric("PassW0rd"))
	fmt.Println(alphanumeric("     "))
	fmt.Println(alphanumeric(""))
	fmt.Println(alphanumeric("\n\t\n"))
	fmt.Println(alphanumeric("ciao\n$$_"))
	fmt.Println(alphanumeric("__ * __"))
	fmt.Println(alphanumeric("&)))((("))
	fmt.Println(alphanumeric("43534h56jmTHHF3k"))
	//fmt.Println(alphanumeric(str))
}
