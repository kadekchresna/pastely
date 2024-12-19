package paste

type GetPasteParams struct {
	Shortlink string
}

func NewGetPasteParams(Shortlink string) GetPasteParams {
	return GetPasteParams{
		Shortlink: Shortlink,
	}
}
