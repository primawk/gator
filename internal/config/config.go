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

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, "workspace/github.com/primawk/gator/gatorconfig.json"), nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create config file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // optional pretty-print
	if err := encoder.Encode(cfg); err != nil {
		return fmt.Errorf("could not encode config to file: %w", err)
	}

	return nil
}


func Read() (Config, error){
	path, err := getConfigFilePath()
	if err != nil {
		return Config{},fmt.Errorf("could not get config path: %w", err)
	}
	// Open the file
	file, err := os.Open(path)
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
	return cfg, nil
}

func (c *Config) SetUser(username string) error {
	c.Username = username
	if err := write(*c); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}
	return nil
}