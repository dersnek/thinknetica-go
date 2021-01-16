// Package hash indexes, stores and retrieves webpages
package hash

import (
	"dersnek/task-8/pkg/index"
	"dersnek/task-8/pkg/store"
	"dersnek/task-8/pkg/store/memstore"
	"reflect"
	"testing"
)

func TestService_Build(t *testing.T) {
	type fields struct {
		index index.Data
	}
	type args struct {
		docs  []store.Doc
		store store.Interface
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantRtrn []store.Doc
		wantData index.Data
	}{
		{
			name:   "Default",
			fields: fields{index: make(index.Data)},
			args: args{
				docs: []store.Doc{
					store.Doc{URL: "https://katata.games", Title: "Katata Games"},
					store.Doc{URL: "https://go.dev", Title: "Go"},
				},
				store: memstore.New(),
			},
			wantRtrn: []store.Doc{
				store.Doc{URL: "https://katata.games", Title: "Katata Games"},
				store.Doc{URL: "https://go.dev", Title: "Go"},
			},
			wantData: index.Data{
				"katata": []uint{0},
				"games":  []uint{0},
				"go":     []uint{0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				index: tt.fields.index,
			}
			got := s.Build(tt.args.docs, tt.args.store)
			if !reflect.DeepEqual(got, tt.wantRtrn) {
				t.Errorf("Service.Build() = %v, want %v", got, tt.wantRtrn)
			}
			if !reflect.DeepEqual(s.index, tt.wantData) {
				t.Errorf("Service.index = %v, want %v", s.index, tt.wantData)
			}
		})
	}
}

func TestService_Restore(t *testing.T) {
	type fields struct {
		index index.Data
	}
	type args struct {
		docs  []store.Doc
		data  index.Data
		store store.Interface
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   index.Data
	}{
		{
			name:   "Default",
			fields: fields{index: make(index.Data)},
			args: args{
				docs: []store.Doc{
					store.Doc{URL: "https://katata.games", Title: "Katata Games"},
					store.Doc{URL: "https://go.dev", Title: "Go"},
				},
				data: map[string][]uint{
					"katata": []uint{3, 5},
					"games":  []uint{4, 6},
				},
				store: memstore.New(),
			},
			want: index.Data{
				"katata": []uint{3, 5},
				"games":  []uint{4, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				index: tt.fields.index,
			}
			s.Restore(tt.args.docs, tt.args.data, tt.args.store)
			if !reflect.DeepEqual(s.index, tt.want) {
				t.Errorf("Service.index = %v, want %v", s.index, tt.want)
			}
		})
	}
}

func TestService_Find(t *testing.T) {
	type fields struct {
		index index.Data
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []uint
		want1  bool
	}{
		{
			name:   "Found",
			fields: fields{index: map[string][]uint{"katata": []uint{0}, "games": []uint{1, 2}}},
			args: args{
				word: "games",
			},
			want:  []uint{1, 2},
			want1: true,
		},
		{
			name:   "Not Found",
			fields: fields{index: map[string][]uint{"katata": []uint{0}, "games": []uint{1, 2}}},
			args: args{
				word: "invalid",
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				index: tt.fields.index,
			}
			got, got1 := s.Find(tt.args.word)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Find() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Service.Find() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
