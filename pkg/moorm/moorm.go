package moorm

func (f *FilterBuilder) Eq(field string, value interface{}) *FilterBuilder {
	return f.addSelector(field, "$eq", value)
}

func (f *FilterBuilder) Ne(field string, value interface{}) *FilterBuilder {
	return f.addSelector(field, "$ne", value)
}

func (f *FilterBuilder) Gt(field string, value interface{}) *FilterBuilder {
	return f.addSelector(field, "$gt", value)
}

func (f *FilterBuilder) Gte(field string, value interface{}) *FilterBuilder {
	return f.addSelector(field, "$gte", value)
}

func (f *FilterBuilder) Regex(field string, value interface{}) *FilterBuilder {
	return f.addSelector(field, "$regex", value)
}

func (f *FilterBuilder) Lt(field string, value interface{}) *FilterBuilder {
	return f.addSelector(field, "$lt", value)
}

func (f *FilterBuilder) Lte(field string, value interface{}) *FilterBuilder {
	return f.addSelector(field, "$lte", value)
}

func (f *FilterBuilder) In(field string, value []interface{}) *FilterBuilder {
	return f.addSelector(field, "$in", value)
}

func (f *FilterBuilder) Nin(field string, value []interface{}) *FilterBuilder {
	return f.addSelector(field, "$nin", value)
}

func Or(filters []Document) map[string]interface{} {
	return map[string]interface{}{"$or": filters}
}

func And(filters []Document) map[string]interface{} {
	return map[string]interface{}{"$and": filters}
}
