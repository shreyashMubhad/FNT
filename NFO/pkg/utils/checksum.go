package utils

import (
	"DC/FnO/pkg/logger"
	"crypto/md5"
	"encoding/json"
)

type Ichecksum interface {
	GetType() int
}

func GetCheckSum(ic Ichecksum) []byte {
	hasher := md5.New()

	ConToStr, err := json.Marshal(ic)
	if err != nil {
		logger.Log().Error("failed to convert oe_reqres to string")
	}

	hasher.Write(ConToStr)
	checksum := hasher.Sum(nil)
	return checksum
}
