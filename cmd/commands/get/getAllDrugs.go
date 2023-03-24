package get

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szmulinho/drugAppCli/internal/model"
	"net/http"
)

var GetAllDrugsCmd = &cobra.Command{
	Use:   "getall",
	Short: "get all drugs",
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("http://localhost:8081/drugs")
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&model.Drugs)
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		fmt.Println("List of drugs:")

		for _, drug := range model.Drugs {
			fmt.Printf("ID: %s\n", drug.DrugID)
			fmt.Printf("Name: %s\n", drug.Name)
			fmt.Printf("Type: %s\n", drug.Type)
			fmt.Printf("Price: %s\n\n", drug.Price)
		}
	},
}
