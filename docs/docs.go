// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/authors": {
            "get": {
                "description": "Возвращает список всех авторов",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authors"
                ],
                "summary": "Получение списка авторов",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Author"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Создаёт нового автора в системе",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authors"
                ],
                "summary": "Создание нового автора",
                "parameters": [
                    {
                        "description": "Автор",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Author"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Author"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/authors/{id}": {
            "put": {
                "description": "Обновляет данные существующего автора",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authors"
                ],
                "summary": "Обновление информации об авторе",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID автора",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные автора",
                        "name": "author",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Author"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Author"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет автора по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authors"
                ],
                "summary": "Удаление автора",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID автора",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/books": {
            "get": {
                "description": "Возвращает список всех книг",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Получение списка книг",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Book"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Создает новую книгу в системе и записывает ее в базу данных",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Создание новой книги",
                "parameters": [
                    {
                        "description": "Книга",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/books/{id}": {
            "put": {
                "description": "Обновляет данные существующей книги",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Обновление информации об книге",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID книги",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные книги",
                        "name": "book",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет книгу по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Удаление книги",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID книги",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/loans": {
            "get": {
                "description": "Возвращает список всех завмов",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "loans"
                ],
                "summary": "Получение списка займов",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Loan"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Создает нового займа в системе",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "loans"
                ],
                "summary": "Создание нового займа",
                "parameters": [
                    {
                        "description": "Займ",
                        "name": "loan",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Loan"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Loan"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/loans/{id}": {
            "put": {
                "description": "Обновляет данные существующего займа",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "loans"
                ],
                "summary": "Обновление информации об займе",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID займа",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные займа",
                        "name": "loans",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Loan"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Loan"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет займ по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "loans"
                ],
                "summary": "Удаление займа",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID займа",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/readers": {
            "get": {
                "description": "Возвращает список всех читателей",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "readers"
                ],
                "summary": "Получение списка читателей",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Reader"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Создает нового читателя в системе",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "readers"
                ],
                "summary": "Создание нового читателя",
                "parameters": [
                    {
                        "description": "Читатель",
                        "name": "reader",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Reader"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Reader"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/api/readers/{id}": {
            "put": {
                "description": "Обновляет данные существующего читателя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "readers"
                ],
                "summary": "Обновление информации об читателе",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID читателя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Данные читателя",
                        "name": "reader",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Reader"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Reader"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Удаляет читателя по ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "readers"
                ],
                "summary": "Удаление читателя",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID читателя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Author": {
            "type": "object",
            "properties": {
                "birth_date": {
                    "description": "BirthDate представляет дату рождения автора\nexample: 1828-09-09",
                    "type": "string"
                },
                "name": {
                    "description": "Name представляет имя автора\nexample: Лев Толстой",
                    "type": "string"
                },
                "nationality": {
                    "description": "Nationality представляет национальность автора\nexample: Русский",
                    "type": "string"
                }
            }
        },
        "models.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "description": "Author представляет автора книги",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Author"
                        }
                    ]
                },
                "author_id": {
                    "description": "AuthorID представляет ID автора книги\nexample: 1",
                    "type": "integer"
                },
                "available": {
                    "description": "Available указывает, доступна ли книга\nexample: true",
                    "type": "boolean"
                },
                "isbn": {
                    "description": "ISBN представляет международный стандартный книжный номер\nexample: 978-3-16-148410-0",
                    "type": "string"
                },
                "published_date": {
                    "description": "PublishedDate представляет дату публикации книги\nexample: 1869-01-01",
                    "type": "string"
                },
                "title": {
                    "description": "Title представляет название книги\nexample: Война и мир",
                    "type": "string"
                }
            }
        },
        "models.Loan": {
            "type": "object",
            "properties": {
                "book": {
                    "description": "Book представляет информацию о книге",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Book"
                        }
                    ]
                },
                "book_id": {
                    "description": "BookID представляет ID книги, которую берут в заем\nexample: 123",
                    "type": "integer"
                },
                "loan_date": {
                    "description": "LoanDate представляет дату выдачи книги\nexample: 2024-04-01",
                    "type": "string"
                },
                "reader": {
                    "description": "Reader представляет информацию о читателе",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Reader"
                        }
                    ]
                },
                "reader_id": {
                    "description": "ReaderID представляет ID читателя, берущего книгу\nexample: 456",
                    "type": "integer"
                },
                "return_date": {
                    "description": "ReturnDate представляет дату возврата книги\nexample: 2024-04-15",
                    "type": "string"
                }
            }
        },
        "models.Reader": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "Email представляет электронную почту читателя\nexample: ivan.ivanov@example.com",
                    "type": "string"
                },
                "membership_date": {
                    "description": "MembershipDate представляет дату вступления читателя в библиотеку\nexample: 2020-01-15",
                    "type": "string"
                },
                "name": {
                    "description": "Name представляет имя читателя\nexample: Иван Иванов",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Library API",
	Description:      "API для управления библиотекой",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}