package main

import (
	"gopkg.in/headzoo/surf.v1"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"github.com/headzoo/surf/browser"
)

func main() {
	bow := surf.NewBrowser()
	err := bow.Open("https://timereport.eworkgroup.com/time/loginLogout.jsp")
	if err != nil {
		panic(err)
	}
	Login("", "", bow)
}

func Login(username string, password string, bow *browser.Browser) {
	forms := bow.Forms()
	sub := forms[1]
	sub.Input("USERNAME", username)
	sub.Input("PASSWORD", password)
	sub.Submit()
	bow.Find("table.headerTable tbody").Each(
		func(_ int, s *goquery.Selection) {
			fmt.Println(strings.TrimSpace(s.Text()))
		})
}