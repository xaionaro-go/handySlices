package handySlices

import (
	"reflect"
)

type KeyStringValuer interface {
	KeyStringValue() string
}

func wrapper(aI, bI interface{}, kernel func(aIE, bIE KeyStringValuer)) {
	aV := reflect.Indirect(reflect.ValueOf(aI))
	bV := reflect.Indirect(reflect.ValueOf(bI))

	bMap := map[string]KeyStringValuer{}

	bLen := bV.Len()
	for i := 0; i < bLen; i++ {
		bE := bV.Index(i)
		bIE := bE.Interface().(KeyStringValuer)
		bMap[bIE.KeyStringValue()] = bIE
	}

	aLen := aV.Len()
	for i := 0; i < aLen; i++ {
		aE := aV.Index(i)
		aIE := aE.Interface().(KeyStringValuer)

		kernel(aIE, bMap[aIE.KeyStringValue()])
	}
}

func GetSubtraction(aI, bI interface{}) interface{} {
	resultV := reflect.MakeSlice(reflect.ValueOf(aI).Type(), 0, 0)
	wrapper(aI, bI, func(aIE, bIE KeyStringValuer){
		if bI != nil {
			return
		}
		resultV = reflect.Append(resultV, reflect.ValueOf(aIE))
	})
	return resultV.Interface()
}

func GetIntersection(aI, bI interface{}) interface{} {
	resultV := reflect.MakeSlice(reflect.ValueOf(aI).Type(), 0, 0)
	wrapper(aI, bI, func(aIE, bIE KeyStringValuer){
		if bI == nil {
			return
		}
		resultV = reflect.Append(resultV, reflect.ValueOf(aIE))
	})
	return resultV.Interface()
}

