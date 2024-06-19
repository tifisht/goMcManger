package routes

import (
	"GMMG/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	//默认模式
	gin.SetMode(utils.AppMode)

	//创建路由
	r := gin.Default()

	r.POST("/login", func(ctx *gin.Context) {
		var user utils.Login
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if user.User == utils.Admin && user.Password == utils.Secret {
			// //正确逻辑,签署并保存到map
			// utils.Session[session] = user.User
			session := utils.JwtSign(user.User)
			ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in", "token": session})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}

	})
	v1 := r.Group("/v1")
	v1.Use(AuthMiddleware())
	{
		//服务器启动路由
		v1.POST("/start", func(ctx *gin.Context) {
			if utils.IfStart {
				ctx.JSON(http.StatusConflict, gin.H{"error": "服务器正在运行！"})
				return
			}
			go utils.StartServer()
			ctx.JSON(http.StatusOK, gin.H{"message": "服务器启动成功！"})
		})
		//历史命令、日志路由
		v1.GET("/log", func(ctx *gin.Context) {
			if !utils.IfStart {
				ctx.JSON(http.StatusConflict, gin.H{"error": "服务器已停止！"})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "获取日志成功！", "log": utils.Stdout.String()})
		})

		//接受命令并执行
		v1.POST("/control", func(ctx *gin.Context) {
			if !utils.IfStart {
				ctx.JSON(http.StatusConflict, gin.H{"error": "服务器已停止！"})
				return
			}

			var command utils.Command
			if err := ctx.Bind(&command); err != nil || command.Context == "" {
				// 处理绑定错误
				ctx.JSON(http.StatusConflict, gin.H{"error": "服务器获取命令失败"})
				println("erroris:", err)
				return
			}

			utils.Input(command.Context)

			ctx.JSON(http.StatusOK, gin.H{"message": "执行成功", "log": utils.Stdout.String()})
		})
	}

	r.Run(utils.HttpPort)
}
