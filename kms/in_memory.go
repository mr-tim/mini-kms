package kms

import (
	"errors"
	"fmt"
)

type InMemoryKms struct {
	metadata map[string]KeyMeta
	material map[string][]byte
}

func notImplemented() error {
	return errors.New("kms: Not implemented")
}

func (k InMemoryKms) CreateKey(desc KeyDesc) (error, *KeyVersion) {
	versionName := fmt.Sprintf("%s/%d", desc.Metadata.Name, 0)
	material := desc.Material

	k.metadata[desc.Metadata.Name] = desc.Metadata
	k.material[versionName] = material

	newKey := KeyVersion{
		versionName,
		material,
	}
	return nil, &newKey
}

func (InMemoryKms) RolloverKey(name string, material string) (error, *KeyVersion) {
	return notImplemented(), nil
}

func (InMemoryKms) DeleteKey(name string) error {
	return notImplemented()
}

func (k InMemoryKms) GetKeyMetadata(name string) (error, *KeyMeta) {
	meta, found := k.metadata[name]
	if found {
		return nil, &meta
	} else {
		return errors.New("kms: key does not exist"), nil
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
	keyNames := make([]string, 0, len(k.metadata))
	for k := range k.metadata {
		keyNames = append(keyNames, k)
	}
	return nil, keyNames
}

func (k InMemoryKms) GetKeysMetadata(names []string) (error, []KeyMeta) {
	result := make([]KeyMeta, 0)
	for _, name := range names {
		meta, found := k.metadata[name]
		if found {
			result = append(result, meta)
		}
	}
	return nil, result
}