# firstProject
My first go project

## TODO LIST
    [DONE] Crear peticion rest simple get
    [DONE] Crear peticion rest post dinamica: que lea el body de la peticion o los parametros
    [DONE] generar modelos simples 
    [TODO] generar test de los modelos
    [DONE] crear conexion con bdd mysql
    [DONE] crear un pequeño CRUD con la MYSQL

## Modulos

Para añadir modulos en el proyecto.
- Ir al raiz y lanzar $>go mod init
- Esto creara el fichero go.mod
- go mod tidy -> para cleanear

### Enable Go modules in a project (GoLand IDE)

A new Go modules project already has Go modules enabled.
If you pulled your Go modules project from Github, you need to enable Go modules manually.

- Open settings by pressing Ctrl+Alt+S and navigate to Go | Go modules.
- Select the Enable Go modules integration checkbox.
- Click OK.

## Links

 - Why to use cmd directory convention:
https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091
   
 - Understanding "methods" in Go: https://tour.golang.org/methods/1

### Creando CRUD

Para probar en local lanzamos un contenedor docker con mysql
```bash
$>docker run --name mysql-for-go -e MYSQL_ROOT_PASSWORD=root-go -d -p 33006:3306 mysql:5.7
```
NOTA: Interfaces en GO son implicitas. Una vez cualquier strcutura defina un metodo con
la misma signatura que la interfaz se entendera que está implementando esa interfaz
(JAVA y c# usan interfaces explicitas)
