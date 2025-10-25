package main

import (
	"fmt"
	"strings"
)

func findDuplicate(paths []string) [][]string {
	contentMap := make(map[string][]string)
	for _, path := range paths {
		parts := strings.Split(path, " ")
		directory := parts[0]
		for i := 1; i < len(parts); i++ {
			file := parts[i]
			openParen := strings.Index(file, "(")
			fileName := file[:openParen]
			content := file[openParen+1 : len(file)-1]
			fullPath := directory + "/" + fileName
			contentMap[content] = append(contentMap[content], fullPath)
		}
	}

	var result [][]string
	for _, filePaths := range contentMap {
		if len(filePaths) > 1 {
			result = append(result, filePaths)
		}
	}
	return result
}

func main() {
	// Example 1
	paths1 := []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"}
	output1 := findDuplicate(paths1)
	fmt.Println("Example 1:")
	fmt.Printf("Input: %v\n", paths1)
	fmt.Printf("Output: %v\n", output1)
	fmt.Println("--------------------")

	// Example 2
	paths2 := []string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)"}
	output2 := findDuplicate(paths2)
	fmt.Println("Example 2:")
	fmt.Printf("Input: %v\n", paths2)
	fmt.Printf("Output: %v\n", output2)
	fmt.Println("--------------------")

	// Edge case: No duplicates
	paths3 := []string{"root/a 1.txt(abcd) 2.txt(efgh)"}
	output3 := findDuplicate(paths3)
	fmt.Println("Edge Case: No Duplicates")
	fmt.Printf("Input: %v\n", paths3)
	fmt.Printf("Output: %v\n", output3)
	fmt.Println("--------------------")

	// Edge case: Empty input
	paths4 := []string{}
	output4 := findDuplicate(paths4)
	fmt.Println("Edge Case: Empty Input")
	fmt.Printf("Input: %v\n", paths4)
	fmt.Printf("Output: %v\n", output4)
	fmt.Println("--------------------")
}