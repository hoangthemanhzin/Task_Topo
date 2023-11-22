package n42

//Data model for activate/ deactivate functiion (api backend control UPFs).

type UpfActivateQuery struct {
	UpfIds []string
}

type UpfDeactivateQuery struct {
	UpfIds []string
}

type UpfActivate struct {
	Msg		map[string] string
}

type UpfDeactivate struct {
	Msg		map[string] string
}
