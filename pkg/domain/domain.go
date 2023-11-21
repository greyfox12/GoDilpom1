package domain

// Для храниния различных типов данных для передачи

// Стртура для хранения УЗ
type TRegister struct {
	Login        string `json:"login"`
	Password     string `json:"password"`
	PasswordHash string
}

// Структура для балланса
type TBallance struct {
	Current   float32 `json:"current"`
	Withdrawn float32 `json:"withdrawn"`
}

// Список нарядoв
type TOrders struct {
	Number     string  `json:"number"`
	Status     string  `json:"status"`
	Accrual    float32 `json:"accrual,omitempty"`
	UploadedAt string  `json:"uploaded_at"`
}

// Список Списаний балов
type TWithdrawals struct {
	Order       string  `json:"order"`
	Sum         float32 `json:"sum"`
	ProcessedAt string  `json:"processed_at"`
}

// Данные для системы начисления баллов
type TAccrualReq struct {
	Order      string  `json:"order"`
	Status     string  `json:"status"`
	Accrual    float32 `json:"accrual"`
	OrdResetCn int
}

type TPostWithdraw struct {
	Order string  `json:"order"`
	Sum   float32 `json:"sum"`
}
