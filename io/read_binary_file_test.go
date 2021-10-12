package iotest

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
	"math"
	"encoding/binary"
	"os"
	"bufio"
)


// EncodeFeatureFloat2Byte 编码浮点数特征为字节特征
func EncodeFeatureFloat2Byte(feature []float32) []byte {
	var data []byte
	for _, v := range feature {
		bits := math.Float32bits(v)
		bytes := make([]byte, 4)
		binary.LittleEndian.PutUint32(bytes, bits)
		data = append(data, bytes...)
	}
	return data
}

// EncodeFeatureByte2Float 编码字节特征为浮点数特征
func EncodeFeatureByte2Float(blob []byte) []float32 {
	off := 0
	var feature []float32
	for i := 0; i < (len(blob) / 4); i++ {
		v := binary.LittleEndian.Uint32(blob[off : off+4])
		feature = append(feature, math.Float32frombits(v))
		off += 4
	}
	return feature
}

// ReadFeature 读取二进制特征文件
func ReadFeature(path string) ([]float32, error) {
	file, err := os.Open(path)
    if err != nil {
        return []float32{}, err
    }
    defer file.Close()

    stats, statsErr := file.Stat()
    if statsErr != nil {
        return nil, statsErr
    }

    var size int64 = stats.Size()
    bytes := make([]byte, size)

    bufr := bufio.NewReader(file)
    _, err = bufr.Read(bytes)

	return EncodeFeatureByte2Float(bytes), nil
}


// TestReadBinaryFile 测试读取特征文件函数
func TestReadBinaryFile(t *testing.T) {
	file := "/data/rtc_align/py_sbct.bin"
	log.Println("test")
	features, _ := ReadFeature(file)
    fmt.Println(features)
}