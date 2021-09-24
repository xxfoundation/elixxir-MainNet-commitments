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

// ExecuteServer adds all child commands to the root command and sets flags
// appropriately.  This is called by main.main(). It only needs to
// happen once to the rootCmd.
func ExecuteServer() {
	if err := serverCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// rootCmd represents the base command when called without any subcommands
var serverCmd = &cobra.Command{
	Use:   "mainnet-commitments-server",
	Short: "Main command for mainnet-commitments server",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		certPath := viper.GetString("certPath")
		keyPath := viper.GetString("keyPath")
		var cert, key []byte
		var err error
		if ep, err := utils.ExpandPath(certPath); err == nil {
			cert, err = utils.ReadFile(ep)
			if err != nil {
				jww.FATAL.Panicf("Failed to read cert file at path %s: %+v", ep, err)
			}
		} else {
			jww.FATAL.Panicf("Failed to expand certificate path: %+v", err)
		}

		if ep, err := utils.ExpandPath(keyPath); err == nil {
			key, err = utils.ReadFile(ep)
			if err != nil {
				jww.FATAL.Panicf("Failed to read key file at path %s: %+v", ep, err)
			}
		} else {
			jww.FATAL.Panicf("Failed to expand key path: %+v", err)
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
			Key:     key,
			Cert:    cert,
			Address: viper.GetString("address"),
			StorageParams: storage.Params{
				Username: viper.GetString("dbUsername"),
				Password: viper.GetString("dbPassword"),
				DBName:   viper.GetString("dbName"),
				Address:  addr,
				Port:     port,
			},
		}
		s, err := server.StartServer(params)
		var stopCh = make(chan bool)
		select {
		case <-stopCh:
			s.Stop()
		}
	},
}
