package services

import (
	"crypto/md5"
	"encoding/hex"
)

type MD5HashService struct {
}

func NewMD5HashService() *MD5HashService {
	return &MD5HashService{}
}

func (M MD5HashService) HashString(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}
