package tosca

import (
	"errors"
	"reflect"
)

// Validates constraints of <value> according to constraints specified in <dt>
func (dt DataType) ValidateConstraints(value interface{}) error {
	var (
		constraint map[Operator]interface{}
		operator   Operator
		arg        interface{}
	)

	for _, constraint = range dt.Constraints {
		if len(constraint) != 1 {
			return errors.New("only one Operator per Constraint allowed")
		}

		// Even though only one entry exists, this for-loop is the easiest way to retrieve the operator and value
		for operator, arg = range constraint {
			switch operator {
			case OperatorEqual:
				// if len(arg) != 1 {
				// 	errors.New("Invalid number of properties/parameters.")
				// }
				if value != arg {
					return errors.New("constraint not fulfilled")
				}
			case OperatorGreaterThan:
				if typedValue, ok := value.(Comparable); ok {
					if typedArg, ok := arg.(Comparable); ok {
						if !(typedValue.GreaterThan(typedArg)) {
							return errors.New("constraint not fulfilled")
						}
					} else {
						return errors.New("invalid comparison")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorGreatorOrEqual:
				if typedValue, ok := value.(Comparable); ok {
					if typedArg, ok := arg.(Comparable); ok {
						if !(typedValue.GreaterOrEqual(typedArg)) {
							return errors.New("constraint not fulfilled")
						}
					} else {
						return errors.New("invalid comparison")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorLessThan:
				if typedValue, ok := value.(Comparable); ok {
					if typedArg, ok := arg.(Comparable); ok {
						if !(typedValue.LessThan(typedArg)) {
							return errors.New("constraint not fulfilled")
						}
					} else {
						return errors.New("invalid comparison")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorLessOrEqual:
				if typedValue, ok := value.(Comparable); ok {
					if typedArg, ok := arg.(Comparable); ok {
						if !(typedValue.LessOrEqual(typedArg)) {
							return errors.New("constraint not fulfilled")
						}
					} else {
						return errors.New("invalid comparison")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorInRange:
				if typedValue, ok := value.(Comparable); ok {
					if typedArg, ok := arg.([]Comparable); ok {
						if reflect.TypeOf(typedArg[0]) == reflect.TypeOf(typedArg[1]) {
							if reflect.TypeOf(typedValue) == reflect.TypeOf(typedArg[0]) {
								if !(typedValue.InRange(typedArg[0], typedArg[1])) {
									return errors.New("constraint not fulfilled")
								}
							} else {
								return errors.New("bounds not of same type as value")
							}
						} else {
							return errors.New("bounds not of same type")
						}
					} else {
						return errors.New("invalid comparison")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorValidValues:
				if typedValue, ok := value.(Equallable); ok {
					if typedArg, ok := arg.([]Equallable); ok {
						for _, element := range typedArg {
							if reflect.TypeOf(typedValue) == reflect.TypeOf(element) {
								if !(typedValue.ContainedIn(typedArg)) {
									return errors.New("constraint not fulfilled")
								}
							} else {
								return errors.New("validvalue not of same type as value")
							}
						}
					} else {
						return errors.New("invalid arg")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorLength:
				if typedValue, ok := value.(Indexable); ok {
					if typedArg, ok := arg.(Comparable); ok {
						if !(typedValue.LengthEquals(typedArg)) {
							return errors.New("constraint not fulfilled")
						}
					} else {
						return errors.New("invalid arg")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorMinLength:
				if typedValue, ok := value.(Indexable); ok {
					if typedArg, ok := arg.(Comparable); ok {
						if !(typedValue.MinLength(typedArg)) {
							return errors.New("constraint not fulfilled")
						}
					} else {
						return errors.New("invalid arg")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorMaxLength:
				if typedValue, ok := value.(Indexable); ok {
					if typedArg, ok := arg.(Comparable); ok {
						if !(typedValue.MaxLength(typedArg)) {
							return errors.New("constraint not fulfilled")
						}
					} else {
						return errors.New("invalid arg")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorPattern:
				if typedValue, ok := value.(String); ok {
					if typedArg, ok := value.(string); ok {
						typedValue.Pattern(&typedArg)
					} else {
						return errors.New("invalid arg")
					}
				} else {
					return errors.New("invalid operator for this type")
				}
			case OperatorSchema:
				return errors.New("operator 'schema' is not implemented yet")
			default:
				return errors.New("invalid operator")
			}
		}
	}

	return nil
}
