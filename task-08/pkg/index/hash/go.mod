module dersnek/task-8/pkg/index/hash

go 1.15

replace dersnek/task-8/pkg/index v0.0.0 => ../
replace dersnek/task-8/pkg/store v0.0.0 => ../../store
replace dersnek/task-8/pkg/store/memstore v0.0.0 => ../../store/memstore
replace dersnek/task-8/pkg/store/file v0.0.0 => ../../store/file

require (
  dersnek/task-8/pkg/index v0.0.0
	dersnek/task-8/pkg/store v0.0.0
  dersnek/task-8/pkg/store/memstore v0.0.0
  dersnek/task-8/pkg/store/file v0.0.0
)
