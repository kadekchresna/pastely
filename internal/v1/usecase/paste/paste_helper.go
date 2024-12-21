package paste

type GetPasteParams struct {
	Shortlink string `query:"shortlink"`
}

func NewGetPasteParams(Shortlink string) GetPasteParams {
	return GetPasteParams{
		Shortlink: Shortlink,
	}
}

type CreatePaste struct {
	ExpirationLengthInMinutes int    `json:"expiration_length_in_minutes"`
	PasteContent              string `json:"paste_content"`
}

func NewCreatePaste(ExpirationLengthInMinutes int, PasteContent string) CreatePaste {
	return CreatePaste{
		ExpirationLengthInMinutes: ExpirationLengthInMinutes,
		PasteContent:              PasteContent,
	}
}
