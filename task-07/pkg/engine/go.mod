module dersnek/task-7/pkg/engine

go 1.15

replace dersnek/task-7/pkg/index v0.0.0 => ../index
replace dersnek/task-7/pkg/index/hash v0.0.0 => ../index/hash
replace dersnek/task-7/pkg/crawler v0.0.0 => ../crawler
replace dersnek/task-7/pkg/crawler/spider v0.0.0 => ../crawler/spider
replace dersnek/task-7/pkg/store v0.0.0 => ../store
replace dersnek/task-7/pkg/store/file v0.0.0 => ../store/file

require dersnek/task-7/pkg/index v0.0.0
require dersnek/task-7/pkg/index/hash v0.0.0
require dersnek/task-7/pkg/crawler v0.0.0
require dersnek/task-7/pkg/crawler/spider v0.0.0
require dersnek/task-7/pkg/store v0.0.0
require dersnek/task-7/pkg/store/file v0.0.0
