package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Insurance struct {
	Id           int64
	ProvinceCode string
	CityCode     string
	Pension      int64 //养老比例
	Medical      float64 //医疗比例
	Unemployment float64 //失业比例
	Provident    int64 //公积金比例
	CityDesc     string
}

func (i *Insurance) TableName() string {
	return "insurance_info"
}

func init() {
	orm.RegisterModel(new(Insurance))
}

func GetInsuranceByCode(provinceCode string, cityCode string) (i []Insurance, err error) {
	var o = orm.NewOrm()
	num, err := o.Raw("SELECT  * "+
		"FROM insurance_info WHERE province_code =? and city_code=?  ",
		provinceCode, cityCode).QueryRows(&i)
	if err == nil && num > 0 {

	}
	return i, nil
}

// AddInsurance insert a new Insurance into database and returns
// last inserted Id on success.
func AddInsurance(m *Insurance) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetInsuranceById retrieves Insurance by Id. Returns error if
// Id doesn't exist
func GetInsuranceById(id int64) (v *Insurance, err error) {
	o := orm.NewOrm()
	v = &Insurance{Id: id}
	if err = o.QueryTable(new(Insurance)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllInsurance retrieves all Insurance matches certain condition. Returns empty list if
// no records exist
func GetAllInsurance(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Insurance))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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

	var l []Insurance
	qs = qs.OrderBy(sortFields...).RelatedSel()
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

// UpdateInsurance updates Insurance by Id and returns error if
// the record to be updated doesn't exist
func UpdateInsuranceById(m *Insurance) (err error) {
	o := orm.NewOrm()
	v := Insurance{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteInsurance deletes Insurance by Id and returns error if
// the record to be deleted doesn't exist
func DeleteInsurance(id int64) (err error) {
	o := orm.NewOrm()
	v := Insurance{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Insurance{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
