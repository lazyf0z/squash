// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Squash debugger api.",
    "title": "Squash Server",
    "contact": {
      "email": "apiteam@solo.io"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "2.0.0"
  },
  "basePath": "/api/v2",
  "paths": {
    "/debugattachment": {
      "get": {
        "description": "Return all debug attachment",
        "tags": [
          "debugattachment"
        ],
        "summary": "Return all debug attachment",
        "operationId": "getDebugAttachments",
        "security": null,
        "parameters": [
          {
            "type": "boolean",
            "description": "wait until there's something to return instead of returning an empty list",
            "name": "wait",
            "in": "query"
          },
          {
            "type": "string",
            "description": "filter by node that the debugattachment is assigned to",
            "name": "node",
            "in": "query"
          },
          {
            "type": "string",
            "description": "filter by the state of debugattachment is assigned to",
            "name": "state",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "description": "Only get a subset of debugattachments",
            "name": "names",
            "in": "query"
          },
          {
            "type": "number",
            "name": "X-Timeout",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/DebugAttachment"
              }
            }
          },
          "408": {
            "description": "Request timed out"
          },
          "422": {
            "description": "Validation exception"
          }
        }
      },
      "post": {
        "description": "A debug attachment instructs squash to attach to a container. Debug attachment is made of\n  - image: The container image we are debugging. this is used for extra validation, as placing breakpoints on the wrong binary can lead to unexpected results. if not provided huerisrtics will be used to identify it.\n  - debugger: Type of debugger to use. \"dlv\" and \"gdb\" are supported now.\n  - match_request: Whether to match this attachment to a debug request. This is used in automated use-cases to guarantee that the attachment will be noticed.\n",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "debugattachment"
        ],
        "summary": "Request squash to attach to a running container.",
        "operationId": "addDebugAttachment",
        "security": null,
        "parameters": [
          {
            "description": "DebugAttachment object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/DebugAttachment"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Debug attachment created",
            "schema": {
              "$ref": "#/definitions/DebugAttachment"
            }
          },
          "400": {
            "description": "Bad request"
          },
          "404": {
            "description": "Not found"
          },
          "422": {
            "description": "Invalid input"
          },
          "503": {
            "description": "Service Unavailable"
          }
        }
      }
    },
    "/debugattachment/{debugAttachmentId}": {
      "get": {
        "description": "Return a debug attachment",
        "tags": [
          "debugattachment"
        ],
        "summary": "Return a debug attachment",
        "operationId": "getDebugAttachment",
        "security": null,
        "parameters": [
          {
            "type": "string",
            "description": "ID of config to return",
            "name": "debugAttachmentId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/DebugAttachment"
            }
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Debug config not found"
          },
          "422": {
            "description": "Validation exception"
          }
        }
      },
      "delete": {
        "description": "Delete a debug attachment. be careful not to delete on during attaching phase.",
        "tags": [
          "debugattachment"
        ],
        "summary": "Delete a debug attachment",
        "operationId": "deleteDebugAttachment",
        "security": null,
        "parameters": [
          {
            "type": "string",
            "description": "ID of config to return",
            "name": "debugAttachmentId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Debug config not found"
          },
          "422": {
            "description": "Validation exception"
          }
        }
      },
      "patch": {
        "description": "Modify an existing attachment.\n",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "debugattachment"
        ],
        "summary": "Modify an existing attachment.",
        "operationId": "patchDebugAttachment",
        "security": null,
        "parameters": [
          {
            "type": "string",
            "description": "ID of config to return",
            "name": "debugAttachmentId",
            "in": "path",
            "required": true
          },
          {
            "description": "DebugAttachment object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/DebugAttachment"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Debug attachment modified",
            "schema": {
              "$ref": "#/definitions/DebugAttachment"
            }
          },
          "400": {
            "description": "Bad request"
          },
          "404": {
            "description": "Not found"
          },
          "409": {
            "description": "Conflict"
          },
          "422": {
            "description": "Invalid input"
          },
          "503": {
            "description": "Service Unavailable"
          }
        }
      }
    },
    "/debugrequest": {
      "get": {
        "description": "Return all debug requests",
        "tags": [
          "debugrequest"
        ],
        "summary": "Return all debug request",
        "operationId": "getDebugRequests",
        "security": null,
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/DebugRequest"
              }
            }
          },
          "422": {
            "description": "Validation exception"
          }
        }
      },
      "post": {
        "description": "Return a debug attachment",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "debugrequest"
        ],
        "summary": "Return a debug attachment",
        "operationId": "createDebugRequest",
        "security": null,
        "parameters": [
          {
            "description": "DebugRequest object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/DebugRequest"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/DebugRequest"
            }
          }
        }
      }
    },
    "/debugrequest/{debugRequestId}": {
      "get": {
        "description": "Get a debug request",
        "tags": [
          "debugrequest"
        ],
        "summary": "Get a debug request",
        "operationId": "getDebugRequest",
        "security": null,
        "parameters": [
          {
            "type": "string",
            "description": "ID of config to return",
            "name": "debugRequestId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/DebugRequest"
            }
          },
          "404": {
            "description": "Not found"
          }
        }
      },
      "delete": {
        "description": "Delete a debug request.",
        "tags": [
          "debugrequest"
        ],
        "summary": "Delete a debug request",
        "operationId": "deleteDebugRequest",
        "security": null,
        "parameters": [
          {
            "type": "string",
            "description": "ID of debug request",
            "name": "debugRequestId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid ID supplied"
          },
          "404": {
            "description": "Debug request not found"
          },
          "422": {
            "description": "Validation exception"
          }
        }
      }
    }
  },
  "definitions": {
    "Attachment": {
      "type": "object"
    },
    "DebugAttachment": {
      "type": "object",
      "required": [
        "spec"
      ],
      "properties": {
        "metadata": {
          "$ref": "#/definitions/ObjectMeta"
        },
        "spec": {
          "$ref": "#/definitions/DebugAttachmentSpec"
        },
        "status": {
          "$ref": "#/definitions/DebugAttachmentStatus"
        }
      }
    },
    "DebugAttachmentSpec": {
      "type": "object",
      "required": [
        "attachment",
        "debugger"
      ],
      "properties": {
        "attachment": {
          "$ref": "#/definitions/Attachment"
        },
        "debugger": {
          "type": "string"
        },
        "image": {
          "type": "string"
        },
        "match_request": {
          "type": "boolean"
        },
        "node": {
          "type": "string"
        }
      }
    },
    "DebugAttachmentStatus": {
      "type": "object",
      "properties": {
        "debug_server_address": {
          "type": "string"
        },
        "state": {
          "type": "string",
          "enum": [
            "none",
            "attaching",
            "attached",
            "error"
          ]
        }
      }
    },
    "DebugRequest": {
      "type": "object",
      "required": [
        "spec"
      ],
      "properties": {
        "metadata": {
          "$ref": "#/definitions/ObjectMeta"
        },
        "spec": {
          "$ref": "#/definitions/DebugRequestSpec"
        },
        "status": {
          "$ref": "#/definitions/DebugRequestStatus"
        }
      }
    },
    "DebugRequestSpec": {
      "type": "object",
      "required": [
        "image",
        "debugger"
      ],
      "properties": {
        "debugger": {
          "type": "string"
        },
        "image": {
          "type": "string"
        }
      }
    },
    "DebugRequestStatus": {
      "type": "object",
      "properties": {
        "debug_attachment_ref": {
          "type": "string"
        }
      }
    },
    "ObjectMeta": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Everything about your debugconfig",
      "name": "debugattachment",
      "externalDocs": {
        "description": "Find out more",
        "url": "http://swagger.io"
      }
    },
    {
      "description": "Everything about your debug sessions",
      "name": "debugsessions"
    }
  ],
  "externalDocs": {
    "description": "Find out more about Squash",
    "url": "https://github.com/solo-io/squash"
  }
}`))
}
