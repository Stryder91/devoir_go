package main

import (
	"devoir10_tran_iyeze/cmd"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Coucou et bienvenue sur le devoir CLI GO migration-tp")
	// stdout, stderr := os.Stdout, os.Stderr
	// fmt.Println(stdout, stderr)
	cmd.Execute()

}
