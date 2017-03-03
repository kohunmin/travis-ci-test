package main

import "testing"

func TestSum(t *testing.T) {

	if sum(1, 2) != 3 {
		t.Error("더하기가 잘못되었습니다")
	}
}
