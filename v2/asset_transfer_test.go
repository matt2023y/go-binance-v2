package binance

import (
	"context"
	"fmt"
	"testing"
)

func TestAssetTransfer(t *testing.T) {
	client := NewClient("", "")
	res, _ := client.NewAssetTransfer().
		Type(AssetTransfer_FUNDING_MAIN).
		Asset("usdt").
		Amount(10).
		FromSymbol("usdt").
		ToSymbol("usdt").Do(context.Background())
	fmt.Println(res)
}
