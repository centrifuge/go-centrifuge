package cmd

import (
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/config"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/identity"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/keytools/ed25519"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/keytools/secp256k1"
	"github.com/spf13/cobra"
)

var centrifugeIdString string
var purpose string

var createIdentityCmd = &cobra.Command{
	Use:   "createidentity",
	Short: "creates identity with signing key as p2p id against ethereum",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		//cmd requires a config file
		readConfigFile()

		defaultBootstrap()
		identityService := identity.EthereumIdentityService{}
		centrifugeId, err := identity.CentrifugeIdStringToSlice(centrifugeIdString)
		if err != nil {
			panic(err)
		}
		_, confirmations, err := identityService.CreateIdentity(centrifugeId)
		watchIdentity := <-confirmations
		if err != nil {
			panic(err)
		}
		log.Infof("Identity created [%s]", watchIdentity.Identity.GetCentrifugeID())
	},
}

//We should support multiple types of keys to add, at the moment only keyPurpose 1 - PeerID/Signature/Encryption
var addKeyCmd = &cobra.Command{
	Use:   "addkey",
	Short: "add a signing key as p2p id against ethereum",
	Long:  "add a signing key as p2p id against ethereum",
	Run: func(cmd *cobra.Command, args []string) {
		//cmd requires a config file
		readConfigFile()

		defaultBootstrap()
		identityService := identity.EthereumIdentityService{}

		var identityConfig *config.IdentityConfig
		var purposeInt int
		var err error

		switch purpose {
		case "p2p":
			identityConfig, err = ed25519.GetIDConfig()
			purposeInt = identity.KeyPurposeP2p
		case "ethauth":
			identityConfig, err = secp256k1.GetIDConfig()
			purposeInt = identity.KeyPurposeEthMsgAuth
		default:
			panic("Option not supported")
		}

		centId, err := identity.NewCentID(identityConfig.ID)
		if err != nil {
			panic(err)
		}
		id, err := identityService.LookupIdentityForID(centId)
		if err != nil {
			panic(err)
		}

		confirmations, err := id.AddKeyToIdentity(purposeInt, identityConfig.PublicKey)
		if err != nil {
			panic(err)
		}
		watchAddedToIdentity := <-confirmations

		lastKey, errLocal := watchAddedToIdentity.Identity.GetLastKeyForPurpose(purposeInt)
		if errLocal != nil {
			panic(err)
		}
		log.Infof("Key [%v] with type [$s] Added to Identity [%s]", lastKey, purpose, watchAddedToIdentity.Identity)
		return
	},
}

func init() {
	createIdentityCmd.Flags().StringVarP(&centrifugeIdString, "centrifugeid", "i", "", "Centrifuge ID")
	addKeyCmd.Flags().StringVarP(&centrifugeIdString, "centrifugeid", "i", "", "Centrifuge ID")
	addKeyCmd.Flags().StringVarP(&purpose, "purpose", "p", "", "Key Purpose [p2p|ethauth]")
	rootCmd.AddCommand(createIdentityCmd)
	rootCmd.AddCommand(addKeyCmd)
}
