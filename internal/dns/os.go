package dns

import (
	"net"
	"strings"

	"github.com/qdm12/private-internet-access-docker/internal/constants"
)

func (c *configurator) SetNameserver(IP net.IP) error {
	c.logger.Info("%s: setting local nameserver to %s", logPrefix, IP.String())
	data, err := c.fileManager.ReadFile(string(constants.ResolvConf))
	if err != nil {
		return err
	}
	s := strings.TrimSuffix(string(data), "\n")
	lines := strings.Split(s, "\n")
	if len(lines) == 1 && lines[0] == "" {
		lines = nil
	}
	found := false
	for i := range lines {
		if strings.HasPrefix(lines[i], "nameserver ") {
			lines[i] = "nameserver " + IP.String()
			found = true
		}
	}
	if !found {
		lines = append(lines, "nameserver "+IP.String())
	}
	data = []byte(strings.Join(lines, "\n"))
	return c.fileManager.WriteToFile(string(constants.ResolvConf), data)
}