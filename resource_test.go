package dokku

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type resourceManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunResourceManagerTestSuite(t *testing.T) {
	suite.Run(t, new(resourceManagerTestSuite))
}

func (s *resourceManagerTestSuite) TestManageAppResources() {
	r := s.Require()
	var err error

	var (
		cpuReserved = &Resource{Type: ResourceCPU, Amount: 1}
		cpuLimit    = &Resource{Type: ResourceCPU, Amount: 2}
		memLimit    = &Resource{Type: ResourceMemoryMegabytes, Amount: 256}
	)

	testAppName := "test-resource-app"

	r.NoError(s.Client.CreateApp(testAppName), "failed to create app")

	r.NoError(s.Client.SetAppResourceReservation(testAppName, cpuReserved.Type, cpuReserved.Amount))
	r.NoError(s.Client.SetAppDefaultResourceLimit(testAppName, cpuLimit.Type, cpuLimit.Amount))
	r.NoError(s.Client.SetAppDefaultResourceLimit(testAppName, memLimit.Type, memLimit.Amount))

	report, err := s.Client.GetAppResourceReport(testAppName)
	r.NoError(err)
	fmt.Printf("report: %+v\n", report)
	r.Equal(cpuLimit, report.Defaults.Limits.CPU)
	r.Equal(cpuReserved, report.Defaults.Reservations.CPU)
	r.Equal(memLimit, report.Defaults.Limits.Memory)

	r.NoError(s.Client.ClearAppDefaultResourceLimit(testAppName, ResourceCPU))
	r.NoError(s.Client.ClearAppResourceReservation(testAppName, ResourceCPU))

	report2, err := s.Client.GetAppResourceReport(testAppName)
	r.NoError(err)

	r.Nil(report2.Defaults.Limits.CPU)
	r.Nil(report2.Defaults.Reservations.CPU)
	r.NotNil(report2.Defaults.Limits.Memory)
}

func (s *resourceManagerTestSuite) TestManageAppProcessResources() {
	r := s.Require()
	var err error

	var (
		cpuReserved = &Resource{Type: ResourceCPU, Amount: 1}
		cpuLimit    = &Resource{Type: ResourceCPU, Amount: 2}
		memLimit    = &Resource{Type: ResourceMemoryMegabytes, Amount: 256}
	)

	testAppName := "test-resource-app"
	processName := "web"

	r.NoError(s.Client.CreateApp(testAppName), "failed to create app")

	r.NoError(s.Client.SetAppProcessResourceReservation(testAppName, processName, cpuReserved.Type, cpuReserved.Amount))
	r.NoError(s.Client.SetAppProcessResourceLimit(testAppName, processName, cpuLimit.Type, cpuLimit.Amount))
	r.NoError(s.Client.SetAppProcessResourceLimit(testAppName, processName, memLimit.Type, memLimit.Amount))

	report, err := s.Client.GetAppResourceReport(testAppName)
	r.NoError(err)

	r.Contains(report.Processes, processName)
	processReport := report.Processes[processName]
	r.Equal(cpuLimit, processReport.Limits.CPU)
	r.Equal(cpuReserved, processReport.Reservations.CPU)
	r.Equal(memLimit, processReport.Limits.Memory)

	r.NoError(s.Client.ClearAppProcessResourceLimit(testAppName, processName, ResourceCPU))
	r.NoError(s.Client.ClearAppProcessResourceReservation(testAppName, processName, ResourceCPU))

	report2, err := s.Client.GetAppResourceReport(testAppName)
	r.NoError(err)

	r.Contains(report2.Processes, processName)
	processReport2 := report2.Processes[processName]
	r.Nil(processReport2.Limits.CPU)
	r.Nil(processReport2.Reservations.CPU)
	r.NotNil(processReport2.Limits.Memory)
}
