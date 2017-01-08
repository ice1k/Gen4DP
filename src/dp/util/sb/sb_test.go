package sb

import (
	"testing"
)

func TestStringBuffer(t *testing.T) {
	stringBuffer := NewStringBuffer()
	stringBuffer.AppendLine("boy next door")
	stringBuffer.AppendLineIndent("loop(1000) {")
	stringBuffer.AppendLine("Ass we can")
	stringBuffer.AppendLineClose("}")
	if stringBuffer.ToString() !=
		`boy next door
loop(1000) {
	Ass we can
}
` {
		t.FailNow()
	}
}

func TestStringBuffer_Clear(t *testing.T) {
	buffer := NewStringBuffer()
	buffer.Append("balabala")
	buffer.Clear()
	buffer.Append("_")
	if buffer.ToString() != "_" {
		t.FailNow()
	}
}
