package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
	"strconv"
	"strings"
)

const (
	resourceReportCmd              = "resource:report"
	resourceReportAppCmd           = "resource:report %s"
	resourceLimitCmd               = "resource:limit %s --%s %s"
	resourceLimitProcessCmd        = "resource:limit %s --process-type %s --%s %s"
	resourceLimitClearCmd          = "resource:limit-clear %s"
	resourceLimitClearProcessCmd   = "resource:limit-clear %s --process-type %s"
	resourceReserveCmd             = "resource:reserve %s --%s %s"
	resourceReserveProcessCmd      = "resource:reserve %s --process-type %s --%s %s"
	resourceReserveClearCmd        = "resource:reserve-clear %s"
	resourceReserveClearProcessCmd = "resource:reserve-clear %s --process-type %s"
)

type ResourceSpec struct {
	Name   string
	Suffix string
}

type Resource struct {
	Type   ResourceSpec
	Amount int
}

type ResourceUnits struct {
	CPU            *Resource
	Memory         *Resource
	MemorySwap     *Resource
	Network        *Resource
	NetworkIngress *Resource
	NetworkEgress  *Resource
	NvidiaGPU      *Resource
}

type ResourceSettings struct {
	Limits       ResourceUnits
	Reservations ResourceUnits
}

type ResourceReport struct {
	Defaults  ResourceSettings
	Processes map[string]ResourceSettings
}

var (
	ResourceCPU                 = ResourceSpec{"cpu", ""}
	ResourceMemoryBytes         = ResourceSpec{"memory", "b"}
	ResourceMemoryKilobytes     = ResourceSpec{"memory", "k"}
	ResourceMemoryMegabytes     = ResourceSpec{"memory", "m"}
	ResourceMemoryGigabytes     = ResourceSpec{"memory", "g"}
	ResourceMemorySwapBytes     = ResourceSpec{"memory-swap", "b"}
	ResourceMemorySwapKilobytes = ResourceSpec{"memory-swap", "k"}
	ResourceMemorySwapMegabytes = ResourceSpec{"memory-swap", "m"}
	ResourceMemorySwapGigabytes = ResourceSpec{"memory-swap", "g"}
	ResourceNetwork             = ResourceSpec{"network", ""}
	ResourceNetworkIngress      = ResourceSpec{"network-ingress", ""}
	ResourceNetworkEgress       = ResourceSpec{"network-egress", ""}
	ResourceNvidiaGPU           = ResourceSpec{"nvidia-gpu", ""}
)

