package core

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

func RegisterTypes(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*govv1beta1.Content)(nil), &CreditTypeProposal{})
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateClass{}, "regen.core/MsgCreateClass", nil)
	cdc.RegisterConcrete(&MsgCreateProject{}, "regen.core/MsgCreateProject", nil)
	cdc.RegisterConcrete(&MsgCancel{}, "regen.core/MsgCancel", nil)
	cdc.RegisterConcrete(&MsgCreateBatch{}, "regen.core/MsgCreateBatch", nil)
	cdc.RegisterConcrete(&MsgMintBatchCredits{}, "regen.core/MsgMintBatchCredits", nil)
	cdc.RegisterConcrete(&MsgSealBatch{}, "regen.core/MsgSealBatch", nil)
	cdc.RegisterConcrete(&MsgRetire{}, "regen.core/MsgRetire", nil)
	cdc.RegisterConcrete(&MsgSend{}, "regen.core/MsgSend", nil)
	cdc.RegisterConcrete(&MsgUpdateClassAdmin{}, "regen.core/MsgUpdateClassAdmin", nil)
	cdc.RegisterConcrete(&MsgUpdateClassMetadata{}, "regen.core/MsgUpdateClassMetadata", nil)
	cdc.RegisterConcrete(&MsgUpdateClassIssuers{}, "regen.core/MsgUpdateClassIssuers", nil)
	cdc.RegisterConcrete(&MsgUpdateProjectAdmin{}, "regen.core/MsgUpdateProjectAdmin", nil)
	cdc.RegisterConcrete(&MsgUpdateProjectMetadata{}, "regen.core/MsgUpdateProjectMetadata", nil)
	cdc.RegisterConcrete(&CreditTypeProposal{}, "regen.core/CreditTypeProposal", nil)
	cdc.RegisterConcrete(&MsgBridgeReceive{}, "regen.core/MsgBridgeReceive", nil)
	cdc.RegisterConcrete(&MsgAddCreditType{}, "regen.core/MsgAddCreditType", nil)
}

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	RegisterLegacyAminoCodec(legacy.Cdc)
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	RegisterLegacyAminoCodec(authzcodec.Amino)
}
