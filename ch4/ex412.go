package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const URL = "https://xkcd.com/%d/info.0.json"
const IndexFile = "ex412.json"

type Xkcd struct {
	Title string `json:"safe_title"`
	Img   string
}

func CreateIndexFile() error {
	_, err := os.Stat(IndexFile)
	if err == nil {
		return err
	}

	items := map[string]Xkcd{}
	for i := 1; i <= 10; i++ {
		resp, err := http.Get(fmt.Sprintf(URL, i))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var item Xkcd
		if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
			return err
		}
		items[strconv.Itoa(i)] = item
	}

	b, err := json.Marshal(items)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ioutil.WriteFile(IndexFile, b, os.ModePerm)
	fmt.Println(err)
	return err
}

func LoadIndex() (map[string]Xkcd, error) {
	err := CreateIndexFile()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(IndexFile)
	if err != nil {
		return nil, err
	}

	var result map[string]Xkcd
	if err := json.NewDecoder(file).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func main() {
	items, err := LoadIndex()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	flag.Parse()
	i := flag.Arg(0)

	item, ok := items[i]
	if !ok {
		fmt.Println("not found")
		return
	}

	fmt.Printf("title:%s\nimg:%s\n", item.Title, item.Img)
}
