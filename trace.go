package rqlite

import (
	"io"

	"github.com/rqlite/gorqlite"
)

// We wrap the tracing functions so that people don't have to import gorqlite.

// TraceOn turns on tracing output to the io.Writer of your choice.
//
// Trace output is very detailed and verbose, as you might expect.
//
// Normally, you should run with tracing off, as it makes absolutely
// no concession to performance and is intended for debugging/dev use.
func TraceOn(w io.Writer) {
	gorqlite.TraceOn(w)
}

// TraceOff turns off tracing output. Once you call TraceOff(), no further
// info is sent to the io.Writer, unless it is TraceOn'd again.
func TraceOff() {
	gorqlite.TraceOff()
}
