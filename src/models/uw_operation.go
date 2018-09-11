package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"library/common"
)

type UwOperation struct {
	Id         int       `orm:"column(id);auto"`
	Fid        int       `orm:"column(fid)"`
	Title      string    `orm:"column(title);size(255);null"`
	Step       int       `orm:"column(step);null"`
	User       string    `orm:"column(user);size(100);null"`
	Hosts      string    `orm:"column(hosts);null"`
	Command    string    `orm:"column(command);null"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null"`
}

func (t *UwOperation) TableName() string {
	return "uw_operation"
}

func init() {
	orm.RegisterModel(new(UwOperation))
}

// AddUwOperation insert a new UwOperation into database and returns
// last inserted Id on success.
func AddUwOperation(m *UwOperation) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUwOperationById retrieves UwOperation by Id. Returns error if
// Id doesn't exist
func GetUwOperationById(id int) (v *UwOperation, err error) {
	o := orm.NewOrm()
	v = &UwOperation{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUwOperation retrieves all UwOperation matches certain condition. Returns empty list if
// no records exist
func GetAllUwOperation(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwOperation))
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

	var l []UwOperation
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

// UpdateUwOperation updates UwOperation by Id and returns error if
// the record to be updated doesn't exist
func UpdateUwOperationById(m *UwOperation) (err error) {
	o := orm.NewOrm()
	v := UwOperation{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUwOperation deletes UwOperation by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUwOperation(id int) (err error) {
	o := orm.NewOrm()
	v := UwOperation{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UwOperation{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// 根据自愈方案ID删除全部自愈执行操作
func DeleteUwOperationByFid(fid int) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("DELETE FROM `uw_operation` WHERE `fid` = ?", fid).Exec()
	if err != nil {
		return err
	}
	return nil
}

func GetOperationsByFid (fid int)  (ids []int, err error) {
	o := orm.NewOrm()
	var l []orm.Params
	num, err := o.Raw("SELECT id FROM `uw_operation` WHERE fid = ? ORDER BY `step` ASC ", fid).Values(&l)
	if num >0 {
		for _, oper := range l {
			ids = append(ids, common.GetInt(oper["id"]))
		}
	}
	return
}
