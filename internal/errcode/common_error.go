package errcode

var (
	ErrorFeatureTypeUnknow  = NewError(2000001, "feature type support int,float,bool,string,date,array,map")
	ErrorTypeConvert        = NewError(2000002, "type convert error")
	ErrorNotSupportOperator = NewError(2000003, "not support operator")
	ErrorNotANumber         = NewError(2000004, "not a number")
	ErrorBooleanValEmpty    = NewError(2000005, "boolean operator value is empty")
	ErrorBooleanValLack     = NewError(2000006, "boolean operator value lack")
)
