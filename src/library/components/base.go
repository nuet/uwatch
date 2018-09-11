package components

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/gaoyue1989/sshexec"
	"library/common"
	"library/ssh"
	"models"
	"regexp"
	"time"
)

const SSHTIMEOUT = 3600
const SSHWorker = 10

const SSHREMOTETIMEOUT = 600

type BaseComponents struct {
	final     int
	duration  int
	operator  string
	autofill  *models.UwAutofill
	task      *models.UwTask
	operation *models.UwOperation
}

func (c *BaseComponents) SetAutofill(autofill *models.UwAutofill) {
	c.autofill = autofill

}

func (c *BaseComponents) SetTask(task *models.UwTask) {
	c.task = task
}

func (c *BaseComponents) SetOperation(operation *models.UwOperation) {
	c.operation = operation
}

func (c *BaseComponents) SetOperator(operator string) {
	c.operator = operator
}

func (c *BaseComponents) SetDuration(duration int) {
	c.duration = duration
}

func (c *BaseComponents) SetFinal(final int) {
	c.final = final
}

/**
* 执行本地宿主机命令
 */
func (c *BaseComponents) runLocalCommand(command string) (sshexec.ExecResult, error) {
	id := c.SaveRecord(command)
	s, err := ssh.CommandLocal(command, SSHTIMEOUT)
	//获取执行时间
	duration := common.GetInt(s.EndTime.Sub(s.StartTime).Seconds())
	createdAt := int(s.StartTime.Unix())
	status := 1
	if s.Error != nil {
		status = 0
	}
	c.SaveRecordRes(id, duration, createdAt, status, s)
	return s, err

}

/**
* 执行远端目标机命令
 */
func (c *BaseComponents) runRemoteCommand(command string, hosts []string) ([]sshexec.ExecResult, error) {
	if len(hosts) == 0 {
		hosts = c.GetHosts()
	}
	id := c.SaveRecord(command)
	start := time.Now()
	createdAt := int(start.Unix())
	sshExecAgent := sshexec.SSHExecAgent{}
	sshExecAgent.Worker = SSHWorker
	sshExecAgent.TimeOut = time.Duration(SSHREMOTETIMEOUT) * time.Second
	s, err := sshExecAgent.SshHostByKey(hosts, c.operation.User, command)
	//获取执行时间
	duration := common.GetInt(time.Now().Sub(start).Seconds())

	status := 1
	if err != nil {
		status = 0
	}
	c.SaveRecordRes(id, duration, createdAt, status, s)
	return s, err

}

/**
 * 获取host
 */
func (c *BaseComponents) GetHosts() []string {
	hostsStr := c.operation.Hosts
	reg := regexp.MustCompile(`(\d+)\.(\d+)\.(\d+)\.(\d+)`)
	hosts := reg.FindAll([]byte(hostsStr), -1)
	res := []string{}
	for _, host := range hosts {
		if !common.InList(string(host), res) {
			res = append(res, string(host))
		}

	}
	return res
}


func (c *BaseComponents) SaveRecord(command string) int {
	re := models.UwRecord{}
	re.Command = command
	re.Operator = c.operator
	re.IsFinal = c.final
	if c.task == nil || c.task.Id == 0 {
		re.TaskId = -99
	}
	if c.operation == nil || c.operation.Id == 0 {
		re.OptId = 0
	}
	re.TaskId = c.task.Id
	re.OptId = c.operation.Id
	re.Status = 1
	id, err := models.AddUwRecord(&re)
	if err != nil {
		beego.Error(err)
	}
	return int(id)
}
func (c *BaseComponents) SaveRecordRes(id int, duration int, createdAt int, status int, value interface{}) {
	beego.Info("test bug")
	beego.Info(value)
	if duration < 0 {
		duration = 0
	}
	re, err := models.GetUwRecordById(id)
	if err != nil {
		beego.Error(err)
		return
	}
	re.Duration = duration
	sResult, _ := json.Marshal(value)
	re.CreatedAt = createdAt
	re.Memo = string(sResult)
	re.Status = int16(status)
	err = models.UpdateUwRecordById(re)
	if err != nil {
		beego.Error(err)
	}
}
