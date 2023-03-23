package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var deleteDrugCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete drug by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		url := fmt.Sprintf("http://localhost:8081/drugs/%s", id)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			fmt.Printf("Drug with ID %s deleted successfully\n", id)
		} else {
			fmt.Printf("Error deleteding drug with ID %s\n", id, resp.Status)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteDrugCmd)
}
