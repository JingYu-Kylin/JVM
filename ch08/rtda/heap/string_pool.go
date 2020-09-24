package heap

import "unicode/utf16"

// 用map来表示字符串池，key是Go字符串，value是Java字符串
var internedStrings = map[string]*Object{}

// todo
// go string -> java.lang.String
// 根据Go字符串返回相应的Java字符串实例
func JString(loader *ClassLoader, goStr string) *Object {
	// 如果Java字符串已经在池中，直接返回即可
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}
	// 否则先把Go字符串（UTF8格式）转换成Java字符数组（UTF16格式），
	// 然后创建一个Java字符实例，把它的value变量设置成刚刚转换而来的字符数组，
	// 最后把Java字符串放入池中
	chars := stringToUtf16(goStr)
	jChars := &Object{loader.LoadClass("[C"), chars}

	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetRefVar("value", "[C", jChars)

	internedStrings[goStr] = jStr
	return jStr
}

// java.lang.String -> go string
// 先拿到String对象的value变量值，然后把字符数组转换成Go字符串。
func GoString(jStr *Object) string {
	charArr := jStr.GetRefVar("value", "[C")
	return utf16ToString(charArr.Chars())
}

// utf8 -> utf16
// Go语言字符串在内存中是UTF8编码的，先把它强制转成UTF32，然后调用utf16包的Encode（）函数编码成UTF16。
func stringToUtf16(s string) []uint16 {
	runes := []rune(s)         // utf32
	return utf16.Encode(runes) // func Encode(s []rune) []uint16
}

// utf16 -> utf8
// 先把UTF16数据转换成UTF8编码，然后强制转换成Go字符串即可。
func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s) // func Decode(s []uint16) []rune
	return string(runes)
}
