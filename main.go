package main

import (
	"fmt"

	"spider/spider"
)

func main() {
	// Instantiate default collector
	c := spider.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		spider.AllowedDomains("www.gszfcg.gansu.gov.cn"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *spider.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Println("Link found: %q -> %s\n", e.Text, link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *spider.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *spider.Response, err error) {

		fmt.Println("Something went wrong:", err)

	})

	// Start scraping on https://hackerspaces.org
	c.Visit("http://www.gszfcg.gansu.gov.cn/web/doSearch.action?op=%271%27&articleSearchInfoVo.title=%25E7%25BD%2591%25E7%25BB%259C&articleSearchInfoVo.bidcode=&articleSearchInfoVo.proj_name=&articleSearchInfoVo.agentname=&articleSearchInfoVo.buyername=&articleSearchInfoVo.division=&articleSearchInfoVo.classname=12804&articleSearchInfoVo.dtype=6209&articleSearchInfoVo.releasestarttime=&articleSearchInfoVo.releaseendtime=&articleSearchInfoVo.tflag=1")
}
