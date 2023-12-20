package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/AnnDutova/bdas/lab1/internal/file"
	"github.com/AnnDutova/bdas/lab1/internal/logic"
)

var (
	clearPath         string
	obfuscationPath   string
	deobfuscationPath string
	forwardFlag       bool
)

func init() {
	flag.StringVar(&clearPath, "path", "./assets/clear.xml", "path to clear xml file")
	flag.StringVar(&obfuscationPath, "obfuscationPath", "./assets/obfuscation.xml", "path to obfuscation xml file")
	flag.StringVar(&deobfuscationPath, "deobfuscationPath", "./assets/deobfuscation.xml", "path to deobfuscation xml file")
	flag.BoolVar(&forwardFlag, "forward", true, "path of obfuscation. true for obfuscation, false for deobfuscation")
}

func main() {
	flag.Parse()
	fmt.Printf("%T", []string{})

	if forwardFlag {
		content, err := file.ReadFile(clearPath)
		if err != nil {
			log.Fatal(err)
		}

		file.WriteFile(logic.Obfuscation(content), obfuscationPath, "obfuscation.xml")

	} else {
		content, err := file.ReadFile(obfuscationPath)
		if err != nil {
			log.Fatal(err)
		}

		file.WriteFile(logic.Deobfuscation(content), deobfuscationPath, "deobfuscation.xml")
	}

}
