package typec

//=========================================================================================================
//	TypeFamily	定义了一个类的集合,这些类都有相同的内部结构表示,仅仅参数不同
type TypeFamily struct {
	Name   string                              //	族名
	Casts  []func(t DataType, s DataType) bool //	类型转换兼容比较函数
	Equal  func(t DataType, s DataType) bool   //	判断两个类型是否相等
	Define func(params []interface{}) DataType //	产生一个新的同族的类型的定义
}

//	DataType	是所有类型的同意接口,他们是TypeFamily实例化后的结果
type DataType interface {
	TypeFamily() *TypeFamily
	DefineParams() []interface{}
}

type DataValue interface {
	DataType() DataType
}

type Var interface {
	Name() string
	Value() DataValue
}

//=========================================================================================================

//	TypesEqual	用于判断两个类型是否相等
func TypesEqual(a DataType, b DataType) bool {
	family := a.TypeFamily()
	return family.Equal(a, b)
}

//	Assignable	用于判断类型s的值是否可以赋值给类型t的变量
func Assignable(t DataType, s DataType) bool {
	family := t.TypeFamily()
	for _, f := range family.Casts {
		if f(t, s) {
			return true
		}
	}

	return false
}

//=========================================================================================================

var familys = map[string]*TypeFamily{

}

type implDataType struct {
	typeFamily   *TypeFamily
	defineParams []interface{}
}

func (t*implDataType) TypeFamily() *TypeFamily {
	return t.typeFamily
}

func (t*implDataType) DefineParams() []interface{} {
	return t.defineParams
}
