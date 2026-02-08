// Code generated — DO NOT EDIT.

package arc_fuse_escrow

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rpc"
	"google.golang.org/protobuf/types/known/emptypb"

	pb2 "github.com/smartcontractkit/chainlink-protos/cre/go/sdk"
	"github.com/smartcontractkit/chainlink-protos/cre/go/values/pb"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm"
	"github.com/smartcontractkit/cre-sdk-go/capabilities/blockchain/evm/bindings"
	"github.com/smartcontractkit/cre-sdk-go/cre"
)

var (
	_ = bytes.Equal
	_ = errors.New
	_ = fmt.Sprintf
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
	_ = emptypb.Empty{}
	_ = pb.NewBigIntFromInt
	_ = pb2.AggregationType_AGGREGATION_TYPE_COMMON_PREFIX
	_ = bindings.FilterOptions{}
	_ = evm.FilterLogTriggerRequest{}
	_ = cre.ResponseBufferTooSmall
	_ = rpc.API{}
	_ = json.Unmarshal
	_ = reflect.Bool
)

var ArcFuseEscrowMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"Release\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"requirements\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"sellerContact\",\"type\":\"string\"}],\"name\":\"ContractRequirementsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"deliveryHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DeliverySubmitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"requirements\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sellerContact\",\"type\":\"string\"}],\"name\":\"setContractRequirements\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"deliveryHash\",\"type\":\"string\"}],\"name\":\"submitDelivery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"getDepositors\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structArcFuseEscrow.DepositorInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"getDepositorCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// Structs
type DepositorInfo struct {
	Depositor common.Address
	Amount    *big.Int
}

// Contract Method Inputs
type DepositInput struct {
	Seller common.Address
}

type DepositsInput struct {
	Arg0 common.Address
}

type GetDepositorCountInput struct {
	Seller common.Address
}

type GetDepositorsInput struct {
	Seller common.Address
}

type RefundInput struct {
	Seller common.Address
	Reason string
}

type ReleaseInput struct {
	Seller        common.Address
	TransactionId [32]byte
}

type SetContractRequirementsInput struct {
	Seller        common.Address
	Requirements  string
	SellerContact string
}

type SubmitDeliveryInput struct {
	DeliveryHash string
}

// Contract Method Outputs

// Errors

// Events
// The <Event>Topics struct should be used as a filter (for log triggers).
// Note: It is only possible to filter on indexed fields.
// Indexed (string and bytes) fields will be of type common.Hash.
// They need to he (crypto.Keccak256) hashed and passed in.
// Indexed (tuple/slice/array) fields can be passed in as is, the Encode<Event>Topics function will handle the hashing.
//
// The <Event>Decoded struct will be the result of calling decode (Adapt) on the log trigger result.
// Indexed dynamic type fields will be of type common.Hash.

type ContractRequirementsSetTopics struct {
	Seller common.Address
}

type ContractRequirementsSetDecoded struct {
	Seller        common.Address
	Requirements  string
	SellerContact string
}

type DeliverySubmittedTopics struct {
	Seller common.Address
}

type DeliverySubmittedDecoded struct {
	Seller       common.Address
	DeliveryHash string
	Timestamp    *big.Int
}

type DepositTopics struct {
	Seller    common.Address
	Depositor common.Address
}

type DepositDecoded struct {
	Seller    common.Address
	Depositor common.Address
	Amount    *big.Int
}

type ReleaseTopics struct {
	Seller        common.Address
	TransactionId [32]byte
}

type ReleaseDecoded struct {
	Seller        common.Address
	Amount        *big.Int
	TransactionId [32]byte
}

type WithdrawalTopics struct {
	Depositor common.Address
}

type WithdrawalDecoded struct {
	Depositor common.Address
	Amount    *big.Int
	Reason    string
}

