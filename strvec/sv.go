// (C) 2016-2017
// Author: jltallon
package strvec

import (
// 	"fmt"
	"strings"
)

type StrVec []string

func New(capacity uint) StrVec {
	return make([]string, 0, int(capacity))
}

func (o *StrVec) Add(s string) {
	*o = append(*o, s)
}

func (o *StrVec) Join(sep string) string {
	return strings.Join(*o, sep)
}

func (o *StrVec) String() string {
	return strings.Join(*o, "\n")
}

func (o *StrVec) Len() int {
	return len(*o)
}

func (o *StrVec) Append(sx ...string) {
	*o = append(*o, sx...)
}

func (o *StrVec) Merge(ss []string) {
	*o = append(*o, ss...)
}


// Utility to merge "DN-like" attribute lists
func (o *StrVec) MergeAttrList(data []string, field string) {

	buf := make([]string, 0, 1)
	for _, v := range data {
		if strings.IndexRune(v, ',') >= 0 {
			v = `"` + v + `"`
		}
// 		buf = append(buf, fmt.Sprintf("%s = %s", field, v))		// 1700ns/op vs 1056ns/op
		buf = append(buf, (field+" = "+v))
	}
	dn := strings.Join(buf, ", ")

	*o = append(*o, dn)
}
