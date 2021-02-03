module dersnek/task-7/cmd/go-search

go 1.15

replace dersnek/task-7/pkg/binary v0.0.0 => ../../pkg/binary
replace dersnek/task-7/pkg/crawler v0.0.0 => ../../pkg/crawler
replace dersnek/task-7/pkg/crawler/spider v0.0.0 => ../../pkg/crawler/spider
replace dersnek/task-7/pkg/engine v0.0.0 => ../../pkg/engine
replace dersnek/task-7/pkg/index v0.0.0 => ../../pkg/index
replace dersnek/task-7/pkg/index/hash v0.0.0 => ../../pkg/index/hash
replace dersnek/task-7/pkg/store v0.0.0 => ../../pkg/store
replace dersnek/task-7/pkg/store/file v0.0.0 => ../../pkg/store/file

require (
	dersnek/task-7/pkg/crawler v0.0.0
	dersnek/task-7/pkg/crawler/spider v0.0.0
	dersnek/task-7/pkg/engine v0.0.0
	dersnek/task-7/pkg/index v0.0.0
	dersnek/task-7/pkg/index/hash v0.0.0
	dersnek/task-7/pkg/store v0.0.0
	dersnek/task-7/pkg/store/file v0.0.0
)
