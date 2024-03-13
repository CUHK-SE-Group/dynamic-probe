package main

import (
	"fmt"
	"kprobe/controller"
	"log"
	"os"

	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
)

func commandGenerate() *cobra.Command {
	var tmplFiles []string
	var executeTmpls []string
	var confFile string
	var headerPath string
	var outputDir string

	var cmdGenerate = &cobra.Command{
		Use:   "generate",
		Short: "generate eBPF C code from the template and configuration",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(tmplFiles)
			ctrl, err := controller.InitContorller(
				controller.WithGenerateCCode(),
				// controller.WithCompileCCode(),
				// controller.WithGenerateCtlCode(),
				controller.WithConfig(confFile),
				controller.WithTemplates(tmplFiles...),
				controller.WithExecuteTemplate(executeTmpls...),
				controller.WithOutputFile(outputDir),
				controller.WithCHeaders(headerPath),
				controller.WithFunc("Transform", strcase.ToCamel),
				controller.WithFunc("AddOne", func(i int) int {
					return i + 1
				}),
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
	cmdGenerate.Flags().StringSliceVar(&executeTmpls, "execute", []string{}, "execute selected templates")
	return cmdGenerate
}

func commandTranslate() *cobra.Command {
	var file string

	var cmdTranslate = &cobra.Command{
		Use:   "translate",
		Short: "translate the ELF file to BTF's Go locale",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmdTranslate.Flags().StringVar(&file, "file", "", "input file name")
	return cmdTranslate
}

func main() {
	log.SetOutput(os.Stderr)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	var rootCmd = &cobra.Command{Use: "kb"}
	rootCmd.AddCommand(commandGenerate())
	rootCmd.AddCommand(commandTranslate())
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
