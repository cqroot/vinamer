package cmd

import (
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/cqroot/edname/internal/renamer"
)

var (
	flagAll              bool
	flagConfig           string
	flagDirectory        bool
	flagDirectoryOnly    bool
	flagEditor           string
	flagWorkingDirectory string

	rootCmd = &cobra.Command{
		Use:   "edname",
		Short: "Use your favorite editor to batch rename files and directories.",
		Long: `Use your favorite editor to batch rename files and directories.

Originally designed for vim, but not just vim.

Notice:
1. Do not add or subtract lines.
2. Unchanged lines are ignored.`,
		Run: RunRootCmd,
	}
)

func init() {
	rootCmd.Flags().BoolVarP(&flagAll, "all", "a", false, "do not ignore entries starting with .")
	rootCmd.Flags().StringVarP(&flagConfig, "config", "c", "", "config file. default $HOME/.config/edname/config.toml")
	rootCmd.Flags().BoolVarP(&flagDirectory, "directory", "d", false, "include directory")
	rootCmd.Flags().BoolVarP(&flagDirectoryOnly, "directory-only", "D", false, "rename directory only")
	rootCmd.Flags().StringVarP(&flagEditor, "editor", "e", "", "")
	rootCmd.Flags().StringVarP(&flagWorkingDirectory, "working-directory", "w", "", "")
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func RunRootCmd(cmd *cobra.Command, args []string) {
	if !path.IsAbs(flagWorkingDirectory) {
		cwd, err := os.Getwd()
		cobra.CheckErr(err)

		if flagWorkingDirectory == "" {
			flagWorkingDirectory = cwd
		} else {
			flagWorkingDirectory = path.Join(cwd, flagWorkingDirectory)
		}
	}

	r := renamer.New(
		flagEditor,
		flagWorkingDirectory,
		flagDirectory,
		flagDirectoryOnly,
		flagAll,
	)

	err := r.Execute()
	cobra.CheckErr(err)
}
