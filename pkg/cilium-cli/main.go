// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package main

import (
	"fmt"
	"os"

	"github.com/cilium/cilium-cli/api"
	"github.com/cilium/cilium-cli/cli"
)

var hooks = api.NopHooks{}

func main() {
	command := cli.NewCiliumCommand(&hooks)
	if err := command.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
