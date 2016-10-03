package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/peterhellberg/giphy"
)

func openCommand() string {
	command := ""
	os := runtime.GOOS

	if os == "linux" {
		command = "gnome-open"
	} else if os == "darwin" {
		command = "open"
	}

	return command
}

func lgtmMarkdown(url string) string {
	return "![LGTM](" + url + ")"
}

var version = "0.0.1"

func main() {
	var lgtmURL = "http://lgtm.herokuapp.com/"
	var maxRetryCount = 3
	var maxContentLength = 2097152

	var showVersion bool
	var lgtmImageURL string
	var random giphy.Random
	var gif giphy.GIF
	var err error

	var tag = flag.String("tag", "cat", "Search query term or phrase.")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()

	if showVersion {
		fmt.Println("version:", version)
		os.Exit(0)
		return
	}

	for i := 0; i < maxRetryCount; i++ {
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
		if fileSize < maxContentLength {
			lgtmImageURL = lgtmURL + random.Data.ImageURL
			break
		}
	}

	if len(lgtmImageURL) == 0 {
		fmt.Printf("File generation fails. Please run the command again.\n")
		os.Exit(1)
	}

	openCommand := openCommand()
	if len(openCommand) != 0 {
		exec.Command(openCommand, lgtmImageURL).Start()
	}

	lgtmMarkdownText := lgtmMarkdown(lgtmImageURL)
	fmt.Println(lgtmMarkdownText)
	clipboard.WriteAll(lgtmMarkdownText)
	os.Exit(0)
}
