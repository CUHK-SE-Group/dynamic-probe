package main

import (
	"fmt"
	"kprobe/controller"
	"log"
	"os"

	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	var tmplFiles []string
	var confFile string
	var headerPath string
	var outputDir string

	var cmdGenerate = &cobra.Command{
		Use:   "generate",
		Short: "generate eBPF C code from the template and configuration",
		Run: func(cmd *cobra.Command, args []string) {
			ctrl, err := controller.InitContorller(
				controller.WithGenerateCCode(),
				controller.WithCompileCCode(),
				controller.WithGenerateCtlCode(),
				controller.WithConfig(confFile),
				controller.WithTemplates(tmplFiles...),
				controller.WithOutputFile(outputDir),
				controller.WithCHeaders(headerPath),
				controller.WithFunc("Transform", strcase.ToCamel),
			)
			if err != nil {
				log.Fatal(err)
			}
			ctrl.Run()
		},
	}

	cmdGenerate.Flags().StringSliceVar(&tmplFiles, "tmpl", []string{}, "template files")
	cmdGenerate.Flags().StringVar(&confFile, "conf", "", "configuration file")
	cmdGenerate.Flags().StringVar(&headerPath, "header", "", "headers for C code")
	cmdGenerate.Flags().StringVar(&outputDir, "output-dir", "", "output-dir for codes")

	var rootCmd = &cobra.Command{Use: "kb"}
	rootCmd.AddCommand(cmdGenerate)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
