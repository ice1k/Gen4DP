package util

import "testing"

func TestIsAlpha(t *testing.T) {
	if !IsAlpha('a') {
		t.FailNow()
	}
	if IsAlpha('"') {
		t.FailNow()
	}
	if IsAlpha('\n') {
		t.FailNow()
	}
	//if IsAlpha('§') {
	//	t.FailNow()
	//}
	//if IsAlpha('№') {
	//	t.FailNow()
	//}
	//if IsAlpha('☆') {
	//	t.FailNow()
	//}
	//if IsAlpha('★') {
	//	t.FailNow()
	//}
	//if IsAlpha('○') {
	//	t.FailNow()
	//}
	//if IsAlpha('●') {
	//	t.FailNow()
	//}
	//if IsAlpha('◎') {
	//	t.FailNow()
	//}
	//if IsAlpha('◇') {
	//	t.FailNow()
	//}
	//if IsAlpha('◆') {
	//	t.FailNow()
	//}
	//if IsAlpha('□') {
	//	t.FailNow()
	//}
	//if IsAlpha('℃') {
	//	t.FailNow()
	//}
	//if IsAlpha('‰') {
	//	t.FailNow()
	//}
	//if IsAlpha('€') {
	//	t.FailNow()
	//}
	//if IsAlpha('■') {
	//	t.FailNow()
	//}
	//if IsAlpha('△') {
	//	t.FailNow()
	//}
	//if IsAlpha('▲') {
	//	t.FailNow()
	//}
	//if IsAlpha('※') {
	//	t.FailNow()
	//}
	//if IsAlpha('→') {
	//	t.FailNow()
	//}
	//if IsAlpha('←') {
	//	t.FailNow()
	//}
	//if IsAlpha('↑') {
	//	t.FailNow()
	//}
	//if IsAlpha('↓') {
	//	t.FailNow()
	//}
	//if IsAlpha('〓') {
	//	t.FailNow()
	//}
	//if IsAlpha('¤') {
	//	t.FailNow()
	//}
	//if IsAlpha('°') {
	//	t.FailNow()
	//}
	//if IsAlpha('＃') {
	//	t.FailNow()
	//}
	//if IsAlpha('＆') {
	//	t.FailNow()
	//}
	//if IsAlpha('＠') {
	//	t.FailNow()
	//}
	//if IsAlpha('＼') {
	//	t.FailNow()
	//}
	//if IsAlpha('︿') {
	//	t.FailNow()
	//}
	//if IsAlpha('\n') {
	//	t.FailNow()
	//}
	//if IsAlpha('￣') {
	//	t.FailNow()
	//}
	//if IsAlpha('―') {
	//	t.FailNow()
	//}
	//if IsAlpha('♂') {
	//	t.FailNow()
	//}
	//if IsAlpha('♀') {
	//	t.FailNow()
	//}
}

func TestIsDigit(t *testing.T) {
	if IsDigit('\n') {
		t.FailNow()
	}
	if IsDigit(']') {
		t.FailNow()
	}
	if !IsDigit('0') {
		t.FailNow()
	}
}

func TestAreAlpha(t *testing.T) {
	if !AreAlpha("boy next door") {
		t.FailNow()
	}
	if AreAlpha("My name is Van♂," +
		"I'm a artist♂," +
		"I'm an performance♂artist." +
		"I'm hired from people to fulfil their fantasies♂." +
		"The Deep♂dark♂fantasies.") {
		t.FailNow()
	}
}

func TestIsLeftBrace(t *testing.T) {
	if IsLeftBrace(']') {
		t.FailNow()
	}
	if !IsLeftBrace('[') {
		t.FailNow()
	}
}

func TestIsOperator(t *testing.T) {
	if IsOperator('G') {
		t.FailNow()
	}
	if !IsOperator('+') {
		t.FailNow()
	}
}

func TestIsRightBrace(t *testing.T) {
	if IsRightBrace('(') {
		t.FailNow()
	}
	if !IsRightBrace(')') {
		t.FailNow()
	}
}
