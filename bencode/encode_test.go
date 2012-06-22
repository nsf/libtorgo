package bencode

import "testing"
import "bytes"

type random_encode_test struct {
	value    interface{}
	expected string
}

type random_struct struct {
	ABC         int    `bencode:"abc"`
	SkipThisOne string `bencode:"-"`
	CDE         string
}

var random_encode_tests = []random_encode_test{
	{int(10), "i10e"},
	{uint(10), "i10e"},
	{"hello, world", "12:hello, world"},
	{true, "i1e"},
	{false, "i0e"},
	{int8(-8), "i-8e"},
	{int16(-16), "i-16e"},
	{int32(32), "i32e"},
	{int64(-64), "i-64e"},
	{uint8(8), "i8e"},
	{uint16(16), "i16e"},
	{uint32(32), "i32e"},
	{uint64(64), "i64e"},
	{random_struct{123, "nono", "hello"}, "d3:CDE5:hello3:abci123ee"},
	{map[string]string{"a": "b", "c": "d"}, "d1:a1:b1:c1:de"},
	{[]byte{1, 2, 3, 4}, "4:\x01\x02\x03\x04"},
	{[4]byte{1, 2, 3, 4}, "li1ei2ei3ei4ee"},
	{nil, ""},
	{[]byte{}, "0:"},
	{"", "0:"},
	{[]int{}, "le"},
	{map[string]int{}, "de"},
}

func TestRandomEncode(t *testing.T) {
	for _, test := range random_encode_tests {
		data, err := Marshal(test.value)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(data, []byte(test.expected)) {
			t.Errorf("got: %s, expected: %s\n",
				string(data), string(test.expected))
		}
	}
}