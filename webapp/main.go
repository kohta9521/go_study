package main
import (
  "log"
  "github.com/gin-gonic/gin"
)
func main() {
  // ルーターを準備 --- (*1)
  router := gin.Default()
  // URIとハンドラを指定 --- (*2)
  router.GET("/", func(ctx *gin.Context) {
    ctx.String(200, "Hello, World!")
  })
  // サーバーを起動 --- (*3)
  err := router.Run("127.0.0.1:8888")
  if err != nil {
    log.Fatal("サーバー起動に失敗", err)
  }
}





// package main

// import (
// 		"log"
// 		"github.com/gin-gonic/gin"
// )

// func main() {
// 	//ルーターを準備
// 	router := gin.Default()

// 	//URIとハンドラを指定
// 	router.GET("/", func(ctx *gin.Context) {
// 		ctx.String(200, "Hello, World!")
// 	})

// 	//サーバーを起動
// 	err := router.Run("127.0.0.1:8888")
// 	if err ! = nil {
// 		log.Fatal("サーバー起動に失敗", err)
// 	}
// }
