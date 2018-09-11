package autofill

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"library/common"
	"library/notice"
	"encoding/json"
	"time"
	"models"
	"strings"
	"library/components"
	"strconv"
)

type FalconInput struct {
	Endpoint  string
	Metric  string
	Status  string
	Time  string
	Tags  string
	Leftvalue  string
	Rightvalue  string
	Counter  string
}

//匹配自愈方案
func RecoverMatch(fv FalconInput) error {
	//捕获panic
	defer func() {
		if panErr := recover(); panErr != nil {
			beego.Info(panErr)
		}
	}()

	autoFillId := MathAutoFill(fv)
	if autoFillId == 0 {
		beego.Info(fv.Counter + "====> 无相匹配自愈方案")
		return nil
	}

	autoFill, _ := models.GetUwAutofillById(autoFillId)

	// 获取操作步骤

	ids, err := models.GetOperationsByFid(autoFill.Id)
	if err != nil {
		return  err
	}
	if len(ids) == 0 {
		beego.Info(fv.Counter + "====> 匹配自愈方案: " + autoFill.Title + " 无具体操作步骤")
		return nil
	}

	// 告警处理收敛，防止由于处理或等待时间过长造成的重复操作
	re := CheckRepeat(fv, autoFillId)
	if re == "running" {
		beego.Info(fv.Counter + "====> 匹配自愈方案: " + autoFill.Title + " 仍在执行中")
		return nil
	}

	// 创建task
	task := models.UwTask{}
	task.Source = fv.Endpoint
	task.Item = fv.Metric + autoFill.Operator + autoFill.Value
	task.Status = 1
	task.AutofillId = autoFillId
	task.Operator = "falcon"
	detail, _ := json.Marshal(fv)
	task.Detail = string(detail)
	taskId, err := models.AddUwTask(&task)
	if err != nil {
		return  err
	}

	start := time.Now()

	// 开启通知
	beginFlag := autoFill.BeginNotice
	failFlag := autoFill.FailNotice
	successFlag := autoFill.SuccNotice
	noticeUser := autoFill.NoticeUser
	noticeTeam := autoFill.NoticeTeam

	if beginFlag == 1 {
		RecoverNotice(noticeUser, noticeTeam, fv.Metric + ":" + fv.Leftvalue + "告警使用自愈模板: 《" + autoFill.Title + "》开始处理")
	}

	// 处理开始
	uwTask, _ := models.GetUwTaskById(int(taskId))
	s := components.BaseComponents{}
	s.SetAutofill(autoFill)
	s.SetTask(uwTask)
	s.SetOperator("falcon")
	s.SetDuration(0)
	s.SetFinal(0)

	for i, oid := range ids {
		oper, _ := models.GetUwOperationById(oid)
		if strings.Contains(oper.Hosts, "{CURRENT}") {
			oper.Hosts = strings.Replace(oper.Hosts, "{CURRENT}", fv.Endpoint, -1)
		}
		s.SetOperation(oper)

		if (i == (len(ids) -1)) {
			s.SetFinal(1)
		}

		// 执行远端自愈处理命令
		err := s.RunAutoFillRemoteCommand()
		if err != nil {
			uwTask.Status = 3
			uwTask.Duration = common.GetInt(time.Now().Sub(start).Seconds())
			uwTask.UpdateTime = time.Now()
			models.UpdateUwTaskById(uwTask)
			if failFlag == 1 {
				RecoverNotice(noticeUser, noticeTeam, fv.Metric + ":" + fv.Leftvalue + "告警使用自愈模板: 《" + autoFill.Title + "》在执行第" + common.GetString(i+1) + "步操作时出错。错误详情：<br>" + err.Error())
			}
			return  err
		}
	}

	// 处理完成
	uwTask.Status = 2
	uwTask.Duration = common.GetInt(time.Now().Sub(start).Seconds())
	uwTask.UpdateTime = time.Now()
	models.UpdateUwTaskById(uwTask)

	if successFlag == 1 {
		RecoverNotice(noticeUser, noticeTeam, fv.Metric + ":" + fv.Leftvalue + "告警使用自愈模板: 《" + autoFill.Title + "》处理完毕，耗时: " + strconv.Itoa(uwTask.Duration) + " 毫秒")
	}

	return nil
}

// 自动匹配自愈方案
func MathAutoFill(fv FalconInput) int {
	o := orm.NewOrm()
	var autoparms []orm.Params
	sql := fmt.Sprintf("SELECT * FROM `uw_autofill` WHERE isvalid = 1 AND metric= '%s'", fv.Metric)
	if fv.Tags != "" {
		sql = sql + fmt.Sprintf("and tag = '%s'", fv.Tags)
	}
	num, err := o.Raw(sql).Values(&autoparms)
	if err != nil || num == 0 {
		return 0
	} else {
		operator := ""
		if common.GetInt(fv.Leftvalue) > common.GetInt(autoparms[0]["value"]) {
			operator = ">"
		} else if common.GetInt(fv.Leftvalue) < common.GetInt(autoparms[0]["value"]) {
			operator = "<"
		} else {
			operator = "="
		}
		//根据operator匹配套餐方案
		sql = sql + fmt.Sprintf("and operator = '%s'", operator)
		num, err := o.Raw(sql).Values(&autoparms)
		if err != nil || num == 0{
			return 0
		}

		return common.GetInt(autoparms[0]["id"])
	}
}

// 检查告警收敛, 存在相同告警正在操作中的自愈任务延后处理
func CheckRepeat(fv FalconInput, fid int) string {
	o := orm.NewOrm()
	var autoparms []orm.Params
	sql := fmt.Sprintf("SELECT 1 FROM `uw_task` WHERE source = '%s' AND autofill_id = %d AND status = %d", fv.Endpoint, fid, 1)
	num, _ := o.Raw(sql).Values(&autoparms)
	if num > 0 {
		return "running"
	}

	return ""
}

//通知
func RecoverNotice(users, feteam, content string) {
	if users == "" && feteam == "" {
		return
	}

	//uwork微信通知
	wxusers := []string{}
	if users != "" {
		user := strings.Split(users, ",")
		for _, u := range user {
			wxusers = append(wxusers, u)
		}
	}

	if feteam != "" {
		//get tid
		o := orm.NewOrm()
		o.Using("uic")

		var team []orm.Params
		var group []orm.Params

		ss, err := o.Raw("SELECT id FROM team WHERE `name` = ?", feteam).Values(&team)
		if err == nil && ss > 0 {
			ss, err := o.Raw("SELECT * FROM rel_team_user WHERE tid = ?", common.GetString(team[0]["id"])).Values(&group)
			if err == nil && ss > 0 {
				for _, uid := range group {
					var username []orm.Params
					ss, err := o.Raw("SELECT name FROM user WHERE id = ?", common.GetInt(uid["uid"])).Values(&username)
					if err == nil && ss > 0 {
						wxusers = append(wxusers, common.GetString(username[0]["name"]))
					}
				}
			}
		}
	}

	for _, username := range wxusers {
		msg := map[string]string{}
		msg["title"] = "Uwatch-告警自愈通知"
		msg["username"] = username
		msg["depid"] = ""
		msg["content"] = content
		notice.Wechat_send(msg)
	}

}


