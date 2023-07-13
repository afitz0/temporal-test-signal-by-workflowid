package testing

import (
	"go.temporal.io/sdk/workflow"
)

func Workflow(ctx workflow.Context) error {
	logger := workflow.GetLogger(ctx)
	info := workflow.GetInfo(ctx)
	logger.Info("Workflow started.", "WorkflowID", info.WorkflowExecution.ID)

	cancelled := false
	selector := workflow.NewSelector(ctx)
	selector.AddReceive(workflow.GetSignalChannel(ctx, "cancel"),
		func(c workflow.ReceiveChannel, _ bool) {
			c.Receive(ctx, nil)
			cancelled = true
		})

	for !cancelled {
		selector.Select(ctx)
	}

	logger.Info("Workflow completed.")
	return nil
}
