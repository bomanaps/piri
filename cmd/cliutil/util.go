package cliutil

import (
	"fmt"
	"os"
	"path"

	"github.com/labstack/gommon/color"
	"github.com/spf13/cobra"
	"github.com/storacha/go-ucanto/did"

	"github.com/storacha/piri/pkg/build"
)

func PrintHero(id did.DID) {
	fmt.Printf(`
▗▄▄▖ ▄  ▄▄▄ ▄  %s
▐▌ ▐▌▄ █    ▄  %s
▐▛▀▘ █ █    █  %s
▐▌   █      █  %s

🔥 %s
🆔 %s
🚀 Ready!
`,
		color.Green(" ▗"),
		color.Red(" █")+color.Red("▌", color.D),
		color.Red("▗", color.B)+color.Red("█")+color.Red("▘", color.D),
		color.Red("▀")+color.Red("▘", color.D),
		build.Version, id.String())
}

func Mkdirp(dirpath ...string) (string, error) {
	dir := path.Join(dirpath...)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return "", fmt.Errorf("creating directory: %s: %w", dir, err)
	}
	return dir, nil
}

func MustGetManageAPI(cmd *cobra.Command) string {
	val, err := cmd.Flags().GetString("manage-api")
	if err != nil {
		panic(err)
	}
	return val
}
