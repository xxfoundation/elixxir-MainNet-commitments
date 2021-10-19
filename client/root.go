///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package client

import (
	"fmt"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"os"
)

var logPath, keyPath, idfPath, contractPath, wallet string

// ExecuteServer adds all child commands to the root command and sets flags
// appropriately.  This is called by server.main(). It only needs to
// happen once to the rootCmd.
func ExecuteClient() {
	if err := clientCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// serverCmd starts a server for mainnet-commitments
var clientCmd = &cobra.Command{
	Use:   "mainnet-commitments-client",
	Short: "Main command for mainnet-commitments client",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		initLog()

		err := SignAndTransmit(keyPath, idfPath, contractPath, wallet, "https://35.81.172.71:11420", "/Users/jonahhusson/gitlab.com/elixxir/deployment/scripts/commitments/commitments.release.elixxir.io.crt")
		if err != nil {
			jww.ERROR.Println(err)
		}
	},
}

// init is the initialization function for Cobra which defines commands
// and flags.
func init() {
	// NOTE: The point of init() is to be declarative.
	// There is one init in each sub command. Do not put variable declarations
	// here, and ensure all the Flags are of the *P variety, unless there's a
	// very good reason not to have them as local params to sub command."

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	clientCmd.Flags().StringVarP(&keyPath, "keyPath", "k",
		"", "Sets a custom key file path")
	clientCmd.Flags().StringVarP(&idfPath, "idfPath", "i",
		"", "Sets a custom id file path")
	clientCmd.Flags().StringVarP(&contractPath, "contract", "c",
		"", "Sets a custom contract file path")
	clientCmd.Flags().StringVarP(&wallet, "wallet", "w",
		"", "Sets a custom wallet")
}

// initLog initializes logging thresholds and the log path.
func initLog() {
	vipLogLevel := 2

	// Check the level of logs to display
	if vipLogLevel > 1 {
		// Set the GRPC log level
		err := os.Setenv("GRPC_GO_LOG_SEVERITY_LEVEL", "info")
		if err != nil {
			jww.ERROR.Printf("Could not set GRPC_GO_LOG_SEVERITY_LEVEL: %+v", err)
		}

		err = os.Setenv("GRPC_GO_LOG_VERBOSITY_LEVEL", "99")
		if err != nil {
			jww.ERROR.Printf("Could not set GRPC_GO_LOG_VERBOSITY_LEVEL: %+v", err)
		}
		// Turn on trace logs
		jww.SetLogThreshold(jww.LevelTrace)
		jww.SetStdoutThreshold(jww.LevelTrace)
	} else if vipLogLevel == 1 {
		// Turn on debugging logs
		jww.SetLogThreshold(jww.LevelDebug)
		jww.SetStdoutThreshold(jww.LevelDebug)
	} else {
		// Turn on info logs
		jww.SetLogThreshold(jww.LevelInfo)
		jww.SetStdoutThreshold(jww.LevelInfo)
	}

	logPath = "/tmp/client.log" //viper.GetString("log")

	logFile, err := os.OpenFile(logPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644)
	if err != nil {
		fmt.Printf("Could not open log file %s!\n", logPath)
	} else {
		jww.SetLogOutput(logFile)
	}
}
