/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	_ "embed"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "commandhooker",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: action,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.commandhooker.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("mode", "m", "basic", "add mode opt")
}

const (
	GitDir   = ".git"
	HooksDir = "hooks"
	HookFile = "prepare-commit-msg"
)

//go:embed message.sh
var message string

func action(cmd *cobra.Command, args []string) {
	flg, path := findGitdir("./")
	if !flg {
		log.Println("git dir not found")
		return
	}
	data, err := createTemplateData(cmd.Flags())
	if err != nil {
		log.Println("git dir not found")
		return
	}
	createHook(path, data)
}

type TemplateData struct {
	CommitMsgString string
	WriteString     string
}

func createTemplateData(flags *pflag.FlagSet) (*TemplateData, error) {
	mode, err := flags.GetString("mode")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	commentType := getCommentType(mode)
	writingTxt := getInsertShell()
	return &TemplateData{
		CommitMsgString: commentType,
		WriteString:     writingTxt,
	}, nil

}

//
func createHook(path string, data *TemplateData) {
	fPath := filepath.Join(path, HooksDir, HookFile)
	file, err := os.Create(fPath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	t, err := template.New("template").Parse(message)
	if err != nil {
		log.Fatal(err)
	}
	if err = t.Execute(file, data); err != nil {
		log.Fatal(err)
	}
}

func getCommentType(mode string) string {
	switch mode {
	case "bname":
		return `
branchPath=$(git symbolic-ref -q HEAD)
commitMsg=${branchPath##*/}
	`
	default:
		return `
branchPath=$(git symbolic-ref -q HEAD)
branchName=${branchPath##*/}
issueNumber=$(echo $branchName | cut -d "_" -f 1)
commitMsg=$(head -n1 $1)
`
	}
}

func getInsertShell() string {
	switch runtime.GOOS {
	case "windows":
		return `sed -i  "1s/^/($commitMsg) \n/" $1`
	default:
		return `sed -i "" "1s/^/($commitMsg) \n/" $1`
	}
}

//
func findGitdir(path string) (bool, string) {
	files, _ := ioutil.ReadDir(path)
	for _, val := range files {
		if val.IsDir() && val.Name() == GitDir {
			return true, filepath.Join(path, val.Name())
		}
	}
	return false, ""
}
