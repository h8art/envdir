package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "envDir",
	Short: "env",
	Long:  `Such interesting desc`,
	Run:   ExecWithParams,
}

func ReadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return text, nil
}
func ExecWithParams(cmd *cobra.Command, args []string) {
	fmt.Println("hi")

	if len(args) == 2 {
		files, err := ioutil.ReadDir(args[0])
		if err != nil {
			cmd.PrintErr("no such dir")
			return
		}
		var flags []string
		flags = append(flags, args[1])
		for _, file := range files {
			flag := file.Name()
			arg, err := ReadFile(args[0] + "/" + flag)
			if err != nil {
				fmt.Println(err)
				cmd.PrintErr("file " + flag + " is empty\n")
			}
			flags = append(flags, "-"+flag+"="+arg)
		}
		c1 := exec.Command(strings.Join(flags, " "))
		fmt.Println(strings.Join(flags, " "))
		stdout, _ := c1.StdoutPipe()
		r := bufio.NewReader(stdout)
		line, _, _ := r.ReadLine()
		fmt.Println(string(line))
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
