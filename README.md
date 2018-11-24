# Let me google that for you CLI

[![Release](https://img.shields.io/github/release/Knappek/lmgtfy.svg?style=flat-square)](https://github.com/Knappek/lmgtfy/releases/latest)
[![Travis](https://img.shields.io/travis/Knappek/lmgtfy.svg?style=flat-square)](https://travis-ci.org/Knappek/lmgtfy)
[![Go Report Card](https://goreportcard.com/badge/Knappek/lmgtfy "Go Report Card")](https://goreportcard.com/report/Knappek/lmgtfy)
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Powered By: GoReleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser)

This is a [lmgtfy](https://github.com/pykler/lmgtfy) CLI written in Go which makes it easy to create Let me google that for you links.

## Install

### MacOS

```shell
brew install Knappek/tap/lmgtfy
```

### Other OS with Go installed

Download with

```shell
go get github.com/Knappek/lmgtfy
```

and install `lmgtfy` CLI by navigating to `$GOPATH/src/github.com/Knappek/lmgtfy/` and run

```shell
go install
```

In order to use `lmgtfy` make sure `$GOPATH/bin` is in your PATH.

### Without Go installed

You can simply download the latest release from the [release](https://github.com/Knappek/lmgtfy/releases) page, extract it and move the binary to your path.


## Usage

<p align="center">
<img src="lmgtfy-recording.gif" alt="lmgtfy-usage" title="lmgtfy Usage (animated gif)" />
</p>

In a shell, simply type

```shell
lmgtfy -g "What is Let me google that for you?"
```

which will print http://lmgtfy.com/?q=What+is+Let+me+google+that+for+you%3F and, if you are on Linux or MacOS, will directly copy it to your clipboard.

### Shorten Url

The lmgtfy-Url is very obvious and you may want to hide this. Simply add a `-s` or `-shorten`, i.e.

```shell
lmgtfy -g "What is Let me google that for you?" -s
```

which will print http://tinyurl.com/3so56ko and, if you are on Linux or MacOS, will directly copy it to your clipboard.

# Next Steps

* code lint
* code coverage
