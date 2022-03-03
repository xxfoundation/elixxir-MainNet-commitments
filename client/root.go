///////////////////////////////////////////////////////////////////////////////
// Copyright © 2021 xx network SEZC                                          //
//                                                                           //
// Use of this source code is governed by a license that can be found in the //
// LICENSE file                                                              //
///////////////////////////////////////////////////////////////////////////////

package client

import (
	"fmt"
	"git.xx.network/elixxir/mainnet-commitments/utils"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"os"
	"strings"
)

var logPath, keyPath, idfPath, paymentWallet string

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
		address := "https://18.185.229.39:11420"
		commitmentCert := `-----BEGIN CERTIFICATE-----
MIIFtjCCA56gAwIBAgIJAI0i//hCyk8BMA0GCSqGSIb3DQEBCwUAMIGMMQswCQYD
VQQGEwJVUzELMAkGA1UECAwCQ0ExEjAQBgNVBAcMCUNsYXJlbW9udDEQMA4GA1UE
CgwHRWxpeHhpcjEUMBIGA1UECwwLRGV2ZWxvcG1lbnQxEzARBgNVBAMMCmVsaXh4
aXIuaW8xHzAdBgkqhkiG9w0BCQEWEGFkbWluQGVsaXh4aXIuaW8wHhcNMjExMDA2
MTczNDIxWhcNMjMxMDA2MTczNDIxWjCBjDELMAkGA1UEBhMCVVMxCzAJBgNVBAgM
AkNBMRIwEAYDVQQHDAlDbGFyZW1vbnQxEDAOBgNVBAoMB0VsaXh4aXIxFDASBgNV
BAsMC0RldmVsb3BtZW50MRMwEQYDVQQDDAplbGl4eGlyLmlvMR8wHQYJKoZIhvcN
AQkBFhBhZG1pbkBlbGl4eGlyLmlvMIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIIC
CgKCAgEAqLpFWmIfM+RJuCEnCjCMcEImmoQeMWCg/XFk5U9M5/hjwvK6la4nroKl
Gaue48d49gQJy8349LoOLN3zxQNj3wavv9Y1u+VkgExF2NQR0NBqbYAuaUAYK2LD
JTWeR5xXoWrLkMP3L+ZMqSOgcrH9ZvrsR8seLaExSYTp/MyodlBfHG7+RQ7fr92p
bnoU8c4D6l78Mn9okht3m5t+oHtAhNBJegsGbvqbhwO+Wyrs2Hw6Uq0xM3KeeP9Z
UfT0uz3HxfX2cuL/ApjUXXJ9Dli2vm0LmKur9JjRQV1IUeh+fkScEewmPhk0vbkI
BGq+ck5gq3oiTV0CUkL2eF8YEFHtJ9dtAX6mSNJD9Dvf8tWB5GkTb5x7XOR58dc9
TWH9Nd3HEF3tskbNLLdcAvI9LKW4K3lls2XkVnVXmu/4C9LWChC6vmNBjXcchCvn
PKI6ByHiZZ7Cz90Hmto/Cyg1o+26DUKfHoKm14QyPpDeNAz0RKuemBuNYKzREAOX
ItWjrlbSmM+dVhJPOi1EnZxUoUrXPdnG61ec80VbZZX9MQIvmyp7w3+bcE6PnnfZ
Nk7M8n4L4AM46JZ53y1XavdylrzglrQj0IT+Os//XEU3eVqlzi3U57KuWl0zJ/4F
uPUZWmYdzUIikIK+NHiNwHPOzf6CgDqDzadkFVAUZiHfxGomDPcCAwEAAaMZMBcw
FQYDVR0RBA4wDIIKZWxpeHhpci5pbzANBgkqhkiG9w0BAQsFAAOCAgEAHbQMj4oN
NrwZMBEkaW8abnpT8Es9fHScfqAiFHpzjO9+Q5/Y5XtL92MO0+czQ2QhtaRd/mja
9IikKgSBf6x2oMdWYvm/sMVTlvgnr7GznGeS8JWMh73XJH9gwASZEuTTXEkAVy6u
J5KEvZqlU1/O7LRNA/seW/LVPL2UxJzNbTcXnmw4AEARZzXMcdL7gUnmf8/gzxh5
ikCoOOnhSQJPHbb6sDaxdDKRzwphacR03WZnQt3ShVoF+F7ffNmyQ27lEHD/r7HC
aOtEum/vqr8mCvCEasPxuxhI+G1lEXoTyA3AB2HaQSF6sXzsC/aQKgAx66oUP4mk
WV+EKijj6Gb2h9J6THEv6Ym9SUlUacnujdjrcnJHTyTTyvz8Oj63Olwi1TzLRVqs
+KNSk3ZyyhutiiktIZO1jUKQPNkYunxVryPYYmPHnrledLIShF/j+C/5VsRxzd/4
4vyz5CHFo4HaxPbkLdYuVQ9NzOux3eh8wF1nwADXh7RJUAQVNge2s5shrnznZ7TO
XBqrZq75vkMpNxz+YFfu8J9DgZZOehUOBgyA3hlA9FdFFslPuzo6SxzfKaWjSTPZ
xHKP3P00TnJNiOMRn94MY2GdUl8pAi8I89n9jPZfa0ANCpyfHluw+lNUfJNrGvwO
Mu7/deeXg4hfNzQoWdZnBhzgaB05MAbJI6E=
-----END CERTIFICATE-----`

		fmt.Println(utils.Contract)
		fmt.Println("Do you accept the contract as it has been laid out? (y/n)")

		var accept string
		_, err := fmt.Scanln(&accept)
		if err != nil {
			jww.FATAL.Fatalf("Failed to read contract acceptance: %+v", err)
		}
		if !(strings.ToUpper(accept) == "Y" || strings.ToUpper(accept) == "YES") {
			jww.FATAL.Fatalf("You must accept the contract to continue")
		}

		err = SignAndTransmit(keyPath, idfPath, paymentWallet, address, commitmentCert, utils.Contract)
		if err != nil {
			jww.FATAL.Fatalf("Failed to sign & transmit commitment: %+v", err)
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
	clientCmd.Flags().StringVarP(&paymentWallet, "paymentWallet", "n",
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
