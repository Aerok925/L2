package dev02

import (
	"testing"
)

func TestUnpackingStr(t *testing.T) {
	testStr := "a4bc2d5e"
	resultStr := "aaaabccddddde"

	testStr1 := "abcd"
	resultStr1 := "abcd"

	testStr2 := "45"
	resultStr2 := ""

	testStr3 := ""
	resultStr3 := ""

	t.Run("simple", func(t *testing.T) {
		realStr, _ := UnpackingStr(testStr1)
		if realStr != resultStr1 {
			t.Error("real string", realStr, "!=", resultStr1)
		}
	})

	t.Run("str with numeric", func(t *testing.T) {
		realStr, _ := UnpackingStr(testStr)
		if realStr != resultStr {
			t.Error("real string", realStr, "!=", resultStr)
		}
		realStr, _ = UnpackingStr("a1s1d1f1g1")
		if realStr != "asdfg" {
			t.Error("real string", realStr, "!=", "asdfg")
		}
	})

	t.Run("russian", func(t *testing.T) {
		realStr, _ := UnpackingStr("фывфывфыв")
		if realStr != "фывфывфыв" {
			t.Error("real string", realStr, "!=", "фывфывфыв")
		}
	})

	t.Run("russian with numeric", func(t *testing.T) {
		realStr, _ := UnpackingStr("ф2ы1в5ф1ы7в1ф1ы3в")
		if realStr != "ффывввввфыыыыыыывфыыыв" {
			t.Error("real string", realStr, "!=", "ффывввввфыыыыыыывфыыыв")
		}
	})

	t.Run("error str", func(t *testing.T) {
		realStr, _ := UnpackingStr(testStr2)
		if realStr != resultStr2 {
			t.Error("real string", realStr, "!=", resultStr2)
		}
	})

	t.Run("void str", func(t *testing.T) {
		realStr, _ := UnpackingStr(testStr3)
		if realStr != resultStr3 {
			t.Error("real string", realStr, "!=", resultStr3)
		}
	})
}
