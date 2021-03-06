package admin

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Cron struct {
	Created time.Time `orm:"column(created);auto_now_add;type(datetime);null" description:"创建时间"`
	Cron    string    `orm:"column(cron);size(255);null" description:"定时器"`
	Id      int       `orm:"column(id);auto" description:"主键"`
	Status  int       `orm:"column(status);null" description:"状态1正常，2删除"`
	Title   string    `orm:"column(title);size(255);null" description:"任务名称"`
	Updated time.Time `orm:"column(updated);auto_now;type(datetime);null" description:"更新时间"`
	Url     string    `orm:"column(url);size(255);null" description:"链接URL"`
}

func (t *Cron) TableName() string {
	return "cron"
}

func init() {
	orm.RegisterModel(new(Cron))
}

// AddCron insert a new Cron into database and returns
// last inserted Id on success.
func AddCron(m *Cron) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCronById retrieves Cron by Id. Returns error if
// Id doesn't exist
func GetCronById(id int) (v *Cron, err error) {
	o := orm.NewOrm()
	v = &Cron{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCron retrieves all Cron matches certain condition. Returns empty list if
// no records exist
func GetAllCron(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Cron))
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

	var l []Cron
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

// UpdateCron updates Cron by Id and returns error if
// the record to be updated doesn't exist
func UpdateCronById(m *Cron) (err error) {
	o := orm.NewOrm()
	v := Cron{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCron deletes Cron by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCron(id int) (err error) {
	o := orm.NewOrm()
	v := Cron{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Cron{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
