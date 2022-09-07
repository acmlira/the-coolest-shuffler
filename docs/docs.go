// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "acmlira",
            "url": "https://github.com/acmlira/the-coolest-shuffler"
        },
        "license": {
            "name": "MIT",
            "url": "https://www.mit.edu/~amini/LICENSE.md"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/deck": {
            "post": {
                "description": "Create new Deck based in predefined cards",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deck"
                ],
                "summary": "Create new Deck",
                "parameters": [
                    {
                        "description": "Deck properties",
                        "name": "Request",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/model.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Deck"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/deck/new": {
            "get": {
                "description": "Create new Deck based in predefined cards",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deck"
                ],
                "summary": "Create new Deck",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "shuffle cards",
                        "name": "shuffle",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "amount of card sets (before filters)",
                        "name": "amount",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "description": "code filter",
                        "name": "codes",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "description": "value filter",
                        "name": "values",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "string"
                        },
                        "description": "suit filter",
                        "name": "suits",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Deck"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/deck/{deckId}": {
            "get": {
                "description": "Show a created deck",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Deck"
                ],
                "summary": "Show a deck",
                "parameters": [
                    {
                        "type": "string",
                        "description": "code filter",
                        "name": "deckId",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Deck"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Card": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "suit": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "model.CreateRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "codes": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "shuffle": {
                    "type": "boolean"
                },
                "suits": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "values": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.Deck": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "cards": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Card"
                    }
                },
                "id": {
                    "type": "string"
                },
                "remaining": {
                    "type": "integer"
                },
                "shuffled": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "http://localhost:8916",
	BasePath:         "/the-coolest-shuffler/v1",
	Schemes:          []string{},
	Title:            "The Coolest Shuffler",
	Description:      "API to handle the deck and cards to be used in any game like Poker or Blackjack",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
