package api

import "../kms"
import (
	"github.com/gin-gonic/gin"
	"fmt"
	"encoding/base64"
	"time"
)

type GinApi struct {
	k kms.Kms
}

var endpoint = "http://localhost:8080"

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

type CreateKeyRequest struct {
	Name string `json:"name"`
	Cipher string `json:"cipher"`
	Length int `json:"length"`
	MaterialBase64 string `json:"material"`
	Description string `json:"description"`
}

func (api GinApi) createKey(c *gin.Context) {
	var createKey CreateKeyRequest
	err := c.BindJSON(&createKey)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	fmt.Printf("Received create key request: %#v\n", createKey)
	material, err := base64.StdEncoding.DecodeString(createKey.MaterialBase64)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	keyDesc := kms.KeyDesc{
		Metadata: kms.KeyMeta{
			Name: createKey.Name,
			Cipher: createKey.Cipher,
			Length: createKey.Length,
			Description: createKey.Description,
			Created: time.Now().UnixNano() / int64(time.Millisecond),
			Versions: 1,
		},
		Material: material,
	}
	err, newKey := api.k.CreateKey(keyDesc)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.Header("Location", fmt.Sprintf("%s/kms/v1/key/%s", endpoint, createKey.Name))
	c.JSON(201, gin.H{
		"name": newKey.VersionName,
		"material": base64.StdEncoding.EncodeToString(newKey.Material),
	})
}

func (api GinApi) getKeyNames(c *gin.Context) {
	err, names := api.k.GetKeyNames()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, names)
}

func (api GinApi) getKeyMetadata(c *gin.Context) {
	keyName := c.Param("keyName")
	err, metadata := api.k.GetKeyMetadata(keyName)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, gin.H{
		"name": metadata.Name,
		"cipher": metadata.Cipher,
		"length": metadata.Length,
		"description": metadata.Description,
		"created": metadata.Created,
		"versions": metadata.Versions,
		})
}