package errcode

var (
	ErrorFeatureTypeUnknow = NewError(2000001, "feature type support int,float,bool,string,date,array,map")
	ErrorTypeConvert       = NewError(2000002, "type convert error")
)
