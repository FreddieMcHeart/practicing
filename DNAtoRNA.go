package main // do not remove
import "fmt"

func DNAtoRNA(dna string) (result string) {

	for _, elem := range dna {
		if string(elem) == "T" {
			result = result + "U"
		} else {
			result = result + string(elem)
		}
	}

	return result
}

func main() {
	var a string
	_, _ = fmt.Scanln(&a)
	fmt.Println(DNAtoRNA(a))
}
