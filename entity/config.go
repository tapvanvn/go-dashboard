package entity

type ConfigHub struct {
	Type     string `json:"Type"`
	Endpoint string `json:"Endpoint"`
}

type ConfigDocumentDB struct {
	Provider         string `json:"Provider"`
	ConnectionString string `json:"ConnectionString"`
	Database         string `json:"Database"`
}

type Config struct {
	Environment string            `json:"Environment"`
	Hub         *ConfigHub        `json:"Hub"`
	DocDB       *ConfigDocumentDB `json:"DocDB"`
}
