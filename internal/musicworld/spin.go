package musicworld

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const (
	// running is the value for a running container's State.Status in a docker inspect object
	running                                   = "running"
	alexmspinMusicworldDevContainerName       = "alexmspina/musicworld-dev:latest"
	musicworldDevContainerName                = "musicworld_dev"
	musicworldDevContainerSuccessfullyStarted = "The musicworld_dev container has successfully been started"
	musicworldDevContainerAlreadyRunning      = "The musicworld_dev container is already running! You're good to run 'docker exec -it musicworld_dev fish'"
	pullingNewMusicworldimage                 = "Pulling down alexmspina/musicworld-dev:latest"
	dockerCommand                             = "docker"
	runningNewMusicworldContainer             = "Preparing to run musicworld_dev using the alexmspina/musicworld-dev image"
	startExistingContainer                    = "The musicworld_dev container has already been built. Attempting to restart the container..."
)

var (
	dockerInspectSlice = []string{"inspect", "-f", "'{{.State.Status}}'", "musicworld_dev"}
	dockerPullSlice    = []string{"pull", "alexmspina/musicworld-dev:latest"}
	dockerRunSlice     = []string{"run", "-itd", "--name", "musicworld_dev", "alexmspina/musicworld-dev"}
	dockerStartSlice   = []string{"start", "musicworld_dev"}
)

var spinLogger *zap.SugaredLogger
var verbose bool

func init() {
	SpinCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "output detailed logging of the spin command")
}

// spinStartingMessage returns a string verifying that the command has started to 'spin-up' a specified service
func spinStartingMessage(service string) string {
	return fmt.Sprintf("spinning up dat %v track!\n\n", service)
}

// SpinCmd is the cobra defintion for the spin subcommand
var SpinCmd = &cobra.Command{
	Use:   "spin SERIVCE",
	Short: "spin up some vinyl!",
	Long: `
Start a service

Services:
	devcontainer		a general rust-alpine image with terminal based dev tools`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Print(`
"musicworld spin" requires exactly one argument.

See 'musicworld spin --help'.

Usage: musicworld spin SERVICE
`)
		} else {
			spinLogger, newErr := InitializeLogger(verbose)
			err := Spin(args[0])
			if err != nil {
				if newErr != nil {
					fmt.Printf("error %v during InitializeLogger", err)
				}
				spinLogger.Debug(err)
			} else {
				fmt.Printf(`%v successfully spun up!

Run 'docker exec -it musicworld_dev fish' to access
the rust development container.
`,
					args[0])
			}

		}
	},
}

// Spin starts a service
func Spin(service string) error {
	spinLogger, err := InitializeLogger(verbose)
	if err != nil {
		fmt.Printf("Spin -> spinLogger %+v", err)
	}
	spinLogger.Info("spin logger initialized")
	spinLogger.Info("spinning-up", "service", service)
	spinStartingMessage(service)

	// attempt to pulldown alexmspina/musicworld-dev:latest image
	fmt.Println(pullingNewMusicworldimage)
	pullDevContainerCmd := exec.Command(dockerCommand, dockerPullSlice...)
	err = pullDevContainerCmd.Run()
	if err != nil {
		return PullingContainerError(err, alexmspinMusicworldDevContainerName)
	}

	// run musicworld_dev with alexmspina/musicworld-dev:latest image
	fmt.Println(runningNewMusicworldContainer)
	runDevContainerCmd := exec.Command(dockerCommand, dockerRunSlice...)
	err = runDevContainerCmd.Run()
	if err != nil {
		if err.Error() == DockerError125 {
			spinLogger.Debugf("Spin -> runDevContainer %+v", err)
			// attempt to restart musicworld_dev
			spinLogger.Info(startExistingContainer)
			restartDevContainerCmd := exec.Command(dockerCommand, dockerStartSlice...)
			err = restartDevContainerCmd.Run()
			if err != nil {
				return fmt.Errorf("Spin -> restartDevContainerCmd %+v", err)
			}
		}
	}

	spinLogger.Info("inspecting container ", musicworldDevContainerName)
	inspectContainer := exec.Command(dockerCommand, dockerInspectSlice...)
	var devContainerStatus bytes.Buffer
	inspectContainer.Stdout = &devContainerStatus
	err = inspectContainer.Run()
	if err != nil {
		return InspectContainerError(err, musicworldDevContainerName)
	}

	if strings.Contains(devContainerStatus.String(), running) {
		spinLogger.Info(musicworldDevContainerSuccessfullyStarted)
		return nil
	}

	return nil
}
