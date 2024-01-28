# Creación de funciones
En Go, las funciones permiten agrupar un conjunto de instrucciones a las que se puede llamar desde otras partes de la aplicación. En lugar de crear un programa con muchas instrucciones, puedes usar funciones para organizar el código y hacer que sea más legible. Cuando el código es más legible, también es más fácil de mantener.

Hasta este punto, hemos llamado a la función `fmt.Println` y hemos escrito código en la función `main()`. En esta sección, exploraremos cómo puedes crear funciones personalizadas. También veremos otras técnicas que puedes usar con las funciones de Go.

## Función main
La función con la que has estado interactuando hasta es la función `main()`. Todos los programas ejecutables que se usan Go tienen esta función porque es el punto inicial del programa. Solo puede haber una función `main()`. Veremos cómo crear paquetes de módulos.

Antes de pasar a los conceptos básicos de la creación de funciones personalizada en Go, echemos un vistazo a un aspecto fundamental de la función `main()`. Como puede haber observado, la función `main()` no tiene ningún parámetro y no devuelve nada. Pero esto no signfica que no pueda ller valores del usuario, como los argumentos de la línea de comandos. Si necesitas acceder a lso argumentos de la línea de comandos en Go, puedes hacerlo con el paquete os y la variable `os.Args`, que contiene todos los argumentos que se pasan al programa.
El código siguiente lee dos números desde la línea de comandos y los suma:

```go
package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    number1, _ := strconv.Atoi(os.Args[1])
    number2, _ := strconv.Atoi(os.Args[2])
    fmt.Println("Sum:", number1+number2)
}
```
La variable `os.Args` contiene todos los argumentos de la línea de comandos que se pasan al programa. Estos valores son de tipo `string`, por lo que debe convertirlos en `int` para sumarlos.

Para ejecutar el programa, usa el siguiente comando:

```go
go run main.go  # 3 5
```
Este es el resultado:

```output
Sum: 8
```
Veamos cómo se puede refactorizar el código anterior y crear una función personalizada.

## Funciones personalizadas

```go
func name(parameters) (results) {
    # body-content
}
```
Observe que se usa la palabra clave `func` para definir una función y, después, se le asigna un nombre. Después del nombre, es necesario especificar la lista de parámetros para la función. Puede haber cero o más parámetros. Además, se pueden definir los tipos de valor devueltos de la función, que también pueden ser cero o más. Hablaremos sobre cómo devolver varios valores en la siguiente sección. Una vez definidos todos esos valores, se escribe el contenido del cuerpo de la función.

Para practicar esta técnica, vamos a refactorizar el código de la sección anterior para sumar los números en una función personalizada. Usaremos el siguiente código:

```go
package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    sum := sum(os.Args[1], os.Args[2])
    fmt.Println("Sum:", sum)
}

func sum(number1 string, number2 string) int {
    int1, _ := strconv.Atoi(number1)
    int2, _ := strconv.Atoi(number2)
    return int1 + int2
}
```
Este código crea una función llamada `sum` que toma dos argumentos de `string`, los convierte en `int` y luego, devuelve el resultado de su suma. Al definir un tipo de valor devuelto, la función debe devolver un valor de dicho tipo.

En Go, es posible establecer un nombre para el valor devuelto de una función, como si fuera una variable. Por ejemplo, la función `sum` se podría refactorizar de la siguiente manera:

```go
func sum(number1 string, number2 string) (result int) {
    int1, _ := strconv.Atoi(number1)
    int2, _ := strconv.Atoi(number2)
    result = int1 + int2
    return
}
```
Observa que ahora se debe incluir el valor del resultado de la función entre paréntesis. También se pueda usar la variable dentro de la función y agregar simplemente una línea `return` al final. Go devolverá los valores actuales de las variables devueltas. La simplicidad de escribir la palabra clave `return` al final de la función es atractiva (especialmente cuando hay más de un valor devuelto). Este enfoque no se recomienda. Puede no estar claro qué devuelve la función.

