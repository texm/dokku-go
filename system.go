package dokku

import "fmt"

func (c *DefaultClient) GetEventLogs() (string, error) {
	output, err := c.exec("events")
	fmt.Println("something")
	return output, err
}

func (c *DefaultClient) SetEventLoggingEnabled(enabled bool) error {
	status := "on"
	if !enabled {
		status = "off"
	}
	_, err := c.exec("events:" + status)
	return err
}

func (c *DefaultClient) ListLoggedEvents() ([]string, error) {
	var events []string

	return events, nil
}
