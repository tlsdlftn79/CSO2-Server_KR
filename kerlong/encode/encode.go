package encode

import (
	"bytes"
	"io/ioutil"

	iconv "github.com/djimenez/iconv-go"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var (
	CVtolocal *iconv.Converter
	CVtoutf8  *iconv.Converter
)

func InitConverter(local string) bool {
	cv, err := iconv.NewConverter("utf-8", local)
	if err != nil {
		panic(err)
	}
	CVtolocal = cv
	cv, err = iconv.NewConverter(local, "utf-8")
	if err != nil {
		panic(err)
	}
	CVtoutf8 = cv
	return true
}

//GbkToUtf8 转换GBK编码到UTF-8编码
func GbkToUtf8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

//Utf8ToGbk 转换UTF-8编码到GBK编码
func Utf8ToGbk(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewEncoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

func Utf8ToLocal(str string) (b string, err error) {
	buf, err := CVtolocal.ConvertString(str)
	return string(buf), err
}

func LocalToUtf8(str string) (b string, err error) {
	buf, err := CVtoutf8.ConvertString(str)
	return string(buf), err
}
