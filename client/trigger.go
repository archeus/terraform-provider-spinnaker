package client

import (
	"errors"
)

// ErrTriggerNotFound trigger not found
var ErrTriggerNotFound = errors.New("could not find trigger")

// Trigger for Pipeline
type Trigger struct {
	ID           string `json:"id"`
	Enabled      bool   `json:"enabled"`
	Job          string `json:"job,omitempty"`
	Master       string `json:"master,omitempty"`
	PropertyFile string `json:"propertyFile,omitempty"`
	RunAsUser    string `json:"runAsUser,omitempty"`
	Type         string `json:"type"`
	Account      string `json:"account,omitempty"`
	Organization string `json:"organization,omitempty"`
	Registry     string `json:"registry,omitempty"`
	Repository   string `json:"repository,omitempty"`
	Tag          string `json:"tag,omitempty""`
}

// GetTrigger by ID
func (p *Pipeline) GetTrigger(triggerID string) (*Trigger, error) {
	for _, trigger := range p.Triggers {
		if trigger.ID == triggerID {
			return trigger, nil
		}
	}
	return nil, ErrTriggerNotFound
}

// AppendTrigger append trigger
func (p *Pipeline) AppendTrigger(trigger *Trigger) {
	p.Triggers = append(p.Triggers, trigger)
}

// UpdateTrigger in pipeline
func (p *Pipeline) UpdateTrigger(trigger *Trigger) error {
	for i, t := range p.Triggers {
		if t.ID == trigger.ID {
			p.Triggers[i] = trigger
			return nil
		}
	}
	return ErrTriggerNotFound
}

// DeleteTrigger in pipeline
func (p *Pipeline) DeleteTrigger(triggerID string) error {
	for i, t := range p.Triggers {
		if t.ID == triggerID {
			p.Triggers = append(p.Triggers[:i], p.Triggers[i+1:]...)
			return nil
		}
	}
	return ErrTriggerNotFound
}
