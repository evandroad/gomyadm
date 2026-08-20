package commands

import (
	"fmt"
	"strings"
)

const (
	bold   = "\033[1m"
	reset  = "\033[0m"
	gray   = "\033[90m"
)

func printTable(headers []string, rows [][]string) {
	widths := make([]int, len(headers))

	for i, header := range headers {
		widths[i] = len(header)
	}

	for _, row := range rows {
		for i, value := range row {
			if len(value) > widths[i] {
				widths[i] = len(value)
			}
		}
	}

	// Header
	fmt.Print(bold)

	for i, header := range headers {
		fmt.Printf("%-*s", widths[i]+2, header)
	}

	fmt.Println(reset)

	// Linha separadora
	for i, width := range widths {
		if i > 0 {
			fmt.Print("  ")
		}

		fmt.Print(strings.Repeat("-", width))
	}

	fmt.Println(reset)

	// Dados
	for _, row := range rows {
		for i, value := range row {
			fmt.Printf("%-*s", widths[i]+2, value)
		}

		fmt.Println()
	}
}