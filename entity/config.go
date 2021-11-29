package entity

type ConfigHub struct {
	Type     string `json:"type"`
	Endpoint string `json:"endpoint"`
}

type ConfigDocumentDB struct {
	Type             string `json:"type"`
	ConnectionString string `json:"connection_string"`
	Database         string `json:"database"`
}

type Config struct {
	Hub        *ConfigHub        `json:"hub"`
	DocumentDB *ConfigDocumentDB `json:"documentdb"`
}
