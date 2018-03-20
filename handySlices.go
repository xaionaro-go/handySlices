package handySlices

import (
	"fmt"
	"reflect"
)

type KeyStringValuer interface {
	KeyStringValue() string
}
type IsEqualToIer interface {
	IsEqualToI(compareTo IsEqualToIer) bool
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
	wrapper(aI, bI, func(aIE, bIE KeyStringValuer) {
		if bIE != nil {
			return
		}
		resultV = reflect.Append(resultV, reflect.ValueOf(aIE))
	})
	return resultV.Interface()
}

func GetIntersection(aI, bI interface{}) interface{} {
	resultV := reflect.MakeSlice(reflect.ValueOf(aI).Type(), 0, 0)
	wrapper(aI, bI, func(aIE, bIE KeyStringValuer) {
		if bIE == nil {
			return
		}
		resultV = reflect.Append(resultV, reflect.ValueOf(aIE))
	})
	return resultV.Interface()
}

func GetDiffedIntersection(aI, bI interface{}) interface{} {
	resultV := reflect.MakeSlice(reflect.ValueOf(aI).Type(), 0, 0)
	wrapper(aI, bI, func(aIE, bIE KeyStringValuer) {
		if bIE == nil {
			return
		}

		isEqual := false
		isEqualToIer, ok := aIE.(IsEqualToIer)
		if ok {
			isEqual = isEqualToIer.IsEqualToI(bIE.(IsEqualToIer))
		} else {
			isEqual = reflect.DeepEqual(aIE, bIE)
		}

		if !isEqual {
			resultV = reflect.Append(resultV, reflect.ValueOf(aIE))
		}
	})
	return resultV.Interface()
}

func MapToSlice(mI interface{}) interface{} {
	m := reflect.ValueOf(mI)
	keys := m.MapKeys()
	slice := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(mI).Elem()), 0, 0)
	for _, key := range keys {
		item := m.MapIndex(key)
		slice = reflect.Append(slice, item)
	}
	return slice.Interface()
}

func valueStringOf(item interface{}) string {
	return fmt.Sprintf("%v", item)
}

func IsEqualCollections(aI, bI interface{}) bool {
	aV := reflect.ValueOf(aI)
	bV := reflect.ValueOf(bI)
	aL := aV.Len()
	bL := bV.Len()

	if aL != bL {
		return false
	}

	isSetMapA := map[string]bool{}
	for i := 0; i < aL; i++ {
		aE := aV.Index(i)
		aK := valueStringOf(aE.Interface())
		isSetMapA[aK] = true
	}

	isSetMapB := map[string]bool{}
	for i := 0; i < bL; i++ {
		bE := bV.Index(i)
		bK := valueStringOf(bE.Interface())
		if !isSetMapA[bK] {
			return false
		}
		isSetMapB[bK] = true
	}

	for i := 0; i < aL; i++ {
		aE := aV.Index(i)
		aK := valueStringOf(aE.Interface())
		if !isSetMapB[aK] {
			return false
		}
	}

	return true
}
