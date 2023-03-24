package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szmulinho/drugAppCli/internal/model"
	"net/http"
)

var AddDrugCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new drug",
	Long:  "Add a new drug to the drug database at localhost:8081/drug",
	Run: func(cmd *cobra.Command, args []string) {
		drugID, _ := cmd.Flags().GetString("drug-id")
		name, _ := cmd.Flags().GetString("name")
		drugType, _ := cmd.Flags().GetString("type")
		price, _ := cmd.Flags().GetString("price")

		newDrug := model.Drug{
			DrugID: drugID,
			Name:   name,
			Type:   drugType,
			Price:  price,
		}

		jsonStr, _ := json.Marshal(newDrug)

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

	AddDrugCmd.Flags().String("drug-id", "", "Drug ID")
	AddDrugCmd.Flags().String("name", "", "Drug Name")
	AddDrugCmd.Flags().String("type", "", "Drug Type")
	AddDrugCmd.Flags().String("price", "", "Drug Price")
}
