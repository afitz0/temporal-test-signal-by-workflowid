package testing

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	id := "custom-workflow-id"
	env.SetStartWorkflowOptions(client.StartWorkflowOptions{
		ID: id,
	})

	env.RegisterDelayedCallback(func() {
		err := env.SignalWorkflowByID(id, "cancel", nil)
		require.NoError(t, err)
	}, time.Second)

	env.ExecuteWorkflow(Workflow)

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
}
