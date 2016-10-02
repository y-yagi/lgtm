package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/peterhellberg/giphy"
)

const LGTM_URL = "http://lgtm.herokuapp.com/"
const MAX_RETRY_COUNT = 3
const MAX_CONTENT_LENGTH = 2097152

func openCommand() string {
	return "gnome-open"
}

func lgtmMarkdown(url string) string {
	return "![LGTM](" + url + ")"
}

func main() {
	var lgtmImageUrl string
	var random giphy.Random
	var gif giphy.GIF
	var err error

	var tag = flag.String("tag", "cat", "Search query term or phrase.")
	flag.Parse()

	for i := 0; i < MAX_RETRY_COUNT; i++ {
		client := giphy.DefaultClient
		random, err = client.Random([]string{*tag})
		if err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}

		gif, err = client.GIF(random.Data.ID)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}

		fileSize, _ := strconv.Atoi(gif.Data.Images.Original.Size)
		if fileSize < MAX_CONTENT_LENGTH {
			lgtmImageUrl = LGTM_URL + random.Data.ImageURL
			break
		}
	}

	if len(lgtmImageUrl) == 0 {
		fmt.Printf("File generation fails. Please run the command again.\n")
		os.Exit(1)
	}

	exec.Command(openCommand(), lgtmImageUrl).Start()

	lgtmMarkdownText := lgtmMarkdown(lgtmImageUrl)
	fmt.Println(lgtmMarkdownText)
	clipboard.WriteAll(lgtmMarkdownText)
	os.Exit(0)
}
