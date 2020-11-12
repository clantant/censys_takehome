package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/clantant/mys_sweepsy/cmd"
)

func main() {
	cmd.Execute()
}
