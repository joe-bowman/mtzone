package msg

import (
    "github.com/cosmos/cosmos-sdk/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
    cdc.RegisterConcrete(TxCreateMarket{}, "microtick/CreateMarket", nil)
    cdc.RegisterConcrete(TxCreateQuote{}, "microtick/CreateQuote", nil)
    cdc.RegisterConcrete(TxCancelQuote{}, "microtick/CancelQuote", nil)
    cdc.RegisterConcrete(TxUpdateQuote{}, "microtick/UpdateQuote", nil)
    cdc.RegisterConcrete(TxDepositQuote{}, "microtick/DepositQuote", nil)
    cdc.RegisterConcrete(TxMarketTrade{}, "microtick/MarketTrade", nil)
    cdc.RegisterConcrete(TxLimitTrade{}, "microtick/LimitTrade", nil)
    cdc.RegisterConcrete(TxSettleTrade{}, "microtick/SettleTrade", nil)
}

// generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
    ModuleCdc = codec.New()
    RegisterCodec(ModuleCdc)
    codec.RegisterCrypto(ModuleCdc)
    ModuleCdc.Seal()
}