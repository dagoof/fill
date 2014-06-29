package fill

import (
	"errors"
	"reflect"
)

func Fill(dst, src interface{}) error {
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

	if !st.ConvertibleTo(dt) {
		return errors.New("src type not convertible to dst type")
	}

	s = s.Convert(dt)
	st = s.Type()

	if !st.AssignableTo(dt) {
		return errors.New("src type not assignable to dst type")
	}

	d.Set(s)
	return nil
}
