Feature: User can access a rust development container

	Scenario: User can access a rust development environment
		Given a running shell
		When user runs 'musicworld spin devcontainer'
		Then there is a rustdev container running
