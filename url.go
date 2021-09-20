package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	f, err := os.Open("son.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	colorRed := "\033[31m"
	colorBlue := "\033[37m"
	colorYellow := "\033[33m"
	colorWhite := "\033[34m"
	colorGreen := "\033[35m"

	fmt.Println(string(colorGreen), "https://github.com/jabir-dev/Admin-Finder-GO.git")

	fmt.Println(string(colorWhite), "--------------------------------------------------------------------------------")

	fmt.Println(string(colorYellow), "HTTP y HTTPS [ http://www.example.com ]  [ https://www.example.com ] ")

	fmt.Println("Enter web url :")

	var input string

	fmt.Scan(&input)
	fmt.Println("WEBSITE: ", input)

	fmt.Println(string(colorWhite), "--------------------------------------------------------------------------------")

	for scanner.Scan() {

		link3 := input + "/" + scanner.Text()

		resp, err := http.Get(link3)

		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
			fmt.Println("ADMIN PAGE ITS FOUND :", string(colorRed), link3, resp.StatusCode, http.StatusText(resp.StatusCode))
		} else {
			fmt.Println("ADMIN PAGE NOT FOUND :", string(colorBlue), link3, resp.StatusCode, http.StatusText(resp.StatusCode))
		}

	}

}
