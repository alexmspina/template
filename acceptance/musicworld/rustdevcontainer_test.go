package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/alexmspina/template/internal/musicworld"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/stretchr/testify/assert"
)

const (
	// running is the value for a running container's State.Status in a docker inspect object
	exited                              = "exited"
	alexmspinMusicworldDevContainerName = "alexmspina/musicworld-dev:latest"
	musicworldDevContainerName          = "musicworld_dev"
	musicworldDevContainerRunning       = "A musicworld_dev container is already running! You're good to run 'docker exec -it musicworld_dev fish'"
	pullingNewMusicworldimage           = "Pulling down alexmspina/musicworld-dev:latest"
	dockerCommand                       = "docker"
	runningNewMusicworldContainer       = "Preparing to run musicworld_dev using the alexmspina/musicworld-dev image"
	startExistingContainer              = "A musicworld_dev container has already been built. Attempting to restart the container..."
	musicworldStillRunning              = "the musicworld dev container is still running"
)

var (
	dockerInspectSlice = []string{"inspect", "-f", "'{{.State.Status}}'", "musicworld_dev"}
	dockerPullSlice    = []string{"pull", "alexmspina/musicworld-dev:latest"}
	dockerRunSlice     = []string{"run", "-itd", "--name", "musicworld_dev", "alexmspina/musicworld-dev"}
	dockerStartSlice   = []string{"start", "musicworld_dev"}
	shellTypes         = []string{"bash", "ash", "zsh", "fish"}
	userShell          bytes.Buffer
	opt                = godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "progress",
	}
)

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
		return fmt.Errorf("aRunningShell -> echoShell cmd %+v", err)
	}

	if userShell.String() == "" {
		noShellErr := errors.New("aRunningShell -> could not find the current shell")
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
		unsupportedShellErr := errors.New("aRunningShell -> the current shell is not supported")
		return unsupportedShellErr
	}

	fmt.Printf("the current shell is %v\n", userShell.String())
	return nil
}

func aUserRunsMusicworldSpinDevcontainer() error {
	buildMusicworld := exec.Command("go", "build", "-o", "musicworld", "../../cmd/musicworld/main.go")
	err := buildMusicworld.Run()
	if err != nil {
		return fmt.Errorf("buildMusicworld cmd %+v", err)
	}

	spinMusicworldDevContainer := exec.Command("./musicworld", "spin", "devcontainer")
	err = spinMusicworldDevContainer.Run()
	if err != nil {
		return fmt.Errorf("aUserRunsMusicworldSpinaDevContainer -> spinMusicworldDevContainer cmd %+v", err)
	}

	return nil
}

func thereIsAMusicworldDevContainerRunning() error {
	dockerPs := exec.Command("docker", "ps")
	var psResult bytes.Buffer
	dockerPs.Stdout = &psResult
	err := dockerPs.Run()
	if err != nil {
		return fmt.Errorf("thereIsAMusicworldDevContainerRunning -> dockerPs cmd %+v", err)
	}
	fmt.Printf("\n%v\n", psResult.String())

	t := &testing.T{}

	godogAssertion := assert.New(t)

	if godogAssertion.Contains(psResult.String(), "exited") {
		containerExitedError := errors.New("thereIsAMusicworldDevContainerRunning -> the musicworld_dev container exited")
		return containerExitedError
	}

	if len(psResult.String()) == 0 {
		containerNotFoundError := errors.New("thereIsAMusicworldDevContainerRunning -> could not find the musicworld_dev container")
		return containerNotFoundError
	}

	execIntoMusicWorld := exec.Command("docker", "exec", "-d", "musicworld_dev", "fish")
	err = execIntoMusicWorld.Run()
	if err != nil {
		return fmt.Errorf("thereIsAMusicworldDevContainerRunning -> execIntoMusicWorld cmd %+v", err)
	}

	return nil
}

func aStoppedMusicworldDevContainer() error {
	stopMusicworldDevCmd := exec.Command("docker", "stop", "musicworld_dev")
	err := stopMusicworldDevCmd.Run()
	if err != nil {
		return fmt.Errorf("StoppedMusicworldDevContainer -> stopMusicworldDevCmd %+v", err)
	}
	inspectContainer := exec.Command(dockerCommand, dockerInspectSlice...)
	var devContainerStatus bytes.Buffer
	inspectContainer.Stdout = &devContainerStatus
	err = inspectContainer.Run()
	if err != nil {
		return musicworld.InspectContainerError(err, musicworldDevContainerName)
	}

	t := &testing.T{}
	godogAssertion := assert.New(t)
	if godogAssertion.Contains(devContainerStatus.String(), "exited") {
		fmt.Println("musicworld_dev exited")
		return nil
	}
	return errors.New(musicworldStillRunning)
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^a running shell$`, aRunningShell)
	s.Step(`^a user runs \'musicworld spin devcontainer\'$`, aUserRunsMusicworldSpinDevcontainer)
	s.Step(`^there is a musicworld dev container running$`, thereIsAMusicworldDevContainerRunning)
	s.Step(`^a stopped musicworld dev container$`, aStoppedMusicworldDevContainer)
}
