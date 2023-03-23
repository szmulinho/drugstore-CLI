package create

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szmulinho/drugAppCli/cmd"
	"github.com/szmulinho/drugAppCli/internal/model"
	"net/http"
	"time"
)

var addPrescCmd = &cobra.Command{
	Use:   "addpresc",
	Short: "add a new prescription",
	Long:  "add a new prescription to the prescription database at localhost:8080/presc",
	Run: func(cmd *cobra.Command, args []string) {
		prescId, _ := cmd.Flags().GetString("presc-id")
		drugs, _ := cmd.Flags().GetStringSlice("drugs")
		expiration, _ := cmd.Flags().GetString("expiration")
		expirationTime, _ := time.Parse("06-01-02-15-04-05", expiration)

		newPresc := model.CreatePrescInput{
			PreId:      prescId,
			Drugs:      drugs,
			Expiration: expirationTime,
		}

		jsonStr, _ := json.Marshal(newPresc)

		url := "http://localhost:8080/presc"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		fmt.Println("Response Status:", resp.Status)
	},
}

func init() {
	cmd.RootCmd.AddCommand(addPrescCmd)
	addPrescCmd.Flags().String("presc-id", "", "Prescription ID")
	addPrescCmd.Flags().StringSlice("drugs", []string{}, "Drugs in prescription ")
	addPrescCmd.Flags().String("expiration", "", "Prescription's expiration")
}
