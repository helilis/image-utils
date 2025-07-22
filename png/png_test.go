package png

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"imageUtils/common"
	"os"
	"path/filepath"
	"testing"
)

func TestProcessorPng_SetMetaData_forTextChunk(t *testing.T) {
	path, err := filepath.Abs("../data/01.png")
	if err != nil {
		t.Fatal("read path fail, err:" + err.Error())
	}
	file, err := os.ReadFile(path)
	if err != nil {
		t.Fatal("read file fail, err:" + err.Error())
	}
	// 测试数据
	value := map[string]string{
		"Label":           "1",
		"ContentProducer": "hejianqiao-producer",
		"ProducerID":      "hejianqiao-AAA",
	}
	v, _ := json.Marshal(value)
	//
	pp := new(ProcessorPng)
	data, err := pp.SetMetaData(bytes.NewReader(file), "AIGC", string(v), common.TextChunkType)
	if err != nil {
		t.Fatal("set metadata fail, err:" + err.Error())
	}
	// 写文件
	outputPath, err := filepath.Abs("../data/01_text.png")
	if err = os.WriteFile(outputPath, data, 0666); nil != err {
		t.Fatal("write file fail, err:" + err.Error())
	}
}

func TestProcessorPng_SetMetaData_forZtxtChunk(t *testing.T) {
	path, err := filepath.Abs("../data/01.png")
	if err != nil {
		t.Fatal("read path fail, err:" + err.Error())
	}
	file, err := os.ReadFile(path)
	if err != nil {
		t.Fatal("read file fail, err:" + err.Error())
	}
	// 测试数据
	value := map[string]string{
		"Label":           "1",
		"ContentProducer": "hejianqiao-producer",
		"ProducerID":      "hejianqiao-AAA",
	}
	v, _ := json.Marshal(value)
	//
	pp := new(ProcessorPng)
	data, err := pp.SetMetaData(bytes.NewReader(file), "AIGC", string(v), common.ZtxtChunkType)
	if err != nil {
		t.Fatal("set metadata fail, err:" + err.Error())
	}
	// 写文件
	outputPath, err := filepath.Abs("../data/01_ztxt.png")
	if err = os.WriteFile(outputPath, data, 0666); nil != err {
		t.Fatal("write file fail, err:" + err.Error())
	}
}

func TestProcessorPng_SetMetaData_forItxtChunk(t *testing.T) {
	path, err := filepath.Abs("../data/01.png")
	if err != nil {
		t.Fatal("read path fail, err:" + err.Error())
	}
	file, err := os.ReadFile(path)
	if err != nil {
		t.Fatal("read file fail, err:" + err.Error())
	}
	// 测试数据
	value := map[string]string{
		"Label":             "1",
		"ContentProducer":   "hejianqiao-producer",
		"ProduceID":         "hejianqiao-AAA",
		"ReservedCode1":     "hejianqiao-ReservedCode1",
		"ContentPropagator": "hejianqiao-ContentPropagator",
		"PropagateID":       "hejianqiao-PropagateID",
		"ReservedCode2":     "hejianqiao-ReservedCode2",
	}
	v, _ := json.Marshal(value)
	//
	pp := new(ProcessorPng)
	data, err := pp.SetMetaData(bytes.NewReader(file), "AIGC", string(v), common.ItxtChunkType)
	if err != nil {
		t.Fatal("set metadata fail, err:" + err.Error())
	}
	// 写文件
	outputPath, err := filepath.Abs("../data/01_itxt.png")
	if err = os.WriteFile(outputPath, data, 0666); nil != err {
		t.Fatal("write file fail, err:" + err.Error())
	}
}

func TestAa(t *testing.T) {
	fmt.Println(binary.LittleEndian.Uint16([]byte{0x08, 0x00, 0x00, 0x00}))
}
