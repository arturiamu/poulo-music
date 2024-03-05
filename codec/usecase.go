package codec

type UseCase interface {
	GetAudioFromUrl(url, output string) error
	GetAudioFromFile(input, output string) error
}
