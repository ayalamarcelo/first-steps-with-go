# Ejercicio: Instalación de Go
Antes de mempezas a crear a aplicaciones con go, debés configurar el entorno de desarrollo.

## Instalación de Go en Windows
Para instalar Go en Windows, descargue el instalador de Go.

**Paso 1**: Descargar el instalador de Go [Link](https://go.dev/dl/)

**Paso 2**: Ejecutar el instalador MSI de Go.

Tras descargar localmente el instalador de Go, podés empezar a instalar Go. Para ello, haz click en el archivo .msi y seguí las instrucciones.

De forma predeterminada, el archivov.msi instala Go en C:\Archivos de programa\Go y la ubicacióm de la carpeta C:\Archivos de programa\Go\bin ahora debe formar parte de la variable de entorno del sistema $PATH.

```cmd
go version
```
Debería ver los detalles de la versión de Go instalada en la estación de trabajo.

## Configuración de un área de trabajo de Go
Go difiere de otros lenguajes de programación en el modo en que organiza los archivos de proyecto. En primer lugar, Go funciona bajo el concepto de *áreas de trabajo*. Un área de trabajo es simplemente una ubicación en la que reside el código fuente de la aplicación. En Go, todos los proyectos comparten la misma área de trabajo. Pero Go ha empezado a cambiar este enfoque a partir de la versión 1.11. No tiene que preocuparse de eso todavía, porque abordaremos las áreas de trabajo en el siguiente módulo. Por ahora, el área de trabajo de Go se encuentra en $HOME/go, pero, si es necesario, se uede configurar otra ubicación para todos los proyectos.

Para establecer el área de trabajo en un ubicación diferente, puede usar la variable de entorno $GOPATH. Esta variable de entorno ayuda a evitar problemas futuros al trabajar con proyectos más complejos.

Para configurar el área de trabajo, especifique la ubicación de la carpeta del proyecto de Go en una variable de entorno de Go local.

1. Crea una carpeta de nivel superior para todos los proyectos de Go. Por ejemplo, C:\Projects\Go.
2. Abra un símbolo del sistema de pw1 y ejecute el siguiente cmdlet para establecer la variable de entorno $GOPATH.

>Reemplace <project-folder> por la carpeta de proyecto de nivel superior que creó en el paso anterior.

```shell
[Environment]::SetEnvironmentVariable("GOPATH", "<project-folder>", "user")
```
En este paso se usa PowerShell, por lo que se puede llamar al cmdlet creado previamente para establecer la variable de entorno.

Después de establecer el valor de $GOPATH, cierre el símbolo del sistema de PowerShell.

3. Confirme que la variable de $GOPATH muetsre la ubicación del área de trabajo correcta. En una nueva ventana del símbolo del sistema, ejecute el comando siguiente:

```shell
go env GOPATH
```
La salida miestra la carpeta del proyecto de nivel superior como la ubicación del área de trabajo actual:

```shell
C:\Projects\Go
```
>[!Note] Si la variable $GOPATH no muestra la carpeta del proyecto de nivel superior, asegúrate de que has abierto una nueva ventana del cmd antes de ejecutar el comando. Después de cambiar el valor de una variable de entorno, cerrar y volver a abrir la ventana, o abrir una nueva ventana para que el cambio de valor surta efecto.

Puedes usar una versión corta del comando para ver todas las variables de entorno que usa Go:

```shell
go env
```
## Adición de carpeta de área de trabajo de Go
Cada área de trabajo de Go tiene tres carpetas básicas:

- bin: contiene archivos ejecutables de aplicaciones.
- src: incluye todo el código fuente de la aplicacióm que reside en la estación de trabajo.
- pkg: contiene versiones compiladas de las bibliotecas disponibles. El compilador puede vincularse a estas bibliotecas sin volver a compilarlas.

Por ejemplo, este ejemplo es el aspecto que podrías tener el árbol de estructura de carpetas de la estacióm de trabajo:

bin/
  hello
  coolapp
pkg/
  github.com/gorilla
    mux.a
src/
  github.com/golang/example/
    .git/
  hello/
     hello.go

Ejecute los siguientes comandos para crear tres subcarpetas para el área de trabajo:

```cmd
cd %GOPATH%
mkdir bin
mkdir src
mkdir pkg
```