package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
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
	if len(args) != 2 {
		log.Fatal("needed only to params")
	}
	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		cmd.PrintErr("no such dir")
		return
	}
	var flags []string
	flags = append(flags)
	for _, file := range files {
		envVar := file.Name()
		arg, err := ReadFile(args[0] + "/" + envVar)
		if err != nil {
			fmt.Println(err)
			cmd.PrintErr("file " + envVar + " is empty\n")
		}
		err = os.Setenv(envVar, arg)
		if err != nil {
			log.Fatal(err)
		}
	}
	c1 := exec.Command(args[1])
	fmt.Println(strings.Join(flags, " "))
	stdout, _ := c1.StdoutPipe()
	r := bufio.NewReader(stdout)
	line, _, _ := r.ReadLine()
	fmt.Println(string(line))
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