## Devolución de varios valores
En Go, una función puede devolver más de un valor. Estos valores se pueden definir de forma similar a la definición de los parámetros de la función. En otras palabras, se específica un tipo y un nombre, aunque el nombre es opcional.

Por ejemplo, supongamos que quiere crear una función que suma dos números, pero que también los multiplica. El código de la función tendría el siguiente aspecto.

```go
func calc(number1 string, number2 string) (sum int, mul int) {
    int1, _ := strconv.Atoi(number1)
    int2, _ := strconv.Atoi(number2)
    sum = int1 + int2
    mul = int1 * int2
    return
}
```
Ahora necesita dos variables par almacenar los resultados de la función, de los contrario, no se compilará. Este es su aspecto:

```go
package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    sum, mul := calc(os.Args[1], os.Args[2])
    fmt.Println("Sum:", sum)
    fmt.Println("Mul:", mul)
}
```
Otra característica interesante de Gi es que, si no necesita uno de los valores que devuelve una función, puede descartarlo asignando el valor devuelto a la variable `_`. La avriable `_` es la forma idiomática que usa Go para omitir los valores devueltos, la cual permite que el programa se compile. Por lo tanto, si solo quisiera el valor de la suma, podría usar este código:

```go
package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    sum, _ := calc(os.Args[1], os.Args[2])
    fmt.Println("Sum:", sum)
}
```
Analizaremos en más detalles cómo omitir los valores devueltos de las funciones cuando exploremos el control de errores en un módulo posterior.

## Cambio de los valores de los parámetros de la función (punteros)
Cuando se pasa un valor a una función, los cambios en esa función no afectan al autor de lllamada. Go es un lenguaje de programación de "paso por valor". Cada vez que se pasa un valor a una función, Go toma ese valor y crea una copia local (una nueva variable en la memoria). Los cambios que realice en esa variable en la función no afectarán a la variable que haya enviado a la función.

Por ejemplo, supongamos que crea una función para adctualizar el nombre de una persona. Observe lo que ocurre cuando se ejecuta este código:

```go
package main

import "fmt"

func main() {
    firstName := "John"
    updateName(firstName)
    fmt.Println(firstName)
}

func updateName(name string) {
    name = "David"
}
```
Aunque ha cambiado el nombre a `David` en la función, la salida sigue siendo "John". La salida no ha cambiado porque el cambio en la función `updateName` modifica únicamente la copia local. Go ha pasado el valor de la variable, no la propia variable.
Si quieres que el cambio realizado en la función `updateName` afecte a la variable `firstName` de la función `main`, debe usar un puntero. Un *puntero* es una variable que contiene la dirección de memoria de otra variable. Cuando se envía un puntero a una función, no se pasa un valor, se pasa una dirección de memoria. Por lo tanto, todos los cambios que se realicen en esa variable afectarán al autor de la llamada.

En Go, hay dos operadores para trabajar con punteros:

- El operador `&` toma la dirección del objeto que la sigue.
- El operador `*` desreferencia un puntero. Le proporciona acceso al objeto en la dirección que contiene el puntero.

Vamos a modificar nuestro ejemplo anterior para aclarar cómo funcionaran los punteros:

```Go
package main

import "fmt"

func main() {
    fistName := "John"
    updateName(&firstName)
    fmt.Println(firsName)
}

func updateName(name *string) {
    *name = "David"
}
```
Ejecuta el código anterior. Fijate en que la salida ahora muestra `David` en lugar de `John`.

Lo primero que tienes que hacer es modificar la signatura de función para indicar que quiere recibir un puntero. Para ello, cambie el tipo de parámetro de `string` a `*string`. Este último tipo sigue siendo una cadena, pero ahora el cambio es *de un puntero a una cadena*. Después, a la hora de asignar un nuevo valor a esa variable, debe agregar la estrella (*) en el lado izquierdo de la variable para proporcionar el valor de esta. De este modo, cuando llama a la función `updateName`, no envía el valor, sino la dirección de memoria de la variable. El símbolo `&` del lado izquierdo de la variable indica la dirección de esta.