package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type UwVoicerecord struct {
	Id             int       `orm:"column(id);auto"`
	ActionId       string    `orm:"column(action_id);size(255);null"`
	CallId         string    `orm:"column(call_id);size(255);null"`
	VoiceMessage   string    `orm:"column(voice_message);size(255)"`
	Mobile         string    `orm:"column(mobile);size(255)"`
	AlertUser      string    `orm:"column(alert_user);size(255)"`
	AlertType      string    `orm:"column(alert_type);size(255)"`
	Result         string    `orm:"column(result);size(255)"`
	FailureReason  string    `orm:"column(failure_reason);size(255)"`
	Count          int       `orm:"column(count)"`
	CallbackStatus int       `orm:"column(callback_status)"`
	CreatedTime    time.Time `orm:"column(created_time);type(datetime);null;auto_now_add"`
}

func (t *UwVoicerecord) TableName() string {
	return "uw_voicerecord"
}

func init() {
	orm.RegisterModel(new(UwVoicerecord))
}

// AddVoiceRecord insert a new VoiceRecord into database and returns
// last inserted Id on success.
func AddVoiceRecord(m *UwVoicerecord) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetVoiceRecordById retrieves VoiceRecord by Id. Returns error if
// Id doesn't exist
func GetVoiceRecordById(id int) (v *UwVoicerecord, err error) {
	o := orm.NewOrm()
	v = &UwVoicerecord{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetVoiceRecordByCallId(id string) (v *UwVoicerecord, err error) {
	o := orm.NewOrm()
	v = &UwVoicerecord{CallId: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllVoiceRecord retrieves all VoiceRecord matches certain condition. Returns empty list if
// no records exist
func GetAllVoiceRecord(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UwVoicerecord))
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

	var l []UwVoicerecord
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

// UpdateVoiceRecord updates VoiceRecord by Id and returns error if
// the record to be updated doesn't exist
func UpdateVoiceRecordById(m *UwVoicerecord) (err error) {
	o := orm.NewOrm()
	v := UwVoicerecord{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteVoiceRecord deletes VoiceRecord by Id and returns error if
// the record to be deleted doesn't exist
func DeleteVoiceRecord(id int) (err error) {
	o := orm.NewOrm()
	v := UwVoicerecord{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UwVoicerecord{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
