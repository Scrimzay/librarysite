package main

import (
	"fmt"
	"github.com/Scrimzay/librarysite/libraryapi"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)
	
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("**/*.html")
    r.Static("/static", "./static")

	r.GET("/", indexHandler)
	r.GET("/general", generalSearchIndexHandler)
	r.POST("/general/search", generalSearchHandler)
	r.GET("/general/:query/info", generalSearchQueryHandler)
	r.GET("/author", authorSearchIndexHandler)
	r.POST("/author/search", authorSearchHandler)
	r.GET("/author/:query/info", authorSearchQueryHandler)
	r.GET("/book", bookSearchIndexHandler)
	r.POST("/book/search", bookSearchHandler)
	r.GET("/book/:query/info", bookSearchQueryHandler)

	err := r.Run(":4000")
	if err != nil {
		log.Fatal(err)
	}
}
	
func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func generalSearchIndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "generalindex.html", nil)
}

func authorSearchIndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "authorindex.html", nil)
}

func bookSearchIndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "bookindex.html", nil)
}

func generalSearchHandler(c *gin.Context) {
	query := c.PostForm("query")
	if query == "" {
        log.Printf("Empty query received")
        c.HTML(http.StatusBadRequest, "generalindex.html", gin.H{
            "Error": "Query cannot be empty",
        })
        return
    }

    _, err := libraryapi.CallGeneralOpenLibraryInformation(query)
    if err != nil {
        log.Printf("Error validating search query %s: %v", query, err)
        c.HTML(http.StatusNotFound, "generalindex.html", gin.H{
            "Error": "No results found for query: " + query,
        })
        return
    }

    redirectURL := "/general/" + url.PathEscape(query) + "/info"
    c.Redirect(http.StatusSeeOther, redirectURL)
}

func generalSearchQueryHandler(c *gin.Context) {
	query := c.Param("query")
	if query == "" {
        log.Printf("Empty query received")
        c.HTML(http.StatusBadRequest, "generalindex.html", gin.H{
            "Error": "Query cannot be empty",
        })
        return
    }

	search, err := libraryapi.CallGeneralOpenLibraryInformation(query)
	if err != nil {
		log.Printf("Error fetching results for query %s: %v", query, err)
		c.HTML(http.StatusOK, "generalindex.html", gin.H{
            "Error": fmt.Sprintf("No results found for query: %s", query),
        })
		return
	}

	c.HTML(http.StatusOK, "generalindexquery.html", gin.H{
		"Search": search,
	})
}

func authorSearchHandler(c *gin.Context) {
	query := c.PostForm("query")
	if query == "" {
        log.Printf("Empty query received")
        c.HTML(http.StatusBadRequest, "authorindex.html", gin.H{
            "Error": "Query cannot be empty",
        })
        return
    }

    _, err := libraryapi.CallAuthorOpenLibraryInformation(query)
    if err != nil {
        log.Printf("Error validating search query %s: %v", query, err)
        c.HTML(http.StatusNotFound, "authorindex.html", gin.H{
            "Error": "No results found for query: " + query,
        })
        return
    }

    redirectURL := "/author/" + url.PathEscape(query) + "/info"
    c.Redirect(http.StatusSeeOther, redirectURL)
}

func authorSearchQueryHandler(c *gin.Context) {
	query := c.Param("query")
	if query == "" {
        log.Printf("Empty query received")
        c.HTML(http.StatusBadRequest, "authorindex.html", gin.H{
            "Error": "Query cannot be empty",
        })
        return
    }

	search, err := libraryapi.CallAuthorOpenLibraryInformation(query)
	if err != nil {
		log.Printf("Error fetching results for query %s: %v", query, err)
		c.HTML(http.StatusOK, "authorindex.html", gin.H{
            "Error": fmt.Sprintf("No results found for query: %s", query),
        })
		return
	}

	c.HTML(http.StatusOK, "authorindexquery.html", gin.H{
		"Search": search,
	})
}

func bookSearchHandler(c *gin.Context) {
	query := c.PostForm("query")
	if query == "" {
        log.Printf("Empty query received")
        c.HTML(http.StatusBadRequest, "bookindex.html", gin.H{
            "Error": "Query cannot be empty",
        })
        return
    }

    _, err := libraryapi.CallBookOpenLibraryInformation(query)
    if err != nil {
        log.Printf("Error validating search query %s: %v", query, err)
        c.HTML(http.StatusNotFound, "bookindex.html", gin.H{
            "Error": "No results found for query: " + query,
        })
        return
    }

    redirectURL := "/book/" + url.PathEscape(query) + "/info"
    c.Redirect(http.StatusSeeOther, redirectURL)
}

func bookSearchQueryHandler(c *gin.Context) {
	query := c.Param("query")
	if query == "" {
        log.Printf("Empty query received")
        c.HTML(http.StatusBadRequest, "bookindex.html", gin.H{
            "Error": "Query cannot be empty",
        })
        return
    }

	search, err := libraryapi.CallBookOpenLibraryInformation(query)
	if err != nil {
		log.Printf("Error fetching results for query %s: %v", query, err)
		c.HTML(http.StatusOK, "bookindex.html", gin.H{
            "Error": fmt.Sprintf("No results found for query: %s", query),
        })
		return
	}

	c.HTML(http.StatusOK, "bookindexquery.html", gin.H{
		"Search": search,
	})
}