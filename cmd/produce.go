/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// produceCmd represents the produce command
var produceCmd = &cobra.Command{
	Use:   "produce",
	Short: "Produce messages to local kafka",
	Long:  `The subcommand will produce messages to local kafka`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			return errors.New("please provid the number of messages to produce")
		} else {
			number_messages, err := strconv.Atoi(args[0])

			if err != nil {
				return errors.New("please provide a valid number for messages to produce")

			}

			produce_messages(number_messages)
			return nil
		}
	},
}

func produce_messages(num_of_messages int) {
	fmt.Printf("Producing %d messages", num_of_messages)
}

func init() {
	rootCmd.AddCommand(produceCmd)
}
