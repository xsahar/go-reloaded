package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 3 {
		fmt.Println("Usage: go run main.go <input-file> <output-file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read the input file
	inputFileObj, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening input file: %s\n", err.Error())
		return
	}
	defer inputFileObj.Close()

	scanner := bufio.NewScanner(inputFileObj)

	var line string
	var temp string
	// Scan the file for lines.
	for scanner.Scan() {
		// Get the current line of text.
		line = scanner.Text()
		temp += line
		// Process the line as needed.
	}

	// Close the file.
	defer inputFileObj.Close()

	// Check for any errors.
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Perform modifications on the text
	modifiedText := temp // Modify the text as per your requirements
	// Replace all instances of (hex) and (bin) with their decimal equivalents
	re := regexp.MustCompile(`(\w+)\s+\((hex|bin)\)`)
	modifiedText = re.ReplaceAllStringFunc(modifiedText, func(match string) string {
		// Get the word before the modification type
		word := strings.Split(match, " ")[0]

		// Convert the word to decimal
		decimalValue := ""
		if match[len(match)-5:] == "(bin)" {
			decimalValue = binToDec(word)
		} else {
			decimalValue = hexToDec(word)
		}

		// Return the modified text
		return fmt.Sprint(decimalValue)
	})

	// Replace all instances of a with an if the next word begins with a vowel or a h
	re = regexp.MustCompile(`(?i)a\s+([aeiouh])`)
	modifiedText = re.ReplaceAllStringFunc(modifiedText, replaceAWithAn)

	ce := regexp.MustCompile(`\b(\w+)\s+\(\s*cap\s*\)`)
	modifiedText = ce.ReplaceAllStringFunc(modifiedText, func(match string) string {
		word := strings.TrimSuffix(match[:len(match)-5], " ")
		capitalizedWord := capitalize(word)
		return capitalizedWord + strings.TrimPrefix(match, word+" (cap)")
	})

	me := regexp.MustCompile(`\b([A-Za-z]+)\s*\(\s*low\s*\)`)
	modifiedText = me.ReplaceAllStringFunc(modifiedText, func(match string) string {
		word := strings.TrimSuffix(match[0:len(match)-5], " ")
		return toLowerCase(word)
	})

	ue := regexp.MustCompile(`(\w+)\s*\(\s*up\s*\)`)
	modifiedText = ue.ReplaceAllStringFunc(modifiedText, func(match string) string {
		word := strings.TrimSuffix(match[0:len(match)-5], " ")
		return toUpperCase(word)
	})

	lowRe := regexp.MustCompile(`(\b[\w\s]+\b)\s+\(low,\s*(\d+)\)`)
	modifiedText = lowRe.ReplaceAllStringFunc(modifiedText, func(match string) string {
		parts := lowRe.FindStringSubmatch(match)
		word := parts[1]
		count := parts[2]
		return toLowerCaseNWords(word, count)
	})

	UpRe := regexp.MustCompile(`(\b[\w\s]+\b)\s+\(up,\s*(\d+)\)`)
	modifiedText = UpRe.ReplaceAllStringFunc(modifiedText, func(match string) string {
		parts := UpRe.FindStringSubmatch(match)
		word := parts[1]
		count := parts[2]
		return toUpCaseNWords(word, count)
	})

	CapRe := regexp.MustCompile(`(\b[\w\s]+\b)\s+\(cap,\s*(\d+)\)`)
	modifiedText = CapRe.ReplaceAllStringFunc(modifiedText, func(match string) string {
		parts := CapRe.FindStringSubmatch(match)
		word := parts[1]
		count := parts[2]
		return toCapCaseNWords(word, count)
	})

	// Write the modified text to the output filemodifyText
	outputFileObj, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %s\n", err.Error())
		return
	}
	defer outputFileObj.Close()

	_, err = outputFileObj.WriteString(modifiedText)
	if err != nil {
		fmt.Printf("Error writing output file: %s\n", err.Error())
		return
	}

	editTextFile(outputFileObj.Name())

	fmt.Println("Text modification completed successfully.")
}

