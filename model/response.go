package model

type Response struct {
	Message string      `json:"message"`
	Status  bool        `json:"status"`
	Data    interface{} `json:"data"`
}

type GetVid struct {
	Videos   []Video `json:"videos"`
	NextPage bool    `json:"next"`
}
