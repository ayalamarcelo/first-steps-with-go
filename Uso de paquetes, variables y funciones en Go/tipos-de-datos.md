# Tipos de datos básicos
Go es un lenguaje *fuertemente tipado*. Todas las variables declaradas se enlazan a un tipo de datos específico y solo se aceptarán los valores que coincidan con este tipo.

En Go hay cuatro categorías de tipos de datos:
- Tipos básicos: números, cadenas y booleanos.
- Tipos agregados: matrices y estructuras.
- Tipos de referencia: punteros, segmentos, mapas, funciones y canales.
- Tipos de interfaz: interfaz.

En este módulo, solo se tratarán los tipos básicos. Para empezar, vamos a explorar los tipos de datos numéricos.

## Números enteros
En términos generales, la palabra clave para definir un tipo entero es `int`. Pero Go también proporciona los tipos `int8`, `int16`, `int32`, `int64`, que son enteros con un tamaño de 8, 16, 32 o 64 bits, respectivamente. Cuando se usa un sistema operativo de 32 bits, si solo se usa `int`, el tamaño suele ser de 32 bits. En sistemas de 64 bits, el tamaño de `int` suele ser de 64 bits. Pero este comportamiento puede diferir de un equipo a otro. Puede usar `uint`. Use este tipo solo si necesita representar un valor como un número sin signo por un motivo determinado. Go también proporciona los tipos `uint8`, `uint16`, `uint32` y `uint64`.

A continuación se muestra un ejemplo de cómo usar los distintos tipos enteros de Go:

```go
var integer8 int8 = 127
var integer16 int16 = 32767
var integer32 int32 = 2147483647
var integer64 int64 = 9223372036854775807
fmt.Println(integer8, integer16, integer32, integer64)
```
Normalmente usarás `int`, pero debes conocer los otros tipos enteros porque, en Go, `int` no es lo mismo que `int32`, aunque el tamaño natural del entero sea 32bits. En otras palabras, deberás convertir explícitamente un tipo de datos en otro cuandos se requieras una conversión. Además, si intentas realizar una operación matemática entre distintos tipos, obtendrás un error. Por ejemplo, supongamos este código:

```go
var integer16 int16 = 127
var integer32 int32 = 32767
fmt.Println(integer16 + integer32)
```
Al ejecutar el programa, recibiras el siguiente error:

`invalid operation: integer16 + integer32 (mismatched types int16 and int32)`

Como puedes ver, al convertir un valor de un tipo a otro en Go, debes especificar explícitamente el nuevo tipo. Hablaremos de cómo convertir tipos correctamente al final del módulo.

A medida que adquieras más conocimientos de Go, es posible que te encuentre con la palabra clave `runes`. `rune` es simplemente un alias para el tipo de datos `int32`. Se usa para representar un carácter Unicode (o un punto de código Unicode). Por ejemplo, supongamos que tiene el siguiente código:

```go
rune := 'G'
fmt.Println(rune)
```
Es posible que espere que el programa imprima `G` en el símbolo del sistema al ejecutar el fragmento de código anterior, pero en su lugar ve el número `71`, que representa el carácter Unicode para `G`.

### Desafío
Establezca otra variable de tipo `int` y use el valor de las variables `integer32` o `integer64` para confirmar el tamaño natura de la variable en el sistema. Si su sistema es de 32 bits y usa un valor superior a 2 147 483 647, obtendrá un error de desbordamiento similar al siguiente: `constant 9223372036854775807 overflows int`.

```go
package main

import "fmt"

func main() {
    var integer32 int = 2147473648
    fmt.Println(integer32)
}
```

### Desafío 2
Declara una variable sin signo, como `uint`, e inicializala con un valor negatico, por ejemplo `-10`. Al intentar ejecutar el programa, debería obtener un error similar a este: `constant -10 overflows uint`

```go
package main

import "fmt"

func main() {
    var integer uint = -10
    fmt.Println(integer)
}
```

## Números de punto flotante
Go proporciona tipos de datos para dos tamaños de números de punto flotante: `float32` y `float64`. Puedes usar estos tipos cuando necesites almacenar números de gran tamaño que no quepan en ninguno de los tupos entero mencionados anteriormente. La diferencia entre estos dos tipos es el tamaño máximo de bits que pueden contener. Consulta las líneas siguientes para ver cómo se usan estos dos tipos:

