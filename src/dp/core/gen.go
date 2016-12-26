package dp

import (
	"../util/sb"
	"fmt"
	"strings"
)

func loopWith(looper string, len string, style codeStyle) string {
	begin := 0
	if style.UseArrayFromOne {
		begin++
	}
	return fmt.Sprintf("for (%s=%d; %s<%s+%d; ++%s) {",
		looper,
		begin,
		looper,
		len,
		begin,
		looper)
}

/// core code generation function
func (info *dyProInfo) GenerateClang(style codeStyle) string {
	ret := sb.NewStringBuffer()
	ret.AppendLine(
		"#include <iostream>\n" +
			"#include <algorithm>\n" +
			"using namespace std;")
	if style.UseDefine {
		ret.AppendFormat("#define SIZE %d\n", info.Detail.MaxLen)
		ret.AppendLine("#define number " + info.Type)
	} else {
		ret.AppendFormat("const int SIZE=%d;\n", info.Detail.MaxLen)
		ret.AppendLine("typedef " + info.Type + "number;")
	}
	ret.Append("number ")
	ret.Append(info.State.Name)
	ret.AppendLine("[SIZE];")
	ret.AppendLine("")
	ret.AppendLineIndent("int main(const int argc, const char *argv[]) {")
	ret.AppendLine("int " + strings.Join(info.State.DimExpr, ", ") + ", n;")
	ret.AppendLine("cin>>n;")
	var dyPro func(int)
	dyPro = func(idx int) {
		if idx >= len(info.State.DimExpr) {
			for index, i := range info.Branches {
				if index >= 1 && index <= len(info.Branches) - 2 {
					ret.AppendLineIndent("else if (" + i.Conditions + ") {")
				} else if index == 0 {
					ret.AppendLineIndent("if (" + i.Conditions + ") {")
				} else {
					ret.AppendLineIndent("else {")
				}
				ret.AppendFormat("%s[%s] = %s;\n", info.State.Name, info.State.DimExpr[0], i.Expression)
				ret.AppendLineClose("}")
			}
			return
		} else {
			ret.AppendLineIndent(loopWith(info.State.DimExpr[idx], "n", style))
			dyPro(idx + 1)
			ret.AppendLineClose("}")
		}
	}
	dyPro(0)
	ret.AppendFormat("cout<<%s[%s-1]<<endl;\n", info.State.Name, "n")
	// add codes in main func
	ret.AppendLine("return 0;")
	ret.AppendLineClose("}")
	return ret.ToString()
}
