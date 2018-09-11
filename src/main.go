package main

import (
	_ "routers"
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"os/signal"
	"syscall"
	"tasks"
	"github.com/astaxie/beego/toolbox"
)

func init() {
	//连接MySQL
	uicDbUser := beego.AppConfig.String("uic_mysqluser")
	uicDbPass := beego.AppConfig.String("uic_mysqlpass")
	uicDbHost := beego.AppConfig.String("uic_mysqlhost")
	uicDbPort := beego.AppConfig.String("uic_mysqlport")
	uicDbName := beego.AppConfig.String("uic_mysqldb")
	maxIdleConn, _ := beego.AppConfig.Int("db_max_idle_conn")
	maxOpenConn, _ := beego.AppConfig.Int("db_max_open_conn")
	uicDbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", uicDbUser, uicDbPass, uicDbHost, uicDbPort, uicDbName) + "&loc=Asia%2FShanghai"
	//utils.Display("dbLink", dbLink)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("uic", "mysql", uicDbLink, maxIdleConn, maxOpenConn)

	dutyDbUser := beego.AppConfig.String("duty_mysqluser")
	dutyDbPass := beego.AppConfig.String("duty_mysqlpass")
	dutyDbHost := beego.AppConfig.String("duty_mysqlhost")
	dutyDbPort := beego.AppConfig.String("duty_mysqlport")
	dutyDbName := beego.AppConfig.String("duty_mysqldb")
	dutyDbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dutyDbUser, dutyDbPass, dutyDbHost, dutyDbPort, dutyDbName) + "&loc=Asia%2FShanghai"
	//utils.Display("dbLink", dbLink)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("duty", "mysql", dutyDbLink, maxIdleConn, maxOpenConn)


	dbUser := beego.AppConfig.String("mysqluser")
	dbPass := beego.AppConfig.String("mysqlpass")
	dbHost := beego.AppConfig.String("mysqlhost")
	dbPort := beego.AppConfig.String("mysqlport")
	dbName := beego.AppConfig.String("mysqldb")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName) + "&loc=Asia%2FShanghai"
	//utils.Display("dbLink", dbLink)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbLink, maxIdleConn, maxOpenConn)

	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
	//设置日志
	fn := "logs/run.log"
	if _, err := os.Stat(fn); err != nil {
		if os.IsNotExist(err) {
			os.Create(fn)
		}
	}
	beego.SetLogger("file", `{"filename":"`+fn+`"}`)
	if beego.BConfig.RunMode == "prod" {
		beego.SetLevel(beego.LevelInformational)
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			beego.Error("Panic error:", err)
		}
	}()

	graceful, _ := beego.AppConfig.Bool("Graceful")
	if graceful {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		go handleSignals(sigs)
	}

	// 定时任务
	// uwatch测试环境跑办公网同步任务
	if beego.BConfig.RunMode == "sit" || beego.BConfig.RunMode == "prod"{
		// 蓝鲸CMDB、Falcon+ 主机组/主机信息同步
		hostGroup_sync := toolbox.NewTask("hostGroup_sync", "0 5 * * * *", func() error {
			err := tasks.HostGroupSync()
			if err != nil {
				beego.Error("定时任务: hHostGroupSync 出错: ", err.Error())
				return err
			}
			return nil
		})
		toolbox.AddTask("hostGroup_sync", hostGroup_sync)
		//go hostGroup_sync.Run()
		// 定时任务
		// uwatch同步ldap到falcon
		userGroup_sync := toolbox.NewTask("userGroup_sync", "1 0 * * * *", func() error {
			err := tasks.UserGroupSync()
			if err != nil {
				beego.Error("定时任务: Ldap_sync 出错: ", err.Error())
				return err
			}
			return nil
		})
		toolbox.AddTask("ldapAllData_sync", userGroup_sync)
		//go userGroup_sync.Run()
	}


	toolbox.StartTask()
	defer toolbox.StopTask()

	beego.Run()
}

func handleSignals(c chan os.Signal) {
	switch <-c {
	case syscall.SIGINT, syscall.SIGTERM:

		beego.Info("Shutdown quickly, bye...")
	case syscall.SIGQUIT:
		beego.Info("Shutdown gracefully, bye...")
	// do graceful shutdown
	}
	os.Exit(0)
}