// MedHash Tools
// Copyright (c) 2023 GHIFARI160
// MIT License

package cmd

// Command is a generic wrapper for all subcommands.
type Command interface {
	Execute() (status int)
}

// GenericCmd is a generic subcommand.
type GenericCmd struct{}

func (c GenericCmd) Execute() (status int) { return }
