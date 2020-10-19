package geoJson

import (
	"os"
	"io/ioutil"
	"encoding/json"
)

func file_get_contents(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadALL(f)
}

func GetPoints(path string) ([]StructPoint, error) {
	var content []byte
	var err error
	var c StructGeo
	var res []StructPoint

	content, err = file_get_contents(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(content), &c)
	if err != nil {
		return nil, err
	}
	for _, points := range c.Features {
		res = append(res, StructPoint{lng: points.Geometry.Coordinates[0], lat: points.Geometry.Coordinates[1]})
	}
	return res, nil
}
