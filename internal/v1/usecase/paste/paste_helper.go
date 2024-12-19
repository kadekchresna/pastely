package paste

type GetPasteParams struct {
	Shortlink string `query:"shortlink"`
}

func NewGetPasteParams(Shortlink string) GetPasteParams {
	return GetPasteParams{
		Shortlink: Shortlink,
	}
}
