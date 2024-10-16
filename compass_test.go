package mongokit

import (
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestPipelineToCompassString(t *testing.T) {
	objectID, _ := primitive.ObjectIDFromHex("670ef82ee2cfc8452bea7023")
	dateTime := primitive.NewDateTimeFromTime(time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC))
	timestamp := primitive.Timestamp{T: 1234567890, I: 1}
	decimal, _ := primitive.ParseDecimal128("1234.5678")
	array := bson.A{"element1", 42, true, nil}

	tests := map[string]struct {
		pl       mongo.Pipeline
		expected string
	}{
		"success_match_group_sort": {
			pl: mongo.Pipeline{
				{{Key: "$match", Value: bson.M{"_id": objectID, "status": "active"}}},
				{{Key: "$group", Value: bson.M{"_id": "$category", "total": bson.M{"$sum": 1}}}},
				{{Key: "$sort", Value: bson.M{"total": -1}}},
			},
			expected: `[{"$match":{"_id":ObjectId("670ef82ee2cfc8452bea7023"),"status":"active"}},{"$group":{"_id":"$category","total":{"$sum":1}}},{"$sort":{"total":-1}}]`,
		},
		"success_with_datetime_and_timestamp": {
			pl: mongo.Pipeline{
				{{Key: "$match", Value: bson.M{"createdAt": dateTime, "updatedAt": timestamp}}},
			},
			expected: `[{"$match":{"createdAt":ISODate("2023-10-10T00:00:00Z"),"updatedAt":Timestamp(1234567890,1)}}]`,
		},
		"success_with_array_and_nil": {
			pl: mongo.Pipeline{
				{{Key: "$project", Value: bson.M{"field": array}}},
			},
			expected: `[{"$project":{"field":["element1",42,true,null]}}]`,
		},
		"success_with_decimal": {
			pl: mongo.Pipeline{
				{{Key: "$match", Value: bson.M{"price": decimal}}},
			},
			expected: `[{"$match":{"price":Decimal128("1234.5678")}}]`,
		},
		"success_with_nested_documents": {
			pl: mongo.Pipeline{
				{{Key: "$match", Value: bson.M{"address": bson.M{"city": "New York", "zipcode": 10001}}}},
			},
			expected: `[{"$match":{"address":{"city":"New York","zipcode":10001}}}]`,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := PipelineToCompassString(tt.pl)
			if got != tt.expected {
				t.Errorf("expected %v, got %v\n", tt.expected, got)
			}
		})
	}
}
