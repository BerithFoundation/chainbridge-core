// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package utils

import (
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls"
	"github.com/ChainSafe/chainbridge-core/chains/evm/calls/evmgaspricer"

	"github.com/spf13/cobra"
)

var UtilsCmd = &cobra.Command{
	Use:   "utils",
	Short: "Set of utility commands",
	Long:  "Set of utility commands",
}

func init() {
	UtilsCmd.AddCommand(simulateCmd)
	UtilsCmd.AddCommand(hashListCmd)
}

type GasPricerWithPostConfig interface {
	calls.GasPricer
	SetClient(client evmgaspricer.LondonGasClient)
	SetOpts(opts *evmgaspricer.GasPricerOpts)
}
