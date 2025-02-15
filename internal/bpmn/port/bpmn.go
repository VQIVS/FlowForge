package port

import (
	"context"

	"github.com/VQIVS/FlowForge/internal/bpmn/domain"
)

type Repo interface {
	SaveWorkflow(ctx context.Context, workflow domain.BPMNWorkflow) error
	FindWorkflowByID(ctx context.Context, ID string) (*domain.BPMNWorkflow, error)
	UpdateWorkflow(ctx context.Context, workflow domain.BPMNWorkflow) error
	DeleteWorkflow(ctx context.Context, ID string) error
	GetAllWorkflows(ctx context.Context) ([]domain.BPMNWorkflow, error)

	SaveStep(ctx context.Context, step domain.BPMNStep) error
	FindStepsByWorkflowID(ctx context.Context, workflowID string) ([]domain.BPMNStep, error)
	DeleteStep(ctx context.Context, stepID string) error

	SaveActor(ctx context.Context, actor domain.BPMNActor) error
	FindActorsByWorkflowID(ctx context.Context, workflowID string) ([]domain.BPMNActor, error)
	DeleteActor(ctx context.Context, actorID string) error
}
