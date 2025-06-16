package code

import "testing"

func TestProcessData(t *testing.T) {
	data := []byte{0xFF, 0xD8, 0x00, 0x01, 0x02, 0x03} // 模拟JPEG数据

	if err := processData(data); err != nil {
		t.Errorf("处理数据时出错: %v", err)
	}

	// 测试不合法的JPEG头
	invalidData := []byte{0x00, 0x01, 0x02}
	if err := processData(invalidData); err == nil {
		t.Error("预期处理不合法JPEG头时返回错误，但没有返回")
	}
}
