package x

//IMP : rs of import files
type IMP struct {
	file    string
	size    int
	rowcont int
	msg     string
	code    int
}

/*
func init() {
	W("b inited.")
}*/

var rb IMP

//ShowImp is func
func ShowImp() {
	W("b", rb.msg)
}
