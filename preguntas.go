package main

import "fmt"

func main() {

	fmt.Println("Bienvenidos a Preguntando Ando!!!")
	fmt.Println("Por favor, ingrese su nombre")
	var name string
	fmt.Scanf("%s", &name)

	fmt.Println("Muy bien", name)
	fmt.Println("Empecemos!!!")
	fmt.Println("Ingrese sus respuestas en MAYúSCULA")

	fmt.Println("1. ¿Cúal es la capital de Colombia")
	fmt.Print("A. Polombia\n", "B. Bogotá\n", "C. Cali\n", "D. Sandander\n" )
	var answer string
	fmt.Scanf("%s", &answer)

	var score int
	score = 0

	if answer == "B" {
		score =+ 1
		fmt.Println("Felicidades, ganaste", score, "puntos!!!")
	} else {
		fmt.Println("¡Casi aciertas, estudia más!\n", "Tu puntuación es:", score, "puntos")
	}

	fmt.Println("2. ¿Cúal es la montaña más alta del mundo?")
	fmt.Print("A. Monte Everest\n", "B. Cho Oyu\n", "C. Shisha Pangma\n", "D. Ninguna de las anteriores\n")
	fmt.Scanf("%s", &answer)

	if answer == "A" {
		score = score + 1
		fmt.Println("Felicidades, ganaste", score, "puntos!!!")
	} else {
		fmt.Println("¡Casi aciertas, esudia un poco más!\n", "Tu puntuación es:", score, "puntos")
	}

	fmt.Println("3. ¿Cúal es la capital de Italia?")
	fmt.Print("A. Milán\n", "B. Roma\n", "C. Venecia\n", "D. París\n")
	fmt.Scanf("%s", &answer)

	if answer == "B" {
		score = score + 1
		fmt.Println("Felicidades, ganaste", score, "puntos!!!")
	} else {
		fmt.Println("¡Casi aciertas, esudia un poco más!\n", "Tu puntuación es:", score, "puntos")
	}

	fmt.Println("4. ¿En qué año se fundó Heinz?")
	fmt.Print("A. 2000\n", "B. 1555\n", "C. 2003\n", "D. 1869\n")
	fmt.Scanf("%s", &answer)

	if answer == "D" {
		score = score + 1
		fmt.Println("Felicidades, ganaste", score, "puntos!!!")
	} else {
		fmt.Println("¡Casi aciertas, esudia un poco más!\n", "Tu puntuación es:", score, "puntos")
	}

	fmt.Println("5. ¿Cúal es la flor nacional de Japón?")
	fmt.Print("A. Tulipán\n", "B. Esmeralda\n", "C. Cerezo\n", "D. Margarita\n")
	fmt.Scanf("%s", &answer)

	if answer == "C" {
		score = score + 1
		fmt.Println("Felicidades, ganaste", score, "puntos!!!")
	} else {
		fmt.Println("¡Casi aciertas, esudia un poco más!\n", "Tu puntuación es:", score, "puntos")
	}

	fmt.Println("6. ¿Cúal es la formula química del agua?")
	fmt.Print("A. H2\n", "B. CO2\n", "C. H2O\n", "D. H30\n")
	fmt.Scanf("%s", &answer)

	if answer == "C" {
		score = score + 1
		fmt.Println("Felicidades, ganaste", score, "puntos!!!")
	} else {
		fmt.Println("¡Casi aciertas, esudia un poco más!\n", "Tu puntuación es:", score, "puntos")
	}

	fmt.Println("7. ¿Qué país tiene más islas en el mundo?")
	fmt.Print("A. Suecia\n", "B. Italia\n", "C. Venecia\n", "D. Alemania\n")
	fmt.Scanf("%s", &answer)

	if answer == "A" {
		score = score + 1
		fmt.Println("Felicidades, ganaste", score, "puntos!!!")
	} else {
		fmt.Println("¡Casi aciertas, esudia un poco más!\n", "Tu puntuación es:", score, "puntos")
	}

	fmt.Println("8. ¿De qué país es la banda BTS?")
	fmt.Print("A. Milán\n", "B. Pakistán\n", "C. Japón\n", "D. Corea del Sur\n")
	fmt.Scanf("%s", &answer)

	if answer == "D" {
		score = score + 1
		fmt.Println("Felicidades, ganaste", score, "puntos!!!")
	} else {
		fmt.Println("¡Casi aciertas, esudia un poco más!\n", "Tu puntuación es:", score, "puntos")
	}

	fmt.Println("9. ¿En qué año murió John F. Kennedy?")
	fmt.Print("A. 2020\n", "B. 1962\n", "C. 1963\n", "D. 1964\n")
	fmt.Scanf("%s", &answer)

	if answer == "D" {
		score = score + 1
		fmt.Println("Felicidades, ganaste", score, "puntos!!!")
	} else {
		fmt.Println("¡Casi aciertas, esudia un poco más!\n", "Tu puntuación es:", score, "puntos")
	}

	fmt.Println("10. ¿Cuál es el símbolo químico del oro?")
	fmt.Print("A. AU\n", "B. UA\n", "C. AV\n", "D. VA\n")
	fmt.Scanf("%s", &answer)

	if answer == "A" {
		score = score + 1
		fmt.Println("Felicidades, ganaste")
		fmt.Println("Tu puntuación final es:", score, "puntos!!!")
	} else {
		fmt.Println("¡Casi aciertas, esudia un poco más!\n", "Tu puntuación final es:", score, "puntos!!!")
	}

	
}