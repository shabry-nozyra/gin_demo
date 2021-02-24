package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shabry-nozyra/gin_demo/handlers"
	"net/http"
)

func main() {
	g := gin.Default()

	c := cors.DefaultConfig()
	c.AllowWildcard = true
	c.AllowOrigins = []string{"*"}
	c.AddAllowHeaders("Authorization", "Content-Type")
	c.AddExposeHeaders("Authorization", "Content-Type")
	g.Use(cors.New(c))

	h := handlers.Context{Gin: g}
	h.Register("")

	s := &http.Server{Addr: "0.0.0.0:8080", Handler: g}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
