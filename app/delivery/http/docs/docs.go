// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/products": {
            "get": {
                "description": "Find return a list of product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Find return a list of product",
                "parameters": [
                    {
                        "maxLength": 1000,
                        "type": "string",
                        "name": "description",
                        "in": "query"
                    },
                    {
                        "maxLength": 255,
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "name": "price",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "minimum": 1,
                        "type": "integer",
                        "description": "number of items per page",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/presenter.Product"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create add new product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Create add new product",
                "parameters": [
                    {
                        "description": "Body of request",
                        "name": "payloadBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.ProductSaveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.Product"
                        }
                    }
                }
            }
        },
        "/products/:id": {
            "put": {
                "description": "Update modify product information by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "products"
                ],
                "summary": "Update modify product information by id",
                "parameters": [
                    {
                        "description": "Body of request",
                        "name": "payloadBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/payload.ProductSaveRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.Product"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "payload.ProductSaveRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 1000
                },
                "name": {
                    "type": "string",
                    "maxLength": 255
                },
                "price": {
                    "type": "number",
                    "minimum": 0
                }
            }
        },
        "presenter.Product": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "pearshop backend",
	Description:      "Pearshop backend api docs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}