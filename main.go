package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/atotto/clipboard"
	"github.com/peterhellberg/giphy"
)

const LGTM_URL = "http://lgtm.herokuapp.com/"

func openCommand() string {
	return "gnome-open"
}

func lgtmMarkdown(url string) string {
	return "![LGTM](" + url + ")"
}

func main() {
	var tag = flag.String("tag", "cat", "Search query term or phrase.")
	flag.Parse()

	client := giphy.DefaultClient
	random, err := client.Random([]string{*tag})

	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	lgtmImageUrl := LGTM_URL + random.Data.ImageURL
	lgtmMarkdownText := lgtmMarkdown(lgtmImageUrl)

	exec.Command(openCommand(), lgtmImageUrl).Start()
	fmt.Println(lgtmMarkdownText)
	clipboard.WriteAll(lgtmMarkdownText)
	os.Exit(0)
}
