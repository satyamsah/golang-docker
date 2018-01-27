
package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"path/filepath"
	"regexp"
	"time"
	"net/http"
	"log"
)

//logic to render webpage
func indexHandler2( w http.ResponseWriter, r *http.Request){

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	fmt.Fprintf(w, "word"+" "+"count")
	fmt.Fprintf(w,"\n----------\n")
	for key, val := range allfilewordcountmap {


		fmt.Println(strings.TrimRight(key, " "), val)
		fmt.Fprintf(w, reg.ReplaceAllString(key, "")+" "+strconv.Itoa(val))
		fmt.Fprintf(w,"\n")

	}

	if err != nil {
		log.Fatal(err)
	}

}


type MapCollector chan chan interface{}


type MapperFunction func(interface{}, chan interface{})


type ReducerFunction func(chan interface{}, chan interface{})



//reading the file
func myReadFile(filename string) chan string {
	output := make(chan string)
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")

	go func() {

		file, err := os.Open(filename)
		if err != nil {
			return
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		// Scan all words from the file.
		for scanner.Scan()  {
			//remove all spaces and special chars
			word := strings.TrimSpace(reg.ReplaceAllString(scanner.Text(),""))
			if len(word) > 0 {
				output <- word
			}
		}
		defer file.Close()

		close(output)


	}()
	return output
}

func getFiles(dirname string) chan interface{} {
	output := make(chan interface{})
	go func() {
		filepath.Walk(dirname, func(path string, f os.FileInfo, err error) error {
			if !f.IsDir() {
				output <- path
			}
			return nil
		})
		close(output)
	}()
	return output
}

//helper for reduce
func reducerDispatcher(collector MapCollector, reducerInput chan interface{}) {

	for output := range collector {
		reducerInput <- <-output
	}
	close(reducerInput)
}

func mapperDispatcher(mapper MapperFunction, input chan interface{}, collector MapCollector) {

	for item := range input {
		taskOutput := make(chan interface{})
		go mapper(item, taskOutput)
		collector <- taskOutput

	}
	close(collector)
}

//single file word count
func mapper(filename interface{}, output chan interface{}) {

	results := map[string]int{}


	for word := range myReadFile(filename.(string)) {


		results[strings.ToLower(word)] += 1

	}

	output <- results
}

//reduce all the words coming from different source files
func reducer(input chan interface{}, output chan interface{}) {

	results := map[string]int{}
	for matches := range input {
		for word,frequency := range matches.(map[string]int)  {
			results[strings.ToLower(word)] += frequency
		}
	}
	output <- results
}


//heavy lifting is done by this function
func mapReduce(mapper MapperFunction, reducer ReducerFunction, input chan interface{}) interface{} {

	reducerInput := make(chan interface{})
	reducerOutput := make(chan interface{})
	MapCollector := make(MapCollector)

	go reducer(reducerInput, reducerOutput)
	go reducerDispatcher(MapCollector, reducerInput)
	go mapperDispatcher(mapper, input, MapCollector)

	return <-reducerOutput
}

var allfilewordcountmap map[string]int
func main() {

	starttime := time.Now()

	input := getFiles("/files/")


	// this will do the heavy lifting of word count
	results := mapReduce(mapper, reducer, input)


	filehandle, err := os.Create("/output"+"/multiplefileswordcountoutput.csv")
	if  err != nil  {
		fmt.Println("Error writing to file: ", err)
		return
	}


	defer filehandle.Close()

	//writing to file
	writer := bufio.NewWriter(filehandle)

	allfilewordcountmap =results.(map[string]int)
	fmt.Println("word", " ","count")
	for word, count := range allfilewordcountmap {
		fmt.Println(word, " ",count)
		fmt.Fprintln(writer, word+","+strconv.Itoa(count))
	}
	writer.Flush()
	filehandle.Close()

	elapsedtime := time.Since(starttime)

	fmt.Println("Time taken:",elapsedtime)

	//calling the web-rendering function
	http.HandleFunc("/multiplefileswordcount", indexHandler2)
	http.ListenAndServe(":9081", nil)

}