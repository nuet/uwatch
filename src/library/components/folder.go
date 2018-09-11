package components

import ()
import (
)

/**
 * 测试ssh连接
 *
 */
func (c *BaseComponents) TestSsh() error {
	cmd := "id"
	_, err := c.runRemoteCommand(cmd, c.GetHosts())
	return err
}

/**
 * 自愈执行远端命令
 *
 */
func (c *BaseComponents) RunAutoFillRemoteCommand() error {
	_, err := c.runRemoteCommand(c.operation.Command, c.GetHosts())
	return err
}


