package tosca

import "reflect"

type EquallableTypeRoot struct {
	Equallable
}

func (value EquallableTypeRoot) Equals(arg Equallable) bool {
	v := reflect.ValueOf(value).Elem()
	a := reflect.ValueOf(arg).Elem()
	if v.NumField() != a.NumField() { // different number of fields makes them unequal
		return false
	}
	if v.Type() != a.Type() { // different types make them unequal
		return false
	}
	for i := 0; i < v.NumField(); i++ { // for each field of value
		varName := v.Type().Field(i).Name
		varType := v.Type().Field(i).Type
		varValue := v.Field(i).Interface()

		if !(varName == "DerivedFrom" || varName == "Version" || varName == "Metadata" || varName == "Description") { // This fields should not be derived and I assume, they are thus not important for comparison // TODO is that true?
			if a.Type().Field(i).Name != varName || // if field has different name
				a.Type().Field(i).Type != varType || // or has different type
				a.Field(i).Interface() != varValue { // or has different value
				return false // they are unequal
			}
		}
	}
	return true
}

func (value EquallableTypeRoot) ContainedIn(list []Equallable) bool {
	for _, element := range list {
		if value.Equals(element) {
			return true
		}
	}
	return false
}

type ComparableTypeRoot struct {
	// Equallable is set indirectly via Comparable
	Comparable
}

func (value ComparableTypeRoot) GreaterOrEqual(arg Comparable) bool {
	return value.Equals(arg) || value.GreaterThan(arg) // if equal or greater
}
func (value ComparableTypeRoot) LessThan(arg Comparable) bool {
	return !value.Equals(arg) && !value.GreaterThan(arg) // if not equal and not greater
}
func (value ComparableTypeRoot) LessOrEqual(arg Comparable) bool {
	return value.Equals(arg) || value.LessThan(arg) // if equal or less
}
func (value ComparableTypeRoot) InRange(lowerBound Comparable, upperBound Comparable) bool { // "inclusive"
	return value.GreaterOrEqual(lowerBound) && value.LessOrEqual(upperBound)
}
