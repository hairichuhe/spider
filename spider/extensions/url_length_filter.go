package extensions

import (
	"spider/spider"
)

// URLLengthFilter filters out requests with URLs longer than URLLengthLimit
func URLLengthFilter(c *spider.Collector, URLLengthLimit int) {
	c.OnRequest(func(r *spider.Request) {
		if len(r.URL.String()) > URLLengthLimit {
			r.Abort()
		}
	})
}
