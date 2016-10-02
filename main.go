package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/atotto/clipboard"
	"github.com/peterhellberg/giphy"
)

const LGTM_URL = "http://lgtm.herokuapp.com/"
const MAX_RETRY_COUNT = 3

func openCommand() string {
	return "gnome-open"
}

func lgtmMarkdown(url string) string {
	return "![LGTM](" + url + ")"
}

func main() {
	var lgtmImageUrl string
	var lgtmFile string

	var tag = flag.String("tag", "cat", "Search query term or phrase.")
	flag.Parse()

	dir, err := ioutil.TempDir("", "lgtm")
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	for i := 0; i < MAX_RETRY_COUNT; i++ {
		client := giphy.DefaultClient
		random, err := client.Random([]string{*tag})

		if err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}

		lgtmImageUrl = LGTM_URL + random.Data.ImageURL
		response, err := http.Get(lgtmImageUrl)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if response.StatusCode != 200 {
			continue
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		lgtmFile = dir + "/" + random.Data.ID + ".gif"
		file, err := os.OpenFile(lgtmFile, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
		}

		defer func() {
			file.Close()
		}()

		file.Write(body)
	}

	if len(lgtmFile) == 0 {
		fmt.Printf("File generation fails. Please run the command again.\n")
	}

	exec.Command(openCommand(), lgtmFile).Start()

	lgtmMarkdownText := lgtmMarkdown(lgtmImageUrl)
	fmt.Println(lgtmMarkdownText)
	clipboard.WriteAll(lgtmMarkdownText)
	os.Exit(0)
}
