package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed package/*
var packageFiles embed.FS

type ProjectData struct {
	ScriptName  string
	ScriptTitle string
}

var rootCmd = &cobra.Command{
	Use:   "typstscript",
	Short: "A CLI tool for creating Typst screenplay projects",
	Long:  "typstscript helps you quickly create new screenplay projects using the Typst template system.",
}

var initCmd = &cobra.Command{
	Use:   "init [script-name]",
	Short: "Initialize a new screenplay project",
	Long:  "Create a new directory with the screenplay template files and a main script file.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		scriptName := args[0]

		if strings.Contains(scriptName, " ") || strings.Contains(scriptName, "/") {
			fmt.Fprintf(os.Stderr, "Error: Script name cannot contain spaces or slashes\n")
			os.Exit(1)
		}

		if _, err := os.Stat(scriptName); !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error: Directory '%s' already exists\n", scriptName)
			os.Exit(1)
		}

		if err := createProject(scriptName); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating project: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✓ Created new screenplay project: %s\n", scriptName)
		fmt.Printf("  cd %s\n", scriptName)
		fmt.Printf("  typst compile --font-path ./fonts %s.typ\n", scriptName)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func createProject(scriptName string) error {
	if err := os.MkdirAll(scriptName, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := copyPackageFiles(scriptName); err != nil {
		return fmt.Errorf("failed to copy template files: %w", err)
	}

	if err := createMainScript(scriptName); err != nil {
		return fmt.Errorf("failed to create main script: %w", err)
	}

	return nil
}

func copyPackageFiles(projectDir string) error {
	templateData, err := packageFiles.ReadFile("package/template.typ")
	if err != nil {
		return fmt.Errorf("failed to read embedded template.typ")
	}

	templatePath := filepath.Join(projectDir, "template.typ")
	if err := os.WriteFile(templatePath, templateData, 0644); err != nil {
		return fmt.Errorf("failed to write template.typ: %w", err)
	}

	fontsDir := filepath.Join(projectDir, "fonts", "Courier_Prime")
	if err := os.MkdirAll(fontsDir, 0755); err != nil {
		return err
	}

	fontFiles := []string{
		"CourierPrime-Regular.ttf",
		"CourierPrime-Bold.ttf",
		"CourierPrime-Italic.ttf",
		"CourierPrime-BoldItalic.ttf",
		"OFL.txt",
	}

	for _, fontFile := range fontFiles {
		embeddedPath := filepath.Join("package/fonts/Courier_Prime", fontFile)
		fontData, err := packageFiles.ReadFile(embeddedPath)
		if err != nil {
			return fmt.Errorf("failed to read embedded font %s %w", fontFile, err)
		}

		dstPath := filepath.Join(fontsDir, fontFile)
		if err := os.WriteFile(dstPath, fontData, 0644); err != nil {
			return fmt.Errorf("failed to write font %s: %w", fontFile, err)
		}
	}

	return nil
}

func createMainScript(scriptName string) error {
	scriptTitle := strings.ToUpper(strings.ReplaceAll(scriptName, "-", " "))

	mainTemplateData, err := packageFiles.ReadFile("package/main.typ")
	if err != nil {
		return fmt.Errorf("failed to read embedded main.typ: %w", err)
	}

	tmpl, err := template.New("main").Parse(string(mainTemplateData))
	if err != nil {
		return err
	}

	file, err := os.Create(filepath.Join(scriptName, scriptName+".typ"))
	if err != nil {
		return err
	}
	defer file.Close()

	data := ProjectData{
		ScriptName:  scriptName,
		ScriptTitle: scriptTitle,
	}

	return tmpl.Execute(file, data)
}
