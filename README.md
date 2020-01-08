# go-rest

## REST Method Descriptions

---

#### 1. Creation of data with POST

With a post request to the container ressource:

https:...

you can create a new resource. The following code example shows the HTTP protocol for the call. The data for the new resource to be created is transferred in the body of the HTTP request. In this example in the JSON format.

POST /URL/ HTTP/1.1

Host: ...

Content-Type: application/json

{
  "column": "value",
  "column": value,
  "column": "value",
  "column": "value"
}

A status code of 201 Created in the answer below informs the client about the successful execution of the request.

HTTP/1.1 201 Created

Content-Type: application/json

Location: url/140

In the location header of the response you will find the URI of the newly created resource. With a GET call to this address you can get a representation of the new product.

---

#### 2. Changing/Updating data with PUT

PUT is usually used to overwrite a resource with the representation in the request.

PUT /url/11 HTTP/1.1

Host: ...

Content-Type: application/json

{
  "name": "Red Grapes",
  "price": 1.79,
  "category_url": "/shop/categories/Fruits",
  "vendor_url": "/shop/vendors/501"
}

If successful, the server replies with the status code 200 OK.

---

#### 3. Changing/Updating data with PATCH

With the HTTP PATCH method, individual properties of a resource can be manipulated in a targeted manner. In the example below, the request contains a new value for the price property.

PATCH /url/70 HTTP/1.1

Host: api.predic8.de

Content-Type: application/json

{
  "price": 1.99
}

The body of the response contains the representation of the changed resource with the new price.

HTTP/1.1 200 OK

Cache-Control: no-cache

Content-Type: application/json; charset=utf-8

{
  "name": "Pears",
  "price": 1.99,
  "photo_url": "/shop/products/70/photo",
  "category_url": "/shop/categories/Fruits",
  "vendor_url": "/shop/vendors/501"
}

---

#### 4. Delete data with DELETE

Resources can be deleted again using the DELETE method.

DELETE /url/142 HTTP/1.1

Host: ...

The server notifies the client of the successful deletion with a status code 200 OK.

HTTP/1.1 200 OK

---