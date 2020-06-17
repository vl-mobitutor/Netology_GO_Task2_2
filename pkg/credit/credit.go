package credit

import "math"

//ФУНКЦИЯ РАСЧЕТА АННУИТЕТА
//Принимаемые параметры - сумма кредита в копейках, срок кредита в месяцах, ставка по кредиту в % годовых
//Возвращаемые значения - ежемесячный аннуитетный платеж, переплата по кредиту, общая выплата по кредиту (все в копейках)
func Calculate (loanAmount int64, loanPeriod, loanPercent int) (monthlyPayment, overPayment, totalPayment int64) {

	var percentageMonthly, annuityRate float64

	//Расчет месячной процентной ставки
	percentageMonthly = float64(loanPercent) / 12 / 100

	//Расчет коэффициента аннуитета
	annuityRate = percentageMonthly * math.Pow((percentageMonthly + 1), float64(loanPeriod)) / (math.Pow((percentageMonthly + 1), float64(loanPeriod)) - 1)

	//Расчет суммы ежемесячного аннуитетного платежа
	monthlyPayment = int64(float64(loanAmount) * annuityRate)
	monthlyPayment = int64(math.Round(float64(monthlyPayment) / 100) * 100) //Математическое округление до целых рублей

	//Расчет общей выплаты по кредиту
	totalPayment = int64(loanPeriod) * monthlyPayment

	//Расчет переплаты по кредиту
	overPayment = totalPayment - loanAmount


	return
}