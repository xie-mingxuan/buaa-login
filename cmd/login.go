package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	login "github.com/wangjq4214/buaa-login"
)

var (
	loginUsername string
	loginPassword string
	retry         int
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login BUAA gateway.",
	Run: func(cmd *cobra.Command, args []string) {
		lm := login.NewLoginManager(login.NewLoginManagerParams{
			Username:           loginUsername,
			Password:           loginPassword,
			N:                  "200",
			AcID:               "67",
			Enc:                "srun_bx1",
			VType:              "1",
			LoginPageURL:       "https://gw.buaa.edu.cn/srun_portal_pc?ac_id=67&theme=buaa",
			GetChallengeApiURL: "https://gw.buaa.edu.cn/cgi-bin/get_challenge",
			LoginApiURL:        "https://gw.buaa.edu.cn/cgi-bin/srun_portal",
		})

		for i := 0; i < retry; i++ {
			err := lm.Login()
			if err != nil {
				fmt.Println("We got an error while login:")
				fmt.Println(err.Error())
			} else {
				os.Exit(0)
			}
		}
	},
}

func init() {
	loginCmd.Flags().StringVarP(&loginUsername, "username", "u", "", "Your buaa gw username.")
	loginCmd.MarkFlagRequired("username")

	loginCmd.Flags().StringVarP(&loginPassword, "password", "p", "", "Your buaa gw password.")
	loginCmd.MarkFlagRequired("password")

	loginCmd.Flags().IntVarP(&retry, "retry", "r", 5, "The login retry times.")

	rootCmd.AddCommand(loginCmd)
}
