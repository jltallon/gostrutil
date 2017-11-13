// (C) 2016-2017
// Author: jltallon
package strbuilder

import (
    "fmt"
	"io"
)


// from bytes.Buffer
// StringWriter is the interface that wraps the WriteString method.
type StringWriter interface {
	WriteString(s string) (n int, err error)
}

type StringBuilder interface {
	io.WriterTo
	StringWriter
	String() string
}

type StrBldr struct {
    buf			[]byte
}




func New(capacity uint) StrBldr {
	var sb StrBldr
	sb.buf = make([]byte,0,int(capacity))
	return sb
}

func FromString(x string) StrBldr {
	var sb StrBldr
	sb.add([]byte(x))
	return sb
}


func (sb *StrBldr) Add(s string) {
	sb.add([]byte(s))
}


func (sb *StrBldr) Fmt(s string, x ...interface{}) {
	sb.add([]byte(fmt.Sprintf(s,x...)))
}

func (sb *StrBldr) Append(x []string) {
	for _,v := range x {
		sb.add([]byte(v))
	}
}

// Grow increases the buffer's capacity, if necessary, making room for another
// n bytes. After Grow(n), at least n bytes can be written to the buffer
// without another allocation.
// If the buffer can't grow it will panic with ErrTooLarge.
func (sb *StrBldr) Grow(n uint) {
	m := sb.doGrow(int(n))	
	sb.buf = sb.buf[0:m]
}


func (sb *StrBldr) String() string {
	if nil == sb.buf {
		// Special case, mostly for debugging purposes
		return "<nil>"
	}
	return string(sb.buf)
}


func (sb *StrBldr) WriteString(s string) (int,error) {
	return sb.add([]byte(s))
}
// 
// func (*sb *StrBldr) Write(x []byte) (int,error) {
// 	return sb,add(x)
// }



// WriteTo writes data to w until the buffer is drained or an error occurs.
// The return value n is the number of bytes written; it always fits into an
// int, but it is int64 to match the io.WriterTo interface. Any error
// encountered during the write is also returned.
func (sb *StrBldr) WriteTo(w io.Writer) (n int64, err error) {
	
	nBytes := len(sb.buf)
	m, e := w.Write(sb.buf)
// 	if m > nBytes {
// 		panic("SB::WriteTo: invalid Write count")
// 	}
	
	n = int64(m)
	if nil != e {
		return n, e
	}
	
	// all bytes should have been written, by definition of
	// Write method in io.Writer
	if m != nBytes {
		return n, io.ErrShortWrite
	}
	
	// Buffer is now empty; reset.
	sb.Reset()
	
	return	
}


func (sb *StrBldr) Len() int {
	return len(sb.buf)
}

func (sb *StrBldr) Cap() int {
	return cap(sb.buf)
}


func (sb *StrBldr) Reset() {
	sb.buf = sb.buf[:0]
}


// MinRead is the minimum slice size passed to a Read call by
// Buffer.ReadFrom. As long as the Buffer has at least MinRead bytes beyond
// what is required to hold the contents of r, ReadFrom will not grow the
// underlying buffer.

const MinRead = 512


// LoadFrom reads data from r until EOF and appends it to the buffer, growing
// the buffer as needed. The return value n is the number of bytes read. Any
// error except io.EOF encountered during the read is also returned. If the
// buffer becomes too large, ReadFrom will panic with ErrTooLarge.

func (sb *StrBldr) LoadFrom(r io.Reader) (n int64, err error) {
	
	for {
		if free := cap(sb.buf) - len(sb.buf); free < MinRead {
			newBuf := _alloc(2*cap(sb.buf) + MinRead)
			copy(newBuf, sb.buf)
			sb.buf = newBuf[:len(sb.buf)]
		}
		
		m, e := r.Read(sb.buf[len(sb.buf):cap(sb.buf)])
		
		sb.buf = sb.buf[0 : len(sb.buf)+m]
		
		n += int64(m)
		
		if e == io.EOF {
			break
		}
		
		if e != nil {
			return n, e
		}
	}
	return n, nil // err is EOF, so return nil explicitly
}



////////////////////////////////////////////////////////////////////////////////
// inlineable fastpath for cases where capacity is enough: the buffer only needs to be resliced.
// Returns the index where bytes should be written and whether it succeeded.
func (sb *StrBldr) tryGrow(n int) (int, bool) {
	if l := len(sb.buf); l+n <= cap(sb.buf) {
		sb.buf = sb.buf[:l+n]
		return l, true
	}
	return 0, false
}


// grow grows the buffer to guarantee space for n more bytes.
// It returns the index where bytes should be written.
// If the buffer can't grow it will panic with ErrTooLarge.
func (sb *StrBldr) doGrow(n int) int {
	m := len(sb.buf)
	// If buffer is empty, reset to recover space.
	if 0 == m  {
		sb.Reset()
	}
	// Try to grow by means of a reslice.
	if i, ok := sb.tryGrow(n); ok {
		return i
	}
	
	c := cap(sb.buf)
	if m+n > c/2 {	// optimize to avoid reallocating too frequently
		// Not enough capacity left, we need to allocate.
		buf := _alloc(2*c + n)
		copy(buf, sb.buf[:])
		sb.buf = buf
	} else {
		sb.buf = sb.buf[:m+n]
	}
	return m	
}


func (sb *StrBldr) add(x []byte) (int,error) {
	l := len(x)
	m, ok := sb.tryGrow(l)
	if !ok {
		m = sb.doGrow(l)
	}
	return copy(sb.buf[m:], x), nil
}


// Allocate a slice of size n. If the allocation fails, it panics
func _alloc(n int) []byte {
	// If the make fails, give a known error.
	defer func() {
		if recover() != nil {			
			panic("Size too large")
		}
	}()
	
	return make([]byte, n)
}
