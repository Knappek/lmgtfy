package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"runtime"

	"github.com/urfave/cli"
    "github.com/subosito/shorturl"
)

func lmgtfy_url(s string) string {
	str_enc := url.QueryEscape(s)
	lmgtfy_string := "http://lmgtfy.com/?q=" + str_enc
	return lmgtfy_string
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

func shortenUrl(url string, provider string) string {
  u, err := shorturl.Shorten(url, provider)
  if err != nil { 
    log.Fatal(err)
  }
  return string(u)
}

func main() {
    var search string
    var shorten_provider string

	app := cli.NewApp()
	app.Name = "lmgtfy"
	app.Usage = "The Let me google that for you CLI"
	app.Version = "0.1.0"

    app.Flags = []cli.Flag {
      cli.StringFlag{
        Name: "google, g",
        Usage: "Google Search Text",
        Destination: &search,
      },
      cli.BoolFlag{
        Name: "shorten, s",
        Usage: "Shorten Url with the Provider given in 'shorten-provider'",
      },
      cli.StringFlag{
        Name: "shorten-provider",
        Value: "tinyurl",
        Usage: "shorten URL in order to hide the lmgtfy.com prefix - defaults to 'tinyurl'. Multiple provider possible, see https://github.com/subosito/shorturl",
        Destination: &shorten_provider,
      },
    }
	app.Action = func(c *cli.Context) error {
      lmgtfy_string := lmgtfy_url(search)
      if c.Bool("shorten"){
        lmgtfy_string = shortenUrl(lmgtfy_string, shorten_provider)
      }
      fmt.Println(lmgtfy_string)

      os := runtime.GOOS
      if (os != "windows" ){
        toClipboard([]byte(lmgtfy_string), os)
      }

      return nil
	}

	app.Run(os.Args)
}
