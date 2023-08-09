package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

// Hacker News' article
type hnArticle struct {
	title   string
	summary string
	link    string
	rank    int
	votes   int
	time    string
}

func formatDuration(duration time.Duration) string {
	if duration < 0 {
		duration = -duration
	}

	switch {
	case duration >= time.Hour:
		hours := duration / time.Hour
		return fmt.Sprintf("%d hours ago", hours)
	case duration >= time.Minute:
		minutes := duration / time.Minute
		return fmt.Sprintf("%d minutes ago", minutes)
	default:
		return "just now"
	}
}

func getHNArticles() []hnArticle {
	articles := []hnArticle{}
	c := colly.NewCollector()

	c.OnHTML("article", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildText(".post-summary"))
		fmt.Println(e.ChildText(".post-title"))
		fmt.Println(e.ChildText(".score"))
		fmt.Println(e.ChildText(".host"))

		articleTime, _ := time.Parse("2006-01-02 15:04:05 MST", e.ChildText("span .last-updated"))
		fmt.Println(formatDuration(time.Now().Sub(articleTime)))
		fmt.Println(time.Now().Sub(articleTime))
		fmt.Println(e.Attr("data-rank"))

	})

	c.Visit("https://hackernews.betacat.io/")

	fmt.Println("Done")

	return articles
}

// func main() {
// 	getHNArticles()
// }
