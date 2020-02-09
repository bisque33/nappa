/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// curlCmd represents the curl command
var curlCmd = &cobra.Command{
	Use:   "curl",
	Short: "Convert curl command format to vegeta input format.",
	Long: `Convert curl command format to vegeta input format.
	
- Arguments and flags conform to curl command.
- Supported flags are listed in Flags.`,
	Args: cobra.MinimumNArgs(1),
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		// Ignore unknown flags
		UnknownFlags: true,
	},
	Run: func(cmd *cobra.Command, args []string) {
		method, err := cmd.PersistentFlags().GetString("request")
		if err != nil {
			fmt.Println("error:", err)
		}
		header := make(map[string][]string)
		headers, err := cmd.PersistentFlags().GetStringArray("header")
		if err != nil {
			fmt.Println("error:", err)
		}
		for _, str := range headers {
			slice := strings.Split(str, ":")
			header[slice[0]] = append(header[slice[0]], strings.TrimSpace(slice[1]))
		}

		var body = ""
		if data, _ := cmd.PersistentFlags().GetString("data"); data != "" {
			body = base64.StdEncoding.EncodeToString([]byte(data))
		}
		if data, _ := cmd.PersistentFlags().GetString("data-ascii"); data != "" {
			body = base64.StdEncoding.EncodeToString([]byte(data))
		}
		if data, _ := cmd.PersistentFlags().GetString("data-binary"); data != "" {
			body = base64.StdEncoding.EncodeToString([]byte(data))
		}
		if data, _ := cmd.PersistentFlags().GetString("data-raw"); data != "" {
			body = base64.StdEncoding.EncodeToString([]byte(data))
		}
		if data, _ := cmd.PersistentFlags().GetString("data-urlencode"); data != "" {
			body = base64.StdEncoding.EncodeToString([]byte(data))
		}

		type JSONSchema struct {
			Method string              `json:"method"`
			URL    string              `json:"url"`
			Header map[string][]string `json:"header,omitempty"` // TODO: custom schema
			Body   string              `json:"body,omitempty"`   // TODO: base64
		}
		group := JSONSchema{
			Method: method,
			URL:    args[0],
			Header: header,
			Body:   body,
		}

		b, err := json.Marshal(group)
		if err != nil {
			fmt.Println("error:", err)
		}
		// Need LF
		fmt.Printf("%s\n", b)
	},
}

func init() {
	curlCmd.PersistentFlags().StringP("request", "X", "GET", "method")
	curlCmd.PersistentFlags().StringArrayP("header", "H", []string{}, "headers")
	curlCmd.PersistentFlags().StringP("data", "d", "", "HTTP POST data")
	curlCmd.PersistentFlags().StringP("data-ascii", "", "", "HTTP POST ASCII data")
	curlCmd.PersistentFlags().StringP("data-binary", "", "", "HTTP POST binary data")
	curlCmd.PersistentFlags().StringP("data-raw", "", "", "HTTP POST data, '@' allowed")
	curlCmd.PersistentFlags().StringP("data-urlencode", "", "", "HTTP POST data url encoded")

	rootCmd.AddCommand(curlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// curlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// curlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
