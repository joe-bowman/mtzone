package main

import (
	"os"
	
	"github.com/cosmos/cosmos-sdk/server"
  svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
  
	"github.com/microtick/mtzone/app"
	"github.com/microtick/mtzone/cmd/mtm/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, app.DefaultHome); err != nil {
    switch e := err.(type) {
    case server.ErrorCode:
      os.Exit(e.Code)
    default:
      os.Exit(1)
    }
	}
}
