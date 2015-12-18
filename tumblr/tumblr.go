package tumblr

// Tumblr struct
type Tumblr struct {
	Token          string
	TokenSecret    string
	consumerSecret string
	args           map[string]string
}

// NewTumblr new tumblr.
func NewTumblr(consumerKey string, consumerSecret string) *Tumblr {
	args := make(map[string]string)
	args["api_kei"] = consumerKey
	return &Tumblr{
		consumerSecret: consumerSecret,
		args:           args,
	}
}
