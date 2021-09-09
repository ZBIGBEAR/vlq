package encode

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"vlq/common"
)

/*
输入：一串字符串，逗号分割
示例：0,11,2,9,7,6,2,4,-9,11,7,-8,5,9

输出：base64编码
示例：AWESOMEITWORKS

验证：https://www.murzwin.com/base64vlq.html
 */
func Encode(input string) (string, error) {
	// 1.检查格式是否正确
	if !checkInput(input) {
		return "", errors.New(fmt.Sprintf("format err:%s", input))
	}

	// 2.转换成数字
	numbers := convert2Number(input)

	// 3.返回结果
	ret := ""
	for _, number := range numbers {
		ret += number.Base64Binary
	}

	return ret, nil
}

func checkInput(input string) bool {
	inputs := strings.Split(input, common.Comma)
	for _, in := range inputs {
		if !common.CheckNumber(in) {
			return false
		}
	}

	return true
}

func convert2Number(input string) ([]*common.Number){
	inputs := strings.Split(input, common.Comma)
	ret := make([]*common.Number,0,len(inputs))
	for i, in := range inputs {
		num,_ := strconv.ParseInt(in,10,64)
		n := num
		positive := true
		if num<0{
			positive=false
			n = -n
		}
		originBinary := common.Ten2Two(n)
		vlqBinary := common.Binary2Vlq(positive, originBinary)
		ret = append(ret, &common.Number{
			Order:        i,
			Positive: positive,
			Str:          in,
			Num:          num,
			OriginBinary: originBinary,
			VLQBinary:    vlqBinary,
			Base64Binary: common.Vlq2Base64(vlqBinary),
		})
	}

	return ret
}
