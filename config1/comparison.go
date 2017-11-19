package config1
//
//import "reflect"
//
//func Equal(x, y interface{}, opts ...Option) bool {
//	s := newState(opts)
//	s.compareAny(reflect.ValueOf(x), reflect.ValueOf(y))
//	return s.result.Equal()
//}
//
//func Diff(x, y interface{}, opts ...Option) string {
//	r := new(defaultReporter)
//	opts = Options{Options(opts), r}
//	eq := Equal(x, y, opts...)
//	d := r.String()
//	if (d == "") != eq {
//		panic("inconsistent difference and equality results")
//	}
//	return d
//}
//
//
//type state struct {
//	// These fields represent the "comparison state".
//	// Calling statelessCompare must not result in observable changes to these.
//	result   diff.Result // The current result of comparison
//	curPath  Path        // The current path in the value tree
//	reporter reporter    // Optional reporter used for difference formatting
//
//	// dynChecker triggers pseudo-random checks for option correctness.
//	// It is safe for statelessCompare to mutate this value.
//	dynChecker dynChecker
//
//	// These fields, once set by processOption, will not change.
//	exporters map[reflect.Type]bool // Set of structs with unexported field visibility
//	opts      Options               // List of all fundamental and filter options
//}
//
//func newState(opts []Option) *state {
//	s := new(state)
//	for _, opt := range opts {
//		s.processOption(opt)
//	}
//	return s
//}

