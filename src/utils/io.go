/*
Copyright 2021 Mirantis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import "io"

// writeCloserWrapper represents a WriteCloser whose closer operation is noop.
type writeCloserWrapper struct {
	Writer io.Writer
}

func (w *writeCloserWrapper) Write(buf []byte) (int, error) {
	return w.Writer.Write(buf)
}

func (w *writeCloserWrapper) Close() error {
	return nil
}

// WriteCloserWrapper returns a writeCloserWrapper.
func WriteCloserWrapper(w io.Writer) io.WriteCloser {
	return &writeCloserWrapper{w}
}

// LimitWriter is a copy of the standard library ioutils.LimitReader,
// applied to the writer interface.
// LimitWriter returns a Writer that writes to w
// but stops with EOF after n bytes.
// The underlying implementation is a *LimitedWriter.
func LimitWriter(w io.Writer, n int64) io.Writer { return &LimitedWriter{w, n} }

// A LimitedWriter writes to W but limits the amount of
// data returned to just N bytes. Each call to Write
// updates N to reflect the new amount remaining.
// Write returns EOF when N <= 0 or when the underlying W returns EOF.
type LimitedWriter struct {
	W io.Writer // underlying writer
	N int64     // max bytes remaining
}

func (l *LimitedWriter) Write(p []byte) (n int, err error) {
	if l.N <= 0 {
		return 0, io.ErrShortWrite
	}
	truncated := false
	if int64(len(p)) > l.N {
		p = p[0:l.N]
		truncated = true
	}
	n, err = l.W.Write(p)
	l.N -= int64(n)
	if err == nil && truncated {
		err = io.ErrShortWrite
	}
	return
}
