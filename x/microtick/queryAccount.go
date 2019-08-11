package microtick 

import (
    "fmt"
    "strings"
    
    "github.com/cosmos/cosmos-sdk/codec"
    sdk "github.com/cosmos/cosmos-sdk/types"
    abci "github.com/tendermint/tendermint/abci/types"
)

type ResponseAccountStatus struct {
    Account string `json:"account"`
    Balance MicrotickCoin `json:"balance"`
    Change MicrotickCoin `json:"change"`
    NumQuotes uint32 `json:"numQuotes"`
    NumTrades uint32 `json:"numTrades"`
    ActiveQuotes []MicrotickId `json:"activeQuotes"`
    ActiveTrades []MicrotickId `json:"activeTrades"`
    QuoteBacking MicrotickCoin `json:"quoteBacking"`
    TradeBacking MicrotickCoin `json:"tradeBacking"`
    SettleBacking MicrotickCoin `json:"settleBacking"`
}

func (ras ResponseAccountStatus) String() string {
    return strings.TrimSpace(fmt.Sprintf(`Account: %s
Balance: %s
Change: %s
Num Quotes: %d
Num Trades: %d
Active Quotes: %v
Active Trades: %v
Quote Backing: %s
Trade Backing: %s
Settle Backing: %s`, ras.Account, 
    ras.Balance.String(),
    ras.Change,
    ras.NumQuotes, 
    ras.NumTrades, 
    ras.ActiveQuotes, ras.ActiveTrades,
    ras.QuoteBacking.String(), ras.TradeBacking.String(),
    ras.SettleBacking.String()))
}

func queryAccountStatus(ctx sdk.Context, path []string, 
    req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
    acct := path[0]
    address, err2 := sdk.AccAddressFromBech32(acct)
    balance := keeper.GetTotalBalance(ctx, address)
    data := keeper.GetAccountStatus(ctx, address)
    if err2 != nil {
        return nil, sdk.ErrInternal(fmt.Sprintf("Could not fetch address information: %s", err2))
    }
    activeQuotes := make([]MicrotickId, len(data.ActiveQuotes.Data))
    activeTrades := make([]MicrotickId, len(data.ActiveTrades.Data))
    for i := 0; i < len(data.ActiveQuotes.Data); i++ {
        activeQuotes[i] = data.ActiveQuotes.Data[i].Id
    }
    for i := 0; i < len(data.ActiveTrades.Data); i++ {
        activeTrades[i] = data.ActiveTrades.Data[i].Id
    }
    response := ResponseAccountStatus {
        Account: acct,
        Balance: balance,
        Change: data.Change,
        NumQuotes: data.NumQuotes,
        NumTrades: data.NumTrades,
        ActiveQuotes: activeQuotes,
        ActiveTrades: activeTrades,
        QuoteBacking: data.QuoteBacking,
        TradeBacking: data.TradeBacking,
        SettleBacking: data.SettleBacking,
    }
    
    bz, err2 := codec.MarshalJSONIndent(keeper.cdc, response)
    if err2 != nil {
        panic("Could not marshal result to JSON")
    }
    
    return bz, nil
}
