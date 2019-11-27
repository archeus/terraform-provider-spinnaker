package provider

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jgramoll/terraform-provider-spinnaker/client"
)

// Trigger for Pipeline
type trigger struct {
	ID           string
	Enabled      bool
	Job          string
	Master       string
	PropertyFile string `mapstructure:"property_file"`
	RunAsUser    string `mapstructure:"run_as_user"`
	Type         string
	// Docker
	Account      string
	Organization string
	Registry     string
	Repository   string
	Tag          string `mapstructure:"tag"`
}

func fromClientTrigger(clientTrigger *client.Trigger) *trigger {
	t := trigger(*clientTrigger)
	return &t
}

func (t *trigger) setResourceData(d *schema.ResourceData) error {
	err := d.Set("enabled", t.Enabled)
	if err != nil {
		return err
	}

	if t.Type == "jenkins" {
		err = d.Set("job", t.Job)
		if err != nil {
			return err
		}
		err = d.Set("master", t.Master)
		if err != nil {
			return err
		}
		err = d.Set("property_file", t.PropertyFile)
		if err != nil {
			return err
		}
	} else if t.Type == "docker" {
		// account      = "docker-hub"
		// organization = "library"
		// registry     = "index.docker.io"
		// repository   = "library/nginx"
		err = d.Set("account", t.Account)
		if err != nil {
			return err
		}
		err = d.Set("organization", t.Organization)
		if err != nil {
			return err
		}
		err = d.Set("registry", t.Registry)
		if err != nil {
			return err
		}
		err = d.Set("repository", t.Repository)
		if err != nil {
			return err
		}
		err = d.Set("tag", t.Tag)
		if err != nil {
			return err
		}
	}
	return d.Set("type", t.Type)
}
