package goUtils


import (
    "unicode/utf8"
    "bytes"
    "fmt"
)

func Temp() {
    fmt.Println("temp")
}

func Compare(a, b string) int {
    if a == b {
        return 0
    } else if a < b {
        return -1
    } else {
        return 1
    }
}


func IndexByte(s string, b byte) int {
    for index := 0; index < len(s); index++ {
        if s[index] == b {
            return index;
        }
    }
    return -1
}


/*
    in package "unicode/utf8"
    const (
            RuneError = '\uFFFD'     // the "error" Rune or "Unicode replacement character"
            RuneSelf  = 0x80         // characters below Runeself are represented as themselves in a single byte.
            MaxRune   = '\U0010FFFF' // Maximum valid Unicode code point.
            UTFMax    = 4            // maximum number of bytes of a UTF-8 encoded Unicode character.
    )
*/
func IndexRune(s string, r rune) int {
    if r < utf8.RuneSelf {
        return IndexByte(s, byte(r))
    } else {
        for k, v := range s {
            if v == r {
                return k
            }
        }
    }
    return -1
}


// only deals with English case
// if input is like Chinese, please call EqualFold
func EqualFoldSimple(s1, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    for i := 0; i < len(s1); i++ {
        if s1[i] < s2[i] {
            if s2[i] != s1[i] + 'a' - 'A' {
                return false
            }
        } else if s1[i] > s2[i] {
            if s1[i] != s2[i] + 'a' - 'A' {
                return false
            }
        }
    }
    return true
}


func Fields(s string, sep rune) []string {
    if len(s) == 0 {
        a := make([]string, 1)
        return a
    }

    inField := false
    var n int = 0

    for _, v := range s {
        if v == sep && !inField {
            inField = true
            n++
        } else if v != sep {
            inField = false
        }
    }

    if s[len(s) - 1] != byte(sep) {
        n++
    }

    strArr := make([]string, n)
    index := 0
    fieldStart := -1

    for k, v := range s {
        if v == sep {
            if fieldStart != -1 { // important while the spe occurs in the first place
                strArr[index] = s[fieldStart:k]
                index++
                fieldStart = -1
            }
        } else if fieldStart == -1 {
            fieldStart = k
        }
    }

    if fieldStart != -1 { // important because last field might end at EOF.
        strArr[index] = s[fieldStart:]
    }

    return strArr
}


func HasPrefix(s, pre string) bool {
    return len(s) >= len(pre) && s[0:len(pre)] == pre
}


func HasSuffix(s, suf string) bool {
    return len(s) >= len(suf) && s[len(s) - len(suf):] == suf
}


// To-Do: need to run performance test between JoinByBuffer and Join
func JoinByBuffer(strArr []string, sep string) string {
    if len(strArr) == 0 {
        return ""
    } else if len(strArr) == 1 {
        return strArr[0]
    }

    var buffer bytes.Buffer

    for k, v := range strArr {
        buffer.WriteString(v)
        if k != len(strArr) - 1 {
            buffer.WriteString(sep)
        }
    }

    return buffer.String()
}


func Join(strArr []string, sep string) string {
    if len(strArr) == 0 {
        return ""
    } else if len(strArr) == 1 {
        return strArr[0]
    }

    n := len(sep) * (len(strArr) - 1)

    for i := 0; i < len(strArr); i++ {
        n += len(strArr[i])
    }

    b := make([]byte, n)
    bp := copy(b, strArr[0])

    for _, s := range strArr[1:] {
        bp += copy(b[bp:], sep)
        bp += copy(b[bp:], s)
    }
    return string(b)
}


func LastIndexByte(s string, b byte) int {
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == b {
            return i
        }
    }
    return -1
}


// Input: "test1 !test2 test3"
// Output: Test1 !test2 Test3
// strings.Title() output: Test1 !Test2 Test3
func Title(s string) string {
    var c rune = ' '
    sArr := []rune(s)
    var sep bool = false
    for k, v := range sArr {
        if v == c {
            sep = true
            continue
        }
        if sep {
            if 'a' <= v && v <= 'z' {
                sArr[k] = v - ('a' - 'A')
            }
            sep = false
            continue
        }
        if k == 0 && 'a' <= v && v <= 'z' {
            sArr[0] = v - ('a' - 'A')
        }
    }
    return string(sArr)
}


func TrimAllSpace(s string) string {
    n := 0
    index := 0
    var space rune = ' '
    for _, v := range s {
        if v == space {
            n++
        }
    }

    rArr := make([]rune, len(s) - n)

    for _, v := range s {
        if v!= space {
            rArr[index] = v
            index++
        }
    }
    return string(rArr)
}






