# My first Go project

In this project I going to practice with Go language. 
I will try to implement all interesting things that can be done with this technology in order to improve new programming skills.

#### Table of Content

 - [TODO LIST](https://github.com/srpepperoni/firstProject#todo-list)
    - [TODO LIST: Implement simple API Rest (Get, POST, Delete...)](https://github.com/srpepperoni/firstProject#todo-list-implement-simple-crud-methods)
    - [TODO LIST: Implement simple CRUD methods](https://github.com/srpepperoni/firstProject#todo-list-implement-simple-api-rest-get-post-delete)
 - [Document support tips](https://github.com/srpepperoni/firstProject#document-support-tips)
    - [Document support tips: Modulos](https://github.com/srpepperoni/firstProject#document-support-tips-modulos)
        - [Enable Go modules in a project](https://github.com/srpepperoni/firstProject#enable-go-modules-in-a-project-goland-ide)
 - [Extra Links](https://github.com/srpepperoni/firstProject#extra-links)   

## TODO LIST
  - [x] Implement simple CRUD methods
  - [x] Implement simple API Rest (Get, POST, Delete...)
  - [x] Create unit test for hnadlers
  - [ ] Create unit test for models
  - [ ] Create benchmark test

### TODO LIST: Implement simple CRUD methods

Para probar en local lanzamos un contenedor docker con mysql
```bash
$>docker run --name mysql-for-go -e MYSQL_ROOT_PASSWORD=root-go -d -p 33006:3306 mysql:5.7
```
NOTA: Interfaces en GO son implicitas. Una vez cualquier strcutura defina un metodo con
la misma signatura que la interfaz se entendera que está implementando esa interfaz
(JAVA y c# usan interfaces explicitas)

### TODO LIST: Implement simple API Rest (Get, POST, Delete...)

TBD

## Document support tips

### Document support tips: Modulos

```bash
$>go mod init # para inicializar proyecto con modulos en go
$>go mod vendor # para crear carpeta vendor en el proyecto
$>go mod tidy # Para organizar las dependencias no usadas y las importadas
```

#### Enable Go modules in a project (GoLand IDE)

A new Go modules project already has Go modules enabled.
If you pulled your Go modules project from Github, you need to enable Go modules manually.

- Open settings by pressing Ctrl+Alt+S and navigate to Go | Go modules.
- Select the Enable Go modules integration checkbox.
- Click OK.

## Extra Links

- [Use of cmd directory](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091)

- [Understanding "methods" in Go](https://tour.golang.org/methods/1)
