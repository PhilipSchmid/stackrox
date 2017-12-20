package containerruntime

import (
	"strings"

	"bitbucket.org/stack-rox/apollo/docker-bench/utils"
	"bitbucket.org/stack-rox/apollo/pkg/api/generated/api/v1"
)

type userDockerExecBenchmark struct{}

func (c *userDockerExecBenchmark) Definition() utils.Definition {
	return utils.Definition{
		CheckDefinition: v1.CheckDefinition{
			Name:        "CIS 5.23",
			Description: "Ensure docker exec commands are not used with user option",
		}, Dependencies: []utils.Dependency{utils.InitContainers},
	}
}

func (c *userDockerExecBenchmark) Run() (result v1.CheckResult) {
	utils.Pass(&result)
	auditLog, err := utils.ReadFile("/var/log/audit/audit.log")
	if err != nil {
		utils.Warn(&result)
		utils.AddNotef(&result, "Error reading /var/log/audit/audit.log: %+v", err)
		return
	}
	lines := strings.Split(auditLog, "\n")
	for _, line := range lines {
		if strings.Contains(line, "exec") && strings.Contains(line, "user") {
			utils.Warn(&result)
			utils.AddNotef(&result, "docker exec was used with the --user option: '%v'", line)
		}
	}
	return
}

// NewUserDockerExecBenchmark implements CIS-5.23
func NewUserDockerExecBenchmark() utils.Check {
	return &userDockerExecBenchmark{}
}
