package main

import (
	"fmt"

	"github.com/techrail/ground/dbCodegen"
)

func main() {
	cnf := dbCodegen.CodegenConfig{
		DbModelPackageName: "mainDb",
		DbModelPackagePath: "main",
		// DbModelPackagePath: "/Volumes/TestVM/other_data/obsidian_docker_root/ground/tmp/mainDb",
		PgDbUrl: "postgres://postgres:@127.0.0.1:5432/gofeed?sslmode=disable",
		// PgDbUrl: "postgres://vkaushal288:vkaushal288@127.0.0.1:5432/ac_dev?sslmode=disable",
	}
	g, e := dbCodegen.NewCodeGenerator(cnf)
	if e.IsNotBlank() {
		fmt.Printf("I#1NPKZR - Some error when creating new codegenerator: %v\n", e)
	}
	errTy := g.Generate()
	if errTy.IsNotBlank() {
		fmt.Printf("I#1NPLCJ - %v\n", errTy)
	}
}
