package errset

// ErrSet represents a list of errors. 
//
// Example:
//
//     errs := ErrSet{}
//     for ... {
//       if err := DoSomething(i); err != nil {
//         errs = append(errs, err)
//       }
//     }
//     return errs.ReturnValue()
//
type ErrSet []error

// ReturnValue returns the ErrSet object if at least one error is present or
// nil if there are no errors
func (es ErrSet) ReturnValue() error {
	if len(es) == 0 {
		return nil
	}
	return es
}

func (es ErrSet) Error() string {
	rv := ""
	for i, err := range es {
		if i != 0 {
			rv += "; "
		}
		rv += err.Error()
	}
	return rv
}
