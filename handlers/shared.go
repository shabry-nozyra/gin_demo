package handlers

import "github.com/gin-gonic/gin"

type Context struct {
	Gin *gin.Engine
}

func (c *Context) Register(group string) {
	public := c.Gin.Group(group)
	{
		public.GET("/paslon/add", c.add)
		public.GET("/tps/gettps", c.get)
		public.POST("/add3", c.add3)
		public.PUT("/update", c.edit)
		public.DELETE("/delete", c.delete)
	}
}
