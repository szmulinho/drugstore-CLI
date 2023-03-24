/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/szmulinho/drugAppCli/cmd/commands/create"
	"github.com/szmulinho/drugAppCli/cmd/commands/get"
	"github.com/szmulinho/drugAppCli/cmd/commands/remove"
	"github.com/szmulinho/drugAppCli/cmd/commands/update"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "drugAppCli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(create.AddDrugCmd)
	rootCmd.AddCommand(create.AddPrescCmd)
	rootCmd.AddCommand(create.CreateTokenCmd)
	rootCmd.AddCommand(get.GetAllDrugsCmd)
	rootCmd.AddCommand(get.GetAllPrescsCmd)
	rootCmd.AddCommand(get.GetDrugCmd)
	rootCmd.AddCommand(get.GetPrescCmd)
	rootCmd.AddCommand(remove.RemoveDrugCmd)
	rootCmd.AddCommand(update.UpdateDrugCmd)
}
