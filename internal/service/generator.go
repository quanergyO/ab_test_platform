package service

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"

	"github.com/quanergyO/ab_test_platform/internal/models"
)

const (
	soPath = "internal/generated/sharedobjects/"
)

type Generator struct {
	// TODO inject logger
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) GeneratePlugin(spec models.ABTestSpec) (string, error) {
	tmpl, err := template.ParseFiles("internal/templates/plugin.tmpl")
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%s_%s", spec.ServiceName, spec.RouteName)

	outputPath := fmt.Sprintf("internal/generated/files/%s", fmt.Sprintf("%s.go", fileName))
	file, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	err = tmpl.Execute(file, spec)
	if err != nil {
		return "", err
	}

	return compileToSO(outputPath, fileName)
}

func compileToSO(inputFile, outputFile string) (string, error) {
	outputFile = fmt.Sprintf("%s/%s.so", soPath, outputFile)
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", outputFile, inputFile)
	log.Println(cmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("ошибка компиляции: %v\n%s", err, output)
	}
	log.Printf("Файл успешно скомпилирован: %s\n", outputFile)

	return outputFile, nil
}
