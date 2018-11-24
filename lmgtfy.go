package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"runtime"

	"github.com/subosito/shorturl"
	"github.com/urfave/cli"
)

func lmgtfyURL(s string) string {
	strEnc := url.QueryEscape(s)
	lmgtfyString := "http://lmgtfy.com/?q=" + strEnc
	return lmgtfyString
}

func toClipboard(output []byte, arch string) {
	var copyCmd *exec.Cmd
	// Mac "OS"
	if arch == "darwin" {
		copyCmd = exec.Command("pbcopy")
	}
	// Linux
	if arch == "linux" {
		copyCmd = exec.Command("xclip", "-selection", "c")
	}
	in, err := copyCmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := copyCmd.Start(); err != nil {
		log.Fatal(err)
	}
	if _, err := in.Write([]byte(output)); err != nil {
		log.Fatal(err)
	}
	if err := in.Close(); err != nil {
		log.Fatal(err)
	}
	copyCmd.Wait()
	fmt.Println("Url copied to clipboard")
}

func shortenURL(url string, provider string) string {
	u, err := shorturl.Shorten(url, provider)
	if err != nil {
		log.Fatal(err)
	}
	return string(u)
}

func main() {
	var search string
	var shortenProvider string

	app := cli.NewApp()
	app.Name = "lmgtfy"
	app.Usage = "The Let me google that for you CLI"
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "google, g",
			Usage:       "Google Search Text",
			Destination: &search,
		},
		cli.BoolFlag{
			Name:  "shorten, s",
			Usage: "Shorten Url with the Provider given in 'shorten-provider'",
		},
		cli.StringFlag{
			Name:        "shorten-provider",
			Value:       "tinyurl",
			Usage:       "shorten URL in order to hide the lmgtfy.com prefix. Multiple provider possible, see https://github.com/subosito/shorturl",
			Destination: &shortenProvider,
		},
	}
	app.Action = func(c *cli.Context) error {
		lmgtfyString := lmgtfyURL(search)
		if c.Bool("shorten") {
			lmgtfyString = shortenURL(lmgtfyString, shortenProvider)
		}
		fmt.Println(lmgtfyString)

		os := runtime.GOOS
		if os != "windows" {
			toClipboard([]byte(lmgtfyString), os)
		}

		return nil
	}

	app.Run(os.Args)
}
