package main

import (
	"github.com/tcosolutions/ptpt/cmd"
	"github.com/tcosolutions/ptpt/internal/config"
	_ "github.com/tcosolutions/ptpt/static"
)

func main() {
	config.Init()
	cmd.Execute()
}
