package jpeg

import (
	"bytes"
	"imageUtils/common"
)

type ProcessorJpeg struct{}

func (ProcessorJpeg) GetMetaData(r *bytes.Reader, key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (ProcessorJpeg) SetMetaData(r *bytes.Reader, key, value string, chunkType common.ImageChunkType) ([]byte, error) {
	// 写入到Exif中
	if chunkType == common.TextChunkType {
	}
	return nil, nil
}
