package where

var provinceMap map[string]bool
var cityMap map[string]string
var areaMap map[string][]string

func parseProvince(placeMap map[string]interface{}) {
	cityMap = make(map[string]string)
	areaMap = make(map[string][]string)
	provinceMap = make(map[string]bool)
	for key, val := range placeMap {
		if key == "北京市" || key == "上海市" || key == "天津市" {
			cityMap[key] = "全国"
		} else {
			switch val.(type) {
			case map[string]interface{}:
				provinceMap[key] = true
				parseCityAndArea(val.(map[string]interface{}), key)
			default:
			}
		}
	}
}

func parseCityAndArea(placeMap map[string]interface{}, pro string) {
	for key, val := range placeMap {
		cityMap[key] = pro
		switch val.(type) {
		case []interface{}:
			for _, a := range val.([]interface{}) {
				areaMap[a.(string)] = append(areaMap[a.(string)], key)
			}
		case map[string]interface{}:
			cityMap[key] = pro
		}
	}
}
