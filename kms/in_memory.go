package kms

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"strings"
)

func notImplemented() error {
	return errors.New("kms: Not implemented")
}

func keyDoesNotExist(name string) error {
	return errors.New(fmt.Sprintf("kms: Key '%s' does not exist", name))
}

func keyAlreadyExists(name string) error {
	return errors.New(fmt.Sprintf("kms: Key '%s' already exists", name))
}

func keyVersionDoesNotExist(versionName string) error {
	return errors.New(fmt.Sprintf("kms: Key version '%s' does not exist", versionName))
}

type InMemoryKms struct {
	store KeyStore
}

func (k InMemoryKms) CreateKey(desc KeyDesc) (error, *KeyVersion) {
	material := desc.Material

	if len(material) == 0 {
		material = k.CreateMaterial(desc.Metadata.Cipher, desc.Metadata.Length)
	}

	return k.store.createKey(desc, material)
}

func (k InMemoryKms) RolloverKey(name string, material []byte) (error, *KeyVersion) {
	metadata, found := k.store.getMetadata(name)
	if !found {
		return keyDoesNotExist(name), nil
	}
	if len(material) == 0 {
		material = k.CreateMaterial(metadata.Cipher, metadata.Length)
	}
	return k.store.createNewKeyVersion(name, material)
}

func (k InMemoryKms) DeleteKey(name string) error {
	return k.store.deleteKey(name)
}

func (k InMemoryKms) GetKeyMetadata(name string) (error, *KeyMeta) {
	if meta, found := k.store.getMetadata(name); !found {
		return keyDoesNotExist(name), nil
	} else {
		return nil, meta
	}
}

func (k InMemoryKms) CurrentVersion(name string) (error, *KeyVersion) {
	error, versions := k.store.getKeyVersions(name)
	if error != nil {
		return error, nil
	}
	return nil, &versions[len(versions)-1]
}

func (InMemoryKms) GenerateEncryptedKeys(name string, keysToGenerate int) (error, []EncryptedKeyVersion) {
	return notImplemented(), nil
}

func (InMemoryKms) DecryptEncryptedKey(versionName string, version EncryptedKeyVersion) error {
	return notImplemented()
}

func (k InMemoryKms) GetKeyVersion(versionName string) (error, *KeyVersion) {
	keyName := strings.Split(versionName, "/")[0]
	error, versions := k.store.getKeyVersions(keyName)
	if error != nil {
		return error, nil
	}
	for _, v := range versions {
		if v.VersionName == versionName {
			return nil, &v
		}
	}
	return keyVersionDoesNotExist(versionName), nil
}

func (k InMemoryKms) GetKeyVersions(keyName string) (error, []KeyVersion) {
	return k.store.getKeyVersions(keyName)
}

func (k InMemoryKms) GetKeyNames() (error, []string) {
	return k.store.getKeyNames()
}

func (k InMemoryKms) GetKeysMetadata(names []string) (error, []KeyMeta) {
	result := make([]KeyMeta, 0)
	for _, name := range names {
		meta, found := k.store.getMetadata(name)
		if found {
			result = append(result, *meta)
		}
	}
	return nil, result
}

func (k InMemoryKms) CreateMaterial(cipher string, length int) []byte {
	key := make([]byte, length/8)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		panic(err)
	}
	return key
}
