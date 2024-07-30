package moorm

type Document map[string]M

type M map[string]interface{}

// FilterBuilder represents builder for a find operations.
type FilterBuilder struct {
	selector Document
}

// Filter returns a new instance of a FilterBuilder.
func Filter() *FilterBuilder {
	return &FilterBuilder{
		selector: make(Document),
	}
}

func (f *FilterBuilder) addSelector(field string, operator string, value interface{}) *FilterBuilder {
	v, ok := f.selector[field]
	if !ok {
		f.selector[field] = M{operator: value}
		return f
	}

	v[operator] = value
	return f
}

func (f *FilterBuilder) addSelectorWithOptions(field string, operator string, value interface{}, options string) *FilterBuilder {
	v, ok := f.selector[field]
	if !ok {
		if options != "" {
			f.selector[field] = M{operator: value, "$options": options}
			return f
		}
		f.selector[field] = M{operator: value}
		return f
	}

	v[operator] = value
	return f
}

// Build returns document for using in mongodb find operations.
func (f *FilterBuilder) Build() Document {
	return f.selector
}

// Exists adds $exists selector.
func (f *FilterBuilder) Exists(field string, value bool) *FilterBuilder {
	return f.addSelector(field, "$exists", value)
}

// Like basically is equal to like in sql but can be extended however. The main
// thing we do is just lower if string and wrap with wild cards. If you want regex
// use regex.
func (f *FilterBuilder) Like(field string, value interface{}) *FilterBuilder {
	switch value.(type) {
	case string:
		return f.addSelectorWithOptions(field, "$regex", ".*"+value.(string)+".*", "i")
	default:
		return f.addSelector(field, "$regex", value)
	}
}

// NotLike regex of not likeness
func (f *FilterBuilder) NotLike(field string, value interface{}) *FilterBuilder {
	// db.inventory.find( { price: { $not: { $gt: 1.99 } } } )
	nm := make(map[string]interface{})
	switch value.(type) {
	case string:
		nm["$not"] = map[string]interface{}{"$regex": ".*" + value.(string) + ".*"}
		f.selector[field] = nm
		return f
	default:
		nm["$not"] = map[string]interface{}{"$regex": value}
		f.selector[field] = nm
		return f
	}
}

// UpdateBuilder represents builder for an update queries.
type UpdateBuilder struct {
	operations Document
}

// Update returns a new instance of a UpdateBuilder.
func Update() *UpdateBuilder {
	return &UpdateBuilder{
		operations: make(Document),
	}
}

func (u *UpdateBuilder) addOperator(operator, field string, value interface{}) *UpdateBuilder {
	op, ok := u.operations[operator]
	if !ok {
		u.operations[operator] = M{field: value}
		return u
	}
	op[field] = value
	return u
}

// Unset adds $unset operator.
func (u *UpdateBuilder) Unset(field string) *UpdateBuilder {
	return u.addOperator("$unset", field, "")
}

// Build returns document for using in mongodb update operations.
func (u *UpdateBuilder) Build() Document {
	return u.operations
}
