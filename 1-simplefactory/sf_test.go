package simplefactory

import "testing"

// TestApplePhoneFactory_CreatePhone 简单工厂测试方法
func TestApplePhoneFactory_CreatePhone(t *testing.T) {
	appleFactory := NewApplePhoneFactory()
	phone := appleFactory.CreatePhone("Apple 16", 1234.5)
	phone.Info()
}
