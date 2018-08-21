package controllers

import (
	"github.com/revel/revel"
	"math"
)

type App struct {
	*revel.Controller
}
func (c App) FormulaTelescopica(M, x, C, N float64) float64{
	var Y float64
	Y = M*x*(math.Pow((x+1), N)) - C * (math.Pow((x+1), N) - 1)
	return Y
}
func (c App) FormulaTelescopicaPrima(M, x, C, N float64) float64{
	var Y float64
	Y = M *(math.Pow((x+1), N)) + N * M * x * (math.Pow((x+1), (N - 1))) - C * N * (math.Pow((x + 1), (N - 1)))
	return Y
}
func (c App) Formula(x float64,C float64, N int) float64{
	var M float64
	for i:=1 ;i<=N ;i++{
		M += (C / (math.Pow((x + 1) ,float64(i))))
	}
	return M
}
func (c App) Newton(M, C, N float64) float64{
	var Xf, Xi float64
	Xi = 1
	e := 0.0001
	for  i:=0; i<=10000; i++{
		Xf = Xi - c.FormulaTelescopica(M,Xi,C,float64(N))/c.FormulaTelescopicaPrima(M,Xi,C,float64(N))
		if c.FormulaTelescopica(M,Xf + e,C,float64(N))*c.FormulaTelescopica(M,Xf - e,C,float64(N)) < 0 && math.Abs(Xf - Xi) <= e{
			return Xf * 100 
			//return c.Formula(Xf, C, N)
		}
		Xi = Xf
	}
	//return c.Formula(Xf,C,N)
	return 0
}
func (c App) Index() revel.Result {
	greeting:="Simulador de Interés mensual y Costo Anual equivalente"
	return c.Render(greeting)
}

func (c App) Cae(M, C, N float64) revel.Result {
		interes:=c.Newton(M,C,N)
		costoa:=c.Newton(M,C,N) * 12
		montot:=(C*N)
		return c.Render(interes, costoa, montot)
}
