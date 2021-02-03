package file

import (
	"go-search/pkg/store"
	"testing"
)

// goos: linux
// goarch: amd64
// 3900x / 3200 MHz CL16 RAM
// BenchmarkService_Add/0_entries-24         	1000000000	         0.000001 ns/op	       0 B/op	       0 allocs/op
// BenchmarkService_Add/1,000_entries-24     	1000000000	         0.000055 ns/op	       0 B/op	       0 allocs/op
// BenchmarkService_Add/100,000_entries-24   	1000000000	         0.00545 ns/op	       0 B/op	       0 allocs/op
func BenchmarkService_Add(b *testing.B) {
	type fields struct {
		tree *tree
	}
	type args struct {
		doc      *store.Doc
		assignID bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "0 entries",
			fields: fields{tree: new(tree)},
			args: args{
				doc:      &store.Doc{URL: "https://katata.games", Title: "Katata Games"},
				assignID: true,
			},
		},
		{
			name:   "1,000 entries",
			fields: fields{tree: prepareTree(1_000)},
			args: args{
				doc:      &store.Doc{URL: "https://katata.games", Title: "Katata Games"},
				assignID: true,
			},
		},
		{
			name:   "100,000 entries",
			fields: fields{tree: prepareTree(100_000)},
			args: args{
				doc:      &store.Doc{URL: "https://katata.games", Title: "Katata Games"},
				assignID: true,
			},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			s := &Service{
				tree: tt.fields.tree,
			}
			_ = s.Add(tt.args.doc, tt.args.assignID)
		})
	}
}

// goos: linux
// goarch: amd64
// 3900x / 3200 MHz CL16 RAM
// BenchmarkService_Find/10_entries-24         	1000000000	         0.000000 ns/op	       0 B/op	       0 allocs/op
// BenchmarkService_Find/1,000_entries-24      	1000000000	         0.000051 ns/op	       0 B/op	       0 allocs/op
// BenchmarkService_Find/100,000_entries-24    	1000000000	         0.00369 ns/op	       0 B/op	       0 allocs/op
func BenchmarkService_Find(b *testing.B) {
	type fields struct {
		tree *tree
	}
	type args struct {
		id uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "10 entries",
			fields: fields{tree: prepareTree(10)},
			args:   args{id: 9},
		},
		{
			name:   "1,000 entries",
			fields: fields{tree: prepareTree(1_000)},
			args:   args{id: 999},
		},
		{
			name:   "100,000 entries",
			fields: fields{tree: prepareTree(100_000)},
			args:   args{id: 99_999},
		},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			s := &Service{
				tree: tt.fields.tree,
			}
			_ = s.Find(tt.args.id)
		})
	}
}

func prepareTree(entries int) *tree {
	s := Service{
		new(tree),
	}
	for i := 0; i < entries; i++ {
		s.Add(&store.Doc{URL: "https://katata.games", Title: "Katata Games"}, true)
	}
	return s.tree
}
