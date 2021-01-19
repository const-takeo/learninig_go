package switchTest

func CanIDrinkTwo(age int) bool {
	// switch文も同じくvariable expressionを使用可能
	// variable expressionを用いて範囲を指定する場合 semi-colon後ろの変数名を消す。
	switch koreanAge := age + 2; {
	case koreanAge < 20:
		return false
	case koreanAge > 50:
		return false
	case koreanAge < 50:
		return true
	}
	return false
}
