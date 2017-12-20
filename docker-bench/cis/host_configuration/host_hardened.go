package hostconfiguration

import (
	"bitbucket.org/stack-rox/apollo/docker-bench/utils"
	"bitbucket.org/stack-rox/apollo/pkg/api/generated/api/v1"
)

type hostHardened struct{}

func (c *hostHardened) Definition() utils.Definition {
	return utils.Definition{
		CheckDefinition: v1.CheckDefinition{
			Name:        "CIS 1.2",
			Description: "Ensure the container host has been Hardened",
		},
	}
}

func (c *hostHardened) Run() (result v1.CheckResult) {
	utils.Note(&result)
	utils.AddNotes(&result, "Ensuring the host is hardened with the latest kernel requires manual introspection")
	return
}

// NewHostHardened implements CIS-1.2
func NewHostHardened() utils.Check {
	return &hostHardened{}
}
