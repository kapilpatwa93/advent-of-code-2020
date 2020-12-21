package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/kapilpatwa93/advent-of-code-2020/util"
	"html/template"
	"log"
	"os"
	"path"
	"strings"
)
type FrequencyMap struct {
	Name     string
	Datatype string
	Package  string
}

func main() {
	templatePath := ""
	outputFilePrefix := ""
	outputPackageName := ""
	outputPath := ""
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("cannot fetch current working directory :%s", err)
	}

	flag.StringVar(&templatePath, "templatePath", path.Join(cwd, "template.txt"), "The name used for the queue being generated. This should start with a capital letter so that it is exported.")
	flag.StringVar(&outputFilePrefix, "filePrefix", "frequency_", "The name used for the queue being generated. This should start with a capital letter so that it is exported.")
	flag.StringVar(&outputPackageName, "package", "main", "The name used for the queue being generated. This should start with a capital letter so that it is exported.")
	flag.StringVar(&outputPath, "filePath", cwd, "The name used for the queue being generated. This should start with a capital letter so that it is exported.")
	flag.Parse()

	types := []string{"int", "string", "float32"}
	for _, t := range types {
		fileName := fmt.Sprintf("%s%s.go", outputFilePrefix, t)
		fm := FrequencyMap{strings.Title(t), t, outputPackageName}
		codeBuffer, err := generateCode(templatePath, fm)
		if err != nil {
			log.Fatalf("cannot write %s file : %s ", fileName, err)
		}
		err = util.WriteFile(codeBuffer.Bytes(), fileName, outputPath, false)
		if err != nil {
			log.Fatalf("cannot write %s file : %s ", fileName, err)
		}
	}

}
func generateCode(templatePath string, t FrequencyMap) (bytes.Buffer, error) {
	var codeBuffer bytes.Buffer
	templateText, err := util.ReadFile(templatePath)
	if err != nil {
		return codeBuffer, err
	}
	codeTemplate, err := template.New(t.Name).Parse(string(templateText))
	if err != nil {
		return codeBuffer, err
	}
	err = codeTemplate.Execute(&codeBuffer, t)
	if err != nil {
		return codeBuffer, err
	}
	return codeBuffer, nil
}
