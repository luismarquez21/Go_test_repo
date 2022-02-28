//Printf en go
/*
Hay algunos operadores de posición que yo considero destacables.

%T, tipo de dato (string, int, boolean, etc).
%v, valor en el formato predeterminado de go
%t, Para booleanos, retorna la palabra false o true
%x, número de base 16
%o, número de base 8
%e, número notación científica
%9.2f, flotante con ancho de 9 y precisión de 2
%.2f, flotante con ancho predeterminado y precisión de 2
%q un string o cadena de texto, entre comillas, previamente escapado
Puedes ver el resto de operadores en la documentación oficial de go.
*/
package main

import "fmt"



func main(){
	texto := "World!"
    numero := 42
    // %s es de string y %d de digit
    fmt.Printf("Hello %s %d", texto, numero)
}