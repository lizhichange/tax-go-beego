package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"reflect"
	"strconv"
	"tax-go-beego/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var list, _ = models.GetProvinceByHot("1")
	var provinceList []ProvinceVO
	for _, v := range list {
		var val = reflect.ValueOf(v)
		var p ProvinceVO
		iv := val.Interface()
		province := iv.(models.Province)

		p.Id = province.Id
		p.Name = province.Name
		p.Code = province.Code
		//根据省份code 查询城市信息
		cs, _ := models.GetCityByProvinceCode(province.Code, "1")
		p.City = cs
		provinceList = append(provinceList, p)
	}
	c.Data["provinceList"] = provinceList
	c.TplName = "index.html"
}

type ProvinceVO struct {
	Id   int64
	Code string
	Name string
	City []models.City
}

const nu = 100

//免征额
const exemption = 5000

//专项扣除
const deduction = 0

/**
 *计算
 **/
func (c *MainController) Calc() {

	var cityCode = c.Ctx.Input.Param(":cityCode")

	cs, _ := models.GetCityByCityCode(cityCode)
	var provinceCode string
	if cs != nil {
		city := cs[0]
		provinceCode = city.ProvinceCode
	}

	//税前金额
	var amount = c.Ctx.Input.Param(":amount")
	//获取城市社保金额比例配置信息
	i, _ := models.GetInsuranceByCode(provinceCode, cityCode)
	in := i[0]
	var afterAmount float64

	//转换float
	var amountFloat, _ = strconv.ParseFloat(amount, 64)
	//医疗保险
	var medical = amountFloat * in.Medical / nu
	//失业保险
	var unemployment = amountFloat * in.Unemployment / nu

	//转换成int64
	amountInt, _ := strconv.ParseInt(amount, 10, 64)
	//养老保险金：
	var pension = amountInt * in.Pension / nu
	//公积金
	var provident = amountInt * in.Provident / nu

	//社保金额&公积金=医疗保险+失业保险+养老保险金+公积金
	var socialAmount = medical + unemployment + float64(pension) + float64(provident)

	//税后工资 = 税前工资 - 社保金额 - 减去公积金-减去免征额-专项扣除fmt.Println(pension, medical, unemployment, provident, socialAmount, afterAmount)

	ml, _ := models.GetAllMonthlyInfo(nil, nil, nil, nil, 0, 0)
	var item models.MonthlyInfo
	for i := len(ml) - 1; i > 0; i-- {
		iv := reflect.ValueOf(ml[i]).Interface()
		info := iv.(models.MonthlyInfo)
		if afterAmount >= float64(info.Monthly) {
			item = info
			break
		}
	}
	//计算纳税额
	var personalIncomeTax = afterAmount*float64(item.Rate)/nu - float64(item.QuickDeduction)

	personalIncomeTaxRate := personalIncomeTax / amountFloat * nu

	//税后工资
	var f = amountFloat - personalIncomeTax - medical - unemployment - float64(pension) - float64(provident)

	afterAmountRate := f / amountFloat * nu
	fmt.Println("纳税额:", personalIncomeTax, "税后工资:", f)

	var result = &CalcResult{}

	//纳税额
	result.PersonalIncomeTax = personalIncomeTax
	//税后工资
	result.AfterAmount = f
	//医疗保险
	result.Medical = medical
	//失业保险
	result.Unemployment = unemployment
	//养老保险金
	result.Pension = pension
	//公积金
	result.Provident = provident
	result.Amount = amount
	result.Rate = item.Rate
	result.Exemption = exemption
	//速算扣除数
	result.QuickDeduction = item.QuickDeduction
	result.SocialAmount = socialAmount

	result.PensionRate = in.Pension           //养老比例率
	result.MedicalRate = in.Medical           //医疗比例率
	result.UnemploymentRate = in.Unemployment //失业比例率
	result.ProvidentRate = in.Provident       //公积金比例率
	result.PersonalIncomeTaxRate = int64(personalIncomeTaxRate)
	result.AfterAmountRate = int64(afterAmountRate)
	c.Data["json"] = result
	c.ServeJSON()
}

type CalcResult struct {
	//纳税额
	PersonalIncomeTax float64
	//纳税额 比例
	PersonalIncomeTaxRate int64
	//税后工资
	AfterAmount float64
	//税后工资比例
	AfterAmountRate int64

	//养老保险金
	Pension int64
	//税率
	Rate int64
	//医疗保险
	Medical float64

	PensionRate      int64   //养老比例率
	MedicalRate      float64 //医疗比例率
	UnemploymentRate float64 //失业比例率
	ProvidentRate    int64   //公积金比例率

	//失业保险
	Unemployment float64
	// 税前工资
	Amount string
	//公积金
	Provident int64
	//免征额
	Exemption int64
	//五险一金
	SocialAmount float64
	// 速算扣除数
	QuickDeduction int64
}
