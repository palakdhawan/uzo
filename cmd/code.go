/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	util "uzo/utils"

	"github.com/spf13/cobra"
)
var File string
// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code <zip_file_name>",
	Short: "It will open the directory in vs code.",
	Long: `It will open the unzipped folder in VS Code.
	VS Code should be installed in order for this command to work.`,
	// Args: cobra.ExactArgs(1),
	Args:func(cmd *cobra.Command, args []string) error {
		if File=="" && len(args)<1 {
			return errors.New("accept(s) 1 argument")	
		}
		return nil
	},
	Example: "uzo code zipSpring.zip",
	Run: func(cmd *cobra.Command, args []string) {
		var fileName string
		var err error
		var argument string

		if File != ""{
			argument=File
		}else{
			argument=args[0]
		}
		

		fileExists,err:=util.FileExists(argument)

		if err!=nil {
			fmt.Println(err.Error())
		}

		if fileExists{
			fileName,err=filepath.Abs(argument)
			if err != nil{
				fmt.Println(err.Error())
			}
		}else
			{
				fmt.Printf("file %v does not exist",argument)
				return
			}
			
			wd,err:=os.Getwd()
			if err!=nil{
				fmt.Println(err.Error())
			}
			util.Unzip(fileName,wd)

			os.Chdir(util.FilenameWithoutExtension(fileName))

			wd,err=os.Getwd()
			if err!=nil{
				fmt.Println(err.Error())
			}
			commandCode:=exec.Command("code",wd)
			err=commandCode.Run()

			if err!=nil{
				fmt.Println("vs code executable file not found in %PATH%")
			}
	},
}

func init() {
	rootCmd.AddCommand(codeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	codeCmd.PersistentFlags().StringVarP(&File,"file","f", "", "A file name to unzip and open in IDE.")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// codeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
