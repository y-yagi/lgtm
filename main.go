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
		command = "xdg-open"
	} else if os == "darwin" {
		command = "open"
	}

	return command
}

func lgtmMarkdown(url string) string {
	return "![LGTM](" + url + ")"
}

const version = "0.0.1"

func main() {
	var lgtmURL = "http://lgtm.herokuapp.com/"
	var maxRetryCount = 5
	var maxContentLength = 2097152

	var showVersion bool
	var tag string
	var lgtmImageURL string
	var random giphy.Random
	var gif giphy.GIF
	var err error

	flag.StringVar(&tag, "tag", "cat", "Search query term or phrase.")
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
		random, err = client.Random([]string{tag})
		if err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}

		gif, err = client.GIF(random.Data.ID)
		if err != nil {
			fmt.Printf("error: %s\n", err)
			os.Exit(1)
		}

		fileSize, _ := strconv.Atoi(gif.Data.Images.Downsized.Size)
		if fileSize < maxContentLength {
			fmt.Printf("%+v\n", gif.Data.Images)
			lgtmImageURL = lgtmURL + gif.Data.Images.Downsized.URL
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