func (c *DefaultClient) SetAppDefaultResourceLimit(appName string, resource ResourceSpec, limit int) error {
	amt := fmt.Sprintf("%d%s", limit, resource.Suffix)
	cmd := fmt.Sprintf(resourceLimitCmd, appName, resource.Name, amt)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) ClearAppDefaultResourceLimit(appName string, resource ResourceSpec) error {
	cmd := fmt.Sprintf(resourceLimitCmd, appName, resource.Name, "clear")
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) ClearAppDefaultResourceLimits(appName string) error {
	cmd := fmt.Sprintf(resourceLimitClearCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) SetAppProcessResourceLimit(appName string, process string, resource ResourceSpec, limit int) error {
	amt := fmt.Sprintf("%d%s", limit, resource.Suffix)
	cmd := fmt.Sprintf(resourceLimitProcessCmd, appName, process, resource.Name, amt)
	fmt.Println(cmd)
	out, err := c.Exec(cmd)
	fmt.Println(out)
	return err
}

func (c *DefaultClient) ClearAppProcessResourceLimit(appName string, process string, resource ResourceSpec) error {
	cmd := fmt.Sprintf(resourceLimitProcessCmd, appName, process, resource.Name, "clear")
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) ClearAppProcessResourceLimits(appName string, process string) error {
	cmd := fmt.Sprintf(resourceLimitClearProcessCmd, appName, process)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) SetAppResourceReservation(appName string, resource ResourceSpec, reserve int) error {
	amt := fmt.Sprintf("%d%s", reserve, resource.Suffix)
	cmd := fmt.Sprintf(resourceReserveCmd, appName, resource.Name, amt)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) ClearAppResourceReservation(appName string, resource ResourceSpec) error {
	cmd := fmt.Sprintf(resourceReserveCmd, appName, resource.Name, "clear")
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) ClearAppResourceReservations(appName string) error {
	cmd := fmt.Sprintf(resourceReserveClearCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) SetAppProcessResourceReservation(appName string, process string, resource ResourceSpec, reserve int) error {
	amt := fmt.Sprintf("%d%s", reserve, resource.Suffix)
	cmd := fmt.Sprintf(resourceReserveProcessCmd, appName, process, resource.Name, amt)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) ClearAppProcessResourceReservation(appName string, process string, resource ResourceSpec) error {
	cmd := fmt.Sprintf(resourceReserveProcessCmd, appName, process, resource.Name, "clear")
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) ClearAppProcessResourceReservations(appName string, process string) error {
	cmd := fmt.Sprintf(resourceReserveClearProcessCmd, appName, process)
	_, err := c.Exec(cmd)
	return err
}

func updateResourceUnitSettings(units *ResourceUnits, resourceType string, rawVal string) error {
	resSpec := ResourceSpec{
		Name: resourceType,
	}

	if resourceType == "memory" || resourceType == "memory-swap" {
		resSpec.Suffix = rawVal[len(rawVal)-1:]
		rawVal = rawVal[:len(rawVal)-1]
	}

	amt, err := strconv.Atoi(rawVal)
	if err != nil {
		return fmt.Errorf("failed to convert resource amount (%s)", rawVal)
	}

	res := &Resource{
		Amount: amt,
		Type:   resSpec,
	}

	switch resourceType {
	case "cpu":
		units.CPU = res
	case "memory":
		units.Memory = res
	case "memory-swap":
		units.MemorySwap = res
	case "network":
		units.Network = res
	case "network-ingress":
		units.NetworkIngress = res
	case "network-egress":
		units.NetworkEgress = res
	case "nvidia-gpu":
		units.NvidiaGPU = res
	default:
		return fmt.Errorf("unknown resource type %s", resourceType)
	}
	return nil
}

func parseAppResourceReport(reportMap map[string]string) (*ResourceReport, error) {
	report := &ResourceReport{
		Defaults: ResourceSettings{
			Limits:       ResourceUnits{},
			Reservations: ResourceUnits{},
		},
		Processes: map[string]ResourceSettings{},
	}

	for k, v := range reportMap {
		kSplit := strings.Split(k, " ")
		proc := kSplit[0]
		mType := kSplit[1]
		resType := kSplit[2]

		var resSettings *ResourceSettings
		isProcess := proc != "_default_"

		if !isProcess {
			resSettings = &report.Defaults
		} else {
			settings, ok := report.Processes[proc]
			if !ok {
				settings = ResourceSettings{
					Limits:       ResourceUnits{},
					Reservations: ResourceUnits{},
				}
			}
			resSettings = &settings
		}

		fmt.Printf("k:%s; v:%s\n", k, v)
		fmt.Printf("before; %+v\n", resSettings)
		var err error
		if mType == "limit" {
			err = updateResourceUnitSettings(&resSettings.Limits, resType, v)
		} else if mType == "reserve" {
			err = updateResourceUnitSettings(&resSettings.Reservations, resType, v)
		} else {
			return nil, fmt.Errorf("unknown resource management '%s'", mType)
		}
		fmt.Printf("after; %+v\n", resSettings)

		if err != nil {
			return nil, err
		}

		if isProcess {
			report.Processes[proc] = *resSettings
		}
	}

	return report, nil
}

func (c *DefaultClient) GetAppResourceReport(appName string) (*ResourceReport, error) {
	cmd := fmt.Sprintf(resourceReportAppCmd, appName)
	output, err := c.Exec(cmd)

	if err != nil {
		return nil, err
	}

	reportMap, err := reports.ParseSingle(output)
	if err != nil {
		return nil, err
	}

	return parseAppResourceReport(reportMap)
}

type ResourcesReport map[string]*ResourceReport

func (c *DefaultClient) GetAllAppResourceReport() (ResourcesReport, error) {
	output, err := c.Exec(resourceReportCmd)

	if err != nil {
		return nil, err
	}

	reportMap, err := reports.ParseMultiple(output)
	if err != nil {
		return nil, err
	}

	appsReport := ResourcesReport{}
	for name, report := range reportMap {
		appReport, err := parseAppResourceReport(report)
		if err != nil {
			return nil, err
		}

		appsReport[name] = appReport
	}

	return appsReport, nil
}
