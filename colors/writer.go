// Copyright 2014 shiena Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package colors

import (
	"io"
)

// Colored creates and initializes a new ansiColorWriter
// using io.Writer w as its initial contents.
// In the console of Windows, which change the foreground and background
// colors of the text by the escape sequence.
// In the console of other systems, which writes to w all text.
func Colored(w io.Writer) io.Writer {
	return colored(w)
}
