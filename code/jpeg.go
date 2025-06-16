package code

import (
	"bytes"
	"fmt"
)

/*
func ProcessData(data []byte) {
    // 1. 未校验data长度导致panic
    header := data[:4]
    // 2. 魔法数字"4"
    if bytes.Equal(header, []byte{0xFF,0xD8}) {
        // 3. JPEG处理未分离责任
    }
}
这段代码进行重构
*/

const (
	JPEGMagicNumber = "\xFF\xD8"
	JPEGHeaderSize  = 2
)

type format int32

const (
	formatUnknown format = iota
	formatJPEG
)

func validateJPEGHeader(data []byte) error {
	if len(data) < JPEGHeaderSize {
		return fmt.Errorf("数据长度不足，需要至少%d字节", JPEGHeaderSize)
	}
	return nil
}

func detectFormat(data []byte) format {
	if len(data) >= 2 && bytes.Equal(data[:2], []byte(JPEGMagicNumber)) {
		return formatJPEG
	}
	return formatUnknown
}

func processJPEG(data []byte) error {
	// 处理JPEG数据的逻辑
	fmt.Println("Processing JPEG data...")
	// 这里可以添加实际的JPEG处理代码
	return nil
}

func processData(data []byte) error {
	if err := validateJPEGHeader(data); err != nil {
		return fmt.Errorf("JPEG头验证失败: %w", err)
	}

	switch detectFormat(data) {
	case formatJPEG:
		if err := processJPEG(data); err != nil {
			return fmt.Errorf("处理JPEG数据失败: %w", err)
		}
	default:
		return fmt.Errorf("未知格式，无法处理数据")
	}

	return nil
}
