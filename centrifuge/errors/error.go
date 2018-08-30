package errors

import (
	"fmt"
	"reflect"

	"github.com/CentrifugeInc/centrifuge-protobufs/gen/go/errors"
	"github.com/CentrifugeInc/go-centrifuge/centrifuge/code"
	"github.com/go-errors/errors"
)

const (
	// RequiredField error when required field is empty
	RequiredField = "Required field"

	// NilDocument error when document passed is Nil
	NilDocument = "Nil document"

	// IdentifierReUsed error when same identifier is re-used
	IdentifierReUsed = "Identifier re-used"

	// NilDocumentData error when document data is Nil
	NilDocumentData = "Nil document data"

	// RequirePositiveNumber error when amount or any such is zero or negative
	RequirePositiveNumber = "Require positive number"
)

// errpb is the type alias for errorspb.Error
type errpb errorspb.Error

// Error implements the error interface
// message format: [code]message: [sub errors if any]
func (err *errpb) Error() string {
	if err.Errors == nil || len(err.Errors) == 0 {
		return fmt.Sprintf("[%d]%s", err.Code, err.Message)
	}

	return fmt.Sprintf("[%d]%s: %v", err.Code, err.Message, err.Errors)
}

// New constructs a new error with code and error message
func New(code code.Code, message string) error {
	return NewWithErrors(code, message, nil)
}

// NewWithErrors constructs a new error with code, error message, and errors
func NewWithErrors(c code.Code, message string, errors map[string]string) error {
	if c == code.Ok {
		return nil
	}

	return &errpb{
		Code:    int32(c),
		Message: message,
		Errors:  errors,
	}
}

// P2PError represents p2p error type
type P2PError struct {
	err *errorspb.Error
}

// FromError constructs and returns errorspb.Error
// if bool true, conversion to P2PError successful
// else failed and returns unknown P2PError
func FromError(err error) (*P2PError, bool) {
	if err == nil {
		return &P2PError{err: &errorspb.Error{Code: int32(code.Ok)}}, true
	}

	errpb, ok := err.(*errpb)
	if !ok {
		return &P2PError{err: &errorspb.Error{Code: int32(code.Unknown), Message: err.Error()}}, false
	}

	return &P2PError{err: (*errorspb.Error)(errpb)}, true
}

// Code returns the error code
func (p2pErr *P2PError) Code() code.Code {
	if p2pErr == nil || p2pErr.err == nil {
		return code.Ok
	}

	return code.To(p2pErr.err.Code)
}

// Message returns error message
func (p2pErr *P2PError) Message() string {
	if p2pErr == nil || p2pErr.err == nil {
		return ""
	}

	return p2pErr.err.Message
}

// Errors returns map errors passed
func (p2pErr *P2PError) Errors() map[string]string {
	if p2pErr == nil || p2pErr.err == nil {
		return nil
	}

	return p2pErr.err.Errors
}

// NilError returns error with Type added to message
func NilError(param interface{}) error {
	return errors.Errorf("NIL %v provided", reflect.TypeOf(param))
}

// Wrap appends msg to errpb.Message if it is of type *errpb
// else appends the msg to error through fmt.Errorf
func Wrap(err error, msg string) error {
	if err == nil {
		return fmt.Errorf(msg)
	}

	errpb, ok := err.(*errpb)
	if !ok {
		return fmt.Errorf("%s: %v", msg, err)
	}

	errpb.Message = fmt.Sprintf("%s: %v", msg, errpb.Message)
	return errpb
}
