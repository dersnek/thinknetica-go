module dersnek/task-6/cmd/go-search

go 1.15

replace dersnek/task-6/pkg/binary v0.0.0 => ../../pkg/binary
replace dersnek/task-6/pkg/crawler v0.0.0 => ../../pkg/crawler
replace dersnek/task-6/pkg/crawler/spider v0.0.0 => ../../pkg/crawler/spider
replace dersnek/task-6/pkg/engine v0.0.0 => ../../pkg/engine
replace dersnek/task-6/pkg/index v0.0.0 => ../../pkg/index
replace dersnek/task-6/pkg/store v0.0.0 => ../../pkg/store
replace dersnek/task-6/pkg/store/file v0.0.0 => ../../pkg/store/file

require dersnek/task-6/pkg/crawler/spider v0.0.0
require dersnek/task-6/pkg/engine v0.0.0
