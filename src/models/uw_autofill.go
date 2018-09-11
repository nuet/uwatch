package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type UwAutofill struct {
	Id          int       `orm:"column(id);auto"`
	Title       string    `orm:"column(title);size(255);null"`
	Author      string    `orm:"column(author);size(20);null"`
	Metric      string    `orm:"column(metric);size(255);null"`
	Tag         string    `orm:"column(tag);size(255);null"`
	Operator    string    `orm:"column(operator);size(20);null"`
	Value       string    `orm:"column(value);size(255);null"`
	Timeout     uint      `orm:"column(timeout);null"`
	BeginNotice int8      `orm:"column(begin_notice);null"`
	SuccNotice  int8      `orm:"column(succ_notice);null"`
	FailNotice  int8      `orm:"column(fail_notice);null"`
	NoticeTeam  string    `orm:"column(notice_team);size(255);null"`
	NoticeUser  string    `orm:"column(notice_user);size(512);null"`
	InsertTime  time.Time `orm:"column(insert_time);type(datetime);null"`
	UpdateTime  time.Time `orm:"column(update_time);type(datetime);null"`
	Isvalid     int8      `orm:"column(isvalid);null"`
}

func (t *UwAutofill) TableName() string {
	return "uw_autofill"
}

func init() {
	orm.RegisterModel(new(UwAutofill))
}

// AddUwAutofill insert a new UwAutofill into database and returns
// last inserted Id on success.
func AddUwAutofill(m *UwAutofill) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUwAutofillById retrieves UwAutofill by Id. Returns error if
// Id doesn't exist
func GetUwAutofillById(id int) (v *UwAutofill, err error) {
	o := orm.NewOrm()
	v = &UwAutofill{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAssetBorrow retrieves all AssetBorrow matches certain condition. Returns empty list if
// no records exist
func ListAllAutoFill(query map[string]string, fields []string, sortby []string,
offset int64, limit int64) (num int64, err error, ml []map[string]interface{}) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwAutofill))
	cond := getAutofillCond(query)
	qs = qs.SetCond(cond)

	if len(sortby) != 0 {
		qs = qs.OrderBy(sortby...)
	}

	var l []UwAutofill
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

func getAutofillCond(condArr map[string]string) *orm.Condition {
	cond := orm.NewCondition()
	if condArr["Query"] != "" {
		cond1 := orm.NewCondition()
		cond1 = cond1.Or("title__icontains", condArr["Query"])
		cond = cond.AndCond(cond1)
	}
	return cond
}
//统计数量
func CountAutofill(query map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwAutofill))
	cond := getAutofillCond(query)
	num, _ := qs.SetCond(cond).Count()
	return num
}

// GetAllUwAutofill retrieves all UwAutofill matches certain condition. Returns empty list if
// no records exist
func GetAllUwAutofill(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwAutofill))
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

	var l []UwAutofill
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

// UpdateUwAutofill updates UwAutofill by Id and returns error if
// the record to be updated doesn't exist
func UpdateUwAutofillById(m *UwAutofill) (err error) {
	o := orm.NewOrm()
	v := UwAutofill{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUwAutofill deletes UwAutofill by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUwAutofill(id int) (err error) {
	o := orm.NewOrm()
	v := UwAutofill{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UwAutofill{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