// Main Binding Type for ArcFuseEscrow
type ArcFuseEscrow struct {
	Address common.Address
	Options *bindings.ContractInitOptions
	ABI     *abi.ABI
	client  *evm.Client
	Codec   ArcFuseEscrowCodec
}

type ArcFuseEscrowCodec interface {
	EncodeDepositMethodCall(in DepositInput) ([]byte, error)
	EncodeDepositsMethodCall(in DepositsInput) ([]byte, error)
	DecodeDepositsMethodOutput(data []byte) (*big.Int, error)
	EncodeGetDepositorCountMethodCall(in GetDepositorCountInput) ([]byte, error)
	DecodeGetDepositorCountMethodOutput(data []byte) (*big.Int, error)
	EncodeGetDepositorsMethodCall(in GetDepositorsInput) ([]byte, error)
	DecodeGetDepositorsMethodOutput(data []byte) ([]DepositorInfo, error)
	EncodeOracleMethodCall() ([]byte, error)
	DecodeOracleMethodOutput(data []byte) (common.Address, error)
	EncodeOwnerMethodCall() ([]byte, error)
	DecodeOwnerMethodOutput(data []byte) (common.Address, error)
	EncodeRefundMethodCall(in RefundInput) ([]byte, error)
	EncodeReleaseMethodCall(in ReleaseInput) ([]byte, error)
	EncodeSetContractRequirementsMethodCall(in SetContractRequirementsInput) ([]byte, error)
	EncodeSubmitDeliveryMethodCall(in SubmitDeliveryInput) ([]byte, error)
	EncodeDepositorInfoStruct(in DepositorInfo) ([]byte, error)
	ContractRequirementsSetLogHash() []byte
	EncodeContractRequirementsSetTopics(evt abi.Event, values []ContractRequirementsSetTopics) ([]*evm.TopicValues, error)
	DecodeContractRequirementsSet(log *evm.Log) (*ContractRequirementsSetDecoded, error)
	DeliverySubmittedLogHash() []byte
	EncodeDeliverySubmittedTopics(evt abi.Event, values []DeliverySubmittedTopics) ([]*evm.TopicValues, error)
	DecodeDeliverySubmitted(log *evm.Log) (*DeliverySubmittedDecoded, error)
	DepositLogHash() []byte
	EncodeDepositTopics(evt abi.Event, values []DepositTopics) ([]*evm.TopicValues, error)
	DecodeDeposit(log *evm.Log) (*DepositDecoded, error)
	ReleaseLogHash() []byte
	EncodeReleaseTopics(evt abi.Event, values []ReleaseTopics) ([]*evm.TopicValues, error)
	DecodeRelease(log *evm.Log) (*ReleaseDecoded, error)
	WithdrawalLogHash() []byte
	EncodeWithdrawalTopics(evt abi.Event, values []WithdrawalTopics) ([]*evm.TopicValues, error)
	DecodeWithdrawal(log *evm.Log) (*WithdrawalDecoded, error)
}

func NewArcFuseEscrow(
	client *evm.Client,
	address common.Address,
	options *bindings.ContractInitOptions,
) (*ArcFuseEscrow, error) {
	parsed, err := abi.JSON(strings.NewReader(ArcFuseEscrowMetaData.ABI))
	if err != nil {
		return nil, err
	}
	codec, err := NewCodec()
	if err != nil {
		return nil, err
	}
	return &ArcFuseEscrow{
		Address: address,
		Options: options,
		ABI:     &parsed,
		client:  client,
		Codec:   codec,
	}, nil
}

type Codec struct {
	abi *abi.ABI
}

func NewCodec() (ArcFuseEscrowCodec, error) {
	parsed, err := abi.JSON(strings.NewReader(ArcFuseEscrowMetaData.ABI))
	if err != nil {
		return nil, err
	}
	return &Codec{abi: &parsed}, nil
}

func (c *Codec) EncodeDepositMethodCall(in DepositInput) ([]byte, error) {
	return c.abi.Pack("deposit", in.Seller)
}

