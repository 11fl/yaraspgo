package telegram

type Updates struct {
	Ok     bool `json:"ok,omitempty"`
	Result []struct {
		UpdateID int `json:"update_id,omitempty"`
		Message  struct {
			MessageID int    `json:"message_id,omitempty"`
			Text      string `json:"text,omitempty"`
			From      struct {
				Id       int    `json:"id,omitempty"`
				Username string `json:"username,omitempty"`
			} `json:"from,omitempty"`
		} `json:"message,omitempty"`
	} `json:"result,omitempty"`
}

type Offset struct {
	Offset int `json:"offset,omitempty"`
}

// type SendMsg struct {
// 	ChatID int    `json:"chat_id,omitempty"`
// 	Text   string `json:"text,omitempty"`
// }

type Message struct {
	Chat_Id int    `json:"chat_id,omitempty"`
	Text    string `json:"text,omitempty"`
}
