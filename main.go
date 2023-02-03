package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file_date, err := os.ReadFile("file.txt")

	if err != nil {
		fmt.Println("Error reading file\n", err)
	}
	str := string(file_date)
	strl := strings.ToLower(str)
	words := strings.Split(strl, " ")
	sort.Strings(words)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	fmt.Printf("%s", "file: ")
	for i := 0; i < len(words); i++ {
		count := 1
		for j := i + 1; j < len(words); j++ {
			if len(words[i]) == len(words[j]) {
				count++
				words[j] = "0"
			}
		}
		if count > 0 && words[i] != "0" {
			if len(words[i]) <= 9 {
				fmt.Print("[", len(words[i]), "] = ", count, ", ")
			} else {
				count++
				fmt.Print("[>9] = ", count)
				break
			}
		}
	}
}
