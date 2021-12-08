package main

import (
	"github.com/spf13/cobra"
)

func main() {
	NewCmds().Execute()
}

func NewCmds() *cobra.Command {
	c := &cobra.Command{
		Use:   "hello",
		Short: "Say hello",
	}

	worldCmd := &cobra.Command{
		Use:   "world",
		Short: "Hello world",
		RunE:  hellowWorldHandler,
	}

	cosmosCmd := &cobra.Command{
		Use:   "cosmos",
		Short: "Hello Cosmos",
		RunE:  hellowCosmosHandler,
	}

	c.AddCommand(worldCmd)
	c.AddCommand(cosmosCmd)
	return c
}

func hellowWorldHandler(cmd *cobra.Command, args []string) error {
	cmd.Print("Hello world")

	return nil
}

func hellowCosmosHandler(cmd *cobra.Command, args []string) error {
	cmd.Print("Hello Cosmos")

	return nil
}
