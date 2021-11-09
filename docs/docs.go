// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
	"tags": [
    	{
      		"name": "Coronavirus Stats",
      		"description": "Statewise and Location wise coronavirus stats"
    	}
  	],
	"paths": {
        "/covid/stats": {
            "get": {
				"tags": [
					"Coronavirus Stats"
				],
                "description": "get the coronavirus stats statewise",
				"produces": [
                    "application/json"
                ],
                "summary": "Show the Coronavirus Stats.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
							"type": "array",
							"items": {
								"$ref": "#/definitions/CovidStats"
							}
                        }
                    }
                }
            }
        },
		"/covid/location/stats": {
            "get": {
				"tags": [
					"Coronavirus Stats"
				],
                "description": "get the coronavirus stats statewise",
				"produces": [
                    "application/json"
                ],
				"parameters": [
          			{
            			"name": "lat",
            			"in": "query",
            			"type": "number"
          			},
          			{
            			"name": "long",
            			"in": "query",
            			"type": "number"
          			}
        		],
                "summary": "Show the Coronavirus Stats.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
							"items": {
								"$ref": "#/definitions/Statewise"
							}
                        }
                    }
                }
            }
        }
	},
	"definitions": {
		"CovidStats": {
		  "type": "object",
		  "properties": {
			"timestamp": {
			  "type": "string",
			  "format": "date-time"
			},
			"data": {
				"type": "array",
				"items": {
					"$ref": "#/definitions/Statewise"
				}
			}
		  },
		  "xml": {
			"name": "CovidStats"
		  }
		},
		"Statewise": {
			"type": "object",
			"properties": {
			  "recovered": {
				"type": "string",
				"format": "date-time"
			  },
			  "activecases": {
				  "type": "string",
			  }
			},
			"xml": {
			  "name": "Statewise"
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
	Host:        "localhost:3000",
	BasePath:    "/",
	Schemes:     []string{"http"},
	Title:       "Coronavirus Stats API",
	Description: "This is a Coronavirus Stats.",
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
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
