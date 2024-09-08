package arrayutils

import "reflect"

func AddToArray(slice interface{}, value interface{}) interface{} {
	s := reflect.ValueOf(slice)
	newSlice := reflect.Append(s, reflect.ValueOf(value))
	return newSlice.Interface()
}

func RemoveFromArray(slice interface{}, value interface{}) interface{} {
	s := reflect.ValueOf(slice)
	newSlice := reflect.MakeSlice(s.Type(), 0, s.Len())

	for i := 0; i < s.Len(); i++ {
		if !reflect.DeepEqual(s.Index(i).Interface(), value) {
			newSlice = reflect.Append(newSlice, s.Index(i))
		}
	}

	return newSlice.Interface()
}
