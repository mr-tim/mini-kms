package kms

import "fmt"

type KeyStore interface {
	createKey(desc KeyDesc, material []byte) (error, *KeyVersion)
	getMetadata(name string) (*KeyMeta, bool)
	createNewKeyVersion(name string, material []byte) (error, *KeyVersion)
	getKeyNames() (error, []string)
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

	newVersionNumber := meta.Versions + 1
	newKeyVersion := KeyVersion{
		s.versionName(keyName, newVersionNumber),
		material,
	}
	s.keyVersionMaterial[keyName] = append(s.keyVersionMaterial[keyName], newKeyVersion)
	meta.Versions = newVersionNumber

	return nil, &newKeyVersion
}

func (InMemoryKeyStore) versionName(name string, version int) string {
	return fmt.Sprintf("%s/%d", name, version)
}

func (s InMemoryKeyStore) getKeyNames() (error, []string) {
	keyNames := make([]string, len(s.keyMetadata))
	for k := range s.keyMetadata {
		keyNames = append(keyNames, k)
	}
	return nil, keyNames
}
