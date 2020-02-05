package task

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"log"
	"reflect"
	"tax-go-beego/models"
	"time"
)

func init() {
	SyncInsuranceTask()
}

func SyncInsuranceTask() {

	cronExpress :=  beego.AppConfig.String("cronExpress")

	tk1 := toolbox.NewTask("SyncTask", cronExpress, SyncTask)
	toolbox.AddTask("SyncTask", tk1)
	toolbox.StartTask()
}

func SyncTask() error {
	log.Println("Task Is Run。")
	var year = time.Now().Year()
	log.Println(year)
	add(year)
	return nil
}

func add(year int) {
	var list, _ = models.GetInsuranceByYear(year)
	if list == nil {
		var newYear = year - 1
		fmt.Println(newYear)
		var newList, _ = models.GetInsuranceByYear(newYear)
		if newList != nil {
			for _, v := range newList {
				var val = reflect.ValueOf(v)
				iv := val.Interface()
				var insurance = iv.(models.Insurance)

				item := new(models.Insurance)
				item.Year = int64(year)
				item.ProvinceCode = insurance.ProvinceCode
				item.CityCode = insurance.CityCode
				item.Pension = insurance.Pension           //养老比例
				item.Medical = insurance.Medical           //医疗比例
				item.Unemployment = insurance.Unemployment //失业比例
				item.Provident = insurance.Provident       //公积金比例
				item.CityDesc = insurance.CityDesc

				item.PensionUpper = insurance.PensionUpper //养老上限

				item.PensionLower = insurance.PensionLower //养老下限

				item.ProvidentUpper = insurance.ProvidentUpper //公积金比例上限
				item.ProvidentLower = insurance.ProvidentLower //公积金比例下限

				id, _ := models.AddInsurance(item)
				log.Println(id)
			}
		}
	}
}
