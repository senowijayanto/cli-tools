package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mytools [path of log file] [flags of type] [type] [flags of output] [output path]",
	Short: "A CLI application for read and convert log files",
	Long:  `A tool to retrieve and convert log files on the Linux file system in the /var/log folder into PlainText or JSON format`,

	Run: func(cmd *cobra.Command, args []string) {
		isType := cmd.Flags().Changed("type")
		isOutput := cmd.Flags().Changed("output")
		getOutput, _ := cmd.Flags().GetString("output")
		getType, _ := cmd.Flags().GetString("type")
		outputString := ""

		if isOutput {
			outputString = getOutput
		}

		if getType == "json" {
			convertFiles(args[0], getType, outputString)
		} else if getType == "text" || !isType {
			convertFiles(args[0], getType, outputString)
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("type", "t", "text", "Type of convert, text or json")
	rootCmd.Flags().StringP("output", "o", "", "Choosing output and location file after convert")
}

type Element struct {
	Data string `json:"data"`
}

func convertFiles(logPath string, fileType string, outputPath string) {
	f, err := os.Open(logPath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var dataSlice = make([]Element, 0)
	var strData []string

	if fileType == "json" {
		for scanner.Scan() {
			var data = scanner.Text()
			dataSlice = append(dataSlice, Element{Data: data})
		}
		bts, err := json.Marshal(dataSlice)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n", bts)
	} else {
		for scanner.Scan() {
			data := fmt.Sprintf(scanner.Text() + "\n")
			strData = append(strData, data)
		}
		fmt.Printf("%s\n", strData)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	if outputPath != "" {
		if fileType == "json" {
			file, _ := json.MarshalIndent(dataSlice, "", " ")
			_ = ioutil.WriteFile(outputPath, file, 0644)
		} else {
			strFile := strings.Join(strData, "\n")
			_ = ioutil.WriteFile(outputPath, []byte(strFile), 0644)
		}
	}
}
