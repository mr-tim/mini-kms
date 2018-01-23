package kms

import "fmt"

type KeyStore interface {
	createKey(desc KeyDesc, material []byte) (error, *KeyVersion)
	getMetadata(name string) (*KeyMeta, bool)
	createNewKeyVersion(name string, material []byte) (error, *KeyVersion)
	getKeyNames() (error, []string)
	deleteKey(keyName string) error
	getKeyVersions(name string) (error, []KeyVersion)
}

type InMemoryKeyStore struct {
	keyMetadata        map[string]KeyMeta
	keyVersionMaterial map[string][]KeyVersion
}

func (s InMemoryKeyStore) createKey(desc KeyDesc, material []byte) (error, *KeyVersion) {
	_, found := s.keyMetadata[desc.Metadata.Name]
	if found {
		return keyAlreadyExists(desc.Metadata.Name), nil
	}
	desc.Metadata.Versions = 0
	s.keyMetadata[desc.Metadata.Name] = desc.Metadata

	return s.createNewKeyVersion(desc.Metadata.Name, material)
}

func (s InMemoryKeyStore) getMetadata(name string) (*KeyMeta, bool) {
	meta, found := s.keyMetadata[name]
	if !found {
		return nil, false
	}
	return &meta, true
}

func (s InMemoryKeyStore) createNewKeyVersion(keyName string, material []byte) (error, *KeyVersion) {
	meta, found := s.getMetadata(keyName)

	if !found {
		return keyDoesNotExist(keyName), nil
	}

	newKeyVersion := KeyVersion{
		s.versionName(keyName, meta.Versions),
		material,
	}
	s.keyVersionMaterial[keyName] = append(s.keyVersionMaterial[keyName], newKeyVersion)
	s.incrementVersionCount(keyName)

	return nil, &newKeyVersion
}

func (InMemoryKeyStore) versionName(name string, version int) string {
	return fmt.Sprintf("%s/%d", name, version)
}

func (s InMemoryKeyStore) getKeyNames() (error, []string) {
	keyNames := make([]string, 0)
	for k := range s.keyMetadata {
		keyNames = append(keyNames, k)
	}
	return nil, keyNames
}

func (s InMemoryKeyStore) deleteKey(keyName string) error {
	delete(s.keyVersionMaterial, keyName)
	delete(s.keyMetadata, keyName)
	return nil
}

func (s InMemoryKeyStore) getKeyVersions(keyName string) (error, []KeyVersion) {
	return nil, s.keyVersionMaterial[keyName]
}

func (s InMemoryKeyStore) incrementVersionCount(keyName string) {
	meta, found := s.keyMetadata[keyName]
	if found {
		meta.Versions += 1
		s.keyMetadata[keyName] = meta
	}
}
