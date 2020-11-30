package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path"
)

type Configuration struct {
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
		println("File exits")
		return ioutil.ReadFile(filename)
	}

	file, err := os.Create(filename)
	defer file.Close()
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", Config.Year, day)
	fmt.Println(url)
	resp, err := client.Get(url)
	if err != nil {
		println(1)
		return nil, err
	}
	defer resp.Body.Close()
	if err != nil {
		println(2)
	}

	io.Copy(file, resp.Body)
	file.Seek(0, 0)
	return ioutil.ReadAll(file)
}

// Config struct
var Config Configuration

func init() {
	file, _ := os.Open("conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config = Configuration{}
	_ = decoder.Decode(&Config)
}

var clearFlag = flag.Bool("clean", false, "clean cache")
var dayFlag = flag.Int("day", 0, "select day")

func main() {
	flag.Parse()

	if *clearFlag {
		os.RemoveAll("cache")
		return
	}

	fmt.Printf("%+v\n", Config)
	jar, err := cookiejar.New(nil)
	if err != nil {
		// TODO: handle
	}
	url, err := url.Parse("https://adventofcode.com")
	jar.SetCookies(url, []*http.Cookie{{Name: "session", Value: Config.Session}})

	client := &http.Client{Jar: jar}

	// Create cache if it doesn't exist
	_ = os.Mkdir("cache", os.ModePerm)
	content, err := fetchInput(client, *dayFlag)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(content)
}
