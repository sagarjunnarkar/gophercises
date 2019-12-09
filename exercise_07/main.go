package main

import (
	"fmt"
	"gophercises/task/cmd"
	"gophercises/task/db"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) int {
	if err != nil {
		fmt.Println(err.Error())
		return -1
	}
	return 1
}
