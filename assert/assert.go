// some useful method to simplify test
package assert

// method used in assert
type TWorker interface {
	Errorf(format string, args ...interface{})
}

func Equals(tw TWorker, expect, actual interface{}) {
	if expect != actual {
		tw.Errorf("Equals fail! expect %v, actual %v", expect, actual)
	}
}

func NotEquals(tw TWorker, expect, actual interface{}) {
	if expect == actual {
		tw.Errorf("NotEquals fail! expect %v, actual %v", expect, actual)
	}
}

func Nil(tw TWorker, element interface{}) {
	if element != nil {
		tw.Errorf("Nil fail! element %v", element)
	}
}

func NotNil(tw TWorker, element interface{}) {
	if element == nil {
		tw.Errorf("NotNil fail! element %v", element)
	}
}
