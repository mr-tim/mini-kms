package api

import "github.com/gin-gonic/gin"

func (api GinApi) Run() {
	r := gin.Default()
	keys := r.Group("/kms/v1/keys")
	{
		keys.POST("", api.createKey)
		keys.GET("/names", api.getKeyNames)
		keys.GET("/metadata", api.getKeysMetadata)
	}

	key := r.Group("/kms/v1/key/:keyName")
	{
		key.POST("", api.rolloverKey)
		key.DELETE("", api.deleteKey)
		key.GET("/_metadata", api.getKeyMetadata)
		key.GET("/_currentVersion", api.getCurrentVersion)
		key.GET("/_versions", api.getAllVersions)
	}

	r.GET("/kms/v1/keyversion/:keyVersionName", api.getKeyVersion)

	r.Run()
}
