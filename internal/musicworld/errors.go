package musicworld

import "fmt"

const (
	// DockerError125 is the stdout result from a docker command that results in exit status code of 125
	DockerError125 = "exit status 125"
)

// InspectContainerError returns an error that is a combination of the given error and container name
func InspectContainerError(err error, container string) error {
	return fmt.Errorf("error %v received while trying to inspect the .State.Status of %v", err, container)
}

// PullingContainerError returns an error that is a combination of the given error and container name
func PullingContainerError(err error, container string) error {
	return fmt.Errorf("error %v received while trying to pulldown %v", err, container)
}

// RunningContainerError returns an error that is a combination of the given error and container name
func RunningContainerError(err error, container string) error {
	return fmt.Errorf("error %v occurred while spinning up the %v container", err, container)
}

// CheckForRunningContainer returns an error that is a combination of teh given error and container name
func CheckForRunningContainer(err error, container string) error {
	return fmt.Errorf("error %v occurred. Check for a running %v container", err, container)
}
