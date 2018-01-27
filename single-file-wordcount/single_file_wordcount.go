package main

import (
"fmt"
"bufio"
"os"
"strings"
"regexp"
"log"
"strconv"
"net/http"

	"time"
)

//to show in the webbroser
func indexHandler1( w http.ResponseWriter, r *http.Request){

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	fmt.Fprintf(w, "word"+" "+"count")
	fmt.Fprintf(w,"\n----------\n")
	for key, val := range wordcountMap {

		fmt.Println(strings.TrimRight(key, " "), val)
		fmt.Fprintf(w, reg.ReplaceAllString(key, "")+" "+strconv.Itoa(val))
		fmt.Fprintf(w,"\n")

	}

	if err != nil {
		log.Fatal(err)
	}

}


var wordcountMap =map[string]int{}

func main() {
	starttime := time.Now()
	//wordcountMap := make(map[string]int)
	inputfile, err := os.Open("/files/moby-000.txt")

	if err != nil {
		fmt.Println(err, "Fine not found")
	}

	filescanner := bufio.NewScanner(inputfile)
	for filescanner.Scan() {
		line := strings.Split(strings.TrimSpace(string(filescanner.Text())), " ")
		for _, word := range line {
			val, ifpresentkey := wordcountMap[word]
			if ifpresentkey {
				wordcountMap[word] = val + 1
			} else {
				wordcountMap[word] = 1
			}
		}
	}

	filehandle, err := os.Create("/output/" + "simplewordcountoutput.csv")
	if err != nil {
		fmt.Println("Error writing to file: ", err)
		return
	}
	defer filehandle.Close()
	writer := bufio.NewWriter(filehandle)
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-key-val")
	for key, val := range wordcountMap {

		fmt.Fprintln(writer, reg.ReplaceAllString(key, "")+","+strconv.Itoa(val))

	}
	elapsedtime := time.Since(starttime)
	fmt.Println("Time taken:",elapsedtime)
	writer.Flush()
	http.HandleFunc("/", indexHandler1)
	http.ListenAndServe(":9080", nil)
	fmt.Println("Time taken:",elapsedtime)
}