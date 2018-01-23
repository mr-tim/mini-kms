package kms

type KeyMeta struct {
	Name        string
	Cipher      string
	Length      int
	Description string
	Created     int64
	Versions    int
}

type KeyDesc struct {
	Metadata KeyMeta
	Material []byte
}

type KeyVersion struct {
	VersionName string
	Material    []byte
}

type EncryptedKeyVersion struct {
	KeyName     string
	VersionName string
	IV          []byte
	Material    []byte
}

type Kms interface {
	CreateKey(desc KeyDesc) (error, *KeyVersion)
	RolloverKey(name string, material string) (error, *KeyVersion)
	DeleteKey(name string) error
	GetKeyMetadata(name string) (error, *KeyMeta)
	CurrentVersion(name string) (error, *KeyVersion)
	GenerateEncryptedKeys(name string, keysToGenerate int) (error, []EncryptedKeyVersion)
	DecryptEncryptedKey(versionName string, version EncryptedKeyVersion) error
	GetKeyVersion(versionName string) (error, *KeyVersion)
	GetKeyVersions(keyName string) (error, []string)
	GetKeyNames() (error, []string)
	GetKeysMetadata(names []string) (error, []KeyMeta)
	CreateMaterial(cipher string, length int) []byte
}

func New() Kms {
	return InMemoryKms{
		store: InMemoryKeyStore{
			keyMetadata: map[string]KeyMeta{},
			keyVersionMaterial: map[string][]KeyVersion{},
		},
	}
}
