package phoneutil

import (
	"bytes"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

// 暂时只考虑了中国的手机号

// 移动手机号前缀
var yiDongPhonePrefix = []string{"134", "135", "136", "137", "138", "139", "147", "150", "151", "152", "157", "158", "159", "165", "172", "178", "182", "183", "184", "187", "188", "198"}

// 联通手机区号
var lianTongPhonePrefix = []string{"130", "131", "132", "145", "155", "156", "166", "171", "175", "176", "185", "186"}

// 电信手机号区号
var dianXinPhonePrefix = []string{"133", "149", "153", "173", "177", "180", "181", "189", "199"}

// 随机手机后8位
func Random8Suffix() string {
	// 字符串拼接使用它更高效
	var buffer bytes.Buffer
	r := randomRand()
	for i := 0; i < 8; i++ {
		number := r.Intn(10)
		buffer.WriteString(strconv.Itoa(number)) //转换数字字面值为字符串，之后将该字符串写入到buffer当中
	}
	return buffer.String()
}

// 随机手机区号前缀
func RandomPrefix() string {
	//	 所有的切片加入到一个总切片
	var allPrefix []string
	allPrefix = append(allPrefix, yiDongPhonePrefix...)
	allPrefix = append(allPrefix, lianTongPhonePrefix...) // 三个点不能省略哦
	allPrefix = append(allPrefix, dianXinPhonePrefix...)
	//	 随机取一个
	len := len(allPrefix)
	r := randomRand()
	index := r.Intn(len)
	return allPrefix[index]
}

// 固定前缀区号，随机8尾号
func FixPrefixRandomSuffix(prefix string) string {
	return prefix + Random8Suffix()
}

// 固定前缀区号，指定个数量的手机号
func FixPrefixFixNumberRandomSuffix(prefix string, count int) (numbers []string) {
	for {
		if len(numbers) == count {
			break
		}
		phone := FixPrefixRandomSuffix(prefix)

		//	 判断里面是否已经包含该手机，不能重复。
		if !exists(numbers, phone) {
			//	 追加
			numbers = append(numbers, phone)
		}
	}
	return
}

// 随机前缀，随机后缀，指定数量手机号
func FixNumberPhones(count int) (numbers []string) {
	for {
		if len(numbers) == count {
			break
		}
		phone := FixPrefixRandomSuffix(RandomPrefix())
		//	 判断里面是否已经包含该手机，不能重复。
		if !exists(numbers, phone) {
			//	 追加
			numbers = append(numbers, phone)
		}
	}
	return
}

// 判断是否是合格的中国11位手机号

func IsLegalChinaPhone(phone string) bool {
	// 是否为11位
	if len(phone) != 11 {
		return false
	}
	// 解析是否为全数字模式
	b, err := regexp.MatchString(`\d{11}`, phone)
	if err != nil {
		log.Println(err.Error())
	}
	if !b {
		return false
	}

	// 前缀是否属于注册的所有手机号, 取出前三位，循环对比即可
	prefix := phone[0:3]

	if IsLegalPrefix(prefix) {
		return true
	}
	return false
}

// 是否为合法的前缀
func IsLegalPrefix(prefix string) bool {
	allPrefix := AllPrefix()

	//遍历
	for _, v := range allPrefix {
		if v == prefix {
			return true
		}
	}

	return false
}

// 所有前缀

func AllPrefix() []string {
	var allPrefix []string
	allPrefix = append(allPrefix, yiDongPhonePrefix...)
	allPrefix = append(allPrefix, lianTongPhonePrefix...)
	allPrefix = append(allPrefix, dianXinPhonePrefix...)
	return allPrefix
}

// 判断切片里面是否包含某元素
func exists(numbers []string, phone string) bool {
	for _, v := range numbers {
		if v == phone {
			return true
		}
	}
	return false
}

// 随机对象
func randomRand() *rand.Rand {
	source := rand.NewSource(time.Now().UnixNano()) // 不指定seed会导致每次结果都一样
	r := rand.New(source)
	return r
}
