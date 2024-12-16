package conversions

import (
	"DC/FnO/pkg/logger"
	"bytes"
	"encoding/binary"
	"strings"
	"unicode/utf8"

	"go.uber.org/zap"
)

func ToLitteEndian(data interface{}, requestType string) ([]byte, error) {
	var (
		buf = new(bytes.Buffer)
	)
	if err := binary.Write(buf, binary.LittleEndian, data); err != nil {
		logger.Log().Error("failed to convert to littel-endian", zap.Error(err), zap.String("requestType", requestType))
		return nil, err
	}
	return buf.Bytes(), nil
}

func FromLittleEndian(datatoRead []byte, dst interface{}, repounceType string) error {
	buff := bytes.NewReader(datatoRead)
	if err := binary.Read(buff, binary.LittleEndian, dst); err != nil {
		logger.Log().Error("failed to convert into object from little-endian", zap.Error(err), zap.String("repounceType", repounceType))
		return err
	}
	return nil
}

func BytesToString(b []byte) string {
    sanitized := strings.TrimRight(string(b), "\x00") // Trim null characters
    if !utf8.ValidString(sanitized) {
        return "" // Or handle invalid UTF-8 as needed
    }
    return sanitized
}
     
