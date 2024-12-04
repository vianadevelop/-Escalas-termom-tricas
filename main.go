package main

import "fmt"

func main() {
	var kelvin float64
	fmt.Print("Digite o valor em Kelvin: ")
	fmt.Scan(&kelvin)

	// Convertendo para Celsius
	// Fórmula: C = K - 273
	celsius := kelvin - 273

	// Exibindo o resultado
	fmt.Printf("O valor em Celsius é: %.2f°C\n", celsius)
}
