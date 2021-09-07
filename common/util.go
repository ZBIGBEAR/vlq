package common

import "strings"

func Ten2Two(n int64) string {
	tmp := n
	result := make([]string,0)
	if tmp == 0 {
		return "0"
	}

	for tmp>0 {
		if tmp%2==0{
			result = append(result, "0")
		}else{
			result = append(result,"1")
		}
		tmp /= 2
	}

	ret := ""
	for i:= len(result)-1;i>=0;i--{
		ret += result[i]
	}

	return ret
}

// 检查输入字符串是不是合法整数
func CheckNumber(str string) bool {
	if len(str) == 0 {
		return false
	}

	if len(str) == 1 {
		return str[0]<='9' && str[0] >= '0'
	}

	if str[0] == '-' {
		return CheckNumber(str[1:])
	}

	for _, c := range str{
		if c>'9' || c<'0'{
			return false
		}
	}

	return true
}

func Two2Ten(str string) int64 {
	var ret int64
	var mult int64 = 1
	length:= len(str)-1
	for i:=length;i>=0;i--{
		if i == length{
			mult = 1
		}else{
			mult *=2
		}
		if str[i]=='1'{
			ret += mult
		}
	}

	return ret
}

func Binary2Vlq(positive bool, binary string) string {
	ret := ""
	length := len(binary)
	if length <=4 {
		ret = "0"+Supplement(binary,4)
		if positive {
			ret+="0"
		}else{
			ret+="1"
		}
		return ret
	}

	begin := length-4
	end := length
	ret = "1"+binary[begin:end]
	if positive {
		ret+="0"
	}else{
		ret+="1"
	}

	step := 5
	end = begin
	begin -=step
	if begin<0{
		begin=0
	}
	for begin >=0{
		subSection := binary[begin:end]
		if begin==0{
			ret+="0"+Supplement(subSection,5)
			break
		}else{
			ret+="1"+Supplement(subSection,5)
		}
		end = begin
		begin-=step
		if begin<0{
			begin=0
		}
	}
	return ret
}

func Supplement(binary string, count int) string {
	if len(binary) < count {
		return strings.Repeat("0", count-len(binary))+binary
	}

	return binary
}

func Vlq2Base64(vlq string) string {
	begin :=0
	step := 6
	end := begin+step
	ret := ""
	for end<=len(vlq){
		subSection := vlq[begin:end]
		num := Two2Ten(subSection)
		base64 := Base64[num]
		ret += base64
		begin = end
		end += step
	}

	return ret
}