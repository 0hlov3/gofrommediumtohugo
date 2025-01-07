package converter

type ConvertFunc func(postsHTMLFolder, hugoContentFolder, contentType string)

type Converter interface {
	Convert(postsHTMLFolder, hugoContentFolder, contentType string)
}

type DefaultConverter struct {
	ConvertFunc func(postsHTMLFolder, hugoContentFolder, contentType string)
}

func NewDefaultConverter() *DefaultConverter {
	return &DefaultConverter{
		ConvertFunc: Convert, // Ensure it defaults to the global Convert function
	}
}

func (d *DefaultConverter) Convert(postsHTMLFolder, hugoContentFolder, contentType string) {
	if d.ConvertFunc == nil {
		panic("DefaultConverter: ConvertFunc is not initialized")
	}
	d.ConvertFunc(postsHTMLFolder, hugoContentFolder, contentType)
}
