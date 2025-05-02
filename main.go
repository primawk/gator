package main

import (
    "os"
	"fmt"
    "github.com/primawk/gator/internal/config"
)

func main() {
    if len(os.Args) < 2 {
		fmt.Println("Error: No command provided.\nUsage: <command> [args...]")
		os.Exit(1) 
	}

    cmd := config.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	cfg, err := config.Read()
    if err != nil {
        fmt.Println("failed to read config:", err)
        os.Exit(1) 
    }

    state := &config.State{
        Config: &cfg,
    }

    cmds := &config.Commands{}
    cmds.Register("login", config.HandlerLogin)
    

    if err := cmds.Run(state, cmd); err != nil {
		fmt.Println("Error:", err)
        os.Exit(1) 
	}





    // err = cfg.SetUser("prima")
    // if err != nil {
    //     fmt.Println("failed to set user:", err)
    // }
	// fmt.Println("Config contents:")
	// fmt.Printf("Current User: %s\n", state.Config.Username)
	// fmt.Printf("Current DB: %s\n", state.Config.Db_url)
}
