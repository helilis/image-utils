package image_utils

import (
	"bytes"
	"imageUtils/common"
)

// Processor 图片处理接口
type Processor interface {
	// GetMetaData 获取指定key的图片元数据
	GetMetaData(r *bytes.Reader, key string) (string, error)
	// SetMetaData 设置图片元数据
	SetMetaData(r *bytes.Reader, key, value string, chunkType common.ImageChunkType) ([]byte, error)
}
