module dersnek/task-5/cmd/go-search

go 1.15

replace dersnek/task-5/pkg/b v0.0.0 => ../../pkg/b
replace dersnek/task-5/pkg/crawler v0.0.0 => ../../pkg/crawler
replace dersnek/task-5/pkg/crawler/spider v0.0.0 => ../../pkg/crawler/spider
replace dersnek/task-5/pkg/engine v0.0.0 => ../../pkg/engine
replace dersnek/task-5/pkg/index v0.0.0 => ../../pkg/index
replace dersnek/task-5/pkg/w v0.0.0 => ../../pkg/w

require dersnek/task-5/pkg/crawler v0.0.0
require dersnek/task-5/pkg/crawler/spider v0.0.0
require dersnek/task-5/pkg/engine v0.0.0
require dersnek/task-5/pkg/index v0.0.0
require dersnek/task-5/pkg/w v0.0.0
