package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type UwGraph struct {
	Id         int       `orm:"column(id);auto"`
	Endpoint   string    `orm:"column(endpoint);size(64);null" description:"主机地址"`
	Counter    string    `orm:"column(counter);size(64);null" description:"Counter"`
	Avg        string    `orm:"column(avg);size(64);null" description:"平均值"`
	Start      string    `orm:"column(start);size(64);null" description:"falcon_tag"`
	End        string    `orm:"column(end);size(64);null" description:"运算符"`
	InsertTime time.Time `orm:"column(insert_time);type(datetime);null;auto_now_add" description:"添加时间"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null;auto_now" description:"修改时间"`
	Status     int      `orm:"column(status);null" description:"开启(0:否，1:是)"`
	Title      string   `orm:"column(title);size(255);null" description:"名称"`
	Min				 string   `orm:"column(min);size(255);null" description:"最小值"`
	Max        string   `orm:"column(max);size(255);null" description:"最大值"`
	Counter_cn string   `orm:"column(counter_cn);size(255);null" description:"counter中文名称"`
	Axislabel string   `orm:"column(axislabel);size(255);null" description:"y轴单位"`
	Dividend string   `orm:"column(dividend);size(255);" description:"被除数"`
}

func (t *UwGraph) TableName() string {
	return "uw_graph"
}

func init() {
	orm.RegisterModel(new(UwGraph))
}

// AddUwGraph insert a new UwGraph into database and returns
// last inserted Id on success.
func AddUwGraph(m *UwGraph) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUwGraphById retrieves UwGraph by Id. Returns error if
// Id doesn't exist
func GetUwGraphById(id int) (v *UwGraph, err error) {
	o := orm.NewOrm()
	v = &UwGraph{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAssetBorrow retrieves all AssetBorrow matches certain condition. Returns empty list if
// no records exist
func ListAllGraph(query map[string]string, fields []string, sortby []string,
offset int64, limit int64) (num int64, err error, ml []map[string]interface{}) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwGraph))
	cond := getGraphCond(query)
	qs = qs.SetCond(cond)

	if len(sortby) != 0 {
		qs = qs.OrderBy(sortby...)
	}

	var l []UwGraph
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
func ListAllGraphData(status int64, query map[string]string, fields []string, sortby []string,
offset int64, limit int64) (num int64, err error, ml []map[string]interface{}) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwGraph))
	cond := getGraphCondData(query, status)
	qs = qs.SetCond(cond)

	if len(sortby) != 0 {
		qs = qs.OrderBy(sortby...)
	}

	var l []UwGraph
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

func getGraphCondData(condArr map[string]string, status int64) *orm.Condition {
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

func getGraphCond(condArr map[string]string) *orm.Condition {
	cond := orm.NewCondition()
	if condArr["Query"] != "" {
		cond1 := orm.NewCondition()
		cond1 = cond1.Or("endpoint__icontains", condArr["Query"])
		cond1 = cond1.Or("title__icontains", condArr["Query"])
		cond = cond.AndCond(cond1)
	}
	return cond
}
//统计数量
func CountGraph(query map[string]string) int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwGraph))
	cond := getGraphCond(query)
	num, _ := qs.SetCond(cond).Count()
	return num
}

// GetAllUwGraph retrieves all UwGraph matches certain condition. Returns empty list if
// no records exist
func GetAllUwGraph(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwGraph))
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

	var l []UwGraph
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

// GetAllSearchKeyword retrieves all SearchKeyword matches certain condition. Returns empty list if
// no records exist
func GetAllSearchGraphData(query map[string]string, fields []string, sortby []string,
offset int64, limit int64) (num int64, err error, ml []map[string]interface{}) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwGraph))
	cond := getSearchGraphCond(query)
	qs = qs.SetCond(cond)

	if len(sortby) != 0 {
		qs = qs.OrderBy(sortby...)
	}

	var l []UwGraph
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

func getSearchGraphCond(condArr map[string]string) *orm.Condition {
	cond := orm.NewCondition()
	if condArr["Query"] != "" {
		cond1 := orm.NewCondition()
		cond1 = cond1.Or("title__icontains", condArr["Query"])
		cond = cond.AndCond(cond1)
	}
	return cond
}

// UpdateUwGraph updates UwGraph by Id and returns error if
// the record to be updated doesn't exist
func UpdateUwGraphById(m *UwGraph) (err error) {
	o := orm.NewOrm()
	v := UwGraph{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUwGraph deletes UwGraph by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUwGraph(id int) (err error) {
	o := orm.NewOrm()
	v := UwGraph{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UwGraph{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
