package storage

import (
	"context"
	"errors"

	"github.com/VQIVS/FlowForge/internal/bpmn/domain"
	"github.com/VQIVS/FlowForge/internal/bpmn/port"
	"gorm.io/gorm"
)

type BPMNRepoPostgres struct {
	db *gorm.DB
}

func NewBPMNRepoPostgres(db *gorm.DB) port.Repo {
	return &BPMNRepoPostgres{db: db}
}

func (r *BPMNRepoPostgres) SaveWorkflow(ctx context.Context, workflow domain.BPMNWorkflow) error {
	err := r.db.WithContext(ctx).Create(&workflow).Error
	if err != nil {
		return errors.New("failed to save BPMN workflow")
	}
	return nil
}

func (r *BPMNRepoPostgres) FindWorkflowByID(ctx context.Context, ID string) (*domain.BPMNWorkflow, error) {
	var workflow domain.BPMNWorkflow
	err := r.db.WithContext(ctx).Preload("Steps").Preload("Actors").First(&workflow, "id = ?", ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("BPMN workflow not found")
		}
		return nil, err
	}
	return &workflow, nil
}

func (r *BPMNRepoPostgres) UpdateWorkflow(ctx context.Context, workflow domain.BPMNWorkflow) error {
	err := r.db.WithContext(ctx).Save(&workflow).Error
	if err != nil {
		return errors.New("failed to update BPMN workflow")
	}
	return nil
}

func (r *BPMNRepoPostgres) DeleteWorkflow(ctx context.Context, ID string) error {
	err := r.db.WithContext(ctx).Where("id = ?", ID).Delete(&domain.BPMNWorkflow{}).Error
	if err != nil {
		return errors.New("failed to delete BPMN workflow")
	}
	return nil
}

func (r *BPMNRepoPostgres) GetAllWorkflows(ctx context.Context) ([]domain.BPMNWorkflow, error) {
	var workflows []domain.BPMNWorkflow
	err := r.db.WithContext(ctx).Preload("Steps").Preload("Actors").Find(&workflows).Error
	if err != nil {
		return nil, errors.New("failed to fetch BPMN workflows")
	}
	return workflows, nil
}

func (r *BPMNRepoPostgres) SaveStep(ctx context.Context, step domain.BPMNStep) error {
	err := r.db.WithContext(ctx).Create(&step).Error
	if err != nil {
		return errors.New("failed to save BPMN step")
	}
	return nil
}

func (r *BPMNRepoPostgres) FindStepsByWorkflowID(ctx context.Context, workflowID string) ([]domain.BPMNStep, error) {
	var steps []domain.BPMNStep
	err := r.db.WithContext(ctx).Where("workflow_id = ?", workflowID).Find(&steps).Error
	if err != nil {
		return nil, errors.New("failed to retrieve BPMN steps")
	}
	return steps, nil
}

func (r *BPMNRepoPostgres) DeleteStep(ctx context.Context, stepID string) error {
	err := r.db.WithContext(ctx).Where("id = ?", stepID).Delete(&domain.BPMNStep{}).Error
	if err != nil {
		return errors.New("failed to delete BPMN step")
	}
	return nil
}

func (r *BPMNRepoPostgres) SaveActor(ctx context.Context, actor domain.BPMNActor) error {
	err := r.db.WithContext(ctx).Create(&actor).Error
	if err != nil {
		return errors.New("failed to save BPMN actor")
	}
	return nil
}

func (r *BPMNRepoPostgres) FindActorsByWorkflowID(ctx context.Context, workflowID string) ([]domain.BPMNActor, error) {
	var actors []domain.BPMNActor
	err := r.db.WithContext(ctx).Where("workflow_id = ?", workflowID).Find(&actors).Error
	if err != nil {
		return nil, errors.New("failed to retrieve BPMN actors")
	}
	return actors, nil
}

func (r *BPMNRepoPostgres) DeleteActor(ctx context.Context, actorID string) error {
	err := r.db.WithContext(ctx).Where("id = ?", actorID).Delete(&domain.BPMNActor{}).Error
	if err != nil {
		return errors.New("failed to delete BPMN actor")
	}
	return nil
}
