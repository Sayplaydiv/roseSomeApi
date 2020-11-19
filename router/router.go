package router

import (
	"github.com/gin-gonic/gin"
	conf "roseSomeApi/config"
	"roseSomeApi/roseApi"
)

func InitRouter(PATH string)  {
	//导入配置文件
	configMap := conf.InitConfig(PATH)
	//路径设置
	http_port:=configMap["http_port"]
	ginMode   :=configMap["ginMode"]
	gin.SetMode(ginMode)
	var router =gin.Default()
	v1 :=router.Group("/api_v2")
	{
		v1.GET("/getStatus", roseApi.GetStatus)
		v1.GET("/getSignerNonce", roseApi.GetSignerNonce)
		v1.GET("/submitTxNoWait",roseApi.SubmitTxNoWait)
	}
	router.Run(http_port)
}

