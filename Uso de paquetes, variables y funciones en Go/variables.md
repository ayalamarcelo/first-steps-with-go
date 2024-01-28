# Declaración y uso de variables
Vamos a comenzar este módulo examinando cómo se declaran y cómo se usan las variables en GO. Hay vaarias maneras de declarar una variable. Veremos cada una de ellas y escogerás la que mejor se adapte a tus necesidades o a su estilo. A medida que exploremos los conceptos básicos de las variables, destacaremos algunas características específicas de Go que no suele presentar otros lenguajes de programación.

En esta sección, se incluyen fragmentos de código que se pueden ejecutar en VSC on el área de juegos de Go, el playground.

## Declaración de variables
Para declarar una variable, se usa la palabra clave `var`:

```go
var firstName string
```
La instrucción anterior declara una variable denomidnada `firstName` de tipo `string`. Hablaremos de los tipos de datos en la próxima sección. Esn este ejemplo se muestra la manera más sencilla de declarar una variable. Si se quiere declarar otra variable, simplemente agregamos una instrucción similar a la anterior. Podemos declarar más de una variable en una sola línea si son del mismo tipo:

```go
var firstName, lastName string
```
Agregar una coma `,`después de un nombre de variable indica que se va a declarar otra variable. En este caso, la instrucción anterior declara dos variables denominadas `firstName` y `lastName` de tipo `string`. Si quiere agregar una tercera variable de tipo `int`, el código tendrá el siguiente aspecto:

```go
var firstName, lasName string
var age int
```
Otra forma de escribir la instrucción es usar paréntesis después de la palabra clave `var`, como si tuviera un bloque dedicado para declarar variables.

```go
var (
    firstName, lastName string
    age int
)
```
## Inicialización de variables
Hasta ahora solo has declarado variables, pero habrá ocasiones en las que necesites tener un valor inicial. En Go, puedes inicializar las variables de varias formas. De este modo, partiendo del ejemplo anterior, podrías inicializar cada variable mediante el siguiente código:

```go
var (
    fistName string = "John"
    lastName string = "Doe"
    age      int    = 32
)
```
Si decides inicializar una variable, no es necesario que especifiques su tipo porque Go lo infiere al inicializar la variable con un valor. Por ejemplo, puedes declarar e inicializar variables de esta manera:

```go
var (
    fistName = "John"
    lastName = "Doe"
    age      = 32
)
```
Go infiere que las variables `firstName` y `lastName` son del tipo string y que la variable `age` es del tipo `int`.

## Diferentes manera de inicializar variables
En Go, puedes declarar e inicializar las variables en una sola linea. Separa cada nombre de variable con una coma y haz lo mismo para cada valor (en el mismo orden), como se muestra a continuación:

```go
var (
    fistName, lastName, age = "John", "Doe", 32
)
```
Hay otra manera de declarar e inicializar variables. Este método es la forma más habitual de hacerlo en Go. El mismo ejemplo que hemos usado anteriormente tendría un aspecto similar a este:

```go
package main

import "fmt"

func main() {
    firstName, lastName := "John", "Doe"
    age := 32
    fmt.Println(firstName, lastName, age)
}
```
>[!Note] Observa la `import "fmt"` instrucción. Usamos la palabra clave `import` para incluir el contenido de un paquete en el ámbito. Vamos a importar el paquete `fmt` para que podamos usar el método `Println` en el código.

Ejecuta el código anterior para confirmar que esta otra manera de declarar e inicializar variables funciona.

Tené en cuenta que debés incluir dos puntos seguidos de un signo igual `:=` y el valor correspondiente justo después de definir el nombre de la variable. Cuando se usan los dos puntos con el signo igual, la *variable que se está declarando tiene que se una nueva*. Si usas dos puntos con un signo igual para una variable que ya se declaró, el programa no compilará. Agrega la edad como constante (la siguiente lección), pero usá los dos puntos con el signo igual `:=` y pruebalo.

Por último, solo puedes usar los dos puntos seguidos de un signo igual dentro de una función. Para declarar variables fuera de una función, se debe usar la palabra clave `var`.

## Declaración de constantes
Habrá ocasiones en las que necesites incluir valores estáticos en el código, conocidos como *constantes*. Go admite el uso de constantes. La palabra clave que se usa para declararlas es `const`.

Por ejemplo, se puede declarar una constante de la siguiente forma:

```go
const HTTPStatusOK = 200
```
Igual que con las variables, el tipo de constante se deduce a partir del valor que se le asigne. En Go, los nombres de las constantes normalmente se escriben con una combinación de mayúsculas y minúsculas o con todas las letras en mayúsculas.

Si necesitas declarar varias constantes en un bloque, puedes hacer de la siguiente manera:

```go
const (
    StatusOK              = 0
    StatusConnectionReset = 1
    StatusOtherError      = 2
)
```
>[!Note] Go tiene un concepto interesante relacionado con las constantes llamado `iota`, el cual no se cubre en este módulo. Tené en cuenta que `iota` es una palabra clave que se usa en Go para simplificar la definición de constantes cuando los valores son secuenciales.

Aunque hay similitudes entre las constantes y las variables, también hay algunas diferencias clave. Por ejemplo, se puede declar constantes sin usarlas y no recibirás un mensaje de error. No se pueden usar los dos puntos con el signo igual para declarar constantes. Si lo haces, go enviará un error.

## Errores de Go por declarar variables y no usarlas
Hay algo fundamental que debes tener en cuenta en Go. Cuando se declara una variable y no se usa, Go genera un error y no una advertencia como en otros lenguajes de programación.

A modo de ilustración, volvamos a uno de los ejemplos anteriores y eliminemos la llamada a `fmt.Println`:

```go
func main() {
    firstName, lastName := "John", "Doe"
    age := 32
}
```
Cuando ejecutes este código en VSC o en playground, tendrás un error en las líneas en las que declaras las variables. Verás los siguientes errores:

```output
./main.go:4:2: # firstName declared but not used
./main.go:4:13: # lastName declared but not used
./main.go:5:2: # age declared but not used
```
Ten en cuenta que debes usar cada variable declarada en Go en algún lugar.