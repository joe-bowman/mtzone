package keeper

import (
    mt "github.com/microtick/mtzone/x/microtick/types"
)

// DataAccountStatus

func NewDataAccountStatus(account mt.MicrotickAccount) DataAccountStatus {
    return DataAccountStatus {
        Account: account,
        ActiveQuotes: NewOrderedList(),
        ActiveTrades: NewOrderedList(),
        PlacedQuotes: 0,
        PlacedTrades: 0,
        QuoteBacking: mt.NewMicrotickCoinFromInt(0),
        TradeBacking: mt.NewMicrotickCoinFromInt(0),
        SettleBacking: mt.NewMicrotickCoinFromInt(0),
    }
}

