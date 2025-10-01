package transformer

type Transformer interface {
	Transform(string) string
}

type SimpleTransformer struct {
	TransformFunction func(string) string
}

func (t SimpleTransformer) Transform(s string) string {
	return t.TransformFunction(s)
}

func Pipe(t1 Transformer, t2 Transformer) Transformer {
	return SimpleTransformer{
		TransformFunction: func(s string) string {
			return t2.Transform(t1.Transform(s))
		},
	}
}
