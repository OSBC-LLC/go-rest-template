// Package docs GENERATED BY SWAG; DO NOT EDIT
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
            "name": "SuperNova",
            "url": "https://gus.lightning.force.com/lightning/r/ADM_Scrum_Team__c/a00EE00000PCtdSYAT/view"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/heartbeat": {
            "get": {
                "description": "Shows if database is online, dyno metadata, and health status",
                "produces": [
                    "application/json"
                ],
                "summary": "Get health status of service",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Heartbeat"
                        }
                    }
                }
            }
        },
        "/api/todo": {
            "get": {
                "description": "Show the full list of todos.",
                "produces": [
                    "application/json"
                ],
                "summary": "Get list of todos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Todo"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Heartbeat": {
            "type": "object",
            "properties": {
                "appName": {
                    "type": "string"
                },
                "databaseOnline": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                },
                "releaseCreatedAt": {
                    "type": "string"
                },
                "releaseVersion": {
                    "type": "string"
                },
                "requestId": {
                    "type": "string"
                },
                "slugCommit": {
                    "type": "string"
                }
            }
        },
        "models.Todo": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "requestId": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "herokuapp.com",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Orch-Rest-Template",
	Description:      "Micro-service written in Golang that reads data from a PostgreSQL database.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
