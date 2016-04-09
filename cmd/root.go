package cmd

import (
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/guillaumebreton/regen/generator"
	"github.com/guillaumebreton/regen/loader"
	"github.com/spf13/cobra"
)


//OutputPath represents to path to which template will be outputed
var OutputPath string

// TemplatePath represents the path into which template will be searched
var TemplatePath string

var l = loader.NewLoader()

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "regen directory [-t template.html] [-o output]",
	Short: "Generate an html resume from toml",
	Long: `The tool generate html template from a set of toml file into
  an output directory. The Html template is using the golang templating system
  and the toml structure is defined in the structure.go file
  `,
	Run: func(cmd *cobra.Command, args []string) {
    dir := "."
		if len(args) > 0 {
      dir = args[0]
		}

		//load templates
		files, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Println("Fail to list file in directory")
			os.Exit(2)
		}

		//create the ouput path
		_, err = os.Open(OutputPath)
		if err != nil {
			os.Mkdir(OutputPath, 0777)
		}
		fileInfo, err := os.Stat(OutputPath)
		if !fileInfo.IsDir() || err != nil {
			fmt.Printf("File %f is not a directory\n", OutputPath)
			os.Exit(3)
		}
		for _, file := range files {
			if !file.IsDir() {
				ext := filepath.Ext(file.Name())
				if ext == ".toml" {
					resume, err := l.Load(dir, file.Name())
					if err != nil {
						fmt.Printf("Fail to load resume : %+v\n", err)
					} else {
						output, err := generator.Execute(TemplatePath, resume)
						if err != nil {
							fmt.Printf("Fail to generate resume : %+v\n", err)
						} else {

							templateExt := filepath.Ext(TemplatePath)
							filename := strings.TrimSuffix(file.Name(), ext) + templateExt
							fp := filepath.Join(OutputPath, filename)
							err = writeString(fp, output)
							if err != nil {
								fmt.Printf("Fail write to file %s\n", fp)
								fmt.Println(err)
							} else {
                fmt.Printf("%s generated to %s\n", file.Name(), fp)
              }

						}
					}
				}
			}

		}
	},
}

//writeString write a string to a file. If the file does
//not exists, it is created, else it's overwritten
func writeString(fp string, data string) error {
	f, err := os.Create(fp)
	if err != nil {
		return err
	}
	_, err = f.WriteString(data)
	return err
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&OutputPath, "output-path", "o", "output", "The destination path for the generated resume")
	RootCmd.PersistentFlags().StringVarP(&TemplatePath, "template", "t", "template.html", "The template to use to generate the resume")
}
