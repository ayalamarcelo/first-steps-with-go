# Ejercicio: Hola mundo
## Configuración del IDE para que se abra desde la CLI
Puede abrir el IDE VSC dese el símbolo del sistema de la CLI y empezar a editar los archivos en el área de trabajo actual. Esta característica del IDE se implementa mediante la variable $PATH del entorno del sistema. Cuando la característica está habilitada, puede usar el comando `code .` de la CLI para abir el IDE y editar archivos en el directorio actual.

Algunas instalaciones de VSC agregan compatibilidad con la CLI a la variable $PATH de forma predeterminada. Es posible que vea esta opción durante el proceso de instalación. Si ya tiene esta opción configurada, está todo listo. De lo contrario, es posible que tenga que seguir estos pasos para usar esta característica.

## Ejecución de la aplicación de Go
Para ejecutar la aplicación de Go, utiliza el siguiente comando en la terminal:

```bash
go run main.go
```
Deberías ver la siguiente salida:

```
Hello World!
```
El comando `go run` hace dos cosas. Compila la apicación y, una vez que se completa correctamente, ejecuta la aplicación.

## Compilación de un archivo ejecutable
Para generar un archivo ejecutable para el programa, usa este comando:

```bash
go build main.go
```

Cuando finaliza el comando `go build`, se genera una aplicación ejecutable que se puede ejecutar en cualquier momento sin más procesamiento. El comando solo genera un archivo ejecutable. no ejecuta el programa como el comando `go run`.

## Revisión del contendio de /src
Este es el aspecto que debería tener el proyecto ahora en la vista Explorador > SRC:

SRC/
  helloworld/
    main(principal)
    main.go

> En la vista Explorador, el nombre del archivo sin la extensióm es el archivo ejecutable que puede usar para ejecutar el programa. (En Windows, este archivo tiene la extensión .exe). Al desarrollar, use el comando `go run`. Para compilar los archivos binarios de la aplicación, use el comando `go build` e implemente el archivo ejecutable binario en un entorno adecuado.

## ¿Qué acaba de escribir en Go?
Ha creado su primera aplicación de Go y se ha asegurado de que se compila y se ejecuta. Examinemos el código línea a línea.

Comenzamos con la primera instrucción del archivo de Go:

```go
package main
```
La instrucción `package main` especifíca cómo se le indica a Go que la aplicación que estamos creando es un programa ejecutable (un archivo que puede ejecutar). Nuestra aplicación "Hola mundo" forma parte del paquete main. Un paquete es un conjunto de archivos de código fuente comunes. Todas las aplicaciones ejecutables tienen esta primera línea, aunque el proyecto o el archivo tenga otro nombre.

Examinaremos más de cerca estos conceptos en el módulo siguiente. Por ahora, es necesario saber que todos los programas ejecutables deben formar parte del paquete `main`.

>[!Important] Todos los programas ejecutables deben formar parte del paquete `main`.

Esta es la siguiente línea de archivo de Go:

```go
import "fmt"
```
La instrucción `import` proporciona acceso al programa a otro código en paquetes diferentes. En este caso, *fmt* es un paquete de biblioteca estándar.
Necesita esta instrucción `import` porque, más adelante, vas a usar una función de este paquete para imprimir un mensaje en la pantalla. Puede incluir tantas intrucciones `import` como quiera o necesite en el programa. Sin embargo, Go es idiomático en este sentido. Si importa un paquete, pero no usa una función correspondiente del paquete, la aplicación no se compilará. Una excelente característica de VSC es que quita automáticamente las importaciones no utilizadas en un programa al guardar el archivo.

```go
func main() {
  fmt.println("Hello World!")
}
```
La instrucción `func` es una palabra reservada que se usa para declarar una función. Esta primera función se denomina `main` porque es el punto inicial de nuestro programa. Solo puede tener una función `main()` en `package main` (el que definió en la primera línea). En la función `Println` desd el paquete `fmt`. Envió un mensaje de texto que quería ver en la pantalla.