// Copyright 2017 GRAIL, Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// +build !linux

package reflowlet

import "time"

func uptime() time.Duration {
	return 5 * time.Minute
}
