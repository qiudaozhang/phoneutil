package phoneutil

import (
	"fmt"
	"testing"
)

// 测试命名规范 Test函数名
func TestRandom8Suffix(t *testing.T) {
	suffix := Random8Suffix()
	fmt.Println(suffix)
}

func TestRandomPrefix(t *testing.T) {
	fmt.Println(RandomPrefix())
}

func TestFixPrefixRandomSuffix(t *testing.T) {
	fmt.Println(FixPrefixRandomSuffix("139"))
}

func TestFixPrefixFixNumberRandomSuffix(t *testing.T) {
	phones := FixPrefixFixNumberRandomSuffix("135", 20000)
	fmt.Println(len(phones))

}

// 在MacBookPro2019上测试，30000个手机号耗时2秒左右，30万个手机号，136秒左右
func TestFixNumberPhones(t *testing.T) {
	FixNumberPhones(300000)
}

func TestIsLegalChinaPhone(t *testing.T) {
	fmt.Println(IsLegalChinaPhone("19912341231"))
}
