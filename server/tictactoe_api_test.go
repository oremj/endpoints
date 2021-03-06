// Copyright 2013 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

const apiConfigJson = `
{
  "extends" : "thirdParty.api",
  "abstract" : false,
  "root" : "https://tictactoe.appspot.com/_ah/api",
  "name" : "tictactoe",
  "version" : "v1",
  "description" : "",
  "defaultVersion" : false,
  "adapter" : {
    "bns" : "http://tictactoe.appspot.com/_ah/spi",
    "type" : "lily"
  },
  "methods" : {
    "tictactoe.scores.get" : {
      "path" : "scores/{key}",
      "httpMethod" : "GET",
      "rosyMethod" : "ScoreEndpoint.get",
      "request" : {
        "parameters" : {
          "key" : {
            "required" : true,
            "type" : "string"
          }
        },
	"parameterOrder" : [ "key" ],
        "body" : "autoTemplate(backendRequest)",
        "bodyName" : "resource"
      },
      "response" : {
        "body" : "autoTemplate(backendResponse)"
      }
    },
    "tictactoe.scores.list" : {
      "path" : "scores",
      "httpMethod" : "GET",
      "rosyMethod" : "ScoreEndpoint.list",
      "request" : {
        "body" : "autoTemplate(backendRequest)",
        "bodyName" : "resource"
      },
      "response" : {
        "body" : "autoTemplate(backendResponse)"
      }
    },
    "tictactoe.scores.insert" : {
      "path" : "scores",
      "httpMethod" : "POST",
      "rosyMethod" : "ScoreEndpoint.insert",
      "request" : {
        "body" : "autoTemplate(backendRequest)",
        "bodyName" : "resource"
      },
      "response" : {
        "body" : "autoTemplate(backendResponse)"
      }
    },
    "tictactoe.scores.echo1" : {
      "path" : "echo1/{value2}",
      "httpMethod" : "GET",
      "rosyMethod" : "ScoreEndpoint.echo1",
      "request" : {
        "parameters" : {
          "value2" : {
            "required" : true,
            "type" : "string"
          }
        },
	"parameterOrder" : [ "value2" ],
        "body" : "autoTemplate(backendRequest)",
        "bodyName" : "resource"
      },
      "response" : {
        "body" : "autoTemplate(backendResponse)"
      }
    },
    "tictactoe.scores.echo2" : {
      "path" : "echo2/{name}",
      "httpMethod" : "GET",
      "rosyMethod" : "ScoreEndpoint.echo2",
      "request" : {
        "parameters" : {
          "name" : {
            "required" : true,
            "type" : "string"
          }
        },
	"parameterOrder" : [ "name" ],
        "body" : "autoTemplate(backendRequest)",
        "bodyName" : "resource"
      },
      "response" : {
        "body" : "autoTemplate(backendResponse)"
      }
    },
    "tictactoe.scores.echo3" : {
      "path" : "echo3/{version}",
      "httpMethod" : "GET",
      "rosyMethod" : "ScoreEndpoint.echo3",
      "request" : {
        "parameters" : {
          "version" : {
            "required" : true,
            "type" : "string"
          }
        },
	"parameterOrder" : [ "version" ],
        "body" : "autoTemplate(backendRequest)",
        "bodyName" : "resource"
      },
      "response" : {
        "body" : "autoTemplate(backendResponse)"
      }
    },
    "tictactoe.scores.echo4" : {
      "path" : "echo4/{version2}",
      "httpMethod" : "GET",
      "rosyMethod" : "ScoreEndpoint.echo4",
      "request" : {
        "parameters" : {
          "version2" : {
            "required" : true,
            "type" : "string"
          }
        },
	"parameterOrder" : [ "version2" ],
        "body" : "autoTemplate(backendRequest)",
        "bodyName" : "resource"
      },
      "response" : {
        "body" : "autoTemplate(backendResponse)"
      }
    },
    "tictactoe.scores.cors" : {
      "path" : "cors/{echo}",
      "httpMethod" : "GET",
      "rosyMethod" : "ScoreEndpoint.cors",
      "request" : {
        "parameters" : {
          "echo" : {
            "required" : true,
            "type" : "string"
          }
        },
	"parameterOrder" : [ "echo" ],
        "body" : "autoTemplate(backendRequest)",
        "bodyName" : "resource"
      },
      "response" : {
        "body" : "autoTemplate(backendResponse)"
      }
    },
    "tictactoe.board.getmove" : {
      "path" : "board",
      "httpMethod" : "POST",
      "rosyMethod" : "BoardEndpoint.getmove",
      "request" : {
        "body" : "autoTemplate(backendRequest)",
        "bodyName" : "resource"
      },
      "response" : {
        "body" : "autoTemplate(backendResponse)"
      }
    }
  },
  "descriptor" : {
    "schemas" : {
      "Score" : {
        "id" : "Score",
        "type" : "object",
        "properties" : {
          "encodedKey" : {
            "type" : "string"
          },
          "outcome" : {
            "type" : "string"
          },
          "played" : {
            "type" : "string",
            "format" : "date"
          }
        }
      },
      "Scores" : {
        "id" : "Scores",
        "type" : "object",
        "properties" : {
          "items" : {
            "type" : "array",
            "items" : {
              "$ref" : "Score"
            }
          }
        }
      },
      "Board" : {
        "id" : "Board",
        "type" : "object",
        "properties" : {
          "state" : {
            "type" : "string"
          }
        }
      }
    },
    "methods" : {
      "ScoreEndpoint.get" : {
        "response" : {
          "$ref" : "Score"
        }
      },
      "ScoreEndpoint.list" : {
        "response" : {
          "$ref" : "Scores"
        }
      },
      "ScoreEndpoint.insert" : {
        "request" : {
          "$ref" : "Score"
        },
        "response" : {
          "$ref" : "Score"
        }
      },
      "ScoreEndpoint.echo1" : {
        "response" : {
          "$ref" : "Score"
        }
      },
      "ScoreEndpoint.echo2" : {
        "response" : {
          "$ref" : "Score"
        }
      },
      "ScoreEndpoint.echo3" : {
        "response" : {
          "$ref" : "Score"
        }
      },
      "ScoreEndpoint.echo4" : {
        "response" : {
          "$ref" : "Score"
        }
      },
      "ScoreEndpoint.cors" : {
        "response" : {
          "$ref" : "Score"
        }
      },
      "BoardEndpoint.getmove" : {
        "request" : {
          "$ref" : "Board"
        },
        "response" : {
          "$ref" : "Board"
        }
      }
    }
  }
}
`
