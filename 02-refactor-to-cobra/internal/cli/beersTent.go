package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CobraStoreFn func(cmd *cobra.Command, args []string)

var stores = map[string]string{
	"01DC9ZAPGKEQJS4P4A48EG3P43": "Mercadona",
	"01DC9ZB23EW0J0ARAER09SJDKC": "Carrefour",
	"01DC9ZB89V1PQD977ZE6QXSQHH": "Alcampo",
}

const storeFlag = "id"

// InitStoreCmd initialize stores command
func InitStoreCmd() *cobra.Command {
	storeCmd := &cobra.Command{
		Use:   "store",
		Short: "Print data about beer stores",
		Run:   runStoreFn(),
	}

	storeCmd.Flags().StringP(storeFlag, "i", "", "id of the beer store")

	return storeCmd
}

func runStoreFn() CobraStoreFn {
	return func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString(storeFlag)

		if id != "" {
			fmt.Println(stores[id])
		} else {
			fmt.Println(stores)
		}
	}
}
