package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// Open connection
	recordFile, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println("An error encountered ::", err)
	}
	// Initialize the reader
	reader := csv.NewReader(recordFile)
	// Read all the records
	records, _ := reader.ReadAll()

	shuffle(&records)

	var correct int = 0
	var wrong int = 0
	var input string

	for i, val := range records {
		fmt.Printf("%d) %s\nAns : ", i+1, strings.TrimSpace(val[0]))
		fmt.Scanln(&input)
		if input == strings.TrimSpace(val[1]) {
			correct++
		} else {
			wrong++
		}
		fmt.Print("\n\n")
	}
	clearConsole()
	fmt.Println("Correct : " + fmt.Sprint(correct) + "\nIncorrect : " + fmt.Sprint(wrong))
}

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func shuffle(data *[][]string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*data), func(i, j int) {
		(*data)[i], (*data)[j] = (*data)[j], (*data)[i]
	})
}
