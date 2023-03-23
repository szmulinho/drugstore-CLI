package get

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szmulinho/drugAppCli/cmd"
	"github.com/szmulinho/drugAppCli/internal/model"
	"io/ioutil"
	"log"
	"net/http"
)

var getDrugCmd = &cobra.Command{
	Use:   "getdrug",
	Short: "Get drug by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		drugID := args[0]
		url := fmt.Sprintf("http://localhost:8081/drugs/%s", drugID)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("Failed to get drug: %v", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read response body: %v", err)
		}

		err = json.Unmarshal(body, &model.D)
		if err != nil {
			log.Fatalf("Failed to unmarshal drug: %v", err)
		}
		fmt.Printf("Drug: %+v\n", model.D)
	},
}

func init() {
	cmd.RootCmd.AddCommand(getDrugCmd)
}
