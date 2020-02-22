Feature: User can access a rust development container

	Scenario: User can access a musicworld dev container
		Given a running shell
		When a user runs 'musicworld spin devcontainer'
		Then there is a musicworld dev container running

	Scenario: User can access an existing, but stopped musicworld dev container
		Given a running shell
		And a stopped musicworld dev container 
		When a user runs 'musicworld spin devcontainer'
		Then there is a musicworld dev container running