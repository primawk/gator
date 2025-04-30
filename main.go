package main

import (
	"fmt"
    "github.com/primawk/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
    if err != nil {
        fmt.Println("failed to read config:", err)
        return
    }

    err = cfg.SetUser("prima")
    if err != nil {
        fmt.Println("failed to set user:", err)
    }
	fmt.Println("Config contents:")
	fmt.Printf("DB URL: %s\n", cfg.Db_url)
	fmt.Printf("Current User: %s\n", cfg.Username)
}
