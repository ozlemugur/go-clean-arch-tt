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
        "/automatic-message-sender": {
            "post": {
                "description": "Starts or stops the automatic message sending scheduler based on the provided action.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "scheduler"
                ],
                "summary": "Control automatic message sender scheduler",
                "operationId": "control-message-scheduler",
                "parameters": [
                    {
                        "description": "Scheduler action (start or stop)",
                        "name": "action",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.AutoMessageSender"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Operation completed successfully",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/messages": {
            "post": {
                "description": "Add a new message to the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "messages"
                ],
                "summary": "Insert a message",
                "operationId": "insert-message",
                "parameters": [
                    {
                        "description": "Message to insert",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Message"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        },
        "/messages/sent": {
            "get": {
                "description": "Retrieve all messages with status \"sent\" from the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "messages"
                ],
                "summary": "Retrieve sent messages",
                "operationId": "get-sent-messages",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/v1.messageResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.AutoMessageSender": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                }
            }
        },
        "entity.Message": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "recipient_phone": {
                    "type": "string"
                }
            }
        },
        "v1.messageResponse": {
            "type": "object",
            "properties": {
                "Messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entity.Message"
                    }
                },
                "count": {
                    "description": "Number of messages in the response",
                    "type": "integer"
                },
                "error": {
                    "description": "Error message, optional field",
                    "type": "string"
                },
                "status": {
                    "description": "Status of the response (e.g., \"success\", \"empty\")",
                    "type": "string"
                }
            }
        },
        "v1.response": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "message"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Automatic Message Sender API",
	Description:      "the system send 2 messages in  every 2 minutes",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
