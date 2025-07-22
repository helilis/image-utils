package png

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"errors"
	"hash/crc32"
	"imageUtils/common"
)

// 添加png图片iTXt chunk，默认加到IHDR块后第一个块
func addItxtChunk(fileData []byte, position int, key, value string, compress bool) ([]byte, error) {
	// 组装chunk数据块
	chunkData := packageItxtChunk(key, value, compress)
	// 组装chunk：length + type + data + CRC
	var chunk bytes.Buffer
	// 写入length
	if err := binary.Write(&chunk, binary.BigEndian, uint32(len(chunkData))); nil != err {
		return nil, errors.New("[PngAddChunk]add itxt chunk length fail, err:" + err.Error())
	}
	// 写入chunk type
	if _, err := chunk.Write([]byte(common.ItxtChunkType)); nil != err {
		return nil, errors.New("[PngAddChunk]add itxt chunk type fail, err:" + err.Error())
	}
	// 写入chunk data
	if _, err := chunk.Write(chunkData); nil != err {
		return nil, errors.New("[PngAddChunk]add itxt chunk data fail, err:" + err.Error())
	}
	// 写入CRC
	crc := crc32.ChecksumIEEE(append([]byte(common.ItxtChunkType), chunkData...))
	if err := binary.Write(&chunk, binary.BigEndian, crc); nil != err {
		return nil, errors.New("[PngAddChunk]add itxt chunk crc fail, err:" + err.Error())
	}
	// 组装新文件
	return append(fileData[:position], append(chunk.Bytes(), fileData[position:]...)...), nil
}

// 组装iTXt chunk
func packageItxtChunk(key, value string, compress bool) []byte {
	var buf bytes.Buffer
	// 写入key
	buf.Write([]byte(key))
	// 写入分割符
	buf.Write([]byte{0x00})
	// 写入压缩标志
	if compress {
		buf.Write([]byte{0x01})
	} else {
		buf.Write([]byte{0x00})
	}
	// 写入压缩方法，只能为zlib(0)
	buf.Write([]byte{0x00})
	// 写入语言标识
	buf.Write(append([]byte("zh-CN"), 0x00))
	// 写入key翻译版本，这里默认写死，可以传入
	buf.Write(append([]byte(key), 0x00))
	// 是否需要压缩
	if compress {
		// 压缩
		var compressBuf bytes.Buffer
		w := zlib.NewWriter(&compressBuf)
		_, _ = w.Write([]byte(value))
		_ = w.Close()
		buf.Write(compressBuf.Bytes())
	} else {
		// 不压缩
		buf.Write([]byte(value))
	}
	//
	return buf.Bytes()
}
