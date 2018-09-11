package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type UwTask struct {
	Id        int       `orm:"column(id);auto"`
	Source    string    `orm:"column(source);size(255);null" description:"触发源"`
	Duration  int       `orm:"column(duration);null" description:"操作时间"`
	Item      string    `orm:"column(item);size(100);null" description:"告警项目"`
	Status    int       `orm:"column(status);null" description:"状态"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null;auto_now_add" description:"产生时间"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null" description:"更新时间"`
	AutofillId    int   `orm:"column(autofill_id);null" `
	Operator  string    `orm:"column(operator);size(20);null" description:"操作人"`
	Detail      string `orm:"column(detail);type(text);null" description:"告警详细"`
}

func (t *UwTask) TableName() string {
	return "uw_task"
}

func init() {
	orm.RegisterModel(new(UwTask))
}

// AddUwTask insert a new UwTask into database and returns
// last inserted Id on success.
func AddUwTask(m *UwTask) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUwTaskById retrieves UwTask by Id. Returns error if
// Id doesn't exist
func GetUwTaskById(id int) (v *UwTask, err error) {
	o := orm.NewOrm()
	v = &UwTask{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUwTask retrieves all UwTask matches certain condition. Returns empty list if
// no records exist
func GetAllUwTask(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwTask))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []UwTask
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateUwTask updates UwTask by Id and returns error if
// the record to be updated doesn't exist
func UpdateUwTaskById(m *UwTask) (err error) {
	o := orm.NewOrm()
	v := UwTask{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUwTask deletes UwTask by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUwTask(id int) (err error) {
	o := orm.NewOrm()
	v := UwTask{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UwTask{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetAllAsset retrieves all Asset matches certain condition. Returns empty list if
// no records exist
func ListAllUwTask(query map[string]string, fields []string, sortby []string,
offset int64, limit int64) (num int64, err error, ml []map[string]interface{}) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwTask))
	cond := getUwTaskCond(query)
	qs = qs.SetCond(cond)

	if len(sortby) != 0 {
		qs = qs.OrderBy(sortby...)
	}

	var l []UwTask
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

func getUwTaskCond(condArr map[string]string) *orm.Condition {
	cond := orm.NewCondition()
	if condArr["status"] != "" {
		cond = cond.And("Status", condArr["status"])
	}
	if condArr["source"] != "" {
		cond = cond.And("Source__icontains", condArr["source"])
	}
	if condArr["start"] != "" {
		cond = cond.And("CreateTime__gte", condArr["start"])
	}
	if condArr["end"] != "" {
		cond = cond.And("CreateTime__lte", condArr["end"])
	}
	return cond
}

//统计数量
func CountUwTask(query map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwTask))
	cond := getUwTaskCond(query)
	num, _ := qs.SetCond(cond).Count()
	return num
}