```go
var float32 float32 = 2147483647
var float64 float64 = 9223372036854775807
fmt.Println(float32, float64)
```
Para encontrar los límites de estos dos tipos, puedes usar las constantes `math.MaxFloat32` y `math.MaxFloat64`, las cuales están disponibles, en el paquete `math`. Usa el código siguiente para imprimir los valores de punto flotante máximos en el cmd:

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.MaxFloat32, mathMaxFloat64)
}
```
Los tipos de punto flotante también resultan útiles cuando se necesita usar números decimales. Por ejemplo, puedes usarlos para escribir código como el siguiente:

```go
const e = 2.71828
const Avogadro = 6.02214129e23
const Planck = 6.62606957e-34
```
Observa que, con el código anterior, Go deduce los tipos de datos a partir de los valores usados.

## Valores Booleanos
Un tipo booleano solamente tiene dos valores posibles: `true` y `false`. Los tipos booleanos se declaran mediante la palabra clave `bool`. Go es diferente a otros lenguajes de programación. En Go, no puedes convertir implícitamente un tipo booleano a 0 y 1. Se debe hacer explícitamente.

Por lo tanto, puedes declarar una variable booleana como la siguiente:

```go
var featureFlag bool = true
```
Usaremos tipos de booleanos en el próximo módulo cuando hablemos sobre las intrucciones de flujo de control en Go. También los usaremos en módulos posteriores.

## Cadenas
Por último, echemos un vistazo al tipo de datos más común en cualquier lenguaje de programación: las cadenas. En Go, la palabra clave `string` se usa para representar un tipo de datos de cadena. Para inicializar una variable de cadena, se debe definir su valor entre comillas dobles `"`. Las comillas simples `''` se usan para los caracteres individuales (y para los "runes", como hemos visto en una sección anterior).

Por ejemplo, en el código siguiente se muestran dos formas de declarar e inicializar una variable de cadena:

```go
var firstName string = "John"
lastName := "Doe"
fmt.Println(firstName, lastName)
```
A veces necesitarás usar caracteres de escape. Para incluirlos en Go, usa una barra inversa `\` antes del carácter. A continuación se muestran los ejemplos más comunes del uso de caracteres de escape:

- \n para lineas nuevas.
- \r para retornos de carro
- \t para pestañas
- \' para comillas simples
- \" para comillas dobles
- \ \ para barras diagonales inversas.

### Ejecuta el siguiente fragmento de código para probar los caracteres de escape

```go
fullName := "John Doe \t(alias \ "Foo\")\n"
```
Deberías ver la sigueinte salida (incluida la nueva línea):

```
John Doe     (alias "Foo")
```
## Valores predeterminados
Pro el momento, caso todas la veces que se ha declarado una variable en este módulo, se ha inicializado con un valor. Pero en Go, a diferencia de otros leguajes de programación, todos los tipos de datos tienen un valor predeterminado cuando no se inicializa una variable. Esta característica es útil porque así no es necesario comprobar si se ha inicializado una variable antes de usarla.

A continuación se presenta una lista con algunos valores predeterminados para los tipos que hemos explorado hasta ahora:

- 0 para los tipos int (y todos sus subtipos, como int64).
- +0.0000000e+000 para los float32 y float64.
- false para los tipos bool.
- Un valor vacío para los tipos string.

Ejecuta el siguiente fragmento de código para confirmar los valores predeterminados enumerados anteriorment:

```output
var defaultInt int
var defaultFloat32 float32
var defaultFloat64 float64
var defaultBool bool
var defaultString string
fmt.Println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)
```
Podría usar este código, por ejemplo, para determinar el valor predeterminado de un tipo de datos que no hemos analizado aquí.

## Conversiones de tipos
En una sección anterior, hemos confirmado que la conversión implícita no funciona en Go. En Go, la conversión debe llevarse a cabo explícitamente. Este lenguaje de programación ofrece algunas maneras nativas de convertir un tipo de datos en otro diferente. Por ejemplo, una de ellas consiste en usar la función integrada para cada tipos de la siguiente forma:

```go
var integer16 int16 = 127
var integer32 int32 = 32767
fmt.Println(int32(integer16) + integer32)
```
Otro enfoque para la conversión en Go es usar el paquete strconv. Por ejemplo, para convertir un tipo string en un tipo int y viceversa, podrías usar este código:

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    i, _ := strconv.Atoi("-42")
    s := strconv.Itoa(-42)
    fmt.Println(i, s)
}
```
Ejecute el código anterior y confirme que funciona e imprime -42 dos veces.

Observe que hay un carácter de subrayado (_) que se usa como el nombre de una variable en el código anterior. En Go, el objeto _ significa que no vamos a usar el valor de esa variable y que queremos omitirlo. De lo contrario, el programa no se compilará porque necesitamos usar todas las variables declaradas. Volveremos a abordar este tema en los próximos módulos y aprenderá qué representa normalmente el carácter _.