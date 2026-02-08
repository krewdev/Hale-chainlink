// Code generated — DO NOT EDIT.

//go:build !wasip1

package arc_fuse_escrow

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	evmmock "github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/mock"
)

var (
	_ = errors.New
	_ = fmt.Errorf
	_ = big.NewInt
	_ = common.Big1
)

// ArcFuseEscrowMock is a mock implementation of ArcFuseEscrow for testing.
type ArcFuseEscrowMock struct {
	Deposits          func(DepositsInput) (*big.Int, error)
	GetDepositorCount func(GetDepositorCountInput) (*big.Int, error)
	GetDepositors     func(GetDepositorsInput) ([]DepositorInfo, error)
	Oracle            func() (common.Address, error)
	Owner             func() (common.Address, error)
}

// NewArcFuseEscrowMock creates a new ArcFuseEscrowMock for testing.
func NewArcFuseEscrowMock(address common.Address, clientMock *evmmock.ClientCapability) *ArcFuseEscrowMock {
	mock := &ArcFuseEscrowMock{}

	codec, err := NewCodec()
	if err != nil {
		panic("failed to create codec for mock: " + err.Error())
	}

	abi := codec.(*Codec).abi
	_ = abi

	funcMap := map[string]func([]byte) ([]byte, error){
		string(abi.Methods["deposits"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Deposits == nil {
				return nil, errors.New("deposits method not mocked")
			}
			inputs := abi.Methods["deposits"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := DepositsInput{
				Arg0: values[0].(common.Address),
			}

			result, err := mock.Deposits(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["deposits"].Outputs.Pack(result)
		},
		string(abi.Methods["getDepositorCount"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetDepositorCount == nil {
				return nil, errors.New("getDepositorCount method not mocked")
			}
			inputs := abi.Methods["getDepositorCount"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetDepositorCountInput{
				Seller: values[0].(common.Address),
			}

			result, err := mock.GetDepositorCount(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getDepositorCount"].Outputs.Pack(result)
		},
		string(abi.Methods["getDepositors"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.GetDepositors == nil {
				return nil, errors.New("getDepositors method not mocked")
			}
			inputs := abi.Methods["getDepositors"].Inputs

			values, err := inputs.Unpack(payload)
			if err != nil {
				return nil, errors.New("Failed to unpack payload")
			}
			if len(values) != 1 {
				return nil, errors.New("expected 1 input value")
			}

			args := GetDepositorsInput{
				Seller: values[0].(common.Address),
			}

			result, err := mock.GetDepositors(args)
			if err != nil {
				return nil, err
			}
			return abi.Methods["getDepositors"].Outputs.Pack(result)
		},
		string(abi.Methods["oracle"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Oracle == nil {
				return nil, errors.New("oracle method not mocked")
			}
			result, err := mock.Oracle()
			if err != nil {
				return nil, err
			}
			return abi.Methods["oracle"].Outputs.Pack(result)
		},
		string(abi.Methods["owner"].ID[:4]): func(payload []byte) ([]byte, error) {
			if mock.Owner == nil {
				return nil, errors.New("owner method not mocked")
			}
			result, err := mock.Owner()
			if err != nil {
				return nil, err
			}
			return abi.Methods["owner"].Outputs.Pack(result)
		},
	}

	evmmock.AddContractMock(address, clientMock, funcMap, nil)
	return mock
}
