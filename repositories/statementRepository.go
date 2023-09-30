package repositories

type StatementId struct {
	OpeningSession string `json:"openingSession"`
	ProviderName   string `json:"providerName"`
	Shift          int    `json:"shift"`
}

type DailyStatements struct {
	Id                        StatementId `bson:"_id"`
	TotalIncomeBet            float32     `json:"totalIncomeBet"`
	TotalIncomeOpened         float32     `json:"totalIncomeOpened"`
	TotalExpense              float32     `json:"totalExpense"`
	TotalWin                  float32     `json:"totalWin"`
	TotalLose                 float32     `json:"totalLose"`
	TotalStatementBet         float32     `json:"totalStatementBet"`
	TotalStatementOpened      float32     `json:"totalStatementOpened"`
	TotalStatementOutstanding float32     `json:"totalStatementOutstanding"`
	CursorId                  string      `json:"cursorId"`
}