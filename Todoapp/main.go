
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", hoge)
    r.Run() // デフォルトが8080ポートなので今回は変更しません
}

func hoge(c *gin.Context) {
    c.JSON(200, "hogeeeeeeee")
}
