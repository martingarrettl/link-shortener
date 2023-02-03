package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/martingarrettl/link-shortener/internal/urlhash"
)

func ShortURL(c *gin.Context) {
	sc := c.Param("shortcode")

	c.JSON(http.StatusOK, gin.H{"shortcode": sc})
}

func NewURL(c *gin.Context) {
	url := c.PostForm("url")
	if url == "" {
		c.String(http.StatusBadRequest, "URL is required")
		return
	}

	shortcode := urlhash.GenerateShortCode(url)

	c.JSON(http.StatusOK, gin.H{"shortcode": shortcode})
}
