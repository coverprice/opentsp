// Copyright 2014 The Sporting Exchange Limited. All rights reserved.
// Use of this source code is governed by a free license that can be
// found in the LICENSE file.

// tsp-forwarder extracts local metrics and forwards them to arbitrary relays.
package main

import "opentsp.org/cmd/tsp-forwarder/submit/cmd"

func main() {
	cmd.Run("forwarder")
}
