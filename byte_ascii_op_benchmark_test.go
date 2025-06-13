package byte_ascii_op_benchmark

import (
	"testing"
)

func BenchmarkGenAesKeyWithMap(b *testing.B) {
	initKey := "pengmingadfdlfuedfklisdlfbfppwufwe"
	var aesKey, key string
	for i := 0; i < b.N; i++ {
		aesKey, key = GenAesKeyWithMap(initKey)
	}
	_, _ = aesKey, key
}

func BenchmarkGenAesKeyWitthSlice(b *testing.B) {
	initKey := "pengmingadfdlfuedfklisdlfbfppwufwe"
	var aesKey, key string
	for i := 0; i < b.N; i++ {
		aesKey, key = GenAesKeyWithSlice(initKey)
	}
	_, _ = aesKey, key
}

func TestGenAesKeyWithMap(t *testing.T) {
	type args struct {
		initKey string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "TestGenAesKeyWithMap",
			args: args{
				initKey: "pengmingadfdlfuedfklisdlfbfppwufwe",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAesKey, gotKey := GenAesKeyWithMap(tt.args.initKey)
			t.Logf("GenAesKeyWithMap() gotAesKey = %v gotKey = %v", gotAesKey, gotKey)
		})
	}
}

func TestGenAesKeyWithSlice(t *testing.T) {
	type args struct {
		initKey string
	}
	tests := []struct {
		name       string
		args       args
		wantAesKey string
		wantKey    string
	}{
		{
			name: "TestGenAesKeyWithSlice",
			args: args{
				initKey: "pengmingadfdlfuedfklisdlfbfppwufwe",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAesKey, gotKey := GenAesKeyWithSlice(tt.args.initKey)
			t.Logf("GenAesKeyWithSlice() gotAesKey = %v gotKey = %v", gotAesKey, gotKey)
		})
	}
}
