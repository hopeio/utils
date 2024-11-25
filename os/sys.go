/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package os

import "os"

// Hostname returns the host name reported by the kernel.
func Hostname() string {
	hostname, _ := os.Hostname()
	return hostname
}
