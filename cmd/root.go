package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nightcord",
	Short: "Generate Dockerfile from .lang files",
	Long: `Automatically generates Dockerfile by concatenating 
all .lang files in the current directory`,
	Run: func(cmd *cobra.Command, args []string) {
		langFiles, err := filepath.Glob("*.lang")
		if err != nil {
			fmt.Printf("查找.lang文件失败: %v\n", err)
			return
		}

		var dockerContent strings.Builder
		dockerContent.WriteString("# Auto-generated Dockerfile - DO NOT EDIT\n\n")
		dockerContent.WriteString("FROM alpine:latest\n\n")

		for _, langFile := range langFiles {
			content, err := os.ReadFile(langFile)
			if err != nil {
				fmt.Printf("读取文件 %s 失败: %v\n", langFile, err)
				continue
			}

			dockerContent.WriteString(fmt.Sprintf("# ==== %s ====\n", langFile))
			dockerContent.Write(content)
			dockerContent.WriteString("\n\n")
		}

		err = os.WriteFile("Dockerfile", []byte(dockerContent.String()), 0644)
		if err != nil {
			fmt.Printf("写入Dockerfile失败: %v\n", err)
			return
		}

		fmt.Println("Dockerfile生成成功")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
