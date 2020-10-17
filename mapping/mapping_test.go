package mapping

import (
	"reflect"
	"testing"
)

type testStruct struct {
	Name string
}

func useInterface() interface{} {
	var i interface{}
	i = &testStruct{Name: "TestName"}
	return i
}

func Test_toMap(t *testing.T) {
	tests := []struct {
		Name interface{}
		Want map[string]interface{}
	}{
		{useInterface(), map[string]interface{}{"Name": "TestName"}},
	}
	for _, tt := range tests {
		t.Run("TestName", func(t *testing.T) {
			got, _ := toMap(tt.Name)

			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("toMap() got = %v, want %v", got, tt.Want)
			}
		})
	}
}
