package data

import (
	"reflect"
	"testing"
)

func TestNewPerson(t *testing.T) {
	type args struct {
		name    string
		surname string
	}
	tests := []struct {
		name    string
		args    args
		want    *Person
		wantErr bool
	}{
		{"standard", args{ "John", "Doe"}, &Person{	"John", "Doe"}, false},
		{"no-name", args{ "", "Doe"}, nil, true},
		{"no-surname", args{ "John", ""}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPerson(tt.args.name, tt.args.surname)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPerson() got = %v, want %v", got, tt.want)
			}
		})
	}
}
