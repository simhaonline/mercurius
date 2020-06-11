package setup

type BlobStore struct {
	Driver string `json:"driver"`
	Path   string `json:"path"`
}

// Settings contains the minimal bootstrapping configuration to reach all data stores which itself contains
// all other configurations
type Settings struct {
	BlobStore   BlobStore `json:"blobStore"`
	Database    Sql       `json:"database"`
	Server      Server    `json:"server"`
	Development bool      `json:"development"`
}

// Server contains the http server specific settings
type Server struct {
	Port    int    `json:"port"`
	Address string `json:"address"`
}

// Sql contains the sql related settings
type Sql struct {
	Driver       string `json:"driver"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DatabaseName string `json:"databaseName"`
	SSLMode      string `json:"sslMode"`
}
