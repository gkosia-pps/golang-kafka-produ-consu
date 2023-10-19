package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// consumeCmd represents the consume command
var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "Consume data from local kafka",
	Long:  `The subcommand will consume messages from local kafka`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) == 0 {
			return errors.New("please provid the number of messages to produce")
		} else {
			number_messages, err := strconv.Atoi(args[0])

			if err != nil {
				return errors.New("please provide a valid number for messages to produce")

			}

			consume_messages(number_messages)
			return nil
		}
	},
}

func consume_messages(num_of_messages int) {
	fmt.Printf("Consuming %d messages", num_of_messages)

	appkafka.consume_from_kafka(num_of_messages)
}

func init() {
	rootCmd.AddCommand(consumeCmd)
}
