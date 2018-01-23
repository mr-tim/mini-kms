package api

import "github.com/gin-gonic/gin"

func (api GinApi) Run() {
	r := gin.Default()
	keys := r.Group("/kms/v1/keys")
	{
		keys.POST("", api.createKey)
		keys.GET("/names", api.getKeyNames)
	}

	key := r.Group("kms/v1/key/:keyName")
	{
		key.GET("/_metadata", api.getKeyMetadata)
	}

	r.Run()
}
