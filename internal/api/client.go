package api

import (
	"github.com/thilob97/hubporter/internal/models"
)

type Client struct {
	accessToken string
}

func NewClient(token string) *Client {
	return &Client{
		accessToken: token,
	}
}

// GetWorkflows returns a list of all workflows (currently mock data)
func (c *Client) GetWorkflows() ([]models.Workflow, error) {
	// TODO: Implement real API call
	return []models.Workflow{
		{
			Name:    "Test Workflow 1",
			ID:      1,
			Type:    "standard",
			Enabled: true,
		},
		{
			Name:    "Test Workflow 2",
			ID:      2,
			Type:    "drip",
			Enabled: false,
		},
	}, nil
}

// GetWorkflowByID returns a single workflow by ID (currently mock data)
func (c *Client) GetWorkflowByID(id int) (*models.Workflow, error) {
	// TODO: Implement real API call
	return &models.Workflow{
		Name:    "Test Workflow",
		ID:      id,
		Type:    "standard",
		Enabled: true,
	}, nil
}
