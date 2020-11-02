package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "goleaning/tool"
)

func main() {
    fmt.Printf("hello \n")

    //1,解析配置文件

    cfg, err := tool.ParseConfig("./config/apps.json")

    if err != nil {
        panic(err)
    }

    app := gin.Default()

    //用户模块注册路由
    Router(app)
    fmt.Println(cfg)
    app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

func Router(router *gin.Engine) {
    router.GET("/api/hello", helloFmt)
}

func helloFmt(content *gin.Context) {
    content.JSON(200, map[string]interface{}{
        "code":    0,
        "message": "处理成功",
        "data":    "",
    })
}
