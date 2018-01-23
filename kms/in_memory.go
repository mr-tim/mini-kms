package kms

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
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

func (InMemoryKms) RolloverKey(name string, material string) (error, *KeyVersion) {
	return notImplemented(), nil
}

func (InMemoryKms) DeleteKey(name string) error {
	return notImplemented()
}

func (k InMemoryKms) GetKeyMetadata(name string) (error, *KeyMeta) {
	if meta, found := k.store.getMetadata(name); !found {
		return keyDoesNotExist(name), nil
	} else {
		return nil, meta
	}
}

func (InMemoryKms) CurrentVersion(name string) (error, *KeyVersion) {
	return notImplemented(), nil
}

func (InMemoryKms) GenerateEncryptedKeys(name string, keysToGenerate int) (error, []EncryptedKeyVersion) {
	return notImplemented(), nil
}

func (InMemoryKms) DecryptEncryptedKey(versionName string, version EncryptedKeyVersion) error {
	return notImplemented()
}

func (InMemoryKms) GetKeyVersion(versionName string) (error, *KeyVersion) {
	return notImplemented(), nil
}

func (InMemoryKms) GetKeyVersions(keyName string) (error, []string) {
	return notImplemented(), nil
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
