package udf

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/internal/util"
	"reflect"
)

//sum return float64
func Sum(v interface{}) (interface{}, error) {
	kind, err := util.GetType(v)
	switch kind {
	case configs.INT:
		return util.ToFloat64(v)
	case configs.FLOAT:
		return util.ToFloat64(v)
	case configs.ARRAY:
		var rs float64
		if reflect.TypeOf(v).Kind() == reflect.Array { //array like [3]int
			varr := reflect.ValueOf(v)
			for i := 0; i < varr.Len(); i++ {
				df, err := util.RVToFloat64(varr.Index(i)) //if array contains nan, return err
				if err != nil {
					return nil, err
				}
				rs += df
			}
		} else { //slice
			vslice := v.([]interface{})
			for _, va := range vslice {
				df, err := util.ToFloat64(va) //if slice contains nan, return err
				if err != nil {
					return nil, err
				}
				rs += df
			}
		}
		return rs, nil
	}
	return nil, err
}
