// model
package gszfcg

import (
	"fmt"
	"math"
	"net/url"
	"spider/spider"
	"strconv"
	"strings"
)

type Params struct {
	Title       string
	City        string
	BiddingSort string
}

type Info struct {
	Title      string
	Url        string
	Start      string
	Publish    string
	Buyer      string
	Agent      string
	Sort       string
	Profession string
}

func Crawls(p *Params) (map[string]interface{}, error) {
	var page = 0
	var err error
	np := 1
	cookie := ""
	results := make(map[string]interface{})
	results["data"] = make([]Info, 0)
	data := make([]Info, 0)

	// 初始化
	c := spider.NewCollector(
		// 设置访问域名
		spider.AllowedDomains("www.gszfcg.gansu.gov.cn"),
	)

	if page == 0 {
		c.OnHTML("span#pe100_page_通用信息列表_普通式", func(e *spider.HTMLElement) {
			str := e.Text
			str = strings.Replace(str, " ", "", -1)
			str = strings.Replace(str, "\t", "", -1)
			str = strings.Replace(str, "\n", "", -1)
			start := strings.Index(str, "总") + 3
			end := strings.Index(str, "篇首页")
			item := SubStr(str, start, end)
			itemn, _ := strconv.Atoi(item)
			results["total"] = itemn
			page = int(math.Ceil(float64(itemn) / float64(20)))

			if page > 1 {
				c.OnResponse(func(r *spider.Response) {
					siteCookies := c.Cookies(r.Request.URL.String())
					cookie = siteCookies[0].String()
				})

				if cookie != "" {
					c.OnRequest(func(r *spider.Request) {
						r.Headers.Set("cookie", cookie)
					})
				}
				if np < page {
					np++
					url := "http://www.gszfcg.gansu.gov.cn/web/doSearch.action?limit=20&start=" + strconv.Itoa((np-1)*20)
					c.Visit(url)
				}

			}
		})
	}

	// On every a element which has href attribute call callback
	c.OnHTML("ul.Expand_SearchSLisi li", func(e *spider.HTMLElement) {
		//数据清理1
		info := washdata(e.ChildText("p:nth-child(2)"), e.ChildText("p:nth-child(3)"))
		info.Title = e.ChildText("a")
		info.Url = "http://www.gszfcg.gansu.gov.cn" + e.ChildAttr("a", "href")

		data = append(data, info)

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *spider.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *spider.Response, errs error) {
		err = errs
	})

	// Start scraping on https://hackerspaces.org
	startUrl := "http://www.gszfcg.gansu.gov.cn/web/doSearch.action?op=%271%27&articleSearchInfoVo.title="
	startUrl = startUrl + url.QueryEscape(p.Title)
	if p.City != "0" {
		startUrl = startUrl + "&articleSearchInfoVo.dtype=" + url.QueryEscape(p.City)
	}
	startUrl = startUrl + "&articleSearchInfoVo.bidcode=&articleSearchInfoVo.proj_name=&articleSearchInfoVo.agentname=&articleSearchInfoVo.buyername=&articleSearchInfoVo.division=&articleSearchInfoVo.classname=" + url.QueryEscape(p.BiddingSort) + "&articleSearchInfoVo.releasestarttime=&articleSearchInfoVo.releaseendtime=&articleSearchInfoVo.tflag=1"
	c.Visit(startUrl)

	results["data"] = data

	return results, err
}

func washdata(s1 string, s2 string) Info {

	//空格处理
	s1 = strings.Replace(s1, " ", "", -1)
	strs := strings.Split(s1, "|")

	var info Info

	//赋值
	info.Start = strs[0]
	info.Publish = strs[1]
	info.Buyer = strs[2]
	info.Agent = strs[3]

	//数据显示校准
	info.Start = strings.Replace(info.Start, "开标时间：", "", -1)
	info.Start = info.Start[0:10] + " " + info.Start[10:]

	info.Publish = strings.Replace(info.Publish, "发布时间：", "", -1)
	info.Publish = info.Publish[0:10] + " " + info.Publish[10:]

	info.Buyer = strings.Replace(info.Buyer, "采购人：", "", -1)

	info.Agent = strings.Replace(info.Agent, "代理机构：", "", -1)

	//数据清理2
	s2 = strings.Replace(s2, " ", "", -1)
	s2 = strings.Replace(s2, "\t", "", -1)
	s2 = strings.Replace(s2, "\n", "", -1)

	strs1 := strings.Split(s2, "|")

	info.Sort = strs1[0]
	info.Profession = strs1[2]

	return info

}

func SubStr(s string, start, end int) string {
	rs := []byte(s)
	rl := len(rs)
	if start < 0 {
		start = 0
	}
	if start > end {
		start, end = end, start
	}
	if end > rl {
		end = rl
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	return string(rs[start:end])
}
