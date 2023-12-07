package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	//"sync"
	//"time"
)

// Definición de la estructura del Bloque de Control de Proceso (BCP)
type BCP struct {
	ID           string
	Estado       string
	ContadorProg int
	E_S          int
}

func pop(slice []BCP) (BCP, []BCP) {
    bcp:= BCP{}
    if len(slice) == 0 {
        // Manejar el caso de un slice vacío
        return bcp, slice
    }
    // Obtener el primer elemento
    poppedElement := slice[0]
    // Crear un nuevo slice sin el primer elemento
    remainingSlice := slice[1:]
    // Devolver el elemento eliminado y el nuevo slice
    return poppedElement, remainingSlice
}



// funcion para escribir archivo de salida
func printS(texto, archivoSalida string) {
	archivo, err := os.Create(archivoSalida)
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer archivo.Close() // Asegúrate de cerrar el archivo al final

	// Crea un escritor para escribir en el archivo
	escritor := bufio.NewWriter(archivo)

	// Escribe en el archivo
	_, err = escritor.WriteString(texto + "\n")
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	// Se vacia el búfer
	escritor.Flush()

	fmt.Println("Se ha escrito en el archivo correctamente.")
}

func LeerProsesos(archivoEntrada string) ([]string, []int) {
	var procesos []string
	var tiempos []int

	archivo, err := os.Open(archivoEntrada)
	if err != nil {
		fmt.Println("Error al abrir el archivo de orden de creacion de procesos:", err)
		return procesos, tiempos
	}
	defer archivo.Close() // Asegúrate de cerrar el archivo al final

	// Crea un nuevo escáner para leer el archivo línea por línea
	escaner := bufio.NewScanner(archivo)

	// Itera sobre cada línea del archivo
	for escaner.Scan() {
		linea := escaner.Text()
		contenidoLinea := strings.Split(linea, " ")
		if contenidoLinea[0] != "#" {
			t, err := strconv.Atoi(contenidoLinea[0])
			if err != nil {
				fmt.Println("Error al convertir string a numero", err)
			}
			for i := 1; i < len(contenidoLinea); i++ {
				q := t + i - 1
				procesos = append(procesos, contenidoLinea[i])
				tiempos = append(tiempos, q)
			}
		}

	}

	// Verifica si hubo errores durante la lectura
	if err := escaner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return procesos, tiempos
	}
	return procesos, tiempos

}

func finalle(probabilidad int) bool {
	ret := true
	numeroAleatorio := rand.Intn(probabilidad) + 1
	if numeroAleatorio == 1 {
		ret = false
	}
	return ret
}

func creacionProceso(link string) BCP {
	bcp := BCP{}
	Link := link + ".txt"
	archivo, err := os.Open(Link)
	if err != nil {
		fmt.Println("Error al abrir el archivo el programa mencionado no existe:", err)
		return bcp
	}
	defer archivo.Close() // Asegúrate de cerrar el archivo al final

	// Crea un nuevo escáner para leer el archivo línea por línea
	escaner := bufio.NewScanner(archivo)

	// Crea un slice para almacenar las líneas del archivo
	var lineas []string

	// Itera sobre cada línea del archivo
	for escaner.Scan() {
		linea := escaner.Text()
		lineas = append(lineas, linea)
	}
	if lineas[0] == link {
		return BCP{
			ID:           link,
			Estado:       "Nuevo",
			ContadorProg: 0,
			E_S:          0,
		}
	}
	return bcp
}

func Push(bcp BCP, procesos []string, tiempos []int) BCP {





var Fin = false

var ListaListos = make([]BCP, 0)
var ListaBloq = make([]BCP, 0)

func DispatcherPull() BCP {
	ultimoElemento := ListaListos[0]
	ListaListos = ListaListos[1:len(ListaListos)]
	return ultimoElemento
}

func Procesador(Procesos []string, CreacionProceso []int) string {
	//Procesando:=false
	contadorPros := 1
	salida := ""
	for {
		salida = salida + strconv.Itoa(contadorPros)
		for i := 0; i < len(Procesos); i++ {

			if CreacionProceso[i] == contadorPros {
				ListaListos = append(ListaListos, creacionProceso(Procesos[i]))
				salida = salida + "PUSH_Listo  " + Procesos[i] + "   Dispatcher" + "   101\n"

			}
		}
		if contadorPros >= 166 {
			break // Salir del bucle
		}
		salida = salida + "\n"
		contadorPros++
	}
	return salida
}

func main() {

	argumentos := os.Args[1:]
	archivoEntrada := argumentos[3]
	P, err := strconv.Atoi(argumentos[2])
	o, err := strconv.Atoi(argumentos[1])
	if err != nil {
		fmt.Println("Error al convertir la cadena a entero:", err)
		return
	}
	fmt.Println(P, o)
	// Crea un slice para almacenar los procesos
	procesos, tiempos := LeerProsesos(archivoEntrada)

	//se simulan los tiempos de CPU

	//condicionDeSalida:=false
	texto := "# Tiempo de CPU Tipo Instrucción Proceso/Despachador Valor CP\n"

	texto = texto + Procesador(procesos, tiempos)
	contadorPros := 1

	for {
		//fmt.Println(contadorPros)
		for i := 0; i < len(procesos); i++ {
			if tiempos[i] == contadorPros {
				//fmt.Println(procesos[i])
			}
		}
		if contadorPros >= 166 {
			break // Salir del bucle
		}
		contadorPros++
	}

	//-------------------------------------------------------------

	// Configuración inicial
	/*var wg sync.WaitGroup
	estados := []string{"Nuevo", "Listo", "Ejecutando", "Bloqueado", "Terminado"}

	// Crear colas para cada estado
	colas := make([]chan BCP, len(estados))
	for i := range colas {
		colas[i] = make(chan BCP, 5) // Tamaño arbitrario de la cola
	}

	// Crear e iniciar procesos
	for i := 1; i <= 3; i++ {
		// Iniciar procesos en el estado "Nuevo"
		colaprox := colas[0]
		wg.Add(1)
		go cambiarEstado(i, "Inicio", &wg, colaprox)
	}
	// Ejecutar simulación
	for i := 0; i < len(estados)-1; i++ {
		for j := 1; j <= 3; j++ {
			wg.Add(1)
			go ejecutarProceso(j, colas[i], colas[i+1], &wg)
		}

		// Wait for processes to finish execution before moving to the next state
		wg.Wait()

		// Move processes to the next state using buffered channels
		for j := 1; j <= 3; j++ {
			wg.Add(1)
			go cambiarEstado(j, estados[i], &wg, colas[i+1])
		}
	}

	// Wait for the final state to complete
	wg.Wait()*/

	printS(texto, argumentos[4])
}
