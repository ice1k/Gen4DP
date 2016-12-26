package dp

type codeStyle struct {
	UseDefine       bool
	UsePrefixInc    bool
	UseArrayFromOne bool
}

func NewCodeStyle() *codeStyle {
	ret := new(codeStyle)
	ret.UseArrayFromOne = false
	ret.UseDefine = true
	ret.UsePrefixInc = true
	return ret
}
