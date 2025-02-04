package rpc

import "errors"

var ErrNotImplemented = errors.New("not implemented")

func tryUnwrapToRPCErr(err error, rpcErrors ...*RPCError) error {
	for _, rpcErr := range rpcErrors {
		if errors.Is(err, rpcErr) {
			return rpcErr
		}
	}

	return err
}

func isErrNoTraceAvailableError(err error) (*RPCError, bool) {
	clientErr, ok := err.(*RPCError)
	if !ok {
		return nil, false
	}
	switch clientErr.code {
	case ErrNoTraceAvailable.code:
		noTraceAvailableError := ErrNoTraceAvailable
		noTraceAvailableError.data = clientErr.data
		return noTraceAvailableError, true
	}
	return nil, false
}

type RPCError struct {
	code    int
	message string
	data    any
}

func (e *RPCError) Error() string {
	return e.message
}

func (e *RPCError) Code() int {
	return e.code
}

func (e *RPCError) Data() any {
	return e.data
}

var (
	ErrFailedToReceiveTxn = &RPCError{
		code:    1,
		message: "Failed to write transaction",
	}
	ErrNoTraceAvailable = &RPCError{
		code:    10,
		message: "No trace available for transaction",
	}
	ErrContractNotFound = &RPCError{
		code:    20,
		message: "Contract not found",
	}
	ErrBlockNotFound = &RPCError{
		code:    24,
		message: "Block not found",
	}
	ErrInvalidTxnHash = &RPCError{
		code:    25,
		message: "Invalid transaction hash",
	}
	ErrInvalidBlockHash = &RPCError{
		code:    26,
		message: "Invalid block hash",
	}
	ErrInvalidTxnIndex = &RPCError{
		code:    27,
		message: "Invalid transaction index in a block",
	}
	ErrClassHashNotFound = &RPCError{
		code:    28,
		message: "Class hash not found",
	}
	ErrHashNotFound = &RPCError{
		code:    29,
		message: "Transaction hash not found",
	}
	ErrPageSizeTooBig = &RPCError{
		code:    31,
		message: "Requested page size is too big",
	}
	ErrNoBlocks = &RPCError{
		code:    32,
		message: "There are no blocks",
	}
	ErrInvalidContinuationToken = &RPCError{
		code:    33,
		message: "The supplied continuation token is invalid or unknown",
	}
	ErrTooManyKeysInFilter = &RPCError{
		code:    34,
		message: "Too many keys provided in a filter",
	}
	ErrContractError = &RPCError{
		code:    40,
		message: "Contract error",
	}
	ErrInvalidContractClass = &RPCError{
		code:    50,
		message: "Invalid contract class",
	}
)
