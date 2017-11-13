// (C) 2016-2017
// Author: jltallon
package strvec

import (
	"testing"
	"fmt"
)

var (
	testData1 = []string{"papa","pepe","pipi","popo","pupu"}
)



func Test(t *testing.T) {
	
	var sv StrVec
	
	
	sv.MergeAttrList(testData1,"cn")
	
	sv.Merge(testData1)
	sv.Append("joe")
	
	fmt.Println(sv)
	
	fmt.Println(sv.Join(";"))
}


func BenchmarkMergeAttrList(b *testing.B) {
	
	var sv StrVec
	
	for i:=0; i<b.N; i++ {
		
		sv.MergeAttrList(testData1,"cn")
		
	}
}
