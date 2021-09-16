package utils

import "reflect"

// IsSameStruct 判断是否同一个结构体
func IsSameStruct(v1, v2 interface{}) bool {
	var (
		f1 reflect.Type
		f2 reflect.Type
		ok bool
	)

	if f1, ok = v1.(reflect.Type); !ok {
		f1 = reflect.TypeOf(v1)
	}

	if f2, ok = v2.(reflect.Type); !ok {
		f2 = reflect.TypeOf(v2)
	}

	return f1.PkgPath() == f2.PkgPath() && f1.Name() == f2.Name()
}