func (c *Codec) EncodeDepositsMethodCall(in DepositsInput) ([]byte, error) {
	return c.abi.Pack("deposits", in.Arg0)
}

func (c *Codec) DecodeDepositsMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["deposits"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetDepositorCountMethodCall(in GetDepositorCountInput) ([]byte, error) {
	return c.abi.Pack("getDepositorCount", in.Seller)
}

func (c *Codec) DecodeGetDepositorCountMethodOutput(data []byte) (*big.Int, error) {
	vals, err := c.abi.Methods["getDepositorCount"].Outputs.Unpack(data)
	if err != nil {
		return *new(*big.Int), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(*big.Int), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result *big.Int
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(*big.Int), fmt.Errorf("failed to unmarshal to *big.Int: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeGetDepositorsMethodCall(in GetDepositorsInput) ([]byte, error) {
	return c.abi.Pack("getDepositors", in.Seller)
}

func (c *Codec) DecodeGetDepositorsMethodOutput(data []byte) ([]DepositorInfo, error) {
	vals, err := c.abi.Methods["getDepositors"].Outputs.Unpack(data)
	if err != nil {
		return *new([]DepositorInfo), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new([]DepositorInfo), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result []DepositorInfo
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new([]DepositorInfo), fmt.Errorf("failed to unmarshal to []DepositorInfo: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeOracleMethodCall() ([]byte, error) {
	return c.abi.Pack("oracle")
}

func (c *Codec) DecodeOracleMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["oracle"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeOwnerMethodCall() ([]byte, error) {
	return c.abi.Pack("owner")
}

func (c *Codec) DecodeOwnerMethodOutput(data []byte) (common.Address, error) {
	vals, err := c.abi.Methods["owner"].Outputs.Unpack(data)
	if err != nil {
		return *new(common.Address), err
	}
	jsonData, err := json.Marshal(vals[0])
	if err != nil {
		return *new(common.Address), fmt.Errorf("failed to marshal ABI result: %w", err)
	}

	var result common.Address
	if err := json.Unmarshal(jsonData, &result); err != nil {
		return *new(common.Address), fmt.Errorf("failed to unmarshal to common.Address: %w", err)
	}

	return result, nil
}

func (c *Codec) EncodeRefundMethodCall(in RefundInput) ([]byte, error) {
	return c.abi.Pack("refund", in.Seller, in.Reason)
}

func (c *Codec) EncodeReleaseMethodCall(in ReleaseInput) ([]byte, error) {
	return c.abi.Pack("release", in.Seller, in.TransactionId)
}

func (c *Codec) EncodeSetContractRequirementsMethodCall(in SetContractRequirementsInput) ([]byte, error) {
	return c.abi.Pack("setContractRequirements", in.Seller, in.Requirements, in.SellerContact)
}

func (c *Codec) EncodeSubmitDeliveryMethodCall(in SubmitDeliveryInput) ([]byte, error) {
	return c.abi.Pack("submitDelivery", in.DeliveryHash)
}

func (c *Codec) EncodeDepositorInfoStruct(in DepositorInfo) ([]byte, error) {
	tupleType, err := abi.NewType(
		"tuple", "",
		[]abi.ArgumentMarshaling{
			{Name: "depositor", Type: "address"},
			{Name: "amount", Type: "uint256"},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tuple type for DepositorInfo: %w", err)
	}
	args := abi.Arguments{
		{Name: "depositorInfo", Type: tupleType},
	}

	return args.Pack(in)
}

func (c *Codec) ContractRequirementsSetLogHash() []byte {
	return c.abi.Events["ContractRequirementsSet"].ID.Bytes()
}

func (c *Codec) EncodeContractRequirementsSetTopics(
	evt abi.Event,
	values []ContractRequirementsSetTopics,
) ([]*evm.TopicValues, error) {
	var sellerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Seller).IsZero() {
			sellerRule = append(sellerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.Seller)
		if err != nil {
			return nil, err
		}
		sellerRule = append(sellerRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		sellerRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeContractRequirementsSet decodes a log into a ContractRequirementsSet struct.
func (c *Codec) DecodeContractRequirementsSet(log *evm.Log) (*ContractRequirementsSetDecoded, error) {
	event := new(ContractRequirementsSetDecoded)
	if err := c.abi.UnpackIntoInterface(event, "ContractRequirementsSet", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["ContractRequirementsSet"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) DeliverySubmittedLogHash() []byte {
	return c.abi.Events["DeliverySubmitted"].ID.Bytes()
}

func (c *Codec) EncodeDeliverySubmittedTopics(
	evt abi.Event,
	values []DeliverySubmittedTopics,
) ([]*evm.TopicValues, error) {
	var sellerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Seller).IsZero() {
			sellerRule = append(sellerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.Seller)
		if err != nil {
			return nil, err
		}
		sellerRule = append(sellerRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		sellerRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeDeliverySubmitted decodes a log into a DeliverySubmitted struct.
func (c *Codec) DecodeDeliverySubmitted(log *evm.Log) (*DeliverySubmittedDecoded, error) {
	event := new(DeliverySubmittedDecoded)
	if err := c.abi.UnpackIntoInterface(event, "DeliverySubmitted", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["DeliverySubmitted"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) DepositLogHash() []byte {
	return c.abi.Events["Deposit"].ID.Bytes()
}

func (c *Codec) EncodeDepositTopics(
	evt abi.Event,
	values []DepositTopics,
) ([]*evm.TopicValues, error) {
	var sellerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Seller).IsZero() {
			sellerRule = append(sellerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.Seller)
		if err != nil {
			return nil, err
		}
		sellerRule = append(sellerRule, fieldVal)
	}
	var depositorRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Depositor).IsZero() {
			depositorRule = append(depositorRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[1], v.Depositor)
		if err != nil {
			return nil, err
		}
		depositorRule = append(depositorRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		sellerRule,
		depositorRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeDeposit decodes a log into a Deposit struct.
func (c *Codec) DecodeDeposit(log *evm.Log) (*DepositDecoded, error) {
	event := new(DepositDecoded)
	if err := c.abi.UnpackIntoInterface(event, "Deposit", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["Deposit"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) ReleaseLogHash() []byte {
	return c.abi.Events["Release"].ID.Bytes()
}

func (c *Codec) EncodeReleaseTopics(
	evt abi.Event,
	values []ReleaseTopics,
) ([]*evm.TopicValues, error) {
	var sellerRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Seller).IsZero() {
			sellerRule = append(sellerRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.Seller)
		if err != nil {
			return nil, err
		}
		sellerRule = append(sellerRule, fieldVal)
	}
	var transactionIdRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.TransactionId).IsZero() {
			transactionIdRule = append(transactionIdRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[2], v.TransactionId)
		if err != nil {
			return nil, err
		}
		transactionIdRule = append(transactionIdRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		sellerRule,
		transactionIdRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeRelease decodes a log into a Release struct.
func (c *Codec) DecodeRelease(log *evm.Log) (*ReleaseDecoded, error) {
	event := new(ReleaseDecoded)
	if err := c.abi.UnpackIntoInterface(event, "Release", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["Release"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Codec) WithdrawalLogHash() []byte {
	return c.abi.Events["Withdrawal"].ID.Bytes()
}

func (c *Codec) EncodeWithdrawalTopics(
	evt abi.Event,
	values []WithdrawalTopics,
) ([]*evm.TopicValues, error) {
	var depositorRule []interface{}
	for _, v := range values {
		if reflect.ValueOf(v.Depositor).IsZero() {
			depositorRule = append(depositorRule, common.Hash{})
			continue
		}
		fieldVal, err := bindings.PrepareTopicArg(evt.Inputs[0], v.Depositor)
		if err != nil {
			return nil, err
		}
		depositorRule = append(depositorRule, fieldVal)
	}

	rawTopics, err := abi.MakeTopics(
		depositorRule,
	)
	if err != nil {
		return nil, err
	}

	return bindings.PrepareTopics(rawTopics, evt.ID.Bytes()), nil
}

// DecodeWithdrawal decodes a log into a Withdrawal struct.
func (c *Codec) DecodeWithdrawal(log *evm.Log) (*WithdrawalDecoded, error) {
	event := new(WithdrawalDecoded)
	if err := c.abi.UnpackIntoInterface(event, "Withdrawal", log.Data); err != nil {
		return nil, err
	}
	var indexed abi.Arguments
	for _, arg := range c.abi.Events["Withdrawal"].Inputs {
		if arg.Indexed {
			if arg.Type.T == abi.TupleTy {
				// abigen throws on tuple, so converting to bytes to
				// receive back the common.Hash as is instead of error
				arg.Type.T = abi.BytesTy
			}
			indexed = append(indexed, arg)
		}
	}
	// Convert [][]byte → []common.Hash
	topics := make([]common.Hash, len(log.Topics))
	for i, t := range log.Topics {
		topics[i] = common.BytesToHash(t)
	}

	if err := abi.ParseTopics(event, indexed, topics[1:]); err != nil {
		return nil, err
	}
	return event, nil
}

func (c ArcFuseEscrow) Deposits(
	runtime cre.Runtime,
	args DepositsInput,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeDepositsMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[*big.Int](*new(*big.Int), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (*big.Int, error) {
		return c.Codec.DecodeDepositsMethodOutput(response.Data)
	})

}

func (c ArcFuseEscrow) GetDepositorCount(
	runtime cre.Runtime,
	args GetDepositorCountInput,
	blockNumber *big.Int,
) cre.Promise[*big.Int] {
	calldata, err := c.Codec.EncodeGetDepositorCountMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[*big.Int](*new(*big.Int), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (*big.Int, error) {
		return c.Codec.DecodeGetDepositorCountMethodOutput(response.Data)
	})

}

func (c ArcFuseEscrow) GetDepositors(
	runtime cre.Runtime,
	args GetDepositorsInput,
	blockNumber *big.Int,
) cre.Promise[[]DepositorInfo] {
	calldata, err := c.Codec.EncodeGetDepositorsMethodCall(args)
	if err != nil {
		return cre.PromiseFromResult[[]DepositorInfo](*new([]DepositorInfo), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) ([]DepositorInfo, error) {
		return c.Codec.DecodeGetDepositorsMethodOutput(response.Data)
	})

}

func (c ArcFuseEscrow) Oracle(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeOracleMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeOracleMethodOutput(response.Data)
	})

}

func (c ArcFuseEscrow) Owner(
	runtime cre.Runtime,
	blockNumber *big.Int,
) cre.Promise[common.Address] {
	calldata, err := c.Codec.EncodeOwnerMethodCall()
	if err != nil {
		return cre.PromiseFromResult[common.Address](*new(common.Address), err)
	}

	var bn cre.Promise[*pb.BigInt]
	if blockNumber == nil {
		promise := c.client.HeaderByNumber(runtime, &evm.HeaderByNumberRequest{
			BlockNumber: bindings.FinalizedBlockNumber,
		})

		bn = cre.Then(promise, func(finalizedBlock *evm.HeaderByNumberReply) (*pb.BigInt, error) {
			if finalizedBlock == nil || finalizedBlock.Header == nil {
				return nil, errors.New("failed to get finalized block header")
			}
			return finalizedBlock.Header.BlockNumber, nil
		})
	} else {
		bn = cre.PromiseFromResult(pb.NewBigIntFromInt(blockNumber), nil)
	}

	promise := cre.ThenPromise(bn, func(bn *pb.BigInt) cre.Promise[*evm.CallContractReply] {
		return c.client.CallContract(runtime, &evm.CallContractRequest{
			Call:        &evm.CallMsg{To: c.Address.Bytes(), Data: calldata},
			BlockNumber: bn,
		})
	})
	return cre.Then(promise, func(response *evm.CallContractReply) (common.Address, error) {
		return c.Codec.DecodeOwnerMethodOutput(response.Data)
	})

}

func (c ArcFuseEscrow) WriteReportFromDepositorInfo(
	runtime cre.Runtime,
	input DepositorInfo,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	encoded, err := c.Codec.EncodeDepositorInfoStruct(input)
	if err != nil {
		return cre.PromiseFromResult[*evm.WriteReportReply](nil, err)
	}
	promise := runtime.GenerateReport(&pb2.ReportRequest{
		EncodedPayload: encoded,
		EncoderName:    "evm",
		SigningAlgo:    "ecdsa",
		HashingAlgo:    "keccak256",
	})

	return cre.ThenPromise(promise, func(report *cre.Report) cre.Promise[*evm.WriteReportReply] {
		return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
			Receiver:  c.Address.Bytes(),
			Report:    report,
			GasConfig: gasConfig,
		})
	})
}

func (c ArcFuseEscrow) WriteReport(
	runtime cre.Runtime,
	report *cre.Report,
	gasConfig *evm.GasConfig,
) cre.Promise[*evm.WriteReportReply] {
	return c.client.WriteReport(runtime, &evm.WriteCreReportRequest{
		Receiver:  c.Address.Bytes(),
		Report:    report,
		GasConfig: gasConfig,
	})
}

func (c *ArcFuseEscrow) UnpackError(data []byte) (any, error) {
	switch common.Bytes2Hex(data[:4]) {
	default:
		return nil, errors.New("unknown error selector")
	}
}

// ContractRequirementsSetTrigger wraps the raw log trigger and provides decoded ContractRequirementsSetDecoded data
type ContractRequirementsSetTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                // Embed the raw trigger
	contract                        *ArcFuseEscrow // Keep reference for decoding
}

// Adapt method that decodes the log into ContractRequirementsSet data
func (t *ContractRequirementsSetTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[ContractRequirementsSetDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeContractRequirementsSet(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode ContractRequirementsSet log: %w", err)
	}

	return &bindings.DecodedLog[ContractRequirementsSetDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *ArcFuseEscrow) LogTriggerContractRequirementsSetLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []ContractRequirementsSetTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[ContractRequirementsSetDecoded]], error) {
	event := c.ABI.Events["ContractRequirementsSet"]
	topics, err := c.Codec.EncodeContractRequirementsSetTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for ContractRequirementsSet: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &ContractRequirementsSetTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *ArcFuseEscrow) FilterLogsContractRequirementsSet(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.ContractRequirementsSetLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// DeliverySubmittedTrigger wraps the raw log trigger and provides decoded DeliverySubmittedDecoded data
type DeliverySubmittedTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                // Embed the raw trigger
	contract                        *ArcFuseEscrow // Keep reference for decoding
}

// Adapt method that decodes the log into DeliverySubmitted data
func (t *DeliverySubmittedTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[DeliverySubmittedDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeDeliverySubmitted(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode DeliverySubmitted log: %w", err)
	}

	return &bindings.DecodedLog[DeliverySubmittedDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *ArcFuseEscrow) LogTriggerDeliverySubmittedLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []DeliverySubmittedTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[DeliverySubmittedDecoded]], error) {
	event := c.ABI.Events["DeliverySubmitted"]
	topics, err := c.Codec.EncodeDeliverySubmittedTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for DeliverySubmitted: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &DeliverySubmittedTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *ArcFuseEscrow) FilterLogsDeliverySubmitted(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.DeliverySubmittedLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// DepositTrigger wraps the raw log trigger and provides decoded DepositDecoded data
type DepositTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                // Embed the raw trigger
	contract                        *ArcFuseEscrow // Keep reference for decoding
}

// Adapt method that decodes the log into Deposit data
func (t *DepositTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[DepositDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeDeposit(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Deposit log: %w", err)
	}

	return &bindings.DecodedLog[DepositDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *ArcFuseEscrow) LogTriggerDepositLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []DepositTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[DepositDecoded]], error) {
	event := c.ABI.Events["Deposit"]
	topics, err := c.Codec.EncodeDepositTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for Deposit: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &DepositTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *ArcFuseEscrow) FilterLogsDeposit(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.DepositLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// ReleaseTrigger wraps the raw log trigger and provides decoded ReleaseDecoded data
type ReleaseTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                // Embed the raw trigger
	contract                        *ArcFuseEscrow // Keep reference for decoding
}

// Adapt method that decodes the log into Release data
func (t *ReleaseTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[ReleaseDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeRelease(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Release log: %w", err)
	}

	return &bindings.DecodedLog[ReleaseDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *ArcFuseEscrow) LogTriggerReleaseLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []ReleaseTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[ReleaseDecoded]], error) {
	event := c.ABI.Events["Release"]
	topics, err := c.Codec.EncodeReleaseTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for Release: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &ReleaseTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *ArcFuseEscrow) FilterLogsRelease(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.ReleaseLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}

// WithdrawalTrigger wraps the raw log trigger and provides decoded WithdrawalDecoded data
type WithdrawalTrigger struct {
	cre.Trigger[*evm.Log, *evm.Log]                // Embed the raw trigger
	contract                        *ArcFuseEscrow // Keep reference for decoding
}

// Adapt method that decodes the log into Withdrawal data
func (t *WithdrawalTrigger) Adapt(l *evm.Log) (*bindings.DecodedLog[WithdrawalDecoded], error) {
	// Decode the log using the contract's codec
	decoded, err := t.contract.Codec.DecodeWithdrawal(l)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Withdrawal log: %w", err)
	}

	return &bindings.DecodedLog[WithdrawalDecoded]{
		Log:  l,        // Original log
		Data: *decoded, // Decoded data
	}, nil
}

func (c *ArcFuseEscrow) LogTriggerWithdrawalLog(chainSelector uint64, confidence evm.ConfidenceLevel, filters []WithdrawalTopics) (cre.Trigger[*evm.Log, *bindings.DecodedLog[WithdrawalDecoded]], error) {
	event := c.ABI.Events["Withdrawal"]
	topics, err := c.Codec.EncodeWithdrawalTopics(event, filters)
	if err != nil {
		return nil, fmt.Errorf("failed to encode topics for Withdrawal: %w", err)
	}

	rawTrigger := evm.LogTrigger(chainSelector, &evm.FilterLogTriggerRequest{
		Addresses:  [][]byte{c.Address.Bytes()},
		Topics:     topics,
		Confidence: confidence,
	})

	return &WithdrawalTrigger{
		Trigger:  rawTrigger,
		contract: c,
	}, nil
}

func (c *ArcFuseEscrow) FilterLogsWithdrawal(runtime cre.Runtime, options *bindings.FilterOptions) (cre.Promise[*evm.FilterLogsReply], error) {
	if options == nil {
		return nil, errors.New("FilterLogs options are required.")
	}
	return c.client.FilterLogs(runtime, &evm.FilterLogsRequest{
		FilterQuery: &evm.FilterQuery{
			Addresses: [][]byte{c.Address.Bytes()},
			Topics: []*evm.Topics{
				{Topic: [][]byte{c.Codec.WithdrawalLogHash()}},
			},
			BlockHash: options.BlockHash,
			FromBlock: pb.NewBigIntFromInt(options.FromBlock),
			ToBlock:   pb.NewBigIntFromInt(options.ToBlock),
		},
	}), nil
}
