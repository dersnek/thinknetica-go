// Package spider реализует сканер содержимого веб-сайтов.
// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его URL.
package spider

import (
	"testing"
)

func TestService_Scan(t *testing.T) {
	type args struct {
		url   string
		depth int
	}
	tests := []struct {
		name string
		s    *Service
		args args
		want int
	}{
		{
			name: "Default",
			args: args{url: "https://katata.games", depth: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			got, err := s.Scan(tt.args.url, tt.args.depth)
			if err != nil {
				t.Errorf("Service.Scan() error = %v", err)
				return
			}
			if len(got) != tt.want {
				t.Errorf("Service.Scan() = %v, want %v", got, tt.want)
			}
		})
	}
}
