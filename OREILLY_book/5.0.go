package main

import "strconv"

func main() {

	var i int = 123

	//整数同士の変換は、カッコで括り型を前置する。
	var f float64 = float64(i)

	//64ビットOSで64ビットのintと、intも明示的な変換が必要です。
	var i64 int64 = int64(i)

	//boolへの変換は比較演算子を用います。
	var b bool = i != 9

	//文字列との変換はstrcovを利用
	in = 12345
	//strconvの数値入力はint64
	s := strconv.FormatInt(int64(in), 10)
}
