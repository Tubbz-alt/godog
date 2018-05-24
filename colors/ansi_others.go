// Copyright 2014 shiena Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// +build !windows

package colors

import (
	"io"
)

func colored(w io.Writer) io.Writer {
	return w
}
