// Lorenzo Reinoso Fuentes 						0212511

package main

// Zona de mis imports donde
import (
	"encoding/json" // Esta me permite usar el encoding de JSON
	"fmt"           // Me permite usar print y scanf, en si es formatear Strings
	"log"
	"net/http" // Uno de los requisitos es que sea un servidor HTTP
	"sync"
)

type Log struct {
	mu      sync.Mutex
	records []Record
}

type Recod struct {
	Value  []byte `json:"value"`
	Offset uint64 `json:"offset"`
}

/*
## Características del servidor simple

Vamos a crear un servidor que escriba y lea nuestro log simple, pero es más fácil decirlo que hacerlo.

- Para esto necesitaremos dos endpoints que tendrán que procesar lo que necesitamos en tres pasos
    - Primero Desmarshalear/Unmarshal el JSON a una estructura
    - Segundo La lógica como tal del endpoint para procesar el dato
    - Tercero Marshalear/Marshal el resultado para dar una respuesta al request
- Al final solo tendremos dos endpoints en nuestro API uno para leer y otro para escribir

/*

/*
## Mí primer servidor general para guardar commit logs y recibir JSON

- Necesito una forma de mandar mensajes
- Necesito una estructura en la cuál mandar estos mensajes

Vamos a usar JSON por peticiones HTTP ya que:

- JSON sirve para cualquier cosa
- Es fácil de estructurar a nivel código y a nivel de HTTP

Por lo que necesitamos un servidor

- Al inicio tendremos dos endpoints
- Uno que consume y otro que produce
    - Uno que escribe y otro que va a leer pue
- Nuestro servidor va a pasar una petición en JSON y la tendrá que reestructurar en un dato que podemos guardar y viceversa
    - A este proceso de transformar datos duros ya sea de red, puertos y bases de datos a una estructura declara se le conoce como **marshalling**
- Pero necesitamos una forma de estructurar nuestros datos, ya que al momento de empezar nuestro sistema principal tendremos mejor manejo de que esperar en los servicios

Y entra a nuestro escenario: →
*/
