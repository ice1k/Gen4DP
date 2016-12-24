package equation

type state struct {
	/// dimension of the dp equ
	dimension rune
	/// this is a little bit difficult to explain
	/// document will be available soon
	relations  rune
}

type stateEquation struct {
	stt state
}

//const (
//	OneDimention = "#include <stdio.h>\n" +
//		"int dp[]\n" +
//		"int main(const int argc, const char *argv[]) {\n" +
//		"\treturn 0;\n" +
//		"}\n" +
//		"\n"
//)
