/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package filememory

import (
	"fmt"
	"testing"
	"time"
)

func TestEthGetBlockNumber(t *testing.T) {

	tw := Client{
		BaseURL: "https://chain.ipfs-spacetime.com/exchange/",
		Debug:   true,
	}

	if r, err := tw.FMGetBlockNumber(); err != nil {
		t.Errorf("GetAccountNet failed: %v\n", err)
	} else {
		t.Logf("GetAccountNet return: \n\t%+v\n", r)
	}
}

func TestFmGetBlockSpecByBlockNum2(t *testing.T) {

	tw := Client{
		BaseURL: "https://chain.ipfs-spacetime.com/exchange/",
		Debug:   true,
	}

	if r, err := tw.fmGetBlockSpecByBlockNum2(1279134, true); err != nil {
		t.Errorf("GetAccountNet failed: %v\n", err)
	} else {
		t.Logf("GetAccountNet return: \n\t%+v\n", r)
	}
}

func TestGetBalance(t *testing.T) {
	tw := Client{
		BaseURL: "https://chain.ipfs-spacetime.com/exchange/",
		Debug:   true,
	}

	address := "FMa8cc6864cbd7f7e06dc4405ce04bb27abb91403b"

	if r, err := tw.GetAddrBalance2(address, "latest"); err != nil {
		t.Errorf("GetBalance failed: %v\n", err)
	} else {
		t.Logf("GetBalance return: \n\t%+v\n", r)
	}
}

func TestFmGetTransactionCount(t *testing.T) {
	tw := Client{
		BaseURL: "https://chain.ipfs-spacetime.com/exchange/",
		Debug:   true,
	}

	address := "FMa8cc6864cbd7f7e06dc4405ce04bb27abb91403b"

	if r, err := tw.fmGetTransactionCount(address); err != nil {
		t.Errorf("Get Transaction count failed: %v\n", err)
	} else {
		t.Logf("GetTrasactionCount return: \n\t%+v\n", r)
	}
}

func TestGenToken(t *testing.T) {
	now := time.Now().Unix()
	fmt.Println(now)
	fmt.Println(GenToken(now))
}
