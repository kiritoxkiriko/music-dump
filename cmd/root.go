package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "music-dump",
	Short: "A application convert encrypted music file to another format",
	Long:  `A application convert encrypted music file to another format`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.music-dump.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	// global flag
	rootCmd.PersistentFlags().StringP("config", "c", "./", "config file (default is $HOME/.music-dump.toml)")

	// local flag
	rootCmd.Flags().StringP("dumper", "d", "auto", "dumper type")
	rootCmd.Flags().StringP("log_level", "l", "info", "log level")
	rootCmd.Flags().StringP("src_path", "i", "./", "source path")
	rootCmd.Flags().StringP("dest_path", "o", "./out", "destination path")

	// Bind flags to viper
	//TODO: add more flags
}
