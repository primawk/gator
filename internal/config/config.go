package config

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
	"encoding/json"
)

type Config struct {
 	Db_url string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func Read() Config {
	// Get the HOME directory
	homeDir, err := os.UserHomeDir()
	  if err != nil {
		  log.Fatal("Unable to get HOME directory:", err)
	}
	// Build full path to the file
	filePath := filepath.Join(homeDir, "workspace/github.com/primawk/gator/gatorconfig.json")

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var cfg Config
    if err := decoder.Decode(&cfg); err != nil {
        panic(err)
    }

	fmt.Printf("Decoded Config: %+v\n", cfg)
	return cfg
}

func (c Config) SetUser() string {

    return p.FirstName + " " + p.LastName
}