package cmd

import (
  "fmt"
  "io/ioutil"
  "path/filepath"
  "os"
  "strings"

	"github.com/spf13/cobra"
  "github.com/guillaumebreton/regen/loader"
  "github.com/guillaumebreton/regen/generator"
)

var l = loader.NewLoader()

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate directory",
	Short: "Generate resumes using template",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
    if len(args) != 1 {
      fmt.Println("Usage is generate directory")
      os.Exit(1)
    }

    //load templates
    dir := args[0]
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
        if ext == ".toml"{
          resume, err := l.Load(dir, file.Name())
          if err != nil {
            fmt.Printf("Fail to load resume : %+v\n", err)
          } else {
            output, err  := generator.Execute(TemplatePath, resume)
            if err != nil {
              fmt.Printf("Fail to generate resume : %+v\n", err)
            } else {

              templateExt := filepath.Ext(TemplatePath)
              filename := strings.TrimSuffix(file.Name(), ext) + templateExt
              fp := filepath.Join(OutputPath, filename)
              fmt.Printf("Generating %s to %s \n", file.Name(), fp)
              err = writeString(fp, output)
              if err != nil {
                fmt.Printf("Fail write to file %s\n", fp)
                fmt.Println(err)
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

func init() {
	RootCmd.AddCommand(generateCmd)

}
