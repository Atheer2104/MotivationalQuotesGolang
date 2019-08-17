package main

// simple steps
// open and read data file
// parse data
// print to command line

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"github.com/fatih/color"
)

type Quotes struct {
	Quotes []Quote
}

type Quote struct {
	QuoteText   string
	QuoteAuthor string
}

func openJSONFile(fileName string) []byte {

	// opening our json file
	jsonFile, err := os.Open(fileName)
	// if there was an error while opening our json file we print it
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("Successfully Opened users.json")

	// reading our json file as an byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

func parseJSON(byteValue []byte) Quotes {
	// initializing our quote array
	var quotes Quotes

	// we unmarshal our byteArray which contains our
	// jsonFile's content and assign it into 'quotes' which we defined above
	// unmarshal takes our json data as bytearray and we give it a pointer to our struct
	json.Unmarshal(byteValue, &quotes)

	return quotes
}

func getRandomQuoteAndPrint(quotes Quotes) {
	// randomly picking a quote

	// randomly generating seed very fast, which the rand funtion uses
	// default of seed is one 1 and when we generate Intn we will get the same output every time
	// so that's why we changed the seed so we get random output every time we run the program
	rand.Seed(time.Now().UnixNano())

	// Intn works by generating a number between 0 and N
	randomQuote := quotes.Quotes[rand.Intn(len(quotes.Quotes)-1)]

	// just printing our quotes and author using github.com/fatih/color
	// to get colors when we print them
	red := color.New(color.FgMagenta, color.Bold)
	red.Println(randomQuote.QuoteText)
	d := color.New(color.FgWhite, color.Bold)

	// for cases when QuoteAuthor is empty
	if randomQuote.QuoteAuthor != "" {
		d.Println("-", randomQuote.QuoteAuthor)
	} else {
		d.Println("- Unkown")
	}
}

func main() {

	var byteValue = openJSONFile("data.json")
	var quotes = parseJSON(byteValue)
	getRandomQuoteAndPrint(quotes)

}
