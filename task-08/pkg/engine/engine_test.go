package engine

import (
	"dersnek/task-8/pkg/index"
	"dersnek/task-8/pkg/index/fakeindex"
	"dersnek/task-8/pkg/store"
	"dersnek/task-8/pkg/store/memstore"
	"reflect"
	"testing"
)

func TestService_Find(t *testing.T) {
	type fields struct {
		ind index.Interface
	}
	type args struct {
		q  string
		st store.Interface
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []store.Doc
	}{
		{
			name:   "One result",
			fields: fields{ind: fakeindex.New()},
			args:   args{q: "a", st: memstore.New()},
			want:   []store.Doc{store.Doc{URL: "https://katata.games", Title: "Katata Games"}},
		},
		{
			name:   "Several results",
			fields: fields{ind: fakeindex.New()},
			args:   args{q: "b", st: memstore.New()},
			want: []store.Doc{
				store.Doc{URL: "https://katata.games", Title: "Katata Games"},
				store.Doc{URL: "https://go.dev", Title: "Go"},
			},
		},
		{
			name:   "No results",
			fields: fields{ind: fakeindex.New()},
			args:   args{q: "z", st: memstore.New()},
			want:   []store.Doc{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				ind: tt.fields.ind,
			}
			if got := s.Find(tt.args.q, tt.args.st); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
