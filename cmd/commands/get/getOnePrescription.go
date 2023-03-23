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

var getPrescCmd = &cobra.Command{
	Use:   "getpresc",
	Short: "get prescription by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		prescID := args[0]
		url := fmt.Sprintf("http://localhost:8080/prescs/%s", prescID)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("Failed to get prescription: %v", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read response body: %v", err)
		}

		err = json.Unmarshal(body, &model.P)
		if err != nil {
			log.Fatalf("Failed to unmarshal prescription: %v", err)
		}
		fmt.Printf("Prescription: %+v\n", model.P)
	},
}

func init() {
	cmd.RootCmd.AddCommand(getPrescCmd)
}
