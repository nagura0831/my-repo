package main

// EvenOrOdd は、引数として与えられた整数が偶数か奇数かを判定し、
// 偶数の場合は "even"、奇数の場合は "odd" を返します。
func EvenOrOdd(number int) string {
	// number を2で割った余りが0であれば偶数と判断する
	if number%2 == 0 {
		return "even" // 偶数の場合は "even" を返す
	} else {
		return "odd" // 奇数の場合は "odd" を返す
	}
}
