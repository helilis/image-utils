package common

// ImageType 图片类型
type ImageType string

const (
	TypeJpeg = "jpeg" // jpeg格式图片
	TypePng  = "png"  // png格式图片
)

// ImageChunkType 图片数据块类型
type ImageChunkType string

const (
	IHDRChunkType = "IHDR" // PNG图片的IHDR数据块
	TextChunkType = "tEXt" // PNG图片的tEXt数据块
	ZtxtChunkType = "zTXt" // PNG图片的zTXt数据块
	ItxtChunkType = "iTXt" // PNG图片的iTXt数据块
	ExifChunkType = "Exif" // JPEG图片的Exif数据块
	XmpChunkType  = "xmp"  // JPEG图片的xmp数据块
)
