package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/okkun-sh/json2table"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "json2table",
	Short: "Convert JSON array of objects to a formatted table",
	Run: func(cmd *cobra.Command, args []string) {
		input, err := io.ReadAll(bufio.NewReader(os.Stdin))
		if err != nil {
			log.Fatalf("Failed to read stdin: %v", err)
		}

		var raw []map[string]interface{}
		if err := json.Unmarshal(input, &raw); err != nil {
			log.Fatalf("Failed to parse JSON: %v", err)
		}

		json2table.PrintTable(raw)
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
