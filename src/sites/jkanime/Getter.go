package jkanime

import (
	"GetAnime/src/structs"
	"fmt"

	"github.com/gocolly/colly"
)

func AnimeScraper() {
	serverList := []structs.Server{}
	collector := colly.NewCollector(colly.AllowedDomains("jkanime.net", "www80.zippyshare.com"))
	collector.OnHTML("div.simplemodal-wrap", func(element *colly.HTMLElement) {

		servers := element.DOM

		server := structs.Server{
			URL:        servers.Find("div.servers").Text(),
			ServerName: "name",
		}
		fmt.Println(server)
		serverList = append(serverList,server)
	})
	collector.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting: ", r.URL)
	})
	collector.Visit("https://jkanime.net/boruto-naruto-next-generations/234/")
}

//https://jkanime.net/jkfembed.php?u=mdg5qu5w8emkrq4
