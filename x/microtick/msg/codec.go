package msg

import (
    "github.com/cosmos/cosmos-sdk/codec"
    sdk "github.com/cosmos/cosmos-sdk/types"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
)

var (
    ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
	_, _, _, _ sdk.Msg = &TxCancelQuote{}, &TxCreateQuote{}, &TxUpdateQuote{}, &TxDepositQuote{}
	_, _, _, _ sdk.Msg = &TxWithdrawQuote{}, &TxMarketTrade{}, &TxPickTrade{}, &TxSettleTrade{}
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.LegacyAmino) {
    cdc.RegisterConcrete(TxCreateQuote{}, "microtick/QuoteCreate", nil)
    cdc.RegisterConcrete(TxCancelQuote{}, "microtick/QuoteCancel", nil)
    cdc.RegisterConcrete(TxUpdateQuote{}, "microtick/QuoteUpdate", nil)
    cdc.RegisterConcrete(TxDepositQuote{}, "microtick/QuoteDeposit", nil)
    cdc.RegisterConcrete(TxWithdrawQuote{}, "microtick/QuoteWithdraw", nil)
    cdc.RegisterConcrete(TxMarketTrade{}, "microtick/TradeMarket", nil)
    cdc.RegisterConcrete(TxPickTrade{}, "mic rotick/TradePick", nil)
    cdc.RegisterConcrete(TxSettleTrade{}, "microtick/TradeSettle", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
  registry.RegisterImplementations((*sdk.Msg)(nil),
    &TxCancelQuote{},
    &TxCreateQuote{},
  )
}
