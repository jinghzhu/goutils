package goUtils

import (
        "testing"
)

func TestCompare(t *testing.T) {
        s1 := ""
        s2 := "Shanghai"
        s3 := "Paris"
        if Compare(s1, s2) {
            t.Error("error1")
        }

}

func TestIndexRune(t *testing.T) {
        s1 := ""
        s2 := "abcd"
        var r rune = 99 // c
        if IndexRune(s1, r) != -1 {
            t.Error("error1")
        }
        if IndexRune(s2, r) != 2 {
            t.Error("error2")
        }
}