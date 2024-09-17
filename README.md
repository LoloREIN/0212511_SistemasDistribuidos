# Proyecto de Materia Cpomputo distribuido

## Practicas - LAB
### 1. Creando un Servidor HTTP para procesar JSON

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
- *Marshalear*
- Es el concepto de convertir una estructura de datos a un JSON, XML u otro formato.

- *Desmarshalear*
 - Es el concepto que un JSON, XML u otro convertir a una estructura de datos del sistema.

 *Servidor Simple* es un servidor que escribe y lea nuestro log
  - Consta de dos Endpoints que tendran que procesar lo que necesitamos en tres pasos (Uno leer y uno para escribir)
    - Primero desmarshalear/Unmarshalear el JSON a una estructura
    - Logica como tal del endpoint para procesar el dato
    - Marshalear/Marshael el resultado para dar una respuesta al request
  - Usar/interactuar dos servicios con un net/rpc sencillo.

### 2. Crear un Commit log para la transacciones del servidor
