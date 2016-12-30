package main

import (
	"fmt"
	"github.com/keisuke-umezawa/gosearch/crawler"
	"github.com/keisuke-umezawa/gosearch/env"
	"github.com/keisuke-umezawa/gosearch/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func main() {
	resp := make(chan crawler.CrawlResponse)
	var seed string = "http://golang.org/"
	env.Init()
	if err := models.Dialdb(env.GetDBHost()); err != nil {
		fmt.Println("Cannot connect to MongoDB")
		return
	}

	go func() {
		go crawler.Crawl(seed, 4, resp)

		for r := range resp {
			log.Printf("%d : %s", r.StatusCode, r.Url)
			untaggedBody := crawler.RemoveTags(string(r.Body))
			models.AddPageToIndex(untaggedBody, r.Url)
		}
	} ()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		keyword := c.Query("keyword")
		urls := models.Search(keyword)
		c.JSON(http.StatusOK, gin.H{
			"results": urls,
		})
	})
	router.Run(":8080")
}
