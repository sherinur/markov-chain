package main

var words []string

func main() {
	// parsing data
	prefixStr, prefixLength, maxWords := ParseData()

	PrintMarkovChain(prefixStr, prefixLength, maxWords)
}
