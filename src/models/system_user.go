package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SystemUser struct {
	Id        int    `orm:"column(SuId);auto"`
	SuName    string `orm:"column(SuName);size(50)"`
	SuRealName    string `orm:"column(SuRealName);size(50);null"`
	SuRoles   string `orm:"column(SuRoles);size(255);null"`
	SuIsAdmin string `orm:"column(SuIsAdmin);size(10)"`
	SuComPers string `orm:"column(SuComPers);size(255);null"`
	SuStatus  string `orm:"column(SuStatus);size(10)"`
}

func (t *SystemUser) TableName() string {
	return "system_user"
}

func init() {
	orm.RegisterModel(new(SystemUser))
}

// AddSystemUser insert a new SystemUser into database and returns
// last inserted Id on success.
func AddSystemUser(m *SystemUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSystemUserById retrieves SystemUser by Id. Returns error if
// Id doesn't exist
func GetSystemUserById(id int) (v *SystemUser, err error) {
	o := orm.NewOrm()
	v = &SystemUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSystemUser retrieves all SystemUser matches certain condition. Returns empty list if
// no records exist
func GetAllSystemUser(query map[string]string, fields []string, sortby []string, order []string,
offset int64, limit int64) (ml []map[string]interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SystemUser))
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

	var l []SystemUser
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		for _, v := range l {
			m := make(map[string]interface{}, 0)
			t := reflect.TypeOf(v)
			v := reflect.ValueOf(v)
			for k := 0; k < t.NumField(); k++ {
				m[t.Field(k).Name] = v.Field(k).Interface()
			}
			ml = append(ml, m)
		}
		return ml, nil
	}
	return nil, err
}

// GetAllSystemUserStandard retrieves all SystemUser matches certain condition. Returns empty list if
// no records exist
func GetAllSystemUserStandard(query map[string]string,  sortby []string, order []string,
offset int64, limit int64) (ml []SystemUser, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SystemUser))
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

	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&ml); err == nil {
		return ml, nil
	}
	return nil, err
}


// UpdateSystemUser updates SystemUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateSystemUserById(m *SystemUser) (err error) {
	o := orm.NewOrm()
	v := SystemUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSystemUser deletes SystemUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSystemUser(id int) (err error) {
	o := orm.NewOrm()
	v := SystemUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SystemUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

// GetAssetByAsCode retrieves Asset by AsCode. Returns error if
// Id doesn't exist
func GetSystemUserByName(name string) (l []SystemUser, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SystemUser))
	qs = qs.Filter("SuName", name)
	if _, err = qs.All(&l); err == nil {
		return l, nil
	} else {
		return l, err
	}
}
