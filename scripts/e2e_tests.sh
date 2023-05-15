#!/usr/bin/env bash
# Copyright 2021 ChainSafe Systems
# SPDX-License-Identifier: LGPL-3.0-only

go test -timeout 30m -p=1 $(go list ./... | grep 'e2e')