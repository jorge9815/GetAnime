package structs

type Anime struct {
	Name               string
	IsCompleted        bool
	DownloadedChapters uint16
}

func (a *Anime) NewAnime(name string, isCompleted bool, downloadedChapters uint16) {
	a.Name = name
	a.IsCompleted = isCompleted
	a.DownloadedChapters = downloadedChapters
}
