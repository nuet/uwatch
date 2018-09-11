package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type UwNav struct {
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(name);size(255);null" description:"一级导航名称"`
	Url        string    `orm:"column(url);size(255);null" description:"导航url"`
	Createtime time.Time `orm:"column(createtime);type(timestamp);null" description:"创建时间"`
	Updatetime time.Time `orm:"column(updatetime);type(timestamp);null" description:"更新时间"`
	Status     int       `orm:"column(status);null" description:"是否开启：1是，0否"`
	Username   string    `orm:"column(username);size(255);null" description:"操作人"`
}

func (t *UwNav) TableName() string {
	return "uw_nav"
}

func init() {
	orm.RegisterModel(new(UwNav))
}

// AddUwNav insert a new UwNav into database and returns
// last inserted Id on success.
func AddUwNav(m *UwNav) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUwNavById retrieves UwNav by Id. Returns error if
// Id doesn't exist
func GetUwNavById(id int) (v *UwNav, err error) {
	o := orm.NewOrm()
	v = &UwNav{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func getNavCond(condArr map[string]string) *orm.Condition {
	cond := orm.NewCondition()
	if condArr["Query"] != "" {
		cond1 := orm.NewCondition()
		cond1 = cond1.And("status__icontains", 1)
		cond1 = cond1.And("name__icontains", condArr["Query"])
		cond1 = cond1.Or("url__icontains", condArr["Query"])
		cond = cond.AndCond(cond1)
	} else {
		cond1 := orm.NewCondition()
		cond1 = cond1.And("status__icontains", 1)
		cond = cond.AndCond(cond1)
	}
	return cond
}
//统计数量
func CountNav(query map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwNav))
	cond := getNavCond(query)
	num, _ := qs.SetCond(cond).Count()
	return num
}

// GetAllUwNav retrieves all UwNav matches certain condition. Returns empty list if
// no records exist
func GetAllUwNav(query map[string]string, fields []string, sortby []string,
	offset int64, limit int64) (num int64, err error, ml []map[string]interface{}) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwNav))
	cond := getNavCond(query)
	beego.Info("cond===+++++===++++++>", cond)
	qs = qs.SetCond(cond)

	if len(sortby) != 0 {
		qs = qs.OrderBy(sortby...)
	}

	var l []UwNav
	if num, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		for _, v := range l {
			m := make(map[string]interface{}, 0)
			t := reflect.TypeOf(v)
			v := reflect.ValueOf(v)
			for k := 0; k < t.NumField(); k++ {
				m[t.Field(k).Name] = v.Field(k).Interface()
			}
			ml = append(ml, m)
		}
		return num, nil, ml
	}

	return num,err,ml
}

// GetAllAssetBorrow retrieves all AssetBorrow matches certain condition. Returns empty list if
// no records exist
func GetAllUwNavData(status int64, query map[string]string, fields []string, sortby []string,
offset int64, limit int64) (num int64, err error, ml []map[string]interface{}) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwNav))
	cond := getGraphCondList(query, status)
	qs = qs.SetCond(cond)

	if len(sortby) != 0 {
		qs = qs.OrderBy(sortby...)
	}

	var l []UwNav
	if num, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		for _, v := range l {
			m := make(map[string]interface{}, 0)
			t := reflect.TypeOf(v)
			v := reflect.ValueOf(v)
			for k := 0; k < t.NumField(); k++ {
				m[t.Field(k).Name] = v.Field(k).Interface()
			}
			ml = append(ml, m)
		}
		return num, nil, ml
	}

	return num,err,ml
}

// UpdateUwNav updates UwNav by Id and returns error if
// the record to be updated doesn't exist
func UpdateUwNavById(m *UwNav) (err error) {
	o := orm.NewOrm()
	v := UwNav{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUwNav deletes UwNav by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUwNav(id int) (err error) {
	o := orm.NewOrm()
	v := UwNav{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UwNav{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
