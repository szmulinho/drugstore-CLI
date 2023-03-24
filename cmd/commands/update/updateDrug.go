package update

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/szmulinho/drugAppCli/internal/model"
	"net/http"

	"github.com/spf13/cobra"
)

var UpdateDrugCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a drug with a specific ID",
	RunE: func(cmd *cobra.Command, args []string) error {

		drug := model.Drug{
			DrugID: model.DrugId,
			Name:   model.Name,
			Type:   model.Type,
			Price:  model.Price,
		}

		url := fmt.Sprintf("http://localhost:8081/drugs/%s", model.DrugId)
		reqBody, err := json.Marshal(drug)
		if err != nil {
			return err
		}

		req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(reqBody))
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to update drug: %s", resp.Status)
		}

		fmt.Println("Drug updated successfully!")
		return nil
	},
}

func init() {
	UpdateDrugCmd.Flags().StringVarP(&model.DrugId, "id", "i", "", "Drug ID to update (required)")
	UpdateDrugCmd.MarkFlagRequired("id")
	UpdateDrugCmd.Flags().StringVarP(&model.Name, "name", "n", "", "New name of the drug")
	UpdateDrugCmd.Flags().StringVarP(&model.Type, "type", "t", "", "New type of the drug")
	UpdateDrugCmd.Flags().StringVarP(&model.Price, "price", "p", "", "New price of the drug")
}
