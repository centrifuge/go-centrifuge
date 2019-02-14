// +build integration unit cmd

package testingutils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/centrifuge/go-centrifuge/bootstrap"

	"github.com/centrifuge/go-centrifuge/config"
	logging "github.com/ipfs/go-log"
	"github.com/savaki/jq"
)

var log = logging.Logger("test-setup")
var isRunningOnCI = len(os.Getenv("TRAVIS")) != 0

// StartPOAGeth runs the proof of authority geth for tests
func StartPOAGeth() {
	// don't run if its already running
	if IsPOAGethRunning() {
		return
	}
	projDir := GetProjectDir()
	gethRunScript := path.Join(projDir, "build", "scripts", "docker", "run.sh")
	o, err := exec.Command(gethRunScript, "dev").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", string(o))
}

// RunSmartContractMigrations migrates smart contracts to localgeth
func RunSmartContractMigrations() {
	if isRunningOnCI {
		return
	}
	projDir := GetProjectDir()
	migrationScript := path.Join(projDir, "build", "scripts", "migrate.sh")
	_, err := exec.Command(migrationScript, projDir).Output()
	if err != nil {
		log.Fatal(err)
	}
}

// GetSmartContractAddresses finds migrated smart contract addresses for localgeth
func GetSmartContractAddresses() *config.SmartContractAddresses {
	dat, err := findContractDeployJSON()
	if err != nil {
		panic(err)
	}
	idFactoryAddrOp := getOpForContract(".contracts.IdentityFactory.address")
	anchorRepoAddrOp := getOpForContract(".contracts.AnchorRepository.address")
	payObAddrOp := getOpForContract(".contracts.PaymentObligation.address")
	return &config.SmartContractAddresses{
		IdentityFactoryAddr:   getOpAddr(idFactoryAddrOp, dat),
		AnchorRepositoryAddr:  getOpAddr(anchorRepoAddrOp, dat),
		PaymentObligationAddr: getOpAddr(payObAddrOp, dat),
	}
}

func findContractDeployJSON() ([]byte, error) {
	projDir := GetProjectDir()
	deployJSONFile := path.Join(projDir, "vendor", "github.com", "centrifuge", "centrifuge-ethereum-contracts", "deployments", "localgeth.json")
	dat, err := ioutil.ReadFile(deployJSONFile)
	if err != nil {
		return nil, err
	}
	return dat, nil
}

func getOpAddr(addrOp jq.Op, dat []byte) string {
	addr, err := addrOp.Apply(dat)
	if err != nil {
		panic(err)
	}

	// remove extra quotes inside the string
	addrStr := string(addr)
	if len(addrStr) > 0 && addrStr[0] == '"' {
		addrStr = addrStr[1:]
	}
	if len(addrStr) > 0 && addrStr[len(addrStr)-1] == '"' {
		addrStr = addrStr[:len(addrStr)-1]
	}
	return addrStr
}

func getOpForContract(selector string) jq.Op {
	addrOp, err := jq.Parse(selector)
	if err != nil {
		panic(err)
	}
	return addrOp
}

func GetProjectDir() string {
	gp := os.Getenv("GOPATH")
	projDir := path.Join(gp, "src", "github.com", "centrifuge", "go-centrifuge")
	return projDir
}

func GetBinaryPath() string {
	gp := os.Getenv("GOPATH")
	projDir := path.Join(gp, "bin", "centrifuge")
	return projDir
}

// IsPOAGethRunning checks if POA geth is running in the background
func IsPOAGethRunning() bool {
	cmd := "docker ps -a --filter \"name=geth-node\" --filter \"status=running\" --quiet"
	o, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		panic(err)
	}
	return len(o) != 0
}

// LoadTestConfig loads configuration for integration tests
func LoadTestConfig() config.Configuration {
	// To get the config location, we need to traverse the path to find the `go-centrifuge` folder
	projDir := GetProjectDir()
	c := config.LoadConfiguration(fmt.Sprintf("%s/build/configs/testing_config.yaml", projDir))
	return c
}

// SetupSmartContractAddresses sets up smart contract addresses on provided config
func SetupSmartContractAddresses(cfg config.Configuration, sca *config.SmartContractAddresses) {
	network := cfg.Get("centrifugeNetwork").(string)
	cfg.SetupSmartContractAddresses(network, sca)
	fmt.Printf("contract addresses %+v\n", sca)
}

// BuildIntegrationTestingContext sets up configuration for integration tests
func BuildIntegrationTestingContext() map[string]interface{} {
	projDir := GetProjectDir()
	StartPOAGeth()
	// RunSmartContractMigrations()
	addresses := GetSmartContractAddresses()
	cfg := LoadTestConfig()
	cfg.Set("keys.p2p.publicKey", fmt.Sprintf("%s/build/resources/p2pKey.pub.pem", projDir))
	cfg.Set("keys.p2p.privateKey", fmt.Sprintf("%s/build/resources/p2pKey.key.pem", projDir))
	cfg.Set("keys.signing.publicKey", fmt.Sprintf("%s/build/resources/signingKey.pub.pem", projDir))
	cfg.Set("keys.signing.privateKey", fmt.Sprintf("%s/build/resources/signingKey.key.pem", projDir))
	cfg.Set("keys.ethauth.publicKey", fmt.Sprintf("%s/build/resources/ethauth.pub.pem", projDir))
	cfg.Set("keys.ethauth.privateKey", fmt.Sprintf("%s/build/resources/ethauth.key.pem", projDir))
	SetupSmartContractAddresses(cfg, addresses)
	cm := make(map[string]interface{})
	cm[bootstrap.BootstrappedConfig] = cfg
	return cm
}
