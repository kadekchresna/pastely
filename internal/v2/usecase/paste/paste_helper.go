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
	PasteURL                  string `json:"paste_url"`
}

func NewCreatePaste(ExpirationLengthInMinutes int, PasteURL string) CreatePaste {
	return CreatePaste{
		ExpirationLengthInMinutes: ExpirationLengthInMinutes,
		PasteURL:                  PasteURL,
	}
}
