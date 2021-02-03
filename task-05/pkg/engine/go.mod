module dersnek/task-5/pkg/engine

go 1.15

replace dersnek/task-5/pkg/b v0.0.0 => ../../pkg/b
replace dersnek/task-5/pkg/index v0.0.0 => ../index
replace dersnek/task-5/pkg/w v0.0.0 => ../w

require dersnek/task-5/pkg/index v0.0.0
require dersnek/task-5/pkg/w v0.0.0
