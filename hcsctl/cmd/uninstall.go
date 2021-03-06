package cmd

import (
	"hypercloud-storage/hcsctl/pkg/cdi"
	"hypercloud-storage/hcsctl/pkg/rook"

	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Short:   "해당 인벤토리의 hypercloud-storage를 제거합니다.",
	PreRunE: validateInventory,
	Run: func(cmd *cobra.Command, args []string) {
		if isCdiExist(inventoryPath) {
			err := cdi.Delete(inventoryPath)
			if err != nil {
				panic(err)
			}
		}

		err := rook.Delete(inventoryPath)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
