package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
)

var (
	defaultPrefixLength int = 2
	defaultMaxWords     int = 100
)

func ParseData() (string, int, int) {
	// flags parsing
	maxWords := flag.Int("w", defaultMaxWords, "set number of maximum words")
	prefixLength := flag.Int("l", defaultPrefixLength, "set prefix length")
	prefixStr := flag.String("p", "", "set starting prefix")

	var prefix []string

	// program custom usage text.
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "Markov Chain text generator.\n")

		fmt.Fprintf(os.Stdout, "\n")
		fmt.Fprintf(os.Stdout, "Usage:\n")
		fmt.Fprintf(os.Stdout, "  markovchain [-w <N>] [-p <S>] [-l <N>]\n")
		fmt.Fprintf(os.Stdout, "  markovchain --help\n")
		fmt.Fprintf(os.Stdout, "\n")
		fmt.Fprintf(os.Stdout, "Options:\n")
		fmt.Fprintf(os.Stdout, "  --help  Show this screen.\n")
		fmt.Fprintf(os.Stdout, "  -w N    Number of maximum words\n")
		fmt.Fprintf(os.Stdout, "  -p S    Starting prefix\n")
		fmt.Fprintf(os.Stdout, "  -l N    Prefix length\n")
	}
	flag.Parse()

	// error handling
	if *maxWords <= 0 || *prefixLength <= 0 {
		fmt.Fprintf(os.Stderr, "Error: given number can't be negative or zero.\n")
		os.Exit(1)
	}
	if *prefixLength > 5 {
		fmt.Fprintf(os.Stderr, "Error: given prefix length can't be greater than 5\n")
		os.Exit(1)
	}
	if *maxWords > 10000 {
		fmt.Fprintf(os.Stderr, "Error: given number can't be more 10,000.\n")
		os.Exit(1)
	}

	// reading from the pipeline
	ReadPipeline()

	if len(*prefixStr) == 0 {
		prefix = words[:*prefixLength]
		*prefixStr = strings.Join(prefix, " ")
		if len(prefix) != *prefixLength {
			fmt.Fprintf(os.Stderr, "Error: the lenght of prefix is incorrect.\n")
			os.Exit(1)
		}
	}

	return *prefixStr, *prefixLength, *maxWords
}

func GetFrequencyMap(words []string, prefixLength int) map[string][]string {
	frequencyMap := make(map[string][]string)
	for i := 0; i < len(words)-prefixLength; i++ {
		if i >= len(words)-1 {
			break
		}
		prefix := words[i : i+prefixLength]
		currentWord := words[i+prefixLength]

		frequencyMap[strings.Join(prefix, " ")] = append(frequencyMap[strings.Join(prefix, " ")], currentWord)

	}
	return frequencyMap
}

func PrintMarkovChain(prefixStr string, prefixLength int, maxWords int) {
	frequencyMap := GetFrequencyMap(words, prefixLength)
	var generatedText string = prefixStr
	prefixArr := strings.Fields(prefixStr)

	if len(prefixArr) != prefixLength {
		fmt.Fprintf(os.Stderr, "Error: prefix length is incorrect.\n")
		os.Exit(1)
	}
	if len(frequencyMap) == 0 {
		fmt.Fprintf(os.Stderr, "Error: frequency map is empty\n")
		os.Exit(1)
	}
	if frequencyMap[prefixStr] == nil {
		fmt.Fprintf(os.Stderr, "Error: suffix doesn't exist\n")
		os.Exit(1)
	}

	currentPrefix := prefixStr

	for i := 0; i < maxWords-len(prefixArr); i++ {
		nextWords := frequencyMap[currentPrefix]
		if len(nextWords) == 0 {
			break
		}

		nextWord := nextWords[rand.Intn(len(nextWords))]
		generatedText += " " + nextWord

		prefixArr = append(prefixArr[1:], nextWord)
		currentPrefix = strings.Join(prefixArr, " ")

		if frequencyMap[prefixStr] == nil {
			break
		}
	}

	fmt.Println(generatedText)
	os.Exit(0)
}

// Reads all lines from the pipeline
func ReadPipeline() {
	// check if pipe is empty
	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: unable to read stdin\n")
		os.Exit(1)
	}

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintf(os.Stderr, "Error: no input text\n")
		os.Exit(1)
	}

	// reads all data from stdin
	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: no input text\n")
		os.Exit(1)
	}

	// splits into fields, gives an array
	inputStr := string(input)
	words = strings.Fields(inputStr)

	// error handling
	if len(words) == 0 {
		fmt.Fprintf(os.Stderr, "Error: no input text\n")
		os.Exit(1)
	}
}
