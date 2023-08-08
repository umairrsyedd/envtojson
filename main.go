// main.go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "envtojson",
		Short: "Convert .env files to JSON",
		Run:   convert,
	}

	rootCmd.Flags().StringP("input", "i", "", "Input .env file path (required)")
	rootCmd.Flags().StringP("output", "o", "output.json", "Output JSON file path (optional)")

	_ = rootCmd.MarkFlagRequired("input")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func convert(cmd *cobra.Command, args []string) {
	inputFile, _ := cmd.Flags().GetString("input")

	if inputFile == "" {
		fmt.Println("Input file is required.")
		return
	}

	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	envMap := createJsonMap(inputData)

	jsonData, err := json.MarshalIndent(envMap, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	outputFile, err := cmd.Flags().GetString("output")
	if err != nil {
		fmt.Println("Error getting output file:", err)
		return
	}

	err = os.WriteFile(outputFile, jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON output:", err)
		return
	}
	fmt.Println("Conversion successful. JSON file:", outputFile)
}

func createJsonMap(inputData []byte) map[string]string {
	envMap := make(map[string]string)
	lines := strings.Split(string(inputData), "\n")
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			envMap[key] = value
		}
	}
	return envMap
}
