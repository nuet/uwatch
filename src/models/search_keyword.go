package models

import (
	"fmt"
	"reflect"

	"github.com/astaxie/beego/orm"
)

type SearchKeyword struct {
	Id      int    `orm:"column(id);auto"`
	Keyword string `orm:"column(keyword);size(255);null" description:"关键词"`
	Url     string `orm:"column(url);size(255);null" description:"grafana1地址"`
}

func (t *SearchKeyword) TableName() string {
	return "search_keyword"
}

func init() {
	orm.RegisterModel(new(SearchKeyword))
}

// AddSearchKeyword insert a new SearchKeyword into database and returns
// last inserted Id on success.
func AddSearchKeyword(m *SearchKeyword) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSearchKeywordById retrieves SearchKeyword by Id. Returns error if
// Id doesn't exist
func GetSearchKeywordById(id int) (v *SearchKeyword, err error) {
	o := orm.NewOrm()
	v = &SearchKeyword{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllSearchKeyword retrieves all SearchKeyword matches certain condition. Returns empty list if
// no records exist
func GetAllSearchKeyword(query map[string]string, fields []string, sortby []string,
offset int64, limit int64) (num int64, err error, ml []map[string]interface{}) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SearchKeyword))
	cond := getSearchKeywordCond(query)
	qs = qs.SetCond(cond)

	if len(sortby) != 0 {
		qs = qs.OrderBy(sortby...)
	}

	var l []SearchKeyword
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

func getSearchKeywordCond(condArr map[string]string) *orm.Condition {
	cond := orm.NewCondition()
	if condArr["Query"] != "" {
		cond1 := orm.NewCondition()
		cond1 = cond1.Or("keyword__icontains", condArr["Query"])
		cond = cond.AndCond(cond1)
	}
	return cond
}

// UpdateSearchKeyword updates SearchKeyword by Id and returns error if
// the record to be updated doesn't exist
func UpdateSearchKeywordById(m *SearchKeyword) (err error) {
	o := orm.NewOrm()
	v := SearchKeyword{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSearchKeyword deletes SearchKeyword by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSearchKeyword(id int) (err error) {
	o := orm.NewOrm()
	v := SearchKeyword{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SearchKeyword{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
