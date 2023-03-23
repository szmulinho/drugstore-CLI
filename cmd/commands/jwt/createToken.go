package cmd

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cobra"
)

var createTokenCmd = &cobra.Command{
	Use:   "jwt",
	Short: "Create a JWT token",
	Long:  "Create a new JWT token and print it to the console",
	Run: func(cmd *cobra.Command, args []string) {
		claims := jwt.MapClaims{}
		claims["username"] = "exampleUser"
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		secret := []byte("mySecret")
		tokenString, err := token.SignedString(secret)
		if err != nil {
			fmt.Println("Error creating token:", err)
			return
		}

		fmt.Println("JWT token:", tokenString)
	},
}

func init() {
	rootCmd.AddCommand(createTokenCmd)
}
