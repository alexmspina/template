package musicworld

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

const (
	DOCKER_ERR_125 = "exit status 125"
	RUNNING        = "running"
)

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
			Spin(args[0])
		}
	},
}

// Spin starts a service
func Spin(service string) {
	fmt.Printf("spinning up dat %v track!\n\n", service)

	execIntoDevContainer := exec.Command("docker", "inpsect", "-f", "'{{.State.Status}}'", "musicworld_dev")
	var devContainerStatus bytes.Buffer
	execIntoDevContainer.Stdout = &devContainerStatus
	err := execIntoDevContainer.Run()
	if err != nil {
		execIntoDevContainerError := fmt.Errorf("error %v received while trying to inspect the .State.Status of musicworld_dev", err)
		fmt.Println(execIntoDevContainerError)
	}

	if devContainerStatus.String() == RUNNING {
		fmt.Println("A musicworld_dev container is already running! You're good to run 'docker exec -it musicworld_dev fish'")
	} else {
		fmt.Println("Trying to build a new musicworld_dev image..")
		fmt.Println("Pulling down alexmspina/musicworld-dev:latest")

		pullDevContainerCmd := exec.Command("docker", "pull", "alexmspina/musicworld-dev:latest")
		err = pullDevContainerCmd.Run()
		if err != nil {
			pullDevContainerError := fmt.Errorf("error %v occurred while attempting to pull alexmspina/musicworld-dev container", err)
			fmt.Println(pullDevContainerError)
		}

		fmt.Println("Preparing to run musicworld_dev using the alexmspina/musicworld-dev image")
		runDevContainerCmd := exec.Command("docker", "run", "-itd", "--name", "musicworld_dev", "alexmspina/musicworld-dev")
		err = runDevContainerCmd.Run()
		if err != nil {
			runDevContainerCmdErr := fmt.Errorf(
				"error %v occurred while spinning up the musicworld_dev container",
				err,
			)

			fmt.Println(runDevContainerCmdErr)

			if err.Error() == DOCKER_ERR_125 {
				fmt.Println(
					"check if a musicworld_dev container is present on your system",
				)
			}
		}
	}

}
