module dersnek/task-7/pkg/index/hash

go 1.15

replace dersnek/task-7/pkg/index v0.0.0 => ../
replace dersnek/task-7/pkg/store v0.0.0 => ../../store
replace dersnek/task-7/pkg/store/file v0.0.0 => ../../store/file

require (
  dersnek/task-7/pkg/index v0.0.0
	dersnek/task-7/pkg/store v0.0.0
  dersnek/task-7/pkg/store/file v0.0.0
)
