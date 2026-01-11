package cmd

import (
	"os"

	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "what-to-go-wear",
	Short: "An app to tell me what to wear.",
	Long: `An app to tell me what to wear based on my closet, the day of the week, and the weather.`,
	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Beige shirt with blue pants.")
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


