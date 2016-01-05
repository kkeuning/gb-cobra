package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	projectroot          = os.Getenv("GB_PROJECT_DIR")
	args        []string = os.Args[0:]
)

func main() {
	fatalf := func(format string, args ...interface{}) {
		fmt.Fprintf(os.Stderr, "FATAL: "+format+"\n", args...)
		os.Exit(1)
	}

	switch {
	case len(args) < 1, args[0] == "-h", args[0] == "-help", args[0] == "help":
		printUsage()
		os.Exit(0)
	case projectroot == "":
		fatalf("don't run this binary directly, it is meant to be run as 'gb cobra ...'")
	default:
	}

	env := mergeEnv(os.Environ(), map[string]string{
		"GOPATH": projectroot,
	})

	path, err := exec.LookPath("cobra")
	if err != nil {
		fatalf("run: unable to locate %q: %v", args[0], err)
	}

	cmd := exec.Cmd{
		Path: path,
		Args: args,
		Env:  env,

		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if err := cmd.Run(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Println("Done")
}

func mergeEnv(env []string, args map[string]string) []string {
	m := make(map[string]string)
	for _, e := range env {
		v := strings.SplitN(e, "=", 2)
		m[v[0]] = v[1]
	}
	for k, v := range args {
		m[k] = v
	}
	env = make([]string, 0, len(m))
	for k, v := range m {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	return env
}

func printUsage() {
	fmt.Println(`gb-cobra, a gb plugin to run the spf13/cobra generator with a GOPATH set to match the current gb project.

Usage:

        gb cobra [arguments]
`)
}
