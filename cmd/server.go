///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package cmd

import (
	"fmt"
	"git.xx.network/elixxir/mainnet-commitments/server"
	"git.xx.network/elixxir/mainnet-commitments/storage"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
	"gitlab.com/xx_network/primitives/utils"
	"net"
	"os"
)

var cfgFile, logPath string
var validConfig bool

// ExecuteServer adds all child commands to the root command and sets flags
// appropriately.  This is called by server.main(). It only needs to
// happen once to the rootCmd.
func ExecuteServer() {
	if err := serverCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// serverCmd starts a server for mainnet-commitments
var serverCmd = &cobra.Command{
	Use:   "mainnet-commitments-server",
	Short: "Main command for mainnet-commitments server",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		initConfig()
		initLog()

		certPath, err := utils.ExpandPath(viper.GetString("certPath"))
		if err != nil {
			jww.FATAL.Fatalf("Failed to expand cert path: %+v", err)
		}
		keyPath, err := utils.ExpandPath(viper.GetString("keyPath"))
		if err != nil {
			jww.FATAL.Fatalf("Failed to expand key path: %+v", err)
		}
		rawAddr := viper.GetString("dbAddress")
		var addr, port string
		if rawAddr != "" {
			addr, port, err = net.SplitHostPort(rawAddr)
			if err != nil {
				jww.FATAL.Panicf("Unable to get database port from %s: %+v", rawAddr, err)
			}
		}

		params := server.Params{
			KeyPath:  keyPath,
			CertPath: certPath,
			Port:     viper.GetString("port"),
			StorageParams: storage.Params{
				Username: viper.GetString("dbUsername"),
				Password: viper.GetString("dbPassword"),
				DBName:   viper.GetString("dbName"),
				Address:  addr,
				Port:     port,
			},
		}
		err = server.StartServer(params)
		var stopCh = make(chan bool)
		select {
		case <-stopCh:
			return
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
	serverCmd.Flags().StringVarP(&cfgFile, "config", "c",
		"", "Sets a custom config file path")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	//Use default config location if none is passed
	var err error
	validConfig = true
	if cfgFile == "" {
		cfgFile, err = utils.SearchDefaultLocations("commitments.yaml", "xxnetwork")
		if err != nil {
			validConfig = false
			jww.FATAL.Panicf("Failed to find config file: %+v", err)
		}
	} else {
		cfgFile, err = utils.ExpandPath(cfgFile)
		if err != nil {
			validConfig = false
			jww.FATAL.Panicf("Failed to expand config file path: %+v", err)
		}
	}

	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Unable to read config file (%s): %+v", cfgFile, err.Error())
		validConfig = false
	}
}

// initLog initializes logging thresholds and the log path.
func initLog() {
	vipLogLevel := viper.GetUint("logLevel")

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

	logPath = viper.GetString("log")

	logFile, err := os.OpenFile(logPath,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644)
	if err != nil {
		fmt.Printf("Could not open log file %s!\n", logPath)
	} else {
		jww.SetLogOutput(logFile)
	}
}
