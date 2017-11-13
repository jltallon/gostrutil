// (C) 2016-2017
// Author: jltallon
package strbuilder

import (
	"os"
	"testing"
)

var (
	data = []string{"Roses","are","red","violets","are","blue"}
)

func TestSBgeneral(t *testing.T) {
	
	sb := FromString("The quick brown dog ")
	
	sb.Add("...jumped over the lazy fox\n")
	
	sb.Fmt("** This is a format test: '%v' **\n", data)
	
	sb.Append(data)
	sb.Add("\n")

	os.Stdout.WriteString(sb.String())
}


// func BenchmarkSBappend(b *testing.B) {
// 	
// 	sb := New(128)
// 	b.ResetTimer()
// 	
// 	for i:=0; i<b.N; i++ {
// 		
// 		sb.Append(data)
// 		
// 	}
// }

// func BenchmarkSBfmt(b *testing.B) {
// 	
// 	sb := New(128)
// 	b.ResetTimer()
// 	for i:=0; i<b.N; i++ {
// 		
// 		sb.Fmt("cnt=%d\n",i)
// 		
// 	}
// }
