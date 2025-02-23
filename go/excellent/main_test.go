package main

import "testing"

// TestEvenOrOdd は、EvenOrOdd 関数が正しく動作するかどうかを検証するためのテストです。
// このテストでは、入力値として 10 を与え、結果が "even" であることを確認します。
func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10) // EvenOrOdd 関数を呼び出し、結果を取得
	if result != "even" {   // 結果が "even" でなければエラーを報告
		t.Errorf("expected: even, actual: %s", result)
	}
}
