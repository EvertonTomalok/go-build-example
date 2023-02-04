package cmd

import (
	"github.com/evertontomalok/go-build-example/app/domain"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var runCmd = &cobra.Command{
	Use:   "command",
	Short: "Run the command",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		lastName, _ := cmd.Flags().GetString("lastName")

		person := domain.Person{
			Name:     name,
			LastName: lastName,
		}
		person.SayMyName()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().String("name", "", "Tell your name.")
	runCmd.PersistentFlags().String("lastName", "", "Tell your last name.")

	_ = viper.BindEnv("SOME_ENV")
}
