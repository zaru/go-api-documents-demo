FORMAT: 1A

# Go API ドキュメント デモ

# Group グループです

## ユーザ [/users]

### 一覧 [GET]

  + Response 200 (application/json)

    + Body
      [
          {
              "question": "Favourite programming language?",
              "choices": [
                  {
                      "choice": "Swift",
                      "url": "/questions/1/choices/1",
                      "votes": 2048
                  }, {
                      "choice": "Python",
                      "url": "/questions/1/choices/2",
                      "votes": 1024
                  }, {
                      "choice": "Objective-C",
                      "url": "/questions/1/choices/3",
                      "votes": 512
                  }, {
                      "choice": "Ruby",
                      "url": "/questions/1/choices/4",
                      "votes": 256
                  }
              ]
          }
      ]

    + Schema

      {
        "type": "array",
        "$schema": "http://json-schema.org/draft-04/schema#",
        "description": "",
        "minItems": 1,
        "uniqueItems": true,
        "items": {
          "type": "object",
          "required": [
            "question",
            "choices"
          ],
          "properties": {
            "question": {
              "type": "string",
              "minLength": 1
            },
            "choices": {
              "type": "array",
              "uniqueItems": true,
              "minItems": 1,
              "items": {
                "required": [
                  "choice",
                  "url",
                  "votes"
                ],
                "properties": {
                  "choice": {
                    "type": "string",
                    "minLength": 1
                  },
                  "url": {
                    "type": "string",
                    "minLength": 1
                  },
                  "votes": {
                    "type": "number"
                  }
                }
              }
            }
          }
        }
      }

### Create a New Question [POST]

You may create your own question using this action. It takes a JSON object
containing a question and a collection of answers in the form of choices.

  + Request (application/json)

    + Body

      {
        "question": "Favourite language?"
        "choices": [
          "Swift",
          "Objective-C"
        ]
      }

    + Schema

      <!-- include(schema.json) -->

### Create a New Question [POST]

You may create your own question using this action. It takes a JSON object
containing a question and a collection of answers in the form of choices.

  + Request (application/json)

    + Attributes

      + question: Favourite Language? (string, required)
      + choices: Swift, `Objective-C` (array, required)
