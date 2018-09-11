package models

import (
	"fmt"
	"reflect"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type UwCounter struct {
	Id            int       `orm:"column(id);auto"`
	Counter    string    `orm:"column(counter);size(255);null" description:"counter"`
	Counter_cn    string    `orm:"column(counter_cn);size(255);null" description:"counter中文"`
	Createtime        time.Time `orm:"column(createtime);type(timestamp);null" description:"创建时间"`
	Updatetime       time.Time `orm:"column(updatetime);type(timestamp);null" description:"更新时间"`
	Username      string    `orm:"column(username);size(255);null" description:"操作人"`
	Status      int      `orm:"column(status);null" description:"开启(0:否，1:是)"`
	Avg        string    `orm:"column(avg);size(64);null" description:"平均线"`
	Axislabel string   `orm:"column(axislabel);size(255);null" description:"纵轴坐标单位"`
	Dividend string   `orm:"column(dividend);size(255);" description:"单位换算被除数"`
}

func (t *UwCounter) TableName() string {
	return "uw_counter"
}

func init() {
	orm.RegisterModel(new(UwCounter))
}

// AddUwCounter insert a new UwCounter into database and returns
// last inserted Id on success.
func AddUwCounter(m *UwCounter) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUwCounterById retrieves UwCounter by Id. Returns error if
// Id doesn't exist
func GetUwCounterById(id int) (v *UwCounter, err error) {
	o := orm.NewOrm()
	v = &UwCounter{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func getCounterCond(condArr map[string]string) *orm.Condition {
	cond := orm.NewCondition()
	if condArr["Query"] != "" {
		cond1 := orm.NewCondition()
		cond1 = cond1.And("status__icontains", 1)
		cond1 = cond1.And("counter__icontains", condArr["Query"])
		cond1 = cond1.Or("counter_cn__icontains", condArr["Query"])
		cond = cond.AndCond(cond1)
	} else {
		cond1 := orm.NewCondition()
		cond1 = cond1.And("status__icontains", 1)
		cond = cond.AndCond(cond1)
	}
	return cond
}
//统计数量
func CountCounter(query map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwCounter))
	cond := getCounterCond(query)
	num, _ := qs.SetCond(cond).Count()
	return num
}
// GetAllUwCounter retrieves all UwCounter matches certain condition. Returns empty list if
// no records exist
func GetAllUwCounter(query map[string]string, fields []string, sortby []string,
offset int64, limit int64) (num int64, err error, ml []map[string]interface{}) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwCounter))
	cond := getCounterCond(query)
	beego.Info("cond===+++++===++++++>", cond)
	qs = qs.SetCond(cond)

	if len(sortby) != 0 {
		qs = qs.OrderBy(sortby...)
	}

	var l []UwCounter
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
func GetAllUwCounterData(status int64, query map[string]string, fields []string, sortby []string,
offset int64, limit int64) (num int64, err error, ml []map[string]interface{}) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwCounter))
	cond := getGraphCondList(query, status)
	qs = qs.SetCond(cond)

	if len(sortby) != 0 {
		qs = qs.OrderBy(sortby...)
	}

	var l []UwCounter
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

func getGraphCondList(condArr map[string]string, status int64) *orm.Condition {
	cond := orm.NewCondition()
	cond1 := orm.NewCondition()
	if condArr["Query"] != "" {
		cond1 = cond.And("status__icontains", status)
		cond1 = cond1.And("endpoint__icontains", condArr["Query"])
		cond1 = cond1.Or("title__icontains", condArr["Query"])
		cond = cond.AndCond(cond1)
	}
	return cond
}


// UpdateUwCounter updates UwCounter by Id and returns error if
// the record to be updated doesn't exist
func UpdateUwCounterById(m *UwCounter) (err error) {
	o := orm.NewOrm()
	v := UwCounter{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUwCounter deletes UwCounter by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUwCounter(id int) (err error) {
	o := orm.NewOrm()
	v := UwCounter{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UwCounter{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
