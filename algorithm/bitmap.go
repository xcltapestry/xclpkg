package algorithm

//Bitmap
//author:Xiong Chuan Liang
//date:2014-1-25

import (
	"fmt"
)

const (
	BitSize = 8 //一个字节8位
)

type Bitmap struct {
	BitArray  []byte
	ArraySize uint32
}

func NewBitmap(max uint32) *Bitmap {
	var r uint32
	switch {
	case max <= BitSize:
		r = 1
	default:
		r = max / BitSize
		if max%BitSize != 0 {
			r += 1
		}
	}

	fmt.Println("数组大小:", r)
	return &Bitmap{BitArray: make([]byte, r), ArraySize: r}
}

func (bitmap *Bitmap) Set(i uint32) {
	idx, pos := bitmap.calc(i)
	bitmap.BitArray[idx] |= 1 << pos
	fmt.Println("set()  value=", i, " idx=", idx, " pos=", pos, ByteToBinaryString(bitmap.BitArray[idx]))
}

func (bitmap *Bitmap) Test(i uint32) byte {
	idx, pos := bitmap.calc(i)
	return bitmap.BitArray[idx] >> pos & 1
}

func (bitmap *Bitmap) Clear(i uint32) {
	idx, pos := bitmap.calc(i)
	bitmap.BitArray[idx] &^= 1 << pos
}

func (bitmap *Bitmap) calc(i uint32) (idx, pos uint32) {

	idx = i >> 3 //相当于i / 8,即字节位置
	if idx >= bitmap.ArraySize {
		panic("数组越界.")
		return
	}
	pos = i % BitSize //位位置
	return
}

//ByteToBinaryString函数来源:
// Go语言版byte变量的二进制字符串表示
// http://www.sharejs.com/codes/go/4357
func ByteToBinaryString(data byte) (str string) {
	var a byte
	for i := 0; i < 8; i++ {
		a = data
		data <<= 1
		data >>= 1

		switch a {
		case data:
			str += "0"
		default:
			str += "1"
		}

		data <<= 1
	}
	return str
}
