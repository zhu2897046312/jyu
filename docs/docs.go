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
        "/User/DeleteUserHandler": {
            "get": {
                "tags": [
                    "删除模块"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "账户",
                        "name": "account",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\", \"masssge\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/User/createUserHandler": {
            "get": {
                "tags": [
                    "用户模块"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "账户",
                        "name": "account",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "确认密码",
                        "name": "repassword",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\", \"masssge\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/User/getUserList": {
            "get": {
                "tags": [
                    "用户模块"
                ],
                "responses": {
                    "200": {
                        "description": "code\", \"masssge\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/index": {
            "get": {
                "tags": [
                    "首页"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
