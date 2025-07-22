package png

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"hash/crc32"
	"imageUtils/common"
)

// 添加一个tEXt chunk到png图片中
func addTextChunk(fileData []byte, position int, key, value string, compress bool) ([]byte, error) {
	// chunk类型
	chunkType := common.TextChunkType
	if compress {
		chunkType = common.ZtxtChunkType
	}
	// 组装chunk数据
	chunkData, err := packageChunkData(key, value, compress)
	if nil != err {
		return nil, err
	}
	// 新文件数据buffer
	var buf bytes.Buffer
	// 写入chunk长度
	if err := binary.Write(&buf, binary.BigEndian, uint32(len(chunkData))); nil != err {
		return nil, errors.New("[PngAddChunk]write chunk length fail, err:" + err.Error())
	}
	// 写入chunk类型
	buf.Write([]byte(chunkType))
	// 写入chunk数据
	buf.Write(chunkData)
	// 写入CRC
	crc := crc32.ChecksumIEEE(append([]byte(chunkType), chunkData...))
	if err := binary.Write(&buf, binary.BigEndian, crc); nil != err {
		return nil, errors.New("[PngAddChunk]write chunk crc fail, err:" + err.Error())
	}
	// 组装新图片文件
	return append(fileData[:position], append(buf.Bytes(), fileData[position:]...)...), nil
}

func packageChunkData(key, value string, compress bool) ([]byte, error) {
	var buf bytes.Buffer
	// 写入 key
	buf.Write([]byte(key))
	// 写入kv分隔符
	buf.Write([]byte{0x00})
	// 写入chunk数据块
	if compress {
		// 写入压缩标识符（0代表zlib）
		buf.Write([]byte{0x00})
		// 需要压缩处理
		var cb bytes.Buffer
		zw := zlib.NewWriter(&cb)
		if _, err := zw.Write([]byte(value)); nil != err {
			return nil, errors.New("[PngAddChunk]compress value fail, err:" + err.Error())
		}
		_ = zw.Close()
		// 写入数据
		buf.Write(cb.Bytes())
	} else {
		// 无需压缩，直接写入
		buf.Write([]byte(value))
	}
	//
	return buf.Bytes(), nil
}
