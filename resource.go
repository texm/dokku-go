package dokku

import (
	"fmt"
)

const (
	resourceReportCmd       = "resource:report"
	resourceReportAppCmd    = "resource:report %s"
	resourceLimitCmd        = "resource:limit [--process-type <process-type>] [RESOURCE_OPTS...] <app>"
	resourceLimitClearCmd   = "resource:limit-clear [--process-type <process-type>] <app>"
	resourceReserveCmd      = "resource:reserve [--process-type <process-type>] [RESOURCE_OPTS...] <app>"
	resourceReserveClearCmd = "resource:reserve-clear [--process-type <process-type>] <app>"
)

func (c *DefaultClient) GetAppResourceReport(appName string) (*ProcessReport, error) {
	cmd := fmt.Sprintf(resourceReportAppCmd, appName)
	output, err := c.Exec(cmd)
	fmt.Printf("output: ;%s;\n", output)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

type ResourceReport map[string]*ProcessReport

func (c *DefaultClient) GetAllAppResourceReport() (*ProcessReport, error) {
	output, err := c.Exec(resourceReportCmd)
	fmt.Printf("output: ;%s;\n", output)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
