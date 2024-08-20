# Proyecto de Materia Cpomputo distribuido





## Steps
### 1. Creando un servidor HTTP para procesar JSON

Conceptos importantes:

- *Log* : es una estructura de datos a la que solo agregamos cosas, por lo que también se lo conoce como append-only.

- *commit log*
  - Lo que tenemos que hacer es agregar una lice de record del log
  - Se usa el índice para revisar la tabla y conseguir el regristo que nos interesa

- *Estructura del log*
  - ``` 
  type Log struct{
    mu sync.Mutex
    records [] Record    
  } 
  ```
- *RECORD* 
  - ```
  type Record struct {
    Value []byte ´json: "value"
    Offset uint64 ´json:"offset"´
  }
  ```
