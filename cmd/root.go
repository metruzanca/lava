package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/metruzanca/lava/internal"
	"github.com/metruzanca/lava/internal/lists"
	"github.com/metruzanca/lava/internal/markdown"
	"github.com/spf13/cobra"
)

type LavaFile struct {
	Path    string
	Content string
}

var rootCmd = &cobra.Command{
	Use:   "lava",
	Short: "Lava is a minimalist static site generator for Obsidian Vaults",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := internal.GetFileList("testdata/obsidian-help/Sandbox")

		if err != nil {
			fmt.Println("Error getting file list:", err)
			return
		}

		mdFiles := lists.Filter(files, func(f string) bool {
			return strings.HasSuffix(f, ".md")
		})

		mdFileContents := lists.Map(mdFiles, func(f string) *LavaFile {
			return &LavaFile{
				Path:    f,
				Content: internal.ReadFile(f),
			}
		})

		nonEmptyMdFiles := lists.Filter(mdFileContents, func(content *LavaFile) bool {
			return content.Content != ""
		})

		for _, file := range nonEmptyMdFiles {
			internal.WriteFile("./build/"+internal.ChangeExtension(file.Path, ".html"), markdown.Markdown2Html(file.Content))
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
