module dersnek/task-7/pkg/crawler/spider

go 1.15

replace dersnek/task-7/pkg/crawler v0.0.0 => ../
replace dersnek/task-7/pkg/store v0.0.0 => ../../store

require dersnek/task-7/pkg/crawler v0.0.0
require dersnek/task-7/pkg/store v0.0.0

require golang.org/x/net v0.0.0-20201031054903-ff519b6c9102
