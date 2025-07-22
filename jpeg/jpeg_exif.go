package jpeg

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

// 往exif中添加指定kv数据
func addExifInfo(r *bytes.Reader, key, value string) ([]byte, error) {
	// 文件数据
	data, err := io.ReadAll(r)
	if nil != err {
		return nil, errors.New("[addExifInfo]read file fail, err:" + err.Error())
	}
	// 查找APP01 segment（exif）
	pos := findAPP1Exif(data)
	if pos == -1 {
		panic("没有找到 Exif APP1 段")
	}
	// 提取 APP1 段并构造新的段
	exifData := extractAPP1Exif(data[pos:])
	newExif := addCustomTag(exifData, 0x9286, "AIGC", "test")
	// 替换 Exif APP1 段
	newJPEG := append(data[:pos], append(newExif, data[pos+len(exifData):]...)...)

	// 写入新图像
	//err = os.WriteFile("output.jpg", newJPEG, 0644)
	//if err != nil {
	//	return nil, errors.New("[addExifInfo]write file fail, err:" + err.Error())
	//}
	return newJPEG, nil
}

// 查找exif块位置
func findAPP1Exif(data []byte) int {
	for i := 0; i < len(data)-10; i++ {
		if data[i] == 0xFF && data[i+1] == 0xE1 && string(data[i+4:i+10]) == "Exif\x00\x00" {
			return i
		}
	}
	return -1
}

// 提取exif块数据
func extractAPP1Exif(data []byte) []byte {
	// data 起始就是 0xFFE1
	length := int(binary.BigEndian.Uint16(data[2:4]))
	return data[:length+2] // 包括 0xFFE1 和长度字段
}

func addCustomTag(exifData []byte, tagID uint16, tagName, value string) []byte {
	// TIFF Header 从 offset 10（"Exif\0\0"之后）
	tiffStart := 10
	//
	var endian binary.ByteOrder
	endian = binary.LittleEndian
	if string(exifData[tiffStart:tiffStart+2]) == "MM" {
		endian = binary.BigEndian
	}

	// TIFF IFD 偏移地址在 tiffStart+4
	ifdOffset := int(endian.Uint32(exifData[tiffStart+4 : tiffStart+8]))

	// IFD 开始位置
	ifdStart := tiffStart + ifdOffset
	numTags := int(endian.Uint16(exifData[ifdStart : ifdStart+2]))
	entryStart := ifdStart + 2
	entryEnd := entryStart + numTags*12

	// 插入自定义 entry
	buf := &bytes.Buffer{}
	buf.Write(exifData[:ifdStart])               // 之前数据
	binary.Write(buf, endian, uint16(numTags+1)) // tag数 +1
	buf.Write(exifData[entryStart:entryEnd])     // 原始 entry

	// 构造自定义 tag
	var entry [12]byte
	endian.PutUint16(entry[0:2], tagID)                // tag id
	endian.PutUint16(entry[2:4], 2)                    // type ASCII
	endian.PutUint32(entry[4:8], uint32(len(value)+1)) // count，包括 \0

	// 偏移地址：数据追加在 APP1 末尾，补全后跟在后面
	valueOffset := uint32(len(exifData)+100) - uint32(tiffStart)
	endian.PutUint32(entry[8:12], valueOffset)
	buf.Write(entry[:])

	// 写入 nextIFDOffset = 0
	buf.Write([]byte{0, 0, 0, 0})

	// 写入原始 APP1 后续内容（如缩略图等）
	buf.Write(exifData[entryEnd+4:]) // 跳过旧的 nextIFDOffset

	// 添加自定义字符串内容（"test\0"）
	padding := make([]byte, 100)
	copy(padding, value)
	buf.Write(padding)

	// 更新 APP1 段长度（总长 - 起始FFE1不算在长度里）
	app1Data := buf.Bytes()
	app1Len := len(app1Data) - 2
	binary.BigEndian.PutUint16(app1Data[2:4], uint16(app1Len))

	return app1Data
}
