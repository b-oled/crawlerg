package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gocolly/colly"
)

var domainFile string
var c *colly.Collector

func init() {
	flag.StringVar(&domainFile, "list", "/opt/crawlerg/top1milliondomains.txt", "File containing domains (one per line)")

	c = colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 1, RandomDelay: 2 * time.Second})

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
}

func main() {
	flag.Parse()
	rand.Seed(1337)

	defer c.Wait()

	file, err := os.Open(domainFile)
	if err != nil {
		log.Fatalln("Couldn't open the txt file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		url := fmt.Sprintf("https://%s", scanner.Text())
		c.Visit(url)
		if i%2 == 0 {
			c.Wait()
		}
		i++
	}
}
