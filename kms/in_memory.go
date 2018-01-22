package kms

import "errors"

type InMemoryKms struct {

}

func notImplemented() error {
	return errors.New("kms: Not implemented")
}

func (InMemoryKms) CreateKey(desc KeyDesc) (error, *KeyVersion) {
	return notImplemented(), nil
}

func (InMemoryKms) RolloverKey(name string, material string) (error, *KeyVersion) {
	return notImplemented(), nil
}

func (InMemoryKms) DeleteKey(name string) error {
	return notImplemented()
}

func (InMemoryKms) GetKeyMetadata(name string) (error, *KeyDesc) {
	return notImplemented(), nil
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

func (InMemoryKms) GetKeyNames() (error, []string) {
	return notImplemented(), nil
}

func (InMemoryKms) GetKeysMetadata(names []string) (error, []KeyMeta) {
	return notImplemented(), nil
}