package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/stretchr/testify/assert"
)

var shellTypes = []string{"bash", "ash", "zsh", "fish"}
var userShell bytes.Buffer

var opt = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress",
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func aRunningShell() error {
	echoShell := exec.Command("echo", os.ExpandEnv("$SHELL"))
	echoShell.Stdout = &userShell
	err := echoShell.Run()
	if err != nil {
		return fmt.Errorf("error %v occurred while checking for the current shell", err)
	}

	if userShell.String() == "" {
		noShellErr := errors.New("could not find the current shell")
		return noShellErr
	}

	t := &testing.T{}

	godogAssertion := assert.New(t)

	var shellFound bool
	for _, shell := range shellTypes {
		if godogAssertion.Contains(userShell.String(), shell) {
			shellFound = true
		}
	}

	if !shellFound {
		unsupportedShellErr := errors.New("the current shell is not supported")
		return unsupportedShellErr
	}

	fmt.Printf("the current shell is %v", userShell.String())
	return nil
}

func userRunsMusicworldSpinDevcontainer() error {

	buildMusicworld := exec.Command("go", "build", "-o", "musicworld", "../../cmd/musicworld/main.go")
	err := buildMusicworld.Run()
	if err != nil {
		return fmt.Errorf("error %v trying to build musicworld", err)
	}

	spinMusicworldDevContainer := exec.Command("./musicworld", "spin", "devcontainer")
	err = spinMusicworldDevContainer.Run()
	if err != nil {
		return fmt.Errorf("error %v occurred while calling 'musicworld spin devcontainer'", err)
	}

	return nil
}

func thereIsARustdevContainerRunning() error {

	checkForMusicworldDevContainer := exec.Command("docker", "exec", "musicworld_dev", "fish")

	err := checkForMusicworldDevContainer.Run()
	if err != nil {
		return fmt.Errorf("error %v while trying to exec into musicworld_dev", err)
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^a running shell$`, aRunningShell)
	s.Step(`^user runs \'musicworld spin devcontainer\'$`, userRunsMusicworldSpinDevcontainer)
	s.Step(`^there is a rustdev container running$`, thereIsARustdevContainerRunning)
}
