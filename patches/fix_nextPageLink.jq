# Go JSON Unmarshaler is not compatible with union struct generated from allOf as in example:
#
#          "nextPageLink": {
#            "allOf": [
#              {
#                "$ref": "#/components/schemas/nextPageLink"
#              },
#              {
#                "type": "string",
#                "example": "https://coda.io/apis/v1/docs/AbCDeFGH/acl?pageToken=eyJsaW1pd"
#              }
#            ]
#          }
# Given that allOf here is needed only to provide link to example which does not affect target code it can be safely removed
# so the allOf operator is not needed anymore.
#
# The following structure will be yielded instead of the above example:
#
#          "nextPageLink": {
#            "$ref": "#/components/schemas/nextPageLink"
#          }

.components.schemas[] |=
    if .properties.nextPageLink then
        . + {properties: (.properties + {nextPageLink: {"$ref": "#/components/schemas/nextPageLink"} }) }
    else
        .
    end