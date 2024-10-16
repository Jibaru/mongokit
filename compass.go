package mongokit

import (
	"bytes"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// bsonToCompassString converts a bson to compass string representation.
func bsonToCompassString(value interface{}) string {
	switch v := value.(type) {
	case primitive.ObjectID:
		return fmt.Sprintf("ObjectId(\"%s\")", v.Hex())
	case primitive.DateTime:
		t := v.Time().UTC()
		return fmt.Sprintf("ISODate(\"%s\")", t.Format(time.RFC3339))
	case primitive.Timestamp:
		return fmt.Sprintf("Timestamp(%d,%d)", v.T, v.I)
	case primitive.Decimal128:
		return fmt.Sprintf("Decimal128(\"%s\")", v.String())
	case bson.A: // Array
		var buffer bytes.Buffer
		buffer.WriteString("[")
		for i, item := range v {
			if i > 0 {
				buffer.WriteString(",")
			}
			buffer.WriteString(bsonToCompassString(item))
		}
		buffer.WriteString("]")
		return buffer.String()
	case bson.D: // Document
		var buffer bytes.Buffer
		buffer.WriteString("{")
		for i, elem := range v {
			if i > 0 {
				buffer.WriteString(",")
			}
			buffer.WriteString(fmt.Sprintf("\"%s\":%s", elem.Key, bsonToCompassString(elem.Value)))
		}
		buffer.WriteString("}")
		return buffer.String()
	case bson.M: // Map
		var buffer bytes.Buffer
		buffer.WriteString("{")
		first := true
		for key, val := range v {
			if !first {
				buffer.WriteString(",")
			}
			first = false
			buffer.WriteString(fmt.Sprintf("\"%s\":%s", key, bsonToCompassString(val)))
		}
		buffer.WriteString("}")
		return buffer.String()
	case string:
		return fmt.Sprintf("\"%s\"", v)
	case int32, int64, float64, bool:
		return fmt.Sprintf("%v", v)
	case nil:
		return "null"
	default:
		return fmt.Sprintf("%v", v)
	}
}

// PipelineToCompassString converts a mongo.Pipeline to compass representation.
func PipelineToCompassString(pipeline mongo.Pipeline) string {
	var buffer bytes.Buffer
	buffer.WriteString("[")

	for i, stage := range pipeline {
		if i > 0 {
			buffer.WriteString(",")
		}
		buffer.WriteString(bsonToCompassString(stage))
	}

	buffer.WriteString("]")
	return buffer.String()
}
