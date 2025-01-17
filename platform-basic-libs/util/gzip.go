package util

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func GzipCompress(data string) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	gzW := gzip.NewWriter(buf)
	_, err := gzW.Write(Str2bytes(data))
	if err != nil {
		return nil, err
	}
	gzW.Close()
	return buf.Bytes(), err
}

func GzipUnCompress(data []byte) (string, error) {
	gzR, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(gzR)
	return Bytes2str(b), err
}
