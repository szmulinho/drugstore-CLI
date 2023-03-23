package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szmulinho/drugAppCli/internal/model"
	"net/http"
)

var addDrugCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new drug",
	Long:  "Add a new drug to the drug database at localhost:8081/drug",
	Run: func(cmd *cobra.Command, args []string) {
		// read in the drug details
		drugID, _ := cmd.Flags().GetString("drug-id")
		name, _ := cmd.Flags().GetString("name")
		drugType, _ := cmd.Flags().GetString("type")
		price, _ := cmd.Flags().GetString("price")

		// create a new drug struct
		newDrug := model.Drug{
			DrugID: drugID,
			Name:   name,
			Type:   drugType,
			Price:  price,
		}

		// convert the drug struct to JSON
		jsonStr, _ := json.Marshal(newDrug)

		// send a POST request to the server to add the drug
		url := "http://localhost:8081/drug"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// print the response from the server
		fmt.Println("Response Status:", resp.Status)
	},
}

func init() {
	rootCmd.AddCommand(addDrugCmd)
	addDrugCmd.Flags().String("drug-id", "", "Drug ID")
	addDrugCmd.Flags().String("name", "", "Drug Name")
	addDrugCmd.Flags().String("type", "", "Drug Type")
	addDrugCmd.Flags().String("price", "", "Drug Price")
}