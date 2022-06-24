# MERCADO FRESCOS
This API Handle MELI Products.
<p>
  <a href="https://github.com/marcoglnd/mercado-fresco-packmain/actions/workflows/test.yml">
    <img src="https://github.com/marcoglnd/mercado-fresco-packmain/actions/workflows/test.yml/badge.svg">
  </a>
</p>
## Version: 1.0

### Terms of service
<https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones>

**Contact information:**
API Support
<https://developers.mercadolibre.com.ar/support>

**License:** [Apache 2.0](http://www.apache.org/licenses/LICENSE-2.0.html)

### `GET: /buyers`
##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [schemes.Buyer](#schemesBuyer)|
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) |

### `POST: /buyers`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| buyer | body | Buyer to create | Yes | [controllers.requestBuyer](#controllersrequestbuyer) |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [schemes.Buyer](#schemesbuyer) |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) |
| 422 | Unprocessable Entity | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) |

### `DELETE: /buyers/{id}`

##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Buyer ID | Yes | integer |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 204 | No Content | [schemes.JSONSuccessResult](#schemesjsonsuccessresult) & object |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `GET: /buyers/{id}`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Buyer ID | Yes | integer |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [schemes.Buyer](#schemesbuyer) |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) |

#### `PATCH: /buyers/{id}`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Buyer ID | Yes | integer |
| buyer | body | Buyer to update | Yes | [controllers.requestBuyer](#controllersrequestbuyer) |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [schemes.Buyer](#schemesbuyer) |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) |

### `GET: /employees`
##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [schemes.JSONSuccessResult](#schemesjsonsuccessresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `POST: /employees`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| employee | body | Employee to create | Yes | [controllers.requestEmployee](#controllersrequestemployee) |
##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [schemes.Employee](#schemesemployee) |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 422 | Unprocessable Entity | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `DELETE: /employees/{id}`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Employee ID | Yes | integer |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 204 | No Content | [schemes.JSONSuccessResult](#schemesjsonsuccessresult) & object |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `GET: /employees/{id}`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Employee ID | Yes | integer |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [schemes.Employee](#schemesemployee) |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `PATCH: /employees/{id}`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Employee ID | Yes | integer |
| employee | body | Employee to update | Yes | [controllers.requestEmployee](#controllersrequestemployee) |
##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [schemes.Employee](#schemesemployee) |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

#### `GET: /products`
##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [schemes.JSONSuccessResult](#schemesjsonsuccessresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `POST: /products`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| product | body | Product to create | Yes | [controllers.requestProducts](#controllersrequestproducts) |
##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [products.Product](#productsproduct) |
| 409 | Conflict | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 422 | Unprocessable Entity | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `DELETE: /products/{id}`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | product ID | Yes | integer |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 204 | No Content | [schemes.JSONSuccessResult](#schemesjsonsuccessresult) & object |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `GET: /products/{id}`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Product ID | Yes | integer |
##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [products.Product](#productsproduct) |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `PATCH: /products/{id}`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Product ID | Yes | integer |
| product | body | Product to update | Yes | [controllers.requestProducts](#controllersrequestproducts) |
##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [products.Product](#productsproduct) |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 422 | Unprocessable Entity | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `GET: /sections`
##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [schemes.JSONSuccessResult](#schemesjsonsuccessresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `POST: /sections`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| section | body | Section to create | Yes | [controllers.requestSection](#controllersrequestsection) |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [schemes.Section](#schemessection) |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 422 | Unprocessable Entity | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
### `DELETE: /sections/{id}`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Section ID | Yes | integer |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 204 | No Content | [schemes.JSONSuccessResult](#schemesjsonsuccessresult) & object |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `GET: /sections/{id}`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Section ID | Yes | integer |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [schemes.Section](#schemessection) |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `PATCH: /sections/{id}`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Section ID | Yes | integer |
| section | body | Section to update | Yes | [controllers.requestSection](#controllersrequestsection) |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [schemes.Section](#schemessection) |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `GET: /sellers`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| token | header | token | Yes | string |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [schemes.JSONSuccessResult](#schemesjsonsuccessresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `POST: /sellers`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| token | header | token | Yes | string |
| Seller | body | seller to create | Yes | [controllers.requestSellers](#controllersrequestsellers) |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [schemes.Seller](#schemesseller) |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 422 | Unprocessable Entity | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `DELETE: /sellers/{id}`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Seller ID | Yes | integer |
| token | header | token | Yes | string |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 204 | No Content | [schemes.JSONSuccessResult](#schemesjsonsuccessresult) & object |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `GET:  /sellers/{id}`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Seller ID | Yes | integer |
| token | header | token | Yes | string |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [schemes.Seller](#schemesseller) |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `PATCH:  /sellers/{id}`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Seller ID | Yes | integer |
| token | header | token | Yes | string |
| seller | body | Seller to update | Yes | [controllers.requestSellers](#controllersrequestsellers) |
##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [schemes.Seller](#schemesseller) |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `GET: /warehouses`
##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [schemes.JSONSuccessResult](#schemesjsonsuccessresult) & object |
| 422 | Unprocessable Entity | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### `POST: /warehouses`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| warehouse | body | Warehouse to create | Yes | [warehouses.CreateWarehouseInput](#warehousescreatewarehouseinput) |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [warehouses.Warehouse](#warehouseswarehouse) |
| 422 | Unprocessable Entity | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
### `DELETE: /warehouses/{id}`
##### `Parameters`

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | warehouse ID | Yes | integer |

##### `Responses`

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 204 | No Content | [schemes.JSONSuccessResult](#schemesjsonsuccessresult) & object |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

#### GET
##### Summary

Warehouse by id

##### Description

get warehouse by it's id

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Warehouse ID | Yes | integer |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [warehouses.Warehouse](#warehouseswarehouse) |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

#### PATCH
##### Summary

Update warehouse

##### Description

Update existing warehouse in list checking for duplicate warehouses code

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ---- |
| id | path | Warehouse ID | Yes | integer |
| warehouse | body | Warehouse to update | Yes | [warehouses.UpdateWarehouseInput](#warehousesupdatewarehouseinput) |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [warehouses.Warehouse](#warehouseswarehouse) |
| 400 | Bad Request | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 404 | Not Found | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |
| 422 | Unprocessable Entity | [schemes.JSONBadReqResult](#schemesjsonbadreqresult) & object |

### Models

#### controllers.requestBuyer

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| card_number_id | string |  | No |
| first_name | string |  | No |
| last_name | string |  | No |

#### controllers.requestEmployee

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| card_number_id | string |  | No |
| first_name | string |  | No |
| last_name | string |  | No |
| warehouse_id | integer |  | No |

#### controllers.requestProducts

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| description | string |  | Yes |
| expiration_rate | integer |  | Yes |
| freezing_rate | integer |  | Yes |
| height | number |  | Yes |
| length | number |  | Yes |
| netweight | number |  | Yes |
| product_code | string |  | Yes |
| product_type_id | integer |  | Yes |
| recommended_freezing_temperature | number |  | Yes |
| seller_id | integer |  | Yes |
| width | number |  | Yes |

#### controllers.requestSection

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| current_capacity | integer |  | Yes |
| current_temperature | integer |  | Yes |
| maximum_capacity | integer |  | Yes |
| minimum_capacity | integer |  | Yes |
| minimum_temperature | integer |  | Yes |
| product_type_id | integer |  | Yes |
| section_number | integer |  | Yes |
| warehouse_id | integer |  | Yes |

#### controllers.requestSellers

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| address | string |  | No |
| cid | integer |  | No |
| company_name | string |  | No |
| telephone | string |  | No |

#### products.Product

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| description | string |  | No |
| expiration_rate | integer |  | No |
| freezing_rate | integer |  | No |
| height | number |  | No |
| id | integer |  | No |
| length | number |  | No |
| netweight | number |  | No |
| product_code | string |  | No |
| product_type_id | integer |  | No |
| recommended_freezing_temperature | number |  | No |
| seller_id | integer |  | No |
| width | number |  | No |

#### schemes.Buyer

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| card_number_id | string |  | No |
| first_name | string |  | No |
| id | integer |  | No |
| last_name | string |  | No |

#### schemes.Employee

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| card_number_id | string |  | No |
| first_name | string |  | No |
| id | integer |  | No |
| last_name | string |  | No |
| warehouse_id | integer |  | No |

#### schemes.JSONBadReqResult

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| error |  |  | No |

#### schemes.JSONSuccessResult

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| data |  |  | No |

#### schemes.Section

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| current_capacity | integer |  | No |
| current_temperature | integer |  | No |
| id | integer |  | No |
| maximum_capacity | integer |  | No |
| minimum_capacity | integer |  | No |
| minimum_temperature | integer |  | No |
| product_type_id | integer |  | No |
| section_number | integer |  | No |
| warehouse_id | integer |  | No |

#### schemes.Seller

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| address | string |  | No |
| cid | integer |  | No |
| company_name | string |  | No |
| id | integer |  | No |
| telephone | string |  | No |

#### warehouses.CreateWarehouseInput

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| address | string |  | Yes |
| minimum_capacity | integer |  | Yes |
| minimum_temperature | integer |  | Yes |
| telephone | string |  | Yes |
| warehouse_code | string |  | Yes |

#### warehouses.UpdateWarehouseInput

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| address | string |  | No |
| minimum_capacity | integer |  | No |
| minimum_temperature | integer |  | No |
| telephone | string |  | No |
| warehouse_code | string |  | Yes |

#### warehouses.Warehouse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| address | string |  | Yes |
| id | integer |  | No |
| minimum_capacity | integer |  | Yes |
| minimum_temperature | integer |  | Yes |
| telephone | string |  | Yes |
| warehouse_code | string |  | Yes |
