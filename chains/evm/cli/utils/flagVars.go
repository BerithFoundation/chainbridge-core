// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/ethereum/go-ethereum/common"
)

//flag vars
var (
	BlockNumber string
	Blocks      string
	TxHash      string
	FromAddress string
)

//processed flag vars
var (
	txHash   common.Hash
	fromAddr common.Address
)
