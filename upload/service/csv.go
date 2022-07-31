package service

import (
	"bytes"
	"encoding/csv"
)

func Convert(data []byte) ([]string, error) {
	buf := bytes.NewBuffer(data)
	r := csv.NewReader(buf)
	sl := make([]string, 0)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	for _, v := range records {
		sl = append(sl, v[0])
	}
	return sl, nil
}
