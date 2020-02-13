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

	runDevContainerCmd := exec.Command("docker", "run", "-itd", "--name",
		"musicworld_dev", "template/rustdev")

	err := runDevContainerCmd.Run()
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
