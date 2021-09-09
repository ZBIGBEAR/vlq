package decode

import (
	"errors"
	"fmt"
	"strings"
	"vlq/common"
)

/*
输入：base64编码
示例：AWESOMEITWORKS

输出：一串数字，逗号分割
示例：0,11,2,9,7,6,2,4,-9,11,7,-8,5,9

验证：https://www.murzwin.com/base64vlq.html
*/
func Decode(input string) (string, error) {
	// 1.转换成数字
	numbers, err := convert2Number(input)
	if err != nil {
		return "", err
	}

	// 3.返回结果
	temNumbers := make([]string,0)
	for _, number := range numbers {
		temNumbers = append(temNumbers, number.Str)
	}

	return strings.Join(temNumbers, ","), nil
}

func convert2Number(input string) ([]*common.Number, error) {
	numbers := make([]*common.Number,0)
	tmpBinaries := make([]string,0)
	numBase64 := ""
	for i, in := range input {
		find := false
		numBase64 += string(in)
		for j, base64 := range common.Base64 {
			if base64 == string(in) {
				find = true
				binary := common.Ten2Two(int64(j))
				binaryVlq := common.Supplement(binary, 6)
				tmpBinaries = append(tmpBinaries, binaryVlq)
				if binaryVlq[0] == '0' {
					number := &common.Number{
						Order:        i,
						Positive:     false,
						Str:          "",
						Num:          0,
						OriginBinary: "",
						VLQBinary:    strings.Join(tmpBinaries,""),
						Base64Binary: numBase64,
					}
					number.Positive, number.OriginBinary = convert2Binary(tmpBinaries)
					number.Num = common.Two2Ten(number.OriginBinary)
					if !number.Positive{
						number.Num = -1*number.Num
					}
					number.Str = fmt.Sprintf("%d", number.Num)
					numbers = append(numbers, number)
					tmpBinaries = make([]string,0)
					numBase64 = ""
				}
				break
			}
		}
		if !find {
			return nil, errors.New(fmt.Sprintf("input:%s,%s invalid", input,string(in)))
		}
	}

	return numbers, nil
}

func convert2Binary(binarys []string) (bool, string) {
	positive := true
	if binarys[0][5] == '1' {
		positive = false
	}
	binary := ""
	for i:=len(binarys)-1;i>=1;i--{
		binary += binarys[i][1:6]
	}

	binary += binarys[0][1:5]
	return positive, binary
}