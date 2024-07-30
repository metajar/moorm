package moorm

import (
	"reflect"
	"testing"
)

func TestFilterBuilder_Regex(t *testing.T) {
	type fields struct {
		selector Document
	}
	type args struct {
		field string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterBuilder
	}{
		{name: "String Test", fields: fields{selector: map[string]M{}}, args: args{
			field: "name",
			value: ".*bob",
		}, want: &FilterBuilder{map[string]M{"name": {"$regex": ".*bob"}}}},
		{name: "Int Test", fields: fields{selector: map[string]M{}}, args: args{
			field: "name",
			value: 10,
		}, want: &FilterBuilder{map[string]M{"name": {"$regex": 10}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FilterBuilder{
				selector: tt.fields.selector,
			}
			if got := f.Regex(tt.args.field, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Eq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterBuilder_Eq(t *testing.T) {
	type fields struct {
		selector Document
	}
	type args struct {
		field string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterBuilder
	}{
		{name: "String Test", fields: fields{selector: map[string]M{}}, args: args{
			field: "name",
			value: "bob",
		}, want: &FilterBuilder{map[string]M{"name": {"$eq": "bob"}}}},
		{name: "Int Test", fields: fields{selector: map[string]M{}}, args: args{
			field: "name",
			value: 10,
		}, want: &FilterBuilder{map[string]M{"name": {"$eq": 10}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FilterBuilder{
				selector: tt.fields.selector,
			}
			if got := f.Eq(tt.args.field, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Eq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterBuilder_Ne(t *testing.T) {
	type fields struct {
		selector Document
	}
	type args struct {
		field string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterBuilder
	}{
		{name: "String Test", fields: fields{selector: map[string]M{}}, args: args{
			field: "name",
			value: "bob",
		}, want: &FilterBuilder{map[string]M{"name": {"$ne": "bob"}}}},
		{name: "Int Test", fields: fields{selector: map[string]M{}}, args: args{
			field: "name",
			value: 10,
		}, want: &FilterBuilder{map[string]M{"name": {"$ne": 10}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FilterBuilder{
				selector: tt.fields.selector,
			}
			if got := f.Ne(tt.args.field, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterBuilder_Gt(t *testing.T) {
	type fields struct {
		selector Document
	}
	type args struct {
		field string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FilterBuilder
	}{
		{name: "String Test", fields: fields{selector: map[string]M{}}, args: args{
			field: "name",
			value: 100,
		}, want: &FilterBuilder{map[string]M{"name": {"$gt": 100}}}},
		{name: "Int Test", fields: fields{selector: map[string]M{}}, args: args{
			field: "name",
			value: 10,
		}, want: &FilterBuilder{map[string]M{"name": {"$gt": 10}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FilterBuilder{
				selector: tt.fields.selector,
			}
			if got := f.Gt(tt.args.field, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Gt() = %v, want %v", got, tt.want)
			}
		})
	}
}
