package main

import "testing"

func TestLmgtfyURL(t *testing.T) {

	var testString string
	testString = "What is let me google that for you"
	var url string
	url = lmgtfyURL(testString)
	if url != "http://lmgtfy.com/?q=What+is+let+me+google+that+for+you" {
		t.Error("URL should be http://lmgtfy.com/?q=What+is+let+me+google+that+for+you, but it is ", url)
	}

	var tinyurl string
	tinyurl = shortenURL(url, "tinyurl")
	if tinyurl != "http://tinyurl.com/yda5cej" {
		t.Error("Shortened URL should be http://tinyurl.com/yda5cej, but it is ", tinyurl)
	}

}
