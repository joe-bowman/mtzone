package msg

import (
    "context"
    
    sdk "github.com/cosmos/cosmos-sdk/types"
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
    
    mt "github.com/mjackson001/mtzone/x/microtick/types"
)

func (querier Querier) Market(c context.Context, req *QueryMarketRequest) (*QueryMarketResponse, error) {
    ctx := sdk.UnwrapSDKContext(c)
    market := req.Market
    data, err := querier.Keeper.GetDataMarket(ctx, market)
    if err != nil {
        return nil, sdkerrors.Wrap(mt.ErrInvalidMarket, market)
    }
    
    var orderbookStatus []*MarketOrderBookStatus
    for k := 0; k < len(data.OrderBooks); k++ {
        if data.OrderBooks[k].SumBacking.Amount.GT(sdk.ZeroDec()) {
            callask, _ := querier.Keeper.GetActiveQuote(ctx, data.OrderBooks[k].CallAsks.First().Id)
            callbid, _ := querier.Keeper.GetActiveQuote(ctx, data.OrderBooks[k].CallBids.Last().Id)
            putask, _ := querier.Keeper.GetActiveQuote(ctx, data.OrderBooks[k].PutAsks.First().Id)
            putbid, _ := querier.Keeper.GetActiveQuote(ctx, data.OrderBooks[k].PutBids.Last().Id)
            CA := callask.CallAsk(data.Consensus)
            CB := callbid.CallBid(data.Consensus)
            PA := putask.PutAsk(data.Consensus)
            PB := putbid.PutBid(data.Consensus)
            orderbookStatus = append(orderbookStatus, &MarketOrderBookStatus {
                Name: data.OrderBooks[k].Name,
                SumBacking: data.OrderBooks[k].SumBacking,
                SumWeight: data.OrderBooks[k].SumWeight,
                InsideAsk: mt.NewMicrotickSpotFromDec(data.Consensus.Amount.Add(CA.Amount).Sub(PB.Amount)),
                InsideBid: mt.NewMicrotickSpotFromDec(data.Consensus.Amount.Sub(PA.Amount).Add(CB.Amount)),
                InsideCallAsk: CA,
                InsideCallBid: CB,
                InsidePutAsk: PA,
                InsidePutBid: PB,
            })
        }
    }
    
    response := QueryMarketResponse {
        Market: data.Market,
        Description: data.Description,
        Consensus: data.Consensus,
        OrderBooks: orderbookStatus,
        TotalBacking: data.TotalBacking,
        TotalWeight: data.TotalWeight,
    }
    
    return &response, nil
}
