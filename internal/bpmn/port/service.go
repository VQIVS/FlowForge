package port

import (
	"context"

	"github.com/VQIVS/FlowForge/internal/bpmn/domain"
)

type Service interface {
	CreateWorkflow(ctx context.Context, workflow domain.BPMNWorkflow) error
	GetWorkflow(ctx context.Context, ID string) (*domain.BPMNWorkflow, error)
	UpdateWorkflow(ctx context.Context, workflow domain.BPMNWorkflow) error
	DeleteWorkflow(ctx context.Context, ID string) error
	ListWorkflows(ctx context.Context) ([]domain.BPMNWorkflow, error)
}