func hexToDec(hex string) string {
	dec := 0
	base := 1

	for i := len(hex) - 1; i >= 0; i-- {
		switch hex[i] {
		case '0':
			dec += 0 * base
		case '1':
			dec += 1 * base
		case '2':
			dec += 2 * base
		case '3':
			dec += 3 * base
		case '4':
			dec += 4 * base
		case '5':
			dec += 5 * base
		case '6':
			dec += 6 * base
		case '7':
			dec += 7 * base
		case '8':
			dec += 8 * base
		case '9':
			dec += 9 * base
		case 'A', 'a':
			dec += 10 * base
		case 'B', 'b':
			dec += 11 * base
		case 'C', 'c':
			dec += 12 * base
		case 'D', 'd':
			dec += 13 * base
		case 'E', 'e':
			dec += 14 * base
		case 'F', 'f':
			dec += 15 * base
		}
		base *= 16
	}

	return fmt.Sprintf("%d", dec)
}

func binToDec(bin string) string {
	dec := 0
	base := 1

	for i := len(bin) - 1; i >= 0; i-- {
		if bin[i] == '1' {
			dec += base
		}
		base *= 2
	}

	return fmt.Sprintf("%d", dec)
}

func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(word[:1]) + word[1:]
}

func toLowerCase(word string) string {
	return strings.ToLower(word)
}

func toUpperCase(word string) string {
	return strings.ToUpper(word)
}

func replaceAWithAn(match string) string {
	if match[1] >= 'A' && match[1] <= 'z' {
		return "An" + match[1:]
	} else {
		return "an" + match[1:]
	}
}

func toLowerCaseNWords(word string, count string) string {
	n, err := strconv.Atoi(count)
	if err != nil {
		// Invalid count, return the original word
		return word
	}

	words := strings.Fields(word)

	length := len(words)

	// Iterate halfway through the list
	for i := 0; i < length/2; i++ {
		// Swap elements from the beginning and end of the list
		words[i], words[length-i-1] = words[length-i-1], words[i]
	}

	if n >= len(words) {
		// If the count is greater than or equal to the number of words, convert all words to lowercase
		for i := range words {
			words[i] = strings.ToLower(words[i])
		}
	} else {
		// Convert the words starting from the specified index to lowercase
		for i := 0; i < n; i++ {
			words[i] = strings.ToLower(words[i])
		}
	}

	// Iterate halfway through the list
	for i := 0; i < length/2; i++ {
		// Swap elements from the beginning and end of the list
		words[i], words[length-i-1] = words[length-i-1], words[i]
	}

	return strings.Join(words, " ")
}

func toUpCaseNWords(word string, count string) string {
	n, err := strconv.Atoi(count)
	if err != nil {
		// Invalid count, return the original word
		return word
	}

	words := strings.Fields(word)

	length := len(words)

	// Iterate halfway through the list
	for i := 0; i < length/2; i++ {
		// Swap elements from the beginning and end of the list
		words[i], words[length-i-1] = words[length-i-1], words[i]
	}

	if n >= len(words) {
		// If the count is greater than or equal to the number of words, convert all words to lowercase
		for i := range words {
			words[i] = strings.ToUpper(words[i])
		}
	} else {
		// Convert the words starting from the specified index to lowercase
		for i := 0; i < n; i++ {
			words[i] = strings.ToUpper(words[i])
		}
	}

	for i := 0; i < length/2; i++ {
		// Swap elements from the beginning and end of the list
		words[i], words[length-i-1] = words[length-i-1], words[i]
	}

	return strings.Join(words, " ")
}

func toCapCaseNWords(word string, count string) string {
	n, err := strconv.Atoi(count)
	if err != nil {
		// Invalid count, return the original word
		return word
	}

	words := strings.Fields(word)

	length := len(words)

	// Iterate halfway through the list
	for i := 0; i < length/2; i++ {
		// Swap elements from the beginning and end of the list
		words[i], words[length-i-1] = words[length-i-1], words[i]
	}

	if n >= len(words) {
		// If the count is greater than or equal to the number of words, convert all words to uppercase
		for i := range words {
			words[i] = strings.Title(words[i])
		}
	} else {
		// Convert the first 'n' words to uppercase
		for i := 0; i < n; i++ {
			words[i] = strings.Title(words[i])
		}
	}

	for i := 0; i < length/2; i++ {
		// Swap elements from the beginning and end of the list
		words[i], words[length-i-1] = words[length-i-1], words[i]
	}

	return strings.Join(words, " ")
}
