package png

import (
	"bytes"
	"encoding/binary"
	"errors"
	"imageUtils/common"
	"io"
)

// ProcessorPng PNG图片处理器
type ProcessorPng struct{}

// GetMetaData 获取图片元数据
func (pp *ProcessorPng) GetMetaData(r *bytes.Reader, key string) (string, error) {
	return "", nil
}

// SetMetaData 设置图片元数据
// chunk位置默认为在IHDR chunk后面第一个
func (pp *ProcessorPng) SetMetaData(r *bytes.Reader, key, value string, chunkType common.ImageChunkType) ([]byte, error) {
	// reader转字节切片
	fileData, err := io.ReadAll(r)
	if nil != err {
		return nil, errors.New("[ProcessorPng]read image data fail, err:" + err.Error())
	}
	// 索引开始位置，跳过PNG文件头部8个字节的标志符
	pos := 8
	found := false // 是否找到IHDR chunk
	// 遍历文件chunk，找到IHDR chunk
	for pos < len(fileData) {
		// 提取chunk长度数据，并更新游标
		chunkLength := binary.BigEndian.Uint32(fileData[pos : pos+4])
		pos += 4
		// 提取chunk类型，并更新游标
		ct := fileData[pos : pos+4]
		pos += 4
		// 跳过chunk数据和CRC
		pos += int(chunkLength) + 4
		// 判断chunk 是否是 IHDR
		if bytes.Equal(ct, []byte(common.IHDRChunkType)) {
			found = true
			break
		}
	}
	// 写入chunk数据
	if found {
		//
		if chunkType == common.TextChunkType {
			// 写入到一个新的tEXt chunk中
			return addTextChunk(fileData, pos, key, value, false)
		} else if chunkType == common.ZtxtChunkType {
			// 写入到一个新的zTXt chunk中
			return addTextChunk(fileData, pos, key, value, true)
		} else if chunkType == common.ItxtChunkType {
			// 写入到一个新的iTXt chunk中
			return addItxtChunk(fileData, pos, key, value, true)
		}
	}
	// 默认返回原始数据
	return fileData, nil
}
