package version

import (
	"fmt"

	"github.com/CentrifugeInc/go-centrifuge/centrifuge/centerrors"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/code"
	"github.com/Masterminds/semver"
	logging "github.com/ipfs/go-log"
)

var log = logging.Logger("centrifuge-version")

// checkMajorCompatibility ensures that a version string matches the major version of
// the app.
func checkMajorCompatibility(versionString string) (match bool, err error) {
	v, err := semver.NewVersion(versionString)
	if err != nil {
		return false, err
	}
	return v.Major() == GetVersion().Major(), nil
}

// checkVersion checks if the peer node version matches with the current node
func CheckVersion(peerVersion string) bool {
	compatible, err := checkMajorCompatibility(peerVersion)
	if err != nil {
		return false
	}

	return compatible
}

// IncompatibleVersionError for any peer with incompatible versions
func IncompatibleVersionError(nodeVersion string) error {
	return centerrors.New(code.VersionMismatch, fmt.Sprintf("Incompatible version: node version: %s, client version: %s", GetVersion(), nodeVersion))
}
