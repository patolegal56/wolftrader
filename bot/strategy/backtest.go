package strategy

import "go.mongodb.org/mongo-driver/bson/primitive"

type Backtest struct {
	Id         primitive.ObjectID `bson:"_id"`
	StrategyId primitive.ObjectID
	Strategy   rawStrategy
}
