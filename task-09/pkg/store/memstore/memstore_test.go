package memstore

import (
	"go-search/pkg/store"
	"testing"
)

// goos: linux
// goarch: amd64
// 3900x / 3200 MHz CL16 RAM
// BenchmarkService_Add/0_entries-24         	1000000000	         0.000001 ns/op	       0 B/op	       0 allocs/op
// BenchmarkService_Add/1,000_entries-24     	1000000000	         0.000000 ns/op	       0 B/op	       0 allocs/op
// BenchmarkService_Add/100,000_entries-24   	1000000000	         0.000000 ns/op	       0 B/op	       0 allocs/op
func BenchmarkService_Add(b *testing.B) {
	type fields struct {
		Interface store.Interface
		Memory    []*store.Doc
	}
	type args struct {
		doc *store.Doc
		in1 bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "0 entries",
			args: args{
				doc: &store.Doc{URL: "https://katata.games", Title: "Katata Games"},
			},
		},
		{
			name:   "1,000 entries",
			fields: fields{Memory: prepareArray(1_000)},
			args: args{
				doc: &store.Doc{URL: "https://katata.games", Title: "Katata Games"},
			},
		},
		{
			name:   "100,000 entries",
			fields: fields{Memory: prepareArray(100_000)},
			args: args{
				doc: &store.Doc{URL: "https://katata.games", Title: "Katata Games"},
			},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			s := &Service{
				Interface: tt.fields.Interface,
				Memory:    tt.fields.Memory,
			}
			_ = s.Add(tt.args.doc, tt.args.in1)
		})
	}
}

// goos: linux
// goarch: amd64
// 3900x / 3200 MHz CL16 RAM
// BenchmarkService_FindInMemory/10_entries-24         	1000000000	         0.000000 ns/op	       0 B/op	       0 allocs/op
// BenchmarkService_FindInMemory/1,000_entries-24      	1000000000	         0.000000 ns/op	       0 B/op	       0 allocs/op
// BenchmarkService_FindInMemory/100,000_entries-24    	1000000000	         0.000000 ns/op	       0 B/op	       0 allocs/op
func BenchmarkService_FindInMemory(b *testing.B) {
	type fields struct {
		Interface store.Interface
		Memory    []*store.Doc
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *store.Doc
	}{
		{
			name:   "10 entries",
			fields: fields{Memory: prepareArray(10)},
			args:   args{id: 9},
		},
		{
			name:   "1,000 entries",
			fields: fields{Memory: prepareArray(1_000)},
			args:   args{id: 999},
		},
		{
			name:   "100,000 entries",
			fields: fields{Memory: prepareArray(100_000)},
			args:   args{id: 99_999},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			s := &Service{
				Interface: tt.fields.Interface,
				Memory:    tt.fields.Memory,
			}
			_ = s.FindInMemory(tt.args.id)
		})
	}
}

func prepareArray(entries int) []*store.Doc {
	s := Service{}
	for i := 0; i < entries; i++ {
		s.Add(&store.Doc{URL: "https://katata.games", Title: "Katata Games"}, true)
	}
	return s.Memory
}
