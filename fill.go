package fill

import (
	"errors"
	"reflect"
)

type option uint64

const (
	convert option = 1 << iota
)

func fill(options option, dst, src interface{}) error {
	d := reflect.ValueOf(dst)
	s := reflect.ValueOf(src)

	if d.Kind() != reflect.Ptr {
		return errors.New("dst is not ptr")
	}

	d = d.Elem()

	if !d.CanSet() {
		return errors.New("cannot set dst")
	}

	st := s.Type()
	dt := d.Type()

	if options&convert > 0 {
		if !st.ConvertibleTo(dt) {
			return errors.New("src type not convertible to dst type")
		}
		s = s.Convert(dt)
		st = s.Type()
	}

	if !st.AssignableTo(dt) {
		return errors.New("src type not assignable to dst type")
	}

	d.Set(s)
	return nil
}

func Fill(dst, src interface{}) error {
	return fill(0, dst, src)
}

func ConvertFill(dst, src interface{}) error {
	return fill(convert, dst, src)
}
