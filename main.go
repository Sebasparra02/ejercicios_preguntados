package main

import "fmt"

func main() {
	puntuacion := 0
	var nombre string
	fmt.Println("*** Bienvenidos a preguntados ***")
	fmt.Println("nos gustaria saber tu nombre")
	fmt.Scanf("%s", &nombre)

	pregunta_1 := `¿cual es el animal mas temido de la selva?
    A) Tigre de vengala
    B) Leon
    C) Leopardo
	D) Puma`
	fmt.Println(pregunta_1)
	_, err := fmt.Scanf("%s", &pregunta_1)
	if err != nil {
		fmt.Println("ups algo ha fallado, intentalo de nuevo")
		panic(err)
	}

	if pregunta_1 == "a" {
		fmt.Println("has respondido correctamente")
		puntuacion += 1
	} else {
		println("perdon,me parese que has repondido mal, sigamos con la siguiente pregunta")
	}

	pregunta_2 := `¿cual es el pais mas grande del mundo?
	A) Brasil
	B) Rusia
	C) Inglaterra
	D) Estados unidos`
	fmt.Println(pregunta_2)
	_, err = fmt.Scanf("%s", &pregunta_2)
	if err != nil {
		fmt.Println("ups algo ha fallado, intentalo de nuevo")
		panic(err)
	}
	if pregunta_2 == "b" {
		fmt.Println("has respondido correctamente")
		puntuacion += 1
	} else {
		println("perdon,me parese que has repondido mal, sigamos con la siguiente pregunta")
	}

	pregunta_3 := `¿Cuál es el mar más extenso de todo el mundo?
	A) El mar Atlantico
	B) El mar Mediterráneo
	C) El mar cantábrico
	D) EL mar pacifico`
	fmt.Println(pregunta_3)
	_, err = fmt.Scanf("%s", &pregunta_3)
	if err != nil {
		fmt.Println("ups algo ha fallado, intentalo de nuevo")
		panic(err)
	}
	if pregunta_3 == "d" {
		fmt.Println("has respondido correctame")
		puntuacion += 1
	} else {
		println("perdon,me parese que has repondido mal, sigamos con la siguiente pregunta")
	}

	pregunta_4 := `¿cual es el animal mas grande del oceano?
	A) Ballena
	B) Tiburon
	C) Pulpo
	D) Calamar`
	fmt.Println(pregunta_4)
	_, err = fmt.Scanf("%s", &pregunta_4)
	if err != nil {
		fmt.Println("ups algo ha fallado, intentalo de nuevo")
		panic(err)
	}
	if pregunta_4 == "a" {
		fmt.Println("has respondido correctame")
		puntuacion += 1
	} else {
		println("perdon,me parese que has repondido mal, sigamos con la siguiente pregunta")
	}

	pregunta_5 := `¿cual es el rio mas grande del mundo?
	A) Amazonas
	B) Obi
	C) Nilo
	D) Ninguno`
	fmt.Println(pregunta_5)
	_, err = fmt.Scanf("%s", &pregunta_5)
	if err != nil {
		fmt.Println("ups algo ha fallado, intentalo de nuevo")
		panic(err)
	}
	if pregunta_5 == "c" {
		fmt.Println("has respondido correctame")
		puntuacion += 1
	} else {
		println("perdon,me parese que has repondido mal, sigamos con la siguiente pregunta")
	}

	pregunta_6 := `Donde esta transilvania?
	A) Rusia
	B) Francia
	C) Rumania
	D) Grecia`
	fmt.Println(pregunta_6)
	_, err = fmt.Scanf("%s", &pregunta_6)
	if err != nil {
		fmt.Println("ups algo ha fallado, intentalo de nuevo")
		panic(err)
	}
	if pregunta_6 == "c" {
		fmt.Println("has respondido correctame")
		puntuacion += 1
	} else {
		println("perdon,me parese que has repondido mal, sigamos con la siguiente pregunta")
	}

	pregunta_7 := `¿En qué año cayó el muro de Berlín?
	A) 1989
	B) 1980
	C) 1985
	D) 1999`
	fmt.Println(pregunta_7)
	_, err = fmt.Scanf("%s", &pregunta_7)
	if err != nil {
		fmt.Println("ups algo ha fallado, intentalo de nuevo")
		panic(err)
	}
	if pregunta_7 == "a" {
		fmt.Println("has respondido correctame")
		puntuacion += 1
	} else {
		println("perdon,me parese que has repondido mal, sigamos con la siguiente pregunta")
	}

	pregunta_8 := `¿Cuántos huesos tiene el cuerpo humano?
	A) 206
	B) 204
	C) 205
	D) 200`
	fmt.Println(pregunta_8)
	_, err = fmt.Scanf("%s", &pregunta_8)
	if err != nil {
		fmt.Println("ups algo ha fallado, intentalo de nuevo")
		panic(err)
	}
	if pregunta_8 == "a" {
		fmt.Println("has respondido correctame")
		puntuacion += 1
	} else {
		println("perdon,me parese que has repondido mal, sigamos con la siguiente pregunta")
	}

	pregunta_9 := `¿Cuántas estrellas hay en la bandera estadounidense?
	A) 30
	B) 45
	C) 50
	D) 10`
	fmt.Println(pregunta_9)
	_, err = fmt.Scanf("%s", &pregunta_9)
	if err != nil {
		fmt.Println("ups algo ha fallado, intentalo de nuevo")
		panic(err)
	}
	if pregunta_9 == "c" {
		fmt.Println("has respondido correctame")
		puntuacion += 1
	} else {
		println("perdon,me parese que has repondido mal, sigamos con la siguiente pregunta")
	}

	pregunta_10 := ` ¿Cuáles son los únicos mamíferos que pueden volar?
	A) Murcielagos
	B) colibri
	C) Aguila
	D) Todas las anteriores`
	fmt.Println(pregunta_10)
	_, err = fmt.Scanf("%s", &pregunta_10)
	if err != nil {
		fmt.Println("ups algo ha fallado, intentalo de nuevo")
		panic(err)
	}
	if pregunta_10 == "a" {
		fmt.Println("has respondido correctame")
		puntuacion += 1
	} else {
		fmt.Println("tu respuesta es incorrecta! \n, tu puntuacion fue de:", puntuacion, "puntos")

	}

	fmt.Println("hola ", nombre, "tu puntuacion final fue de: ", puntuacion, "puntos")

}
