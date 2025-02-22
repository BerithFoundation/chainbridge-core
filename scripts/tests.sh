#!/usr/bin/env bash
# Copyright 2021 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only


CVPKG=$(go list ./... | grep -v 'e2e\|generated\|bindata\|mock\|main.go\|' | tr '\n' ',')
go test -coverpkg=$CVPKG -coverprofile=cover.out -p=1 $(go list ./... | grep -v 'cbcli\|e2e')
