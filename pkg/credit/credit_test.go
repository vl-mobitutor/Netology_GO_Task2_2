package credit_test

import (
	"fmt"
	"github.com/vl-mobitutor/Netology_GO_Task2_2/pkg/credit"
)

func ExampleCalculate() {

	//1-й тест  - выходные данные по расчетной формуле на сайте https://www.banki.ru/wikibank/raschet_annuitetnogo_plateja/
	//Все суммы приведены в копейках
	fmt.Println(credit.Calculate(1_000_000_00, 36, 20))


	//2-й тест  - выходные данные по калькулятору на сайте https://www.banki.ru/services/calculators/credits/?amount=1000000&currency=RUB&period=1095&rate=20
	//Все суммы приведены в копейках
	fmt.Println(credit.Calculate(1_000_000_00, 36, 20))

	// Output:
	// 3718400 33862300 133862300
	// 3716400 33790400 133790400
}