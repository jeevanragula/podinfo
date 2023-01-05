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
            "name": "Source Code",
            "url": "https://github.com/stefanprodan/podinfo"
        },
        "license": {
            "name": "MIT License",
            "url": "https://github.com/stefanprodan/podinfo/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "renders podinfo UI",
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Index",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/echo": {
            "post": {
                "description": "forwards the call to the backend service and echos the posted content",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Echo",
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/api.MapResponse"
                        }
                    }
                }
            }
        },
        "/api/info": {
            "get": {
                "description": "returns the runtime information",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Runtime information",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.RuntimeResponse"
                        }
                    }
                }
            }
        },
        "/cache/{key}": {
            "get": {
                "description": "returns the content from cache if key exists",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Get payload from cache",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Key to load from cache",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "writes the posted content in cache",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Save payload in cache",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Key to save to",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted"
                    }
                }
            },
            "delete": {
                "description": "deletes the key and its value from cache",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Delete payload from cache",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Key to delete",
                        "name": "key",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted"
                    }
                }
            }
        },
        "/chunked/{seconds}": {
            "get": {
                "description": "uses transfer-encoding type chunked to give a partial response and then waits for the specified period",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Chunked transfer encoding",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "seconds to wait for",
                        "name": "seconds",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MapResponse"
                        }
                    }
                }
            }
        },
        "/delay/{seconds}": {
            "get": {
                "description": "waits for the specified period",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Delay",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "seconds to wait for",
                        "name": "seconds",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MapResponse"
                        }
                    }
                }
            }
        },
        "/env": {
            "get": {
                "description": "returns the environment variables as a JSON array",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Environment",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/headers": {
            "get": {
                "description": "returns a JSON array with the request HTTP headers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Headers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/healthz": {
            "get": {
                "description": "used by Kubernetes liveness probe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Liveness check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/metrics": {
            "get": {
                "description": "returns HTTP requests duration and Go runtime metrics",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Prometheus metrics",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/panic": {
            "get": {
                "description": "crashes the process with exit code 255",
                "tags": [
                    "HTTP API"
                ],
                "summary": "Panic",
                "responses": {}
            }
        },
        "/readyz": {
            "get": {
                "description": "used by Kubernetes readiness probe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Readiness check",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/readyz/disable": {
            "post": {
                "description": "signals the Kubernetes LB to stop sending requests to this instance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Disable ready state",
                "responses": {
                    "202": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/readyz/enable": {
            "post": {
                "description": "signals the Kubernetes LB that this instance is ready to receive traffic",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Kubernetes"
                ],
                "summary": "Enable ready state",
                "responses": {
                    "202": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/status/{code}": {
            "get": {
                "description": "sets the response status code to the specified code",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Status code",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "status code to return",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MapResponse"
                        }
                    }
                }
            }
        },
        "/store": {
            "post": {
                "description": "writes the posted content to disk at /data/hash and returns the SHA1 hash of the content",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Upload file",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MapResponse"
                        }
                    }
                }
            }
        },
        "/store/{hash}": {
            "get": {
                "description": "returns the content of the file /data/hash if exists",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Download file",
                "parameters": [
                    {
                        "type": "string",
                        "description": "hash value",
                        "name": "hash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "file",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/token": {
            "post": {
                "description": "issues a JWT token valid for one minute",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Generate JWT token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.TokenResponse"
                        }
                    }
                }
            }
        },
        "/token/validate": {
            "post": {
                "description": "validates the JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Validate JWT token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.TokenValidationResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "description": "returns podinfo version and git commit hash",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Version",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.MapResponse"
                        }
                    }
                }
            }
        },
        "/ws/echo": {
            "post": {
                "description": "echos content via websockets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "HTTP API"
                ],
                "summary": "Echo over websockets",
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/api.MapResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.MapResponse": {
            "type": "object",
            "additionalProperties": {
                "type": "string"
            }
        },
        "api.RuntimeResponse": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "goarch": {
                    "type": "string"
                },
                "goos": {
                    "type": "string"
                },
                "hostname": {
                    "type": "string"
                },
                "logo": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "num_cpu": {
                    "type": "string"
                },
                "num_goroutine": {
                    "type": "string"
                },
                "revision": {
                    "type": "string"
                },
                "runtime": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "api.TokenResponse": {
            "type": "object",
            "properties": {
                "expires_at": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "api.TokenValidationResponse": {
            "type": "object",
            "properties": {
                "expires_at": {
                    "type": "string"
                },
                "token_name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2.0",
	Host:             "localhost:9898",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "Podinfo API",
	Description:      "Go microservice template for Kubernetes.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
