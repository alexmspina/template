// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Build

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	//Change to root/portapp to access go.mod and go.sum
	fmt.Println("Changing to root/portapp")
	os.Chdir("../..")

	// Install dependencies with go mod download
	mg.Deps(InstallDeps)

	// Build executable from cmd/portapp.go
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "./cmd/producer/producer.go")

	// Execute target commmand
	return cmd.Run()
}

// A custom install step if you need your bin someplace other than go/bin
func Install() error {
	mg.Deps(Build)
	fmt.Println("Installing...")
	return os.Rename("./producer", "/usr/bin/producer")
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing Deps...")
	cmd := exec.Command("go", "mod", "download")
	return cmd.Run()
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("producer")
}
