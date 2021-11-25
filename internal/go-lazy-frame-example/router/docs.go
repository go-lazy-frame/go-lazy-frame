
package router

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var SwaggerDoc = `{
    "swagger": "2.0",
    "info": {
        "description": "Go-Lazy_frame-Example",
        "title": "Go-Lazy_frame-Example",
        "contact": {
            "name": "qianlong",
            "email": "642321251@qq.com"
        },
        "version": "1.0"
    },
    "host": "localhost:8890",
    "basePath": "/api/v1",
    "paths": {
        "/login": {
            "post": {
                "description": "登陆",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登陆注销"
                ],
                "summary": "登陆",
                "operationId": "AuthLoginUsingPOST",
                "parameters": [
                    {
                        "description": "创建",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.LoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "description": "注销",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登陆注销"
                ],
                "summary": "注销",
                "operationId": "AuthLogoutUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_log/create": {
            "post": {
                "description": "操作日志创建",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "操作日志"
                ],
                "summary": "操作日志创建",
                "operationId": "RbacLogCreateUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "创建",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacLogCreateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_log/find_by_id": {
            "get": {
                "description": "操作日志查询ById",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "操作日志"
                ],
                "summary": "操作日志ById",
                "operationId": "RbacLogFindByIdUsingGET",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "查询条件id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_log/query": {
            "post": {
                "description": "操作日志查询记录\u003cbr\u003e请求示例：\u003cbr\u003e{}\u003cbr\u003e其他查询条件，可根据以下条件字段酌情添加查询条件\u003cbr\u003e注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "操作日志"
                ],
                "summary": "操作日志查询",
                "operationId": "RbacLogQueryUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "条件查询",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacLogQueryDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_log/query_page": {
            "post": {
                "description": "操作日志分页查询\u003cbr\u003e请求示例：\u003cbr\u003e{\u003cbr\u003e\u0026nbsp;\u0026nbsp;\"start\": 0,\u003cbr\u003e\u0026nbsp;\u0026nbsp;\"limit\": 20\u003cbr\u003e}\u003cbr\u003e其他查询条件，可根据以下条件字段酌情添加查询条件\u003cbr\u003e注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "操作日志"
                ],
                "summary": "操作日志分页查询",
                "operationId": "RbacLogQueryPageUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "条件查询",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacLogPageDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_log/update": {
            "post": {
                "description": "操作日志更新，根据 id 更新",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "操作日志"
                ],
                "summary": "操作日志更新",
                "operationId": "RbacLogUpdateUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "更新",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacLogUpdateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_permissions/create": {
            "post": {
                "description": "权限创建",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限"
                ],
                "summary": "权限创建",
                "operationId": "RbacPermissionsCreateUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "创建",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacPermissionsCreateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_permissions/find_by_id": {
            "get": {
                "description": "权限查询ById",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限"
                ],
                "summary": "权限ById",
                "operationId": "RbacPermissionsFindByIdUsingGET",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "查询条件id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_permissions/find_by_permission": {
            "get": {
                "description": "权限查询ByPermission\u003cbr/\u003e字符串类型字段：空串或不传都代表忽略该查询，如果要指定空串为条件，则指定为 STRING__BLANK，例如 { \"name\": \"STRING__BLANK\" }",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限"
                ],
                "summary": "权限ByPermission",
                "operationId": "RbacPermissionsFindByPermissionUsingGET",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "查询条件值",
                        "name": "permission",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_permissions/query": {
            "post": {
                "description": "权限查询记录\u003cbr\u003e请求示例：\u003cbr\u003e{}\u003cbr\u003e其他查询条件，可根据以下条件字段酌情添加查询条件\u003cbr\u003e注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限"
                ],
                "summary": "权限查询",
                "operationId": "RbacPermissionsQueryUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "条件查询",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacPermissionsQueryDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_permissions/query_page": {
            "post": {
                "description": "权限分页查询\u003cbr\u003e请求示例：\u003cbr\u003e{\u003cbr\u003e\u0026nbsp;\u0026nbsp;\"start\": 0,\u003cbr\u003e\u0026nbsp;\u0026nbsp;\"limit\": 20\u003cbr\u003e}\u003cbr\u003e其他查询条件，可根据以下条件字段酌情添加查询条件\u003cbr\u003e注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限"
                ],
                "summary": "权限分页查询",
                "operationId": "RbacPermissionsQueryPageUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "条件查询",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacPermissionsPageDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_permissions/update": {
            "post": {
                "description": "权限更新，根据 id 更新",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "权限"
                ],
                "summary": "权限更新",
                "operationId": "RbacPermissionsUpdateUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "更新",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacPermissionsUpdateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_role/create": {
            "post": {
                "description": "角色创建",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "角色创建",
                "operationId": "RbacRoleCreateUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "创建",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacRoleCreateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_role/find_by_id": {
            "get": {
                "description": "角色查询ById",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "角色ById",
                "operationId": "RbacRoleFindByIdUsingGET",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "查询条件id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_role/find_by_role_name": {
            "get": {
                "description": "角色查询ByRoleName\u003cbr/\u003e字符串类型字段：空串或不传都代表忽略该查询，如果要指定空串为条件，则指定为 STRING__BLANK，例如 { \"name\": \"STRING__BLANK\" }",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "角色ByRoleName",
                "operationId": "RbacRoleFindByRoleNameUsingGET",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "查询条件值",
                        "name": "roleName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_role/query": {
            "post": {
                "description": "角色查询记录\u003cbr\u003e请求示例：\u003cbr\u003e{}\u003cbr\u003e其他查询条件，可根据以下条件字段酌情添加查询条件\u003cbr\u003e注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "角色查询",
                "operationId": "RbacRoleQueryUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "条件查询",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacRoleQueryDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_role/query_page": {
            "post": {
                "description": "角色分页查询\u003cbr\u003e请求示例：\u003cbr\u003e{\u003cbr\u003e\u0026nbsp;\u0026nbsp;\"start\": 0,\u003cbr\u003e\u0026nbsp;\u0026nbsp;\"limit\": 20\u003cbr\u003e}\u003cbr\u003e其他查询条件，可根据以下条件字段酌情添加查询条件\u003cbr\u003e注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "角色分页查询",
                "operationId": "RbacRoleQueryPageUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "条件查询",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacRolePageDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_role/update": {
            "post": {
                "description": "角色更新，根据 id 更新",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "角色更新",
                "operationId": "RbacRoleUpdateUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "更新",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacRoleUpdateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_token/create": {
            "post": {
                "description": "Token创建",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Token创建",
                "operationId": "RbacTokenCreateUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "创建",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacTokenCreateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_token/find_by_id": {
            "get": {
                "description": "Token查询ById",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "TokenById",
                "operationId": "RbacTokenFindByIdUsingGET",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "查询条件id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_token/find_by_token": {
            "get": {
                "description": "Token查询ByToken\u003cbr/\u003e字符串类型字段：空串或不传都代表忽略该查询，如果要指定空串为条件，则指定为 STRING__BLANK，例如 { \"name\": \"STRING__BLANK\" }",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "TokenByToken",
                "operationId": "RbacTokenFindByTokenUsingGET",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "查询条件值",
                        "name": "token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_token/query": {
            "post": {
                "description": "Token查询记录\u003cbr\u003e请求示例：\u003cbr\u003e{}\u003cbr\u003e其他查询条件，可根据以下条件字段酌情添加查询条件\u003cbr\u003e注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Token查询",
                "operationId": "RbacTokenQueryUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "条件查询",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacTokenQueryDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_token/query_page": {
            "post": {
                "description": "Token分页查询\u003cbr\u003e请求示例：\u003cbr\u003e{\u003cbr\u003e\u0026nbsp;\u0026nbsp;\"start\": 0,\u003cbr\u003e\u0026nbsp;\u0026nbsp;\"limit\": 20\u003cbr\u003e}\u003cbr\u003e其他查询条件，可根据以下条件字段酌情添加查询条件\u003cbr\u003e注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Token分页查询",
                "operationId": "RbacTokenQueryPageUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "条件查询",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacTokenPageDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_token/update": {
            "post": {
                "description": "Token更新，根据 id 更新",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Token"
                ],
                "summary": "Token更新",
                "operationId": "RbacTokenUpdateUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "更新",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacTokenUpdateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_user/create": {
            "post": {
                "description": "用户创建",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户创建",
                "operationId": "RbacUserCreateUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "创建",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.AddRbacUserDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_user/find_by_id": {
            "get": {
                "description": "用户查询ById",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户ById",
                "operationId": "RbacUserFindByIdUsingGET",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "查询条件id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_user/find_by_login_name": {
            "get": {
                "description": "用户查询ByLoginName\u003cbr/\u003e字符串类型字段：空串或不传都代表忽略该查询，如果要指定空串为条件，则指定为 STRING__BLANK，例如 { \"name\": \"STRING__BLANK\" }",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户ByLoginName",
                "operationId": "RbacUserFindByLoginNameUsingGET",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "查询条件值",
                        "name": "loginName",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_user/query": {
            "post": {
                "description": "用户查询记录\u003cbr\u003e请求示例：\u003cbr\u003e{}\u003cbr\u003e其他查询条件，可根据以下条件字段酌情添加查询条件\u003cbr\u003e注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户查询",
                "operationId": "RbacUserQueryUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "条件查询",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacUserQueryDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_user/query_page": {
            "post": {
                "description": "用户分页查询\u003cbr\u003e请求示例：\u003cbr\u003e{\u003cbr\u003e\u0026nbsp;\u0026nbsp;\"start\": 0,\u003cbr\u003e\u0026nbsp;\u0026nbsp;\"limit\": 20\u003cbr\u003e}\u003cbr\u003e其他查询条件，可根据以下条件字段酌情添加查询条件\u003cbr\u003e注意：字段赋值只能为【字符串】或【数字】（包括数组内元素），不能赋值为请求示例中的 【{}】",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户分页查询",
                "operationId": "RbacUserQueryPageUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "条件查询",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacUserPageDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/rbac_user/update": {
            "post": {
                "description": "用户更新，根据 id 更新",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户更新",
                "operationId": "RbacUserUpdateUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "更新",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.RbacUserUpdateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/self_modify_password": {
            "post": {
                "description": "密码修改",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登陆用户"
                ],
                "summary": "密码修改",
                "operationId": "AuthSelfModifyPasswordUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "资料修改",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.SelfModifyPasswordDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        },
        "/self_modify_profile": {
            "post": {
                "description": "资料修改",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登陆用户"
                ],
                "summary": "资料修改",
                "operationId": "AuthSelfModifyProfileUsingPOST",
                "parameters": [
                    {
                        "type": "string",
                        "description": "登陆成功后的授权 Token，后续的所有接口header，都要带上 token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "资料修改",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth_rbac.SelfModifyProfileDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.ResponseResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth_rbac.AddRbacUserDto": {
            "type": "object",
            "required": [
                "loginName",
                "loginPswd",
                "repeatLoginPswd"
            ],
            "properties": {
                "admin": {
                    "description": "是否管理员",
                    "type": "boolean"
                },
                "loginName": {
                    "description": "登录名",
                    "type": "string"
                },
                "loginPswd": {
                    "description": "登陆密码",
                    "type": "string"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "phone": {
                    "description": "手机号",
                    "type": "string"
                },
                "repeatLoginPswd": {
                    "description": "重复登陆密码",
                    "type": "string"
                },
                "superAdmin": {
                    "description": "是否超级管理员",
                    "type": "boolean"
                }
            }
        },
        "auth_rbac.LoginDto": {
            "type": "object",
            "properties": {
                "loginName": {
                    "description": "登录名",
                    "type": "string"
                },
                "loginPassword": {
                    "description": "登陆密码",
                    "type": "string"
                }
            }
        },
        "auth_rbac.RbacLogCreateDto": {
            "type": "object",
            "required": [
                "ip",
                "loginName",
                "url",
                "urlParams"
            ],
            "properties": {
                "body": {
                    "description": "请求Body体参数",
                    "type": "string"
                },
                "id": {
                    "description": "Id",
                    "type": "integer"
                },
                "ip": {
                    "description": "IP地址",
                    "type": "string"
                },
                "loginName": {
                    "description": "登陆用户",
                    "type": "string"
                },
                "status": {
                    "description": "状态 0：正常 1：鉴权失败",
                    "type": "integer"
                },
                "url": {
                    "description": "访问地址",
                    "type": "string"
                },
                "urlParams": {
                    "description": "地址参数",
                    "type": "string"
                }
            }
        },
        "auth_rbac.RbacLogPageDto": {
            "type": "object",
            "properties": {
                "body": {
                    "description": "【body】请求Body体参数 全匹配"
                },
                "bodyIn": {
                    "description": "【body】请求Body体参数 in 查询",
                    "type": "array",
                    "items": {}
                },
                "bodyLeft": {
                    "description": "【body】请求Body体参数 左匹配(xxx%)"
                },
                "bodyMiddle": {
                    "description": "【body】请求Body体参数 模糊匹配(%xxx%，查询有性能影响)"
                },
                "bodyRight": {
                    "description": "【body】请求Body体参数 右匹配(%xxx，查询有性能影响)"
                },
                "createdAtBetween": {
                    "description": "【created_at】CreatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "id": {
                    "description": "【id】Id 全匹配"
                },
                "idBetween": {
                    "description": "【id】Id 范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "idGt": {
                    "description": "【id】Id 大于"
                },
                "idGte": {
                    "description": "【id】Id 大于等于"
                },
                "idIn": {
                    "description": "【id】Id in 查询",
                    "type": "array",
                    "items": {}
                },
                "idLt": {
                    "description": "【id】Id 小于"
                },
                "idLte": {
                    "description": "【id】Id 小于等于"
                },
                "ip": {
                    "description": "【ip】IP地址 全匹配"
                },
                "ipIn": {
                    "description": "【ip】IP地址 in 查询",
                    "type": "array",
                    "items": {}
                },
                "ipLeft": {
                    "description": "【ip】IP地址 左匹配(xxx%)"
                },
                "ipMiddle": {
                    "description": "【ip】IP地址 模糊匹配(%xxx%，查询有性能影响)"
                },
                "ipRight": {
                    "description": "【ip】IP地址 右匹配(%xxx，查询有性能影响)"
                },
                "loginName": {
                    "description": "【login_name】登陆用户 全匹配"
                },
                "loginNameIn": {
                    "description": "【login_name】登陆用户 in 查询",
                    "type": "array",
                    "items": {}
                },
                "loginNameLeft": {
                    "description": "【login_name】登陆用户 左匹配(xxx%)"
                },
                "loginNameMiddle": {
                    "description": "【login_name】登陆用户 模糊匹配(%xxx%，查询有性能影响)"
                },
                "loginNameRight": {
                    "description": "【login_name】登陆用户 右匹配(%xxx，查询有性能影响)"
                },
                "orderBy": {
                    "description": "排序，例如：[\"id desc\", \"name asc\"] 字段名为每个字段说明后的开头中【】内的内容，方式只能为 desc 或 asc。支持指定多个字段排序",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "pageNum": {
                    "description": "分页开始位置从 1 开始",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "分页大小（默认值 20），必须大于 0",
                    "type": "integer"
                },
                "status": {
                    "description": "【status】状态 0：正常 1：鉴权失败 全匹配"
                },
                "statusBetween": {
                    "description": "【status】状态 0：正常 1：鉴权失败 范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "statusGt": {
                    "description": "【status】状态 0：正常 1：鉴权失败 大于"
                },
                "statusGte": {
                    "description": "【status】状态 0：正常 1：鉴权失败 大于等于"
                },
                "statusIn": {
                    "description": "【status】状态 0：正常 1：鉴权失败 in 查询",
                    "type": "array",
                    "items": {}
                },
                "statusLt": {
                    "description": "【status】状态 0：正常 1：鉴权失败 小于"
                },
                "statusLte": {
                    "description": "【status】状态 0：正常 1：鉴权失败 小于等于"
                },
                "total": {
                    "description": "总条数",
                    "type": "integer"
                },
                "updatedAtBetween": {
                    "description": "【updated_at】UpdatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "url": {
                    "description": "【url】访问地址 全匹配"
                },
                "urlIn": {
                    "description": "【url】访问地址 in 查询",
                    "type": "array",
                    "items": {}
                },
                "urlLeft": {
                    "description": "【url】访问地址 左匹配(xxx%)"
                },
                "urlMiddle": {
                    "description": "【url】访问地址 模糊匹配(%xxx%，查询有性能影响)"
                },
                "urlParams": {
                    "description": "【url_params】地址参数 全匹配"
                },
                "urlParamsIn": {
                    "description": "【url_params】地址参数 in 查询",
                    "type": "array",
                    "items": {}
                },
                "urlParamsLeft": {
                    "description": "【url_params】地址参数 左匹配(xxx%)"
                },
                "urlParamsMiddle": {
                    "description": "【url_params】地址参数 模糊匹配(%xxx%，查询有性能影响)"
                },
                "urlParamsRight": {
                    "description": "【url_params】地址参数 右匹配(%xxx，查询有性能影响)"
                },
                "urlRight": {
                    "description": "【url】访问地址 右匹配(%xxx，查询有性能影响)"
                }
            }
        },
        "auth_rbac.RbacLogQueryDto": {
            "type": "object",
            "properties": {
                "body": {
                    "description": "【body】请求Body体参数 全匹配"
                },
                "bodyIn": {
                    "description": "【body】请求Body体参数 in 查询",
                    "type": "array",
                    "items": {}
                },
                "bodyLeft": {
                    "description": "【body】请求Body体参数 左匹配(xxx%)"
                },
                "bodyMiddle": {
                    "description": "【body】请求Body体参数 模糊匹配(%xxx%，查询有性能影响)"
                },
                "bodyRight": {
                    "description": "【body】请求Body体参数 右匹配(%xxx，查询有性能影响)"
                },
                "createdAtBetween": {
                    "description": "【created_at】CreatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "id": {
                    "description": "【id】Id 全匹配"
                },
                "idBetween": {
                    "description": "【id】Id 范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "idGt": {
                    "description": "【id】Id 大于"
                },
                "idGte": {
                    "description": "【id】Id 大于等于"
                },
                "idIn": {
                    "description": "【id】Id in 查询",
                    "type": "array",
                    "items": {}
                },
                "idLt": {
                    "description": "【id】Id 小于"
                },
                "idLte": {
                    "description": "【id】Id 小于等于"
                },
                "ip": {
                    "description": "【ip】IP地址 全匹配"
                },
                "ipIn": {
                    "description": "【ip】IP地址 in 查询",
                    "type": "array",
                    "items": {}
                },
                "ipLeft": {
                    "description": "【ip】IP地址 左匹配(xxx%)"
                },
                "ipMiddle": {
                    "description": "【ip】IP地址 模糊匹配(%xxx%，查询有性能影响)"
                },
                "ipRight": {
                    "description": "【ip】IP地址 右匹配(%xxx，查询有性能影响)"
                },
                "loginName": {
                    "description": "【login_name】登陆用户 全匹配"
                },
                "loginNameIn": {
                    "description": "【login_name】登陆用户 in 查询",
                    "type": "array",
                    "items": {}
                },
                "loginNameLeft": {
                    "description": "【login_name】登陆用户 左匹配(xxx%)"
                },
                "loginNameMiddle": {
                    "description": "【login_name】登陆用户 模糊匹配(%xxx%，查询有性能影响)"
                },
                "loginNameRight": {
                    "description": "【login_name】登陆用户 右匹配(%xxx，查询有性能影响)"
                },
                "orderBy": {
                    "description": "排序，例如：[\"id desc\", \"name asc\"] 字段名为每个字段说明后的开头中【】内的内容，方式只能为 desc 或 asc。支持指定多个字段排序",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "status": {
                    "description": "【status】状态 0：正常 1：鉴权失败 全匹配"
                },
                "statusBetween": {
                    "description": "【status】状态 0：正常 1：鉴权失败 范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "statusGt": {
                    "description": "【status】状态 0：正常 1：鉴权失败 大于"
                },
                "statusGte": {
                    "description": "【status】状态 0：正常 1：鉴权失败 大于等于"
                },
                "statusIn": {
                    "description": "【status】状态 0：正常 1：鉴权失败 in 查询",
                    "type": "array",
                    "items": {}
                },
                "statusLt": {
                    "description": "【status】状态 0：正常 1：鉴权失败 小于"
                },
                "statusLte": {
                    "description": "【status】状态 0：正常 1：鉴权失败 小于等于"
                },
                "updatedAtBetween": {
                    "description": "【updated_at】UpdatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "url": {
                    "description": "【url】访问地址 全匹配"
                },
                "urlIn": {
                    "description": "【url】访问地址 in 查询",
                    "type": "array",
                    "items": {}
                },
                "urlLeft": {
                    "description": "【url】访问地址 左匹配(xxx%)"
                },
                "urlMiddle": {
                    "description": "【url】访问地址 模糊匹配(%xxx%，查询有性能影响)"
                },
                "urlParams": {
                    "description": "【url_params】地址参数 全匹配"
                },
                "urlParamsIn": {
                    "description": "【url_params】地址参数 in 查询",
                    "type": "array",
                    "items": {}
                },
                "urlParamsLeft": {
                    "description": "【url_params】地址参数 左匹配(xxx%)"
                },
                "urlParamsMiddle": {
                    "description": "【url_params】地址参数 模糊匹配(%xxx%，查询有性能影响)"
                },
                "urlParamsRight": {
                    "description": "【url_params】地址参数 右匹配(%xxx，查询有性能影响)"
                },
                "urlRight": {
                    "description": "【url】访问地址 右匹配(%xxx，查询有性能影响)"
                }
            }
        },
        "auth_rbac.RbacLogUpdateDto": {
            "type": "object",
            "properties": {
                "body": {
                    "description": "请求Body体参数"
                },
                "id": {
                    "description": "Id"
                },
                "ip": {
                    "description": "IP地址"
                },
                "loginName": {
                    "description": "登陆用户"
                },
                "status": {
                    "description": "状态 0：正常 1：鉴权失败"
                },
                "url": {
                    "description": "访问地址"
                },
                "urlParams": {
                    "description": "地址参数"
                }
            }
        },
        "auth_rbac.RbacPermissionsCreateDto": {
            "type": "object",
            "required": [
                "description",
                "permission"
            ],
            "properties": {
                "description": {
                    "description": "权限描述",
                    "type": "string"
                },
                "id": {
                    "description": "Id",
                    "type": "integer"
                },
                "permission": {
                    "description": "权限值",
                    "type": "string"
                }
            }
        },
        "auth_rbac.RbacPermissionsPageDto": {
            "type": "object",
            "properties": {
                "createdAtBetween": {
                    "description": "【created_at】CreatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "description": {
                    "description": "【description】权限描述 全匹配"
                },
                "descriptionIn": {
                    "description": "【description】权限描述 in 查询",
                    "type": "array",
                    "items": {}
                },
                "descriptionLeft": {
                    "description": "【description】权限描述 左匹配(xxx%)"
                },
                "descriptionMiddle": {
                    "description": "【description】权限描述 模糊匹配(%xxx%，查询有性能影响)"
                },
                "descriptionRight": {
                    "description": "【description】权限描述 右匹配(%xxx，查询有性能影响)"
                },
                "id": {
                    "description": "【id】Id 全匹配"
                },
                "idBetween": {
                    "description": "【id】Id 范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "idGt": {
                    "description": "【id】Id 大于"
                },
                "idGte": {
                    "description": "【id】Id 大于等于"
                },
                "idIn": {
                    "description": "【id】Id in 查询",
                    "type": "array",
                    "items": {}
                },
                "idLt": {
                    "description": "【id】Id 小于"
                },
                "idLte": {
                    "description": "【id】Id 小于等于"
                },
                "orderBy": {
                    "description": "排序，例如：[\"id desc\", \"name asc\"] 字段名为每个字段说明后的开头中【】内的内容，方式只能为 desc 或 asc。支持指定多个字段排序",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "pageNum": {
                    "description": "分页开始位置从 1 开始",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "分页大小（默认值 20），必须大于 0",
                    "type": "integer"
                },
                "permission": {
                    "description": "【permission】权限值 全匹配"
                },
                "permissionIn": {
                    "description": "【permission】权限值 in 查询",
                    "type": "array",
                    "items": {}
                },
                "permissionLeft": {
                    "description": "【permission】权限值 左匹配(xxx%)"
                },
                "permissionMiddle": {
                    "description": "【permission】权限值 模糊匹配(%xxx%，查询有性能影响)"
                },
                "permissionRight": {
                    "description": "【permission】权限值 右匹配(%xxx，查询有性能影响)"
                },
                "total": {
                    "description": "总条数",
                    "type": "integer"
                },
                "updatedAtBetween": {
                    "description": "【updated_at】UpdatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                }
            }
        },
        "auth_rbac.RbacPermissionsQueryDto": {
            "type": "object",
            "properties": {
                "createdAtBetween": {
                    "description": "【created_at】CreatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "description": {
                    "description": "【description】权限描述 全匹配"
                },
                "descriptionIn": {
                    "description": "【description】权限描述 in 查询",
                    "type": "array",
                    "items": {}
                },
                "descriptionLeft": {
                    "description": "【description】权限描述 左匹配(xxx%)"
                },
                "descriptionMiddle": {
                    "description": "【description】权限描述 模糊匹配(%xxx%，查询有性能影响)"
                },
                "descriptionRight": {
                    "description": "【description】权限描述 右匹配(%xxx，查询有性能影响)"
                },
                "id": {
                    "description": "【id】Id 全匹配"
                },
                "idBetween": {
                    "description": "【id】Id 范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "idGt": {
                    "description": "【id】Id 大于"
                },
                "idGte": {
                    "description": "【id】Id 大于等于"
                },
                "idIn": {
                    "description": "【id】Id in 查询",
                    "type": "array",
                    "items": {}
                },
                "idLt": {
                    "description": "【id】Id 小于"
                },
                "idLte": {
                    "description": "【id】Id 小于等于"
                },
                "orderBy": {
                    "description": "排序，例如：[\"id desc\", \"name asc\"] 字段名为每个字段说明后的开头中【】内的内容，方式只能为 desc 或 asc。支持指定多个字段排序",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "permission": {
                    "description": "【permission】权限值 全匹配"
                },
                "permissionIn": {
                    "description": "【permission】权限值 in 查询",
                    "type": "array",
                    "items": {}
                },
                "permissionLeft": {
                    "description": "【permission】权限值 左匹配(xxx%)"
                },
                "permissionMiddle": {
                    "description": "【permission】权限值 模糊匹配(%xxx%，查询有性能影响)"
                },
                "permissionRight": {
                    "description": "【permission】权限值 右匹配(%xxx，查询有性能影响)"
                },
                "updatedAtBetween": {
                    "description": "【updated_at】UpdatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                }
            }
        },
        "auth_rbac.RbacPermissionsUpdateDto": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "权限描述"
                },
                "id": {
                    "description": "Id"
                },
                "permission": {
                    "description": "权限值"
                }
            }
        },
        "auth_rbac.RbacRoleCreateDto": {
            "type": "object",
            "required": [
                "roleName"
            ],
            "properties": {
                "id": {
                    "description": "Id",
                    "type": "integer"
                },
                "roleDesc": {
                    "description": "角色描述",
                    "type": "string"
                },
                "roleName": {
                    "description": "角色名",
                    "type": "string"
                },
                "valid": {
                    "description": "是否可用",
                    "type": "boolean"
                }
            }
        },
        "auth_rbac.RbacRolePageDto": {
            "type": "object",
            "properties": {
                "createdAtBetween": {
                    "description": "【created_at】CreatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "id": {
                    "description": "【id】Id 全匹配"
                },
                "idBetween": {
                    "description": "【id】Id 范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "idGt": {
                    "description": "【id】Id 大于"
                },
                "idGte": {
                    "description": "【id】Id 大于等于"
                },
                "idIn": {
                    "description": "【id】Id in 查询",
                    "type": "array",
                    "items": {}
                },
                "idLt": {
                    "description": "【id】Id 小于"
                },
                "idLte": {
                    "description": "【id】Id 小于等于"
                },
                "orderBy": {
                    "description": "排序，例如：[\"id desc\", \"name asc\"] 字段名为每个字段说明后的开头中【】内的内容，方式只能为 desc 或 asc。支持指定多个字段排序",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "pageNum": {
                    "description": "分页开始位置从 1 开始",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "分页大小（默认值 20），必须大于 0",
                    "type": "integer"
                },
                "roleDesc": {
                    "description": "【role_desc】角色描述 全匹配"
                },
                "roleDescIn": {
                    "description": "【role_desc】角色描述 in 查询",
                    "type": "array",
                    "items": {}
                },
                "roleDescLeft": {
                    "description": "【role_desc】角色描述 左匹配(xxx%)"
                },
                "roleDescMiddle": {
                    "description": "【role_desc】角色描述 模糊匹配(%xxx%，查询有性能影响)"
                },
                "roleDescRight": {
                    "description": "【role_desc】角色描述 右匹配(%xxx，查询有性能影响)"
                },
                "roleName": {
                    "description": "【role_name】角色名 全匹配"
                },
                "roleNameIn": {
                    "description": "【role_name】角色名 in 查询",
                    "type": "array",
                    "items": {}
                },
                "roleNameLeft": {
                    "description": "【role_name】角色名 左匹配(xxx%)"
                },
                "roleNameMiddle": {
                    "description": "【role_name】角色名 模糊匹配(%xxx%，查询有性能影响)"
                },
                "roleNameRight": {
                    "description": "【role_name】角色名 右匹配(%xxx，查询有性能影响)"
                },
                "total": {
                    "description": "总条数",
                    "type": "integer"
                },
                "updatedAtBetween": {
                    "description": "【updated_at】UpdatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "valid": {
                    "description": "【valid】是否可用 全匹配"
                },
                "validIn": {
                    "description": "【valid】是否可用 in 查询",
                    "type": "array",
                    "items": {}
                }
            }
        },
        "auth_rbac.RbacRoleQueryDto": {
            "type": "object",
            "properties": {
                "createdAtBetween": {
                    "description": "【created_at】CreatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "id": {
                    "description": "【id】Id 全匹配"
                },
                "idBetween": {
                    "description": "【id】Id 范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "idGt": {
                    "description": "【id】Id 大于"
                },
                "idGte": {
                    "description": "【id】Id 大于等于"
                },
                "idIn": {
                    "description": "【id】Id in 查询",
                    "type": "array",
                    "items": {}
                },
                "idLt": {
                    "description": "【id】Id 小于"
                },
                "idLte": {
                    "description": "【id】Id 小于等于"
                },
                "orderBy": {
                    "description": "排序，例如：[\"id desc\", \"name asc\"] 字段名为每个字段说明后的开头中【】内的内容，方式只能为 desc 或 asc。支持指定多个字段排序",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "roleDesc": {
                    "description": "【role_desc】角色描述 全匹配"
                },
                "roleDescIn": {
                    "description": "【role_desc】角色描述 in 查询",
                    "type": "array",
                    "items": {}
                },
                "roleDescLeft": {
                    "description": "【role_desc】角色描述 左匹配(xxx%)"
                },
                "roleDescMiddle": {
                    "description": "【role_desc】角色描述 模糊匹配(%xxx%，查询有性能影响)"
                },
                "roleDescRight": {
                    "description": "【role_desc】角色描述 右匹配(%xxx，查询有性能影响)"
                },
                "roleName": {
                    "description": "【role_name】角色名 全匹配"
                },
                "roleNameIn": {
                    "description": "【role_name】角色名 in 查询",
                    "type": "array",
                    "items": {}
                },
                "roleNameLeft": {
                    "description": "【role_name】角色名 左匹配(xxx%)"
                },
                "roleNameMiddle": {
                    "description": "【role_name】角色名 模糊匹配(%xxx%，查询有性能影响)"
                },
                "roleNameRight": {
                    "description": "【role_name】角色名 右匹配(%xxx，查询有性能影响)"
                },
                "updatedAtBetween": {
                    "description": "【updated_at】UpdatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "valid": {
                    "description": "【valid】是否可用 全匹配"
                },
                "validIn": {
                    "description": "【valid】是否可用 in 查询",
                    "type": "array",
                    "items": {}
                }
            }
        },
        "auth_rbac.RbacRoleUpdateDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "Id"
                },
                "roleDesc": {
                    "description": "角色描述"
                },
                "roleName": {
                    "description": "角色名"
                },
                "valid": {
                    "description": "是否可用"
                }
            }
        },
        "auth_rbac.RbacTokenCreateDto": {
            "type": "object",
            "required": [
                "token",
                "userId"
            ],
            "properties": {
                "id": {
                    "description": "Id",
                    "type": "integer"
                },
                "token": {
                    "description": "Token",
                    "type": "string"
                },
                "userId": {
                    "description": "用户ID",
                    "type": "integer"
                }
            }
        },
        "auth_rbac.RbacTokenPageDto": {
            "type": "object",
            "properties": {
                "createdAtBetween": {
                    "description": "【created_at】CreatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "id": {
                    "description": "【id】Id 全匹配"
                },
                "idBetween": {
                    "description": "【id】Id 范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "idGt": {
                    "description": "【id】Id 大于"
                },
                "idGte": {
                    "description": "【id】Id 大于等于"
                },
                "idIn": {
                    "description": "【id】Id in 查询",
                    "type": "array",
                    "items": {}
                },
                "idLt": {
                    "description": "【id】Id 小于"
                },
                "idLte": {
                    "description": "【id】Id 小于等于"
                },
                "orderBy": {
                    "description": "排序，例如：[\"id desc\", \"name asc\"] 字段名为每个字段说明后的开头中【】内的内容，方式只能为 desc 或 asc。支持指定多个字段排序",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "pageNum": {
                    "description": "分页开始位置从 1 开始",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "分页大小（默认值 20），必须大于 0",
                    "type": "integer"
                },
                "token": {
                    "description": "【token】Token 全匹配"
                },
                "tokenIn": {
                    "description": "【token】Token in 查询",
                    "type": "array",
                    "items": {}
                },
                "tokenLeft": {
                    "description": "【token】Token 左匹配(xxx%)"
                },
                "tokenMiddle": {
                    "description": "【token】Token 模糊匹配(%xxx%，查询有性能影响)"
                },
                "tokenRight": {
                    "description": "【token】Token 右匹配(%xxx，查询有性能影响)"
                },
                "total": {
                    "description": "总条数",
                    "type": "integer"
                },
                "updatedAtBetween": {
                    "description": "【updated_at】UpdatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "userId": {
                    "description": "【user_id】用户ID 全匹配"
                },
                "userIdIn": {
                    "description": "【user_id】用户ID in 查询",
                    "type": "array",
                    "items": {}
                }
            }
        },
        "auth_rbac.RbacTokenQueryDto": {
            "type": "object",
            "properties": {
                "createdAtBetween": {
                    "description": "【created_at】CreatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "id": {
                    "description": "【id】Id 全匹配"
                },
                "idBetween": {
                    "description": "【id】Id 范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "idGt": {
                    "description": "【id】Id 大于"
                },
                "idGte": {
                    "description": "【id】Id 大于等于"
                },
                "idIn": {
                    "description": "【id】Id in 查询",
                    "type": "array",
                    "items": {}
                },
                "idLt": {
                    "description": "【id】Id 小于"
                },
                "idLte": {
                    "description": "【id】Id 小于等于"
                },
                "orderBy": {
                    "description": "排序，例如：[\"id desc\", \"name asc\"] 字段名为每个字段说明后的开头中【】内的内容，方式只能为 desc 或 asc。支持指定多个字段排序",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "token": {
                    "description": "【token】Token 全匹配"
                },
                "tokenIn": {
                    "description": "【token】Token in 查询",
                    "type": "array",
                    "items": {}
                },
                "tokenLeft": {
                    "description": "【token】Token 左匹配(xxx%)"
                },
                "tokenMiddle": {
                    "description": "【token】Token 模糊匹配(%xxx%，查询有性能影响)"
                },
                "tokenRight": {
                    "description": "【token】Token 右匹配(%xxx，查询有性能影响)"
                },
                "updatedAtBetween": {
                    "description": "【updated_at】UpdatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "userId": {
                    "description": "【user_id】用户ID 全匹配"
                },
                "userIdIn": {
                    "description": "【user_id】用户ID in 查询",
                    "type": "array",
                    "items": {}
                }
            }
        },
        "auth_rbac.RbacTokenUpdateDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "Id"
                },
                "token": {
                    "description": "Token"
                },
                "userId": {
                    "description": "用户ID"
                }
            }
        },
        "auth_rbac.RbacUserPageDto": {
            "type": "object",
            "properties": {
                "admin": {
                    "description": "【admin】是否管理员 全匹配"
                },
                "adminIn": {
                    "description": "【admin】是否管理员 in 查询",
                    "type": "array",
                    "items": {}
                },
                "createdAtBetween": {
                    "description": "【created_at】CreatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "id": {
                    "description": "【id】Id 全匹配"
                },
                "idBetween": {
                    "description": "【id】Id 范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "idGt": {
                    "description": "【id】Id 大于"
                },
                "idGte": {
                    "description": "【id】Id 大于等于"
                },
                "idIn": {
                    "description": "【id】Id in 查询",
                    "type": "array",
                    "items": {}
                },
                "idLt": {
                    "description": "【id】Id 小于"
                },
                "idLte": {
                    "description": "【id】Id 小于等于"
                },
                "loginName": {
                    "description": "【login_name】登录名 全匹配"
                },
                "loginNameIn": {
                    "description": "【login_name】登录名 in 查询",
                    "type": "array",
                    "items": {}
                },
                "loginNameLeft": {
                    "description": "【login_name】登录名 左匹配(xxx%)"
                },
                "loginNameMiddle": {
                    "description": "【login_name】登录名 模糊匹配(%xxx%，查询有性能影响)"
                },
                "loginNameRight": {
                    "description": "【login_name】登录名 右匹配(%xxx，查询有性能影响)"
                },
                "loginPswd": {
                    "description": "【login_pswd】登陆密码 全匹配"
                },
                "loginPswdIn": {
                    "description": "【login_pswd】登陆密码 in 查询",
                    "type": "array",
                    "items": {}
                },
                "loginPswdLeft": {
                    "description": "【login_pswd】登陆密码 左匹配(xxx%)"
                },
                "loginPswdMiddle": {
                    "description": "【login_pswd】登陆密码 模糊匹配(%xxx%，查询有性能影响)"
                },
                "loginPswdRight": {
                    "description": "【login_pswd】登陆密码 右匹配(%xxx，查询有性能影响)"
                },
                "nickname": {
                    "description": "【nickname】昵称 全匹配"
                },
                "nicknameIn": {
                    "description": "【nickname】昵称 in 查询",
                    "type": "array",
                    "items": {}
                },
                "nicknameLeft": {
                    "description": "【nickname】昵称 左匹配(xxx%)"
                },
                "nicknameMiddle": {
                    "description": "【nickname】昵称 模糊匹配(%xxx%，查询有性能影响)"
                },
                "nicknameRight": {
                    "description": "【nickname】昵称 右匹配(%xxx，查询有性能影响)"
                },
                "orderBy": {
                    "description": "排序，例如：[\"id desc\", \"name asc\"] 字段名为每个字段说明后的开头中【】内的内容，方式只能为 desc 或 asc。支持指定多个字段排序",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "pageNum": {
                    "description": "分页开始位置从 1 开始",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "分页大小（默认值 20），必须大于 0",
                    "type": "integer"
                },
                "phone": {
                    "description": "【phone】手机号 全匹配"
                },
                "phoneIn": {
                    "description": "【phone】手机号 in 查询",
                    "type": "array",
                    "items": {}
                },
                "phoneLeft": {
                    "description": "【phone】手机号 左匹配(xxx%)"
                },
                "phoneMiddle": {
                    "description": "【phone】手机号 模糊匹配(%xxx%，查询有性能影响)"
                },
                "phoneRight": {
                    "description": "【phone】手机号 右匹配(%xxx，查询有性能影响)"
                },
                "salt": {
                    "description": "【salt】密码加盐 全匹配"
                },
                "saltIn": {
                    "description": "【salt】密码加盐 in 查询",
                    "type": "array",
                    "items": {}
                },
                "saltLeft": {
                    "description": "【salt】密码加盐 左匹配(xxx%)"
                },
                "saltMiddle": {
                    "description": "【salt】密码加盐 模糊匹配(%xxx%，查询有性能影响)"
                },
                "saltRight": {
                    "description": "【salt】密码加盐 右匹配(%xxx，查询有性能影响)"
                },
                "status": {
                    "description": "【status】状态 1:账号正常 0:账号禁用 -1:账号违规 全匹配"
                },
                "statusBetween": {
                    "description": "【status】状态 1:账号正常 0:账号禁用 -1:账号违规 范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "statusGt": {
                    "description": "【status】状态 1:账号正常 0:账号禁用 -1:账号违规 大于"
                },
                "statusGte": {
                    "description": "【status】状态 1:账号正常 0:账号禁用 -1:账号违规 大于等于"
                },
                "statusIn": {
                    "description": "【status】状态 1:账号正常 0:账号禁用 -1:账号违规 in 查询",
                    "type": "array",
                    "items": {}
                },
                "statusLt": {
                    "description": "【status】状态 1:账号正常 0:账号禁用 -1:账号违规 小于"
                },
                "statusLte": {
                    "description": "【status】状态 1:账号正常 0:账号禁用 -1:账号违规 小于等于"
                },
                "superAdmin": {
                    "description": "【super_admin】是否超级管理员 全匹配"
                },
                "superAdminIn": {
                    "description": "【super_admin】是否超级管理员 in 查询",
                    "type": "array",
                    "items": {}
                },
                "total": {
                    "description": "总条数",
                    "type": "integer"
                },
                "updatedAtBetween": {
                    "description": "【updated_at】UpdatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                }
            }
        },
        "auth_rbac.RbacUserQueryDto": {
            "type": "object",
            "properties": {
                "admin": {
                    "description": "【admin】是否管理员 全匹配"
                },
                "adminIn": {
                    "description": "【admin】是否管理员 in 查询",
                    "type": "array",
                    "items": {}
                },
                "createdAtBetween": {
                    "description": "【created_at】CreatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "id": {
                    "description": "【id】Id 全匹配"
                },
                "idBetween": {
                    "description": "【id】Id 范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "idGt": {
                    "description": "【id】Id 大于"
                },
                "idGte": {
                    "description": "【id】Id 大于等于"
                },
                "idIn": {
                    "description": "【id】Id in 查询",
                    "type": "array",
                    "items": {}
                },
                "idLt": {
                    "description": "【id】Id 小于"
                },
                "idLte": {
                    "description": "【id】Id 小于等于"
                },
                "loginName": {
                    "description": "【login_name】登录名 全匹配"
                },
                "loginNameIn": {
                    "description": "【login_name】登录名 in 查询",
                    "type": "array",
                    "items": {}
                },
                "loginNameLeft": {
                    "description": "【login_name】登录名 左匹配(xxx%)"
                },
                "loginNameMiddle": {
                    "description": "【login_name】登录名 模糊匹配(%xxx%，查询有性能影响)"
                },
                "loginNameRight": {
                    "description": "【login_name】登录名 右匹配(%xxx，查询有性能影响)"
                },
                "loginPswd": {
                    "description": "【login_pswd】登陆密码 全匹配"
                },
                "loginPswdIn": {
                    "description": "【login_pswd】登陆密码 in 查询",
                    "type": "array",
                    "items": {}
                },
                "loginPswdLeft": {
                    "description": "【login_pswd】登陆密码 左匹配(xxx%)"
                },
                "loginPswdMiddle": {
                    "description": "【login_pswd】登陆密码 模糊匹配(%xxx%，查询有性能影响)"
                },
                "loginPswdRight": {
                    "description": "【login_pswd】登陆密码 右匹配(%xxx，查询有性能影响)"
                },
                "nickname": {
                    "description": "【nickname】昵称 全匹配"
                },
                "nicknameIn": {
                    "description": "【nickname】昵称 in 查询",
                    "type": "array",
                    "items": {}
                },
                "nicknameLeft": {
                    "description": "【nickname】昵称 左匹配(xxx%)"
                },
                "nicknameMiddle": {
                    "description": "【nickname】昵称 模糊匹配(%xxx%，查询有性能影响)"
                },
                "nicknameRight": {
                    "description": "【nickname】昵称 右匹配(%xxx，查询有性能影响)"
                },
                "orderBy": {
                    "description": "排序，例如：[\"id desc\", \"name asc\"] 字段名为每个字段说明后的开头中【】内的内容，方式只能为 desc 或 asc。支持指定多个字段排序",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "phone": {
                    "description": "【phone】手机号 全匹配"
                },
                "phoneIn": {
                    "description": "【phone】手机号 in 查询",
                    "type": "array",
                    "items": {}
                },
                "phoneLeft": {
                    "description": "【phone】手机号 左匹配(xxx%)"
                },
                "phoneMiddle": {
                    "description": "【phone】手机号 模糊匹配(%xxx%，查询有性能影响)"
                },
                "phoneRight": {
                    "description": "【phone】手机号 右匹配(%xxx，查询有性能影响)"
                },
                "salt": {
                    "description": "【salt】密码加盐 全匹配"
                },
                "saltIn": {
                    "description": "【salt】密码加盐 in 查询",
                    "type": "array",
                    "items": {}
                },
                "saltLeft": {
                    "description": "【salt】密码加盐 左匹配(xxx%)"
                },
                "saltMiddle": {
                    "description": "【salt】密码加盐 模糊匹配(%xxx%，查询有性能影响)"
                },
                "saltRight": {
                    "description": "【salt】密码加盐 右匹配(%xxx，查询有性能影响)"
                },
                "status": {
                    "description": "【status】状态 1:账号正常 0:账号禁用 -1:账号违规 全匹配"
                },
                "statusBetween": {
                    "description": "【status】状态 1:账号正常 0:账号禁用 -1:账号违规 范围（包含边界）",
                    "type": "array",
                    "items": {}
                },
                "statusGt": {
                    "description": "【status】状态 1:账号正常 0:账号禁用 -1:账号违规 大于"
                },
                "statusGte": {
                    "description": "【status】状态 1:账号正常 0:账号禁用 -1:账号违规 大于等于"
                },
                "statusIn": {
                    "description": "【status】状态 1:账号正常 0:账号禁用 -1:账号违规 in 查询",
                    "type": "array",
                    "items": {}
                },
                "statusLt": {
                    "description": "【status】状态 1:账号正常 0:账号禁用 -1:账号违规 小于"
                },
                "statusLte": {
                    "description": "【status】状态 1:账号正常 0:账号禁用 -1:账号违规 小于等于"
                },
                "superAdmin": {
                    "description": "【super_admin】是否超级管理员 全匹配"
                },
                "superAdminIn": {
                    "description": "【super_admin】是否超级管理员 in 查询",
                    "type": "array",
                    "items": {}
                },
                "updatedAtBetween": {
                    "description": "【updated_at】UpdatedAt 时间范围（包含边界）",
                    "type": "array",
                    "items": {}
                }
            }
        },
        "auth_rbac.RbacUserUpdateDto": {
            "type": "object",
            "properties": {
                "admin": {
                    "description": "是否管理员"
                },
                "id": {
                    "description": "Id"
                },
                "loginName": {
                    "description": "登录名"
                },
                "loginPswd": {
                    "description": "登陆密码"
                },
                "nickname": {
                    "description": "昵称"
                },
                "phone": {
                    "description": "手机号"
                },
                "salt": {
                    "description": "密码加盐"
                },
                "status": {
                    "description": "状态 1:账号正常 0:账号禁用 -1:账号违规"
                },
                "superAdmin": {
                    "description": "是否超级管理员"
                }
            }
        },
        "auth_rbac.SelfModifyPasswordDto": {
            "type": "object",
            "properties": {
                "old_password": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "repeat_password": {
                    "type": "string"
                }
            }
        },
        "auth_rbac.SelfModifyProfileDto": {
            "type": "object",
            "properties": {
                "nickname": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "vo.ResponseResult": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "响应码 0为成功，其他都为错误码",
                    "type": "string"
                },
                "data": {
                    "description": "响应的结果数据"
                },
                "message": {
                    "description": "错误信息",
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8890",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "LazyFrame",
	Description: "后端系统接口",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(SwaggerDoc)
	if err != nil {
		return SwaggerDoc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return SwaggerDoc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
