package main

import (
	"context"
	"log"

	fcap "github.com/Factor8Solutions/fc-api-provider"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/actors"
	marketactor "github.com/filecoin-project/lotus/chain/actors/builtin/market"
	"github.com/filecoin-project/lotus/chain/types"
)

func main() {

	// addBalance transfer info
	var targetMiner = "f0xxxxx"
	var fromWallet = "t1xxxxxxxx"
	// amount in attoFIL - this is 5 FIL
	var amount = types.NewInt(5000000000000000000)

	// API info

	var apiIp = "1.2.3.4"
	var apiPort = "1234"
	// token with sign rights --> admin permissions
	var apiToken = "eyxxxxxx"

	lfnc, err := fcap.NewLotusFullNodeClient(context.Background(), "Filecoin API", apiIp, apiPort, apiToken, fcap.Admin)

	if err != nil {
		log.Fatalf("Error creating LotusFullNodeClient: %s", err)
	}

	targetAddr, err := address.NewFromString(targetMiner)
	if err != nil {
		log.Fatalf("Error calling address.NewFromString: %s", err)
	}

	fromAddr, err := address.NewFromString(fromWallet)
	if err != nil {
		log.Fatalf("Error calling address.NewFromString: %s", err)
	}

	params, err := actors.SerializeParams(&targetAddr)
	if err != nil {
		log.Fatalf("Error calling address.actors.SerializeParams: %s", err)
	}

	smsg, err := lfnc.FullNodeApi.MpoolPushMessage(context.Background(), &types.Message{
		To:     marketactor.Address,
		From:   fromAddr,
		Value:  amount,
		Method: marketactor.Methods.AddBalance,
		Params: params,
	}, nil)

	if err != nil {
		log.Fatalf("Error calling MpoolPushMessage - %s", err)
	}

	log.Println(smsg.Cid())
}
