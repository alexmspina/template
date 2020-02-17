package musicworld

import (
	"fmt"
	"os/exec"
)

const (
	dockerErr125 = "exit status 125"
)

// Spin starts a service
func Spin(service string) {
	fmt.Printf("spinning up dat %v track!\n\n", service)

	pullDevContainerCmd := exec.Command("docker", "pull", "alexmspina/musicworld-dev:latest")
	err := pullDevContainerCmd.Run()
	if err != nil {
		pullDevContainerError := fmt.Errorf("error %v occurred while attempting to pull alexmspina/musicworld-dev container", err)
		fmt.Println(pullDevContainerError)
	}

	runDevContainerCmd := exec.Command("docker", "run", "-itd", "--name", "musicworld_dev", "alexmspina/musicworld-dev")

	err = runDevContainerCmd.Run()
	if err != nil {
		runDevContainerCmdErr := fmt.Errorf(
			"error %v occurred while spinning up the musicworld_dev container",
			err,
		)

		fmt.Println(runDevContainerCmdErr)

		if err.Error() == dockerErr125 {
			fmt.Println(
				"check if a musicworld_dev container is present on your system",
			)
		}
	}

}
