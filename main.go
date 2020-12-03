package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"

	"github.com/kanales/advent-of-code-2020/days"
)

type configuration struct {
	Year    int
	Session string
}

// Config struct
var Config configuration

func configInit(filename string) error {
	file, _ := os.Open(filename)
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config = configuration{}
	return decoder.Decode(&Config)
}

var clearFlag = flag.Bool("clean", false, "clean cache")
var dayFlag = flag.Int("day", 0, "select day")

func main() {
	configInit("conf.json")
	flag.Parse()
	log.SetOutput(os.Stderr)
	log.SetFlags(0)

	if *clearFlag {
		os.RemoveAll("cache")
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		// TODO: handle
	}
	url, err := url.Parse("https://adventofcode.com")
	jar.SetCookies(url, []*http.Cookie{{Name: "session", Value: Config.Session}})

	client := &http.Client{Jar: jar}

	// Create cache if it doesn't exist
	_ = os.Mkdir("cache", os.ModePerm)

	if *dayFlag != 0 {
		res := runDay(client, *dayFlag)
		outputResult(res)
	} else {
		runAll(client)
	}

}

func runDay(client *http.Client, day int) days.DayResult {
	// Run day
	content, err := days.FetchInput(client, Config.Year, day)
	if err != nil {
		panic(err)
	}

	dayFun := days.DayMap[day-1]
	result := dayFun(content)
	return result
}

func runAll(client *http.Client) {
	c := make(chan days.DayResult)
	defer close(c)

	for d := range days.DayMap {
		go func(d int) {
			res := runDay(client, d)
			c <- res
		}(d + 1)
	}

	counter := len(days.DayMap)
	for res := range c {
		outputResult(res)
		counter--
		if counter == 0 {
			break
		}
	}

}

func outputResult(res days.DayResult) {
	log.Printf("== Day %v ==\n", res.Day)
	log.Print("First part:")
	fmt.Println(res.First)
	log.Print("Second part:")
	fmt.Println(res.Second)
}
