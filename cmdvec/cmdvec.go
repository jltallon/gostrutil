package cmdvec


func SplitMulti(input []string, delim string) [][]string {

	var result [][]string
	var temp []string

	for _,s := range input {
		
		if delim == s {
			result=append(result,temp)
			temp = nil
		} else {
			temp = append(temp,s)
		}
	}
	result=append(result,temp)	// append temp buffer to result, even if no 'delim'
	
	return result
}
