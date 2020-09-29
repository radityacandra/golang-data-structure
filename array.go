package datastructure

type arrayInstance []interface{}

func NewArray(size int) arrayInstance {
	arrayInstance := make(arrayInstance, 0, size)

	return arrayInstance
}

func (arr arrayInstance) Get(index int) interface{} {
	return arr[index]
}

func (arr arrayInstance) Set(value interface{}) {
	// check size
	if cap(arr) == len(arr) {
		arr.Resize()
	}

	arr = append(arr, value)
}

func (arr arrayInstance) Insert(value interface{}, index int) {
	// check size
	if cap(arr) == len(arr) {
		arr.Resize()
	}

	rescue := arr[index:]
	arr[index] = value

	for _, v := range rescue {
		arr[index+1] = v
		index++
	}
}

func (arr arrayInstance) Delete(index int) {
	for i := index; i < len(arr) - 1; i++ {
		arr[i] = arr[i+1]
	}
}

func (arr arrayInstance) IsEmpty() bool {
	if len(arr) > 0 {
		return false
	}

	return true
}

func (arr arrayInstance) Contains(value interface{}) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}

	return false
}

func (arr arrayInstance) Resize() {
	arrayInstance2 := make(arrayInstance, 0, len(arr) * 2)
	arrayInstance2 = arr
	arr = arrayInstance2
}

func (arr arrayInstance) Size() int {
	return len(arr)
}
