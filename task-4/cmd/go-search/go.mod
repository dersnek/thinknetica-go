module dersnek/task-4/cmd/go-search

go 1.15

replace dersnek/task-4/pkg/crawler v0.0.0 => ../../pkg/crawler
replace dersnek/task-4/pkg/crawler/spider v0.0.0 => ../../pkg/crawler/spider
replace dersnek/task-4/pkg/engine v0.0.0 => ../../pkg/engine
replace dersnek/task-4/pkg/index v0.0.0 => ../../pkg/index

require dersnek/task-4/pkg/crawler v0.0.0
require dersnek/task-4/pkg/crawler/spider v0.0.0
require dersnek/task-4/pkg/engine v0.0.0
require dersnek/task-4/pkg/index v0.0.0
