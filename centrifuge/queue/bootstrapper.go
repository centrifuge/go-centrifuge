package queue

import "errors"

const BootstrappedQueuedTasks string = "BootstrappedQueuedTasks"

type Bootstrapper struct {
}

func (*Bootstrapper) Bootstrap(context map[string]interface{}) error {
	// to see how BootstrappedQueuedTasks get populated check usages of InstallQueuedTask
	if queuedTasks, ok := context[BootstrappedQueuedTasks]; ok {
		if queuedTasksTyped, ok := queuedTasks.([]QueuedTask); ok {
			InitQueue(queuedTasksTyped)
			return nil
		}
	}
	return errors.New("could not find the list of " + BootstrappedQueuedTasks)
}

func (b *Bootstrapper) TestBootstrap(context map[string]interface{}) error {
	return b.Bootstrap(context)
}

func InstallQueuedTask(context map[string]interface{}, taskCreator func() QueuedTask) error {
	// the following code will add a queued task to the context so that when the queue initializes it can update it self
	// with different tasks types queued in the node
	if queuedTasks, ok := context[BootstrappedQueuedTasks]; ok {
		if queuedTasksTyped, ok := queuedTasks.([]QueuedTask); ok {
			context[BootstrappedQueuedTasks] = append(queuedTasksTyped, taskCreator())
			return nil
		} else {
			return errors.New(BootstrappedQueuedTasks + " is of an unexpected type")
		}
	} else {
		context[BootstrappedQueuedTasks] = []QueuedTask{taskCreator()}
		return nil
	}
}
