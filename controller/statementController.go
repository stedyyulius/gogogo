package controller

import (
	"context"
	"iseng/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDailyStatements(c *gin.Context, col *mongo.Collection) {

	pipeline := mongo.Pipeline{
		// $match stage
		{
			{Key: "$match", Value: bson.D{
				{Key: "$expr", Value: bson.D{
					{Key: "$and", Value: bson.A{
						bson.D{
							{Key: "$gte", Value: bson.A{
								bson.D{
									{Key: "$dateFromString", Value: bson.D{
										{Key: "dateString", Value: "$openingSession"},
										{Key: "format", Value: "%d-%m-%Y"},
									}},
								},
								bson.D{
									{Key: "$dateFromString", Value: bson.D{
										{Key: "dateString", Value: "15-09-2023"},
										{Key: "format", Value: "%d-%m-%Y"},
									}},
								},
							}},
						},
						bson.D{
							{Key: "$lte", Value: bson.A{
								bson.D{
									{Key: "$dateFromString", Value: bson.D{
										{Key: "dateString", Value: "$openingSession"},
										{Key: "format", Value: "%d-%m-%Y"},
									}},
								},
								bson.D{
									{Key: "$dateFromString", Value: bson.D{
										{Key: "dateString", Value: "15-09-2023"},
										{Key: "format", Value: "%d-%m-%Y"},
									}},
								},
							}},
						},
					}},
				}},
			}},
		},
		// $group stage
		{
			{Key: "$group", Value: bson.D{
				{Key: "_id", Value: bson.D{
					{Key: "openingSession", Value: "$openingSession"},
					{Key: "providerName", Value: "$providerName"},
					{Key: "shift", Value: "$shift"},
				}},
				{Key: "totalIncomeBet", Value: bson.D{
					{Key: "$sum", Value: bson.D{
						{Key: "$cond", Value: bson.A{
							bson.D{
								{Key: "if", Value: bson.D{
									{Key: "$eq", Value: bson.A{"$status", "bet"}},
								}},
							},
							bson.D{
								{Key: "then", Value: "$netAmount"},
							},
							bson.D{
								{Key: "else", Value: 0},
							},
						}},
					}},
				}},
				{Key: "totalIncomeOpened", Value: bson.D{
					{Key: "$sum", Value: bson.D{
						{Key: "$cond", Value: bson.A{
							bson.D{
								{Key: "if", Value: bson.D{
									{Key: "$eq", Value: bson.A{"$status", "opened"}},
								}},
							},
							bson.D{
								{Key: "then", Value: "$netAmount"},
							},
							bson.D{
								{Key: "else", Value: 0},
							},
						}},
					}},
				}},
				{Key: "totalExpense", Value: bson.D{
					{Key: "$sum", Value: bson.D{
						{Key: "$cond", Value: bson.A{
							bson.D{
								{Key: "if", Value: bson.D{
									{Key: "$eq", Value: bson.A{"$wl", "win"}},
								}},
							},
							bson.D{
								{Key: "then", Value: "$winAmount"},
							},
							bson.D{
								{Key: "else", Value: 0},
							},
						}},
					}},
				}},
				{Key: "totalWin", Value: bson.D{
					{Key: "$sum", Value: bson.D{
						{Key: "$cond", Value: bson.A{
							bson.D{
								{Key: "if", Value: bson.D{
									{Key: "$eq", Value: bson.A{"$wl", "win"}},
								}},
							},
							bson.D{
								{Key: "then", Value: 1},
							},
							bson.D{
								{Key: "else", Value: 0},
							},
						}},
					}},
				}},
				{Key: "totalLose", Value: bson.D{
					{Key: "$sum", Value: bson.D{
						{Key: "$cond", Value: bson.A{
							bson.D{
								{Key: "if", Value: bson.D{
									{Key: "$eq", Value: bson.A{"$wl", "lost"}},
								}},
							},
							bson.D{
								{Key: "then", Value: 1},
							},
							bson.D{
								{Key: "else", Value: 0},
							},
						}},
					}},
				}},
				{Key: "totalStatementBet", Value: bson.D{
					{Key: "$sum", Value: bson.D{
						{Key: "$cond", Value: bson.A{
							bson.D{
								{Key: "if", Value: bson.D{
									{Key: "$eq", Value: bson.A{"$status", "bet"}},
								}},
							},
							bson.D{
								{Key: "then", Value: 1},
							},
							bson.D{
								{Key: "else", Value: 0},
							},
						}},
					}},
				}},
				{Key: "totalStatementOpened", Value: bson.D{
					{Key: "$sum", Value: bson.D{
						{Key: "$cond", Value: bson.A{
							bson.D{
								{Key: "if", Value: bson.D{
									{Key: "$eq", Value: bson.A{"$status", "opened"}},
								}},
							},
							bson.D{
								{Key: "then", Value: 1},
							},
							bson.D{
								{Key: "else", Value: 0},
							},
						}},
					}},
				}},
				{Key: "totalStatementOutstanding", Value: bson.D{
					{Key: "$sum", Value: bson.D{
						{Key: "$cond", Value: bson.A{
							bson.D{
								{Key: "$eq", Value: bson.A{"$status", "bet"}},
							},
							bson.D{
								{Key: "then", Value: 1},
							},
							bson.D{
								{Key: "else", Value: -1},
							},
						}},
					}},
				}},
				{Key: "cursorId", Value: bson.D{
					{Key: "$last", Value: "$_id"},
				}},
			}},
		},
		// $sort stage
		{
			{Key: "$sort", Value: bson.D{
				{Key: "_id.openingSession", Value: 1},
				{Key: "_id.providerName", Value: 1},
			}},
		},
	}
	ctx := context.TODO()
	var data, err = col.Aggregate(ctx, pipeline)

	if err != nil {
		log.Fatal(err)
	}

	var response []repositories.DailyStatements
	
	for data.Next(ctx) {
		var result repositories.DailyStatements
		err := data.Decode(&result)
		if err != nil {
			log.Println("Error decoding document:", err)
			continue // Continue processing other documents
		}
		response = append(response, result)
	}


	if err = data.All(context.TODO(), &response); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, response)
}