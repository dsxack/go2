package collect

import "reflect"

// The SliceMapWithKeys func iterates through the slice and passes each key and value to the given mapFunc.
// The mapFunc should return a single key / value pair
func SliceMapWithKeys(value interface{}, mapFunc interface{}) interface{} {
	refFunc := reflect.ValueOf(mapFunc)
	refSlice := reflect.ValueOf(value)
	resMap := reflect.MakeMap(
		reflect.MapOf(
			refFunc.Type().Out(0),
			refFunc.Type().Out(1),
		),
	)

	for i := 0; i < refSlice.Len(); i += 1 {
		value := refSlice.Index(i)
		refResults := refFunc.Call([]reflect.Value{reflect.ValueOf(i), value})
		resMap.SetMapIndex(refResults[0], refResults[1])
	}

	return resMap.Interface()
}
