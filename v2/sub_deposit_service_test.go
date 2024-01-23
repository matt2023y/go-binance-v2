package binance

import (
	"context"
	"fmt"
	"testing"
)

func TestSubDepositsTransfer(t *testing.T) {
	client := NewClient("key", "sec")
	res, err := client.NewListSubDepositsService().Email("test@test.test").TxID("sssss").Do(context.Background())

	//client.NewListDepositsService().
	fmt.Println(res, err)
}
