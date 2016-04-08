// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
  "fmt"
  "io/ioutil"
  "path/filepath"
  "os"
  "strings"

	"github.com/spf13/cobra"
  "github.com/guillaumebreton/resume-generator/loader"
  "github.com/guillaumebreton/resume-generator/generator"
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
              filename := strings.TrimSuffix(file.Name(), ext) + ".html"
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