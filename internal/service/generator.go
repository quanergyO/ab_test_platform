package service

import (
	"fmt"
	"html/template"
	"os"

	"github.com/quanergyO/ab_test_platform/internal/models"
)

type Generator struct {
	// TODO inject logger
}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) GeneratePlugin(spec models.ABTestSpec) error {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return err
	}

	fmt.Println("Текущая директория:", dir)

	_, err = template.ParseFiles("generator/templates/plugin.tmpl")
	if err != nil {
		return err
	}

	// outputPath = fmt.Sprintf("generated/%s", )

	// file, err := os.Create(outputPath)
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()

	// err = tmpl.Execute(file, spec)
	// if err != nil {
	// 	return err
	// }

	return fmt.Errorf("GeneratePlugin not implemented")
}
