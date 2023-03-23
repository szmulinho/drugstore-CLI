package get

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szmulinho/drugAppCli/cmd"
	"github.com/szmulinho/drugAppCli/internal/model"
	"net/http"
)

var getAllPrescsCmd = &cobra.Command{
	Use:   "getprescs",
	Short: "get the list of all prescriptions",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://localhost:8080/prescs")
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&model.Prescs)
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		fmt.Println("List of drugs:")

		for _, presc := range model.Prescs {
			fmt.Printf("ID: %s\n", presc.PreId)
			fmt.Printf("Drugs: %s\n", presc.Drugs)
			fmt.Printf("Expiration: %s\n", presc.Expiration)
		}
	},
}

func init() {
	cmd.RootCmd.AddCommand(getAllPrescsCmd)
}
