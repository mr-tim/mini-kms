package api

import "../kms"
import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type GinApi struct {
	k kms.Kms
}

var endpoint = "http://localhost:8080"

type CreateKeyRequest struct {
	Name           string `json:"name"`
	Cipher         string `json:"cipher"`
	Length         int    `json:"length"`
	MaterialBase64 string `json:"material"`
	Description    string `json:"description"`
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
			Name:        createKey.Name,
			Cipher:      createKey.Cipher,
			Length:      createKey.Length,
			Description: createKey.Description,
			Created:     time.Now().UnixNano() / int64(time.Millisecond),
			Versions:    1,
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
		"name":     newKey.VersionName,
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

	c.JSON(200, metadataToJson(metadata))
}
func metadataToJson(metadata *kms.KeyMeta) gin.H {
	return gin.H{
		"name":        metadata.Name,
		"cipher":      metadata.Cipher,
		"length":      metadata.Length,
		"description": metadata.Description,
		"created":     metadata.Created,
		"versions":    metadata.Versions,
	}
}

func (api GinApi) getKeysMetadata(c *gin.Context) {
	keyNames := c.QueryArray("key")
	error, metadata := api.k.GetKeysMetadata(keyNames)
	if error != nil {
		c.AbortWithError(500, error)
		return
	}

	metaSlice := make([]gin.H, len(metadata))
	for _, m := range metadata {
		metaSlice = append(metaSlice, metadataToJson(&m))
	}

	c.JSON(200, metaSlice)
}

type NewKeyMaterial struct {
	MaterialBase64 string `json:"material"`
}

func (api GinApi) rolloverKey(c *gin.Context) {
	keyName := c.Param("keyName")
	var m NewKeyMaterial
	c.BindJSON(&m)

	material, err := base64.StdEncoding.DecodeString(m.MaterialBase64)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	err, version := api.k.RolloverKey(keyName, material)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, serializeVersion(version))
}

func serializeVersion(version *kms.KeyVersion) gin.H {
	return gin.H{
		"name":     version.VersionName,
		"material": base64.StdEncoding.EncodeToString(version.Material),
	}
}

func (api GinApi) deleteKey(c *gin.Context) {
	keyName := c.Param("keyName")
	err := api.k.DeleteKey(keyName)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.Status(200)
}

func (api GinApi) getCurrentVersion(c *gin.Context) {
	keyName := c.Param("keyName")
	err, version := api.k.CurrentVersion(keyName)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, serializeVersion(version))
}

func (api GinApi) getAllVersions(c *gin.Context) {
	keyName := c.Param("keyName")
	err, versions := api.k.GetKeyVersions(keyName)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	result := make([]gin.H, 0)
	for _, v := range versions {
		result = append(result, serializeVersion(&v))
	}
	c.JSON(200, result)
}

func (api GinApi) getKeyVersion(c *gin.Context) {
	keyVersionName := c.Param("keyVersionName")
	err, version := api.k.GetKeyVersion(keyVersionName)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, serializeVersion(version))
}
