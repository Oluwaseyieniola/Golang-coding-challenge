package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"path/filepath"
	"time"
)

var current_time int
var current_year int
var current_date int

var birth_month int
var birth_year int

func main() {

	//filepath := flag.String("c", "", "path to the file")

	//flag.Parse()

	//if *filepath == "" {
	//	log.Fatal("please enter  a valid filepath")
	//}
	//fileinfo, err := os.Stat(*filepath)

	//if err != nil {
	//log.Fatal(err)
	//}

	//fileBytesize := int(fileinfo.Size())
	//fmt.Printf("The size of the file is %d bytes\n", fileBytesize)
	current_year = time.Now().Year()
	fmt.Println("Please enter your year of birth")
	fmt.Scanln(&birth_year)

	calculate_age(current_year, birth_year)
	// define the c and l flags
	linesCount := flag.String("l", "", "count the number of lines in the file ")
	wordsCount := flag.String("w", "", "count the number of words in the file")
	charactercount := flag.String("m", "", "count the number of strings int the file")
	filename := filepath.Base(*linesCount)
	// parse the flags
	flag.Parse()

	if *charactercount != "" {
		file, err := os.Open(*charactercount)
		if err != nil {
			log.Fatal("error opening file %v", err)
		}
		defer file.Close()

		content, err := os.ReadFile(*charactercount)
		if err != nil {
			log.Fatal("error reading the file %v", err)
		}
		charcount := len(content)

		fmt.Printf("The file '%s' includes %d characters", *charactercount, charcount)
	}
	if *wordsCount != "" {
		file, err := os.Open(*wordsCount)
		if err != nil {
			log.Fatal("error opening file %v", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		wordsCount := 0
		for scanner.Scan() {
			wordsCount++
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading the file %v", err)
		}

		fmt.Printf("the file '%s' include %d words", *&wordsCount, wordsCount)
	}
	// check if the  flag is provided then open file for reading
	if *linesCount != "" {
		file, err := os.Open(*linesCount)
		if err != nil {
			log.Fatal("error opening file %v", err)
		}
		defer file.Close()
		// use scanner to read the lines in each file
		scanner := bufio.NewScanner(file)
		linesCount := 0
		for scanner.Scan() {
			linesCount++
		}
		// check if any errror ocurred during reading

		if err := scanner.Err(); err != nil {
			fmt.Printf("error reading the file %v", err)
		}
		// output the number of line
		fmt.Printf("The file '%s', include %d lines", filename, linesCount)
	}

	if *linesCount == "" {
		log.Fatal("Please provide a filepath using -l")
	}

	// i neither of the tags are provided throw an error

	file, err := os.Open("")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var count int

	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("there are", count, "lines in the file")
}

func calculate_age(current_year int, birth_year int) int {
	age := current_year - birth_year
	fmt.Println("Your age is:", age)
	return age
}

//  output wc command
