package execa

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os/exec"

	"github.com/immanoj16/execa/types"
)

func setExtraOptions(cmd *exec.Cmd, options types.Exec) {
	cmd.Dir = options.Dir
	cmd.Env = options.Env.ToString()
	cmd.Stdin = options.Stdin
	cmd.Stdout = options.Stdout
	cmd.Stderr = options.Stderr
}

func showOutput(cmd *exec.Cmd) {
	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

// Run is responsible to run the command with the available options
func Run(name string, args []string, options types.Exec) *exec.Cmd {
	cmd := exec.Command(name, args...)
	setExtraOptions(cmd, options)
	if options.Pipe {
		showOutput(cmd)
	}
	return cmd
}

// RunContext is responsible to run the command with context with the available options
func RunContext(ctx context.Context, name string, args []string, options types.Exec) *exec.Cmd {
	cmd := exec.CommandContext(ctx, name, args...)
	setExtraOptions(cmd, options)
	if options.Pipe {
		showOutput(cmd)
	}
	return cmd
}
