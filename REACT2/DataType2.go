package yy

type TypeFamily interface {
}

type DataType interface {
}

type DataValue interface {
}

//================================================================================

type implTypeFamily struct {
	base TypeFamily
}

func This(f TypeFamily) TypeFamily {
	return f
}

func Base(f TypeFamily) TypeFamily {
	return f.(implTypeFamily).base
}

func Instance(f TypeFamily, p []DataType) DataType {
	return nil
}

//================================================================================

type implDataType struct {
	family TypeFamily
	params []DataType
}

func ParamsOf(t DataType) []DataType {
	return t.(implDataType).params
}

func DeclareOf(t DataType) TypeFamily {
	return t.(implDataType).family
}

//================================================================================

type implDataValue struct {
	dataType DataType
}

func TypeOf(v DataValue) DataType {
	return v.(implDataValue).dataType
}

func test(){
	//replace(this(decl(T)), params(T)) == T
}