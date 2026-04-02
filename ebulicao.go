package main

import "fmt"

// Função para calcular a temperatura de ebulição da água em Kelvin
const temperaturaEbuliçãoKelvin = 373.15
const temperaturaFusaoKelvin = 273.15

func main() {
	tempK := temperaturaEbuliçãoKelvin
	tempC := tempK - 273.15
	fmt.Printf("A temperatura de ebulição da água em K é: %g(%T)\nA temperatura de ebulição da água em °C é: %g(%T)", tempK, tempK, tempC, tempC)

	fusaoK := temperaturaFusaoKelvin
	fusaoC := fusaoK - 273.15
	fmt.Printf("\nA temperatura de fusão da água em K é: %g(%T)\nA temperatura de fusão da água em °C é: %g(%T)", fusaoK, fusaoK, fusaoC, fusaoC)

}