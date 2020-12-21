package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	rand.Seed(1337)

	file, err := os.Open("/opt/crawlerg/top10milliondomains.txt")
	if err != nil {
		log.Fatalln("Couldn't open the txt file", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		url := fmt.Sprintf("https://%s", scanner.Text())
		response, err := http.Get(url)
		if err != nil {
			log.Printf("Acessing %s FAILED", url)
			continue
		}
		defer response.Body.Close()

		log.Printf("HTTP GET: %s, CODE: %s", url, response.Status)

		time.Sleep(time.Duration(500+rand.Intn(2000)) * time.Millisecond)
	}
}
