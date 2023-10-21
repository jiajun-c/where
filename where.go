package where

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
)

func init() {
	m := map[string]interface{}{}
	jsonFile, err := os.Open("./data/area.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &m)
	parseProvince(m)
}

type Place struct {
	Province string
	City     string
	Area     string
	level    string
}

var (
	LevelProvince = "Province"
	LevelCity     = "City"
	LevelArea     = "Area"
)

func IsProvince(req string) ([]Place, error) {
	ans := make([]Place, 0)
	if provinceMap[req] {
		ans = append(ans, Place{
			Province: req,
			City:     "",
			Area:     "",
			level:    LevelProvince,
		})
		return ans, nil
	}
	AddProv := req + "省"
	if provinceMap[AddProv] {
		ans = append(ans, Place{
			Province: AddProv,
			City:     "",
			Area:     "",
			level:    LevelProvince,
		})
		return ans, nil
	}
	return nil, errors.New("no place")
}

func IsCity(req string) ([]Place, error) {
	ans := make([]Place, 0)
	if prov, ok := cityMap[req]; ok {
		ans = append(ans, Place{
			Province: prov,
			City:     req,
			Area:     "",
			level:    LevelCity,
		})
		return ans, nil
	}
	AddCity := req + "市"
	if prov, ok := cityMap[AddCity]; ok {
		ans = append(ans, Place{
			Province: prov,
			City:     AddCity,
			Area:     "",
			level:    LevelCity,
		})
	}
	return nil, errors.New("no place")
}

func IsArea(req string) ([]Place, error) {
	ans := make([]Place, 0)
	if citys, ok := areaMap[req]; ok {
		for _, c := range citys {
			ans = append(ans, Place{
				Province: cityMap[c],
				City:     c,
				Area:     req,
				level:    LevelArea,
			})
		}
		return ans, nil
	}
	AddArea := req + "区"
	if citys, ok := areaMap[AddArea]; ok {
		for _, c := range citys {
			ans = append(ans, Place{
				Province: cityMap[c],
				City:     c,
				Area:     AddArea,
				level:    LevelArea,
			})
		}
		return ans, nil
	}
	return nil, errors.New("no place")
}
func Where(req string) ([]Place, error) {
	ans, err := IsProvince(req)
	if err == nil {
		return ans, nil
	}
	ans, err = IsCity(req)
	if err == nil {
		return ans, nil
	}
	ans, err = IsArea(req)
	if err == nil {
		return ans, nil
	}
	return nil, errors.New("no places")
}
