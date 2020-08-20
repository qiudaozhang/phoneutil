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

func TestFixNumberPhones(t *testing.T) {
	fmt.Println(FixNumberPhones(232))
}

func TestIsLegalChinaPhone(t *testing.T) {
	fmt.Println(IsLegalChinaPhone("19912341231"))
}
