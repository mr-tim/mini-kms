package api

import "../kms"
import (
	"github.com/gin-gonic/gin"
	"fmt"
	"encoding/base64"
)

type GinApi struct {
	k kms.Kms
}

var endpoint = "http://localhost:8080"

func (api GinApi) Run() {
	r := gin.Default()
	v1 := r.Group("/kms/v1")
	{
		v1.POST("/keys", api.createKey)
		v1.GET("/keys/names", api.getKeyNames)
	}

	r.Run()
}

type CreateKeyRequest struct {
	Name string `json:"name"`
	Cipher string `json:"cipher"`
	Length int `json:"cipher"`
	MaterialBase64 string `json:"material"`
	Description string `json:"description"`
}

func (api GinApi) createKey(c *gin.Context) {
	var createKey CreateKeyRequest
	err := c.BindJSON(&createKey)
	if err != nil {
		c.Abort()
		return
	}
	fmt.Printf("Received create key request: %#v\n", createKey)
	material, err := base64.StdEncoding.DecodeString(createKey.MaterialBase64)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	keyDesc := kms.KeyDesc{
		Metadata:kms.KeyMeta{
			createKey.Name,
			createKey.Cipher,
			createKey.Length,
			createKey.Description,
			-1,
			-1,
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