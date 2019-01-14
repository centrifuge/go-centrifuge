package receiver

import (
	"fmt"

	"github.com/centrifuge/go-centrifuge/crypto"
	"github.com/golang/protobuf/proto"

	"github.com/centrifuge/centrifuge-protobufs/gen/go/p2p"
	"github.com/centrifuge/go-centrifuge/centerrors"
	"github.com/centrifuge/go-centrifuge/code"
	"github.com/centrifuge/go-centrifuge/errors"
	"github.com/centrifuge/go-centrifuge/version"
)

// Validator defines method that must be implemented by any validator type.
type Validator interface {
	// Validate validates p2p requests
	Validate(envelope *p2ppb.Envelope) error
}

// ValidatorGroup implements Validator for validating a set of validators.
type ValidatorGroup []Validator

// Validate will execute all group specific atomic validations
func (group ValidatorGroup) Validate(envelope *p2ppb.Envelope) (errs error) {
	for _, v := range group {
		if err := v.Validate(envelope); err != nil {
			errs = errors.AppendError(errs, err)
		}
	}
	return errs
}

// ValidatorFunc implements Validator and can be used as a adaptor for functions
// with specific function signature
type ValidatorFunc func(envelope *p2ppb.Envelope) error

// Validate passes the arguments to the underlying validator
// function and returns the results
func (vf ValidatorFunc) Validate(envelope *p2ppb.Envelope) error {
	return vf(envelope)
}

func versionValidator() Validator {
	return ValidatorFunc(func(envelope *p2ppb.Envelope) error {
		if envelope == nil || envelope.Header == nil {
			return errors.New("nil envelope/header")
		}
		if !version.CheckVersion(envelope.Header.NodeVersion) {
			return version.IncompatibleVersionError(envelope.Header.NodeVersion)
		}
		return nil
	})
}

func networkValidator(networkID uint32) Validator {
	return ValidatorFunc(func(envelope *p2ppb.Envelope) error {
		if envelope == nil || envelope.Header == nil {
			return errors.New("nil envelope/header")
		}
		if networkID != envelope.Header.NetworkIdentifier {
			return incompatibleNetworkError(networkID, envelope.Header.NetworkIdentifier)
		}
		return nil
	})
}

func signatureValidator() Validator {
	return ValidatorFunc(func(envelope *p2ppb.Envelope) error {
		if envelope == nil || envelope.Header == nil {
			return errors.New("nil envelope/header")
		}
		envData := &p2ppb.Envelope{
			Header: &p2ppb.Header{
				NetworkIdentifier: envelope.Header.NetworkIdentifier,
				NodeVersion:       envelope.Header.NodeVersion,
				Type:              envelope.Header.Type,
				SenderId:          envelope.Header.SenderId,
			},
			Body: envelope.Body,
		}

		data, err := proto.Marshal(envData)
		if err != nil {
			return err
		}

		if envelope.Header.Signature == nil {
			return errors.New("signature header missing")
		}

		valid := crypto.VerifyMessage(envelope.Header.Signature.PublicKey, data, envelope.Header.Signature.Signature, crypto.CurveEd25519, false)
		if !valid {
			return errors.New("signature validation failure")
		}
		return nil
	})
}

// HandshakeValidator validates the p2p handshake details
func HandshakeValidator(networkID uint32) ValidatorGroup {
	return ValidatorGroup{
		versionValidator(),
		networkValidator(networkID),
		signatureValidator(),
	}
}

func incompatibleNetworkError(configNetwork uint32, nodeNetwork uint32) error {
	return centerrors.New(code.NetworkMismatch, fmt.Sprintf("Incompatible network id: node network: %d, client network: %d", configNetwork, nodeNetwork))
}
