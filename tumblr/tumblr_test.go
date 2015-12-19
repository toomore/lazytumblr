package tumblr

import (
	"log"
	"os"
	"testing"
)

func getTumblr() *Tumblr {
	return NewTumblr(os.Getenv("TUMBLRCONSUMERKEY"), os.Getenv("TUMBLRCONSUMERSECRET"))
}

func TestTumblr_Post(*testing.T) {
	t := getTumblr()
	t.BaseHost = os.Getenv("TUMBLRUSERBASEHOST")
	t.Token = os.Getenv("TUMBLRUSERTOKEN")
	t.TokenSecret = os.Getenv("TUMBLRUSERSECRET")

	args := make(map[string]string)
	args["format"] = "html"
	args["title"] = "Test from tumblr API"
	args["body"] = "<b>Toomore</b> You are the great!!"

	log.Printf("%+v", t.Post(args))

}
