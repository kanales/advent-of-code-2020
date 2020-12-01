package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path"

	"github.com/kanales/advent-of-code-2020/days"
)

type configuration struct {
	Year    int
	Session string
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func fetchInput(client *http.Client, day int) ([]byte, error) {
	filename := path.Join(".", "cache", fmt.Sprintf("input_%d.txt", day))

	if fileExists(filename) {
		return ioutil.ReadFile(filename)
	}

	file, err := os.Create(filename)
	defer file.Close()
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", Config.Year, day)
	l.Printf("Fetching input from %s...", url)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	l.Println("OK")
	defer resp.Body.Close()

	io.Copy(file, resp.Body)
	file.Seek(0, 0)
	return ioutil.ReadAll(file)
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
var l = log.New(os.Stderr, "", 0)

func main() {
	configInit("conf.json")
	flag.Parse()

	if *clearFlag {
		os.RemoveAll("cache")
	}

	if *dayFlag == 0 {
		l.Fatal("Expected day flag")
		os.Exit(1)
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

	// Run day
	content, err := fetchInput(client, *dayFlag)
	if err != nil {
		panic(err)
	}

	dayFun := days.DayMap[*dayFlag-1]
	first, second := dayFun(string(content))
	l.Print("First part:")
	fmt.Println(first)
	l.Print("Second part:")
	fmt.Println(second)
}
