# Cleverti Test
Cleverti group challenge con Go.

## Problema
Bender es fanático de las cervezas, y quiere tener un registro de todas las cervezas que prueba y cómo calcular el precio que necesita para comprar una caja de algún tipo especifico de cervezas. Para esto necesita una API REST con esta información que posteriormente compartir con sus amigos.
## Descripción
Se solicita crear un API REST basándonos en la definición que se encuentra en el [archivo.](https://bitbucket.org/lgaetecl/microservices-test/src/master/openapi.yaml)

## Funcionalidad
GET /beers: Lista todas las cervezas que se encuentran en el sistema

POST /beers/{beerID}: Lista un detalle de una cerveza especifica

GET /beers/{beerID}/boxprice: Entrega el valor que cuesta una caja especifica de cerveza dependiendo de los parámetros ingresados, esto quiere decir que multiplique el precio por la cantidad una vez se homologa la moneda a la que se ingresó por parámetro. Quanty: Cantidad de cervezas a comprar (valor por defecto 6) Currency: Tipo de moneda con que desea pagar, para este caso se recomienda que utilice esta [API.](https://currencylayer.com)

## Requisitos
- La API debe ser RESTful
- Usar Docker y Docker Compose para los servicios
- Pruebas unitarias con Go
- Agregar mas funcionalidades si se desea

## Desarrollo
- [Repositorio en Github.](https://github.com/whitejokeer/clevertitest)







