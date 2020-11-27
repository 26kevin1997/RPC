package main

import (
	"fmt"
	"net/rpc"
)
func cliente()  {
	c,err:=rpc.Dial("tcp","127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op, result int64
	var nombre, materia,calificacion string
	cont := 1
	for cont != 0 {
		op = 0
		fmt.Println("\t- Menu -")
		fmt.Println("1.- Agregar Calificación")
		fmt.Println("2.- Mostrar Promedio Alumno")
		fmt.Println("3.- Promedio General")
		fmt.Println("4.- Promedio Materia")
		fmt.Println("0.- Salir")
		fmt.Scan(&op)
		switch op{
		case 1:
			fmt.Println("Agregar Nombre Alumno: ")
			fmt.Scan(&nombre)
			fmt.Println("Agregar Nombre Materia: ")
			fmt.Scan(&materia)
			fmt.Println("Agregar Calificación Alumno: ")
			fmt.Scan(&calificacion)
			final := nombre+"/"+materia+"/"+calificacion
			err = c.Call("Server.Agregar", final,&result)
			if err != nil {
				fmt.Println(err)
			}
		case 2:
			fmt.Println("Nombre Alumno:")
			fmt.Scan(&nombre)
			err = c.Call("Server.Promedio_Alumno", nombre,&result)
			if err != nil {
				fmt.Println("El Promedio es: ",err)
			}
			
		case 3:
			err = c.Call("Server.Promedio_General", "0",&result)
			if err != nil {
				fmt.Println("El Promedio Genereal es: ",err)
			}
			
		case 4:
			fmt.Println("Nombre Materia:")
			fmt.Scan(&materia)
			err = c.Call("Server.Promedio_Materia", materia,&result)
			if err != nil {
				fmt.Println("El Promedio por la materia: ",materia," es: ",err)
			}
		default:
			fmt.Println("Fin")
			cont = 0
		
	}
}
}

func main()  {
	cliente()
}