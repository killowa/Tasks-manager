package services

// import "tasks_manager/models"

// type Executor struct {
// 	workerPool WorkerPool
// 	tasks      []models.Task
// }

// func (r *Executor) Start() {
// 	go func() {
// 		for {
// 			worker, error := r.workerPool.GetResource()

// 			if len(r.tasks) != 0 && error == nil {
// 				worker.Add(r.tasks[0])
// 				r.tasks = r.tasks[1:]
// 			}
// 		}
// 	}()
// }
