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
	args["type"] = "photo"
	args["format"] = "html"
	args["title"] = "Test from tumblr API"
	args["body"] = "<b>Toomore</b> You are the great!!"
	args["caption"] = "<b>Toomore</b> You are the great!!"
	// Form source work!
	args["source"] = "https://c1.staticflickr.com/1/759/23392334809_5721ac4bab_o.jpg"
	log.Printf("%+v", t.Post(args, nil))

	// From data work!
	//files := url.Values{}
	//files.Add("data", readFilesBin("/Volumes/RamDisk/45070001_1.jpg").String())
	//files.Add("data", readFilesBin("/Volumes/RamDisk/45070002_1.jpg").String())
	//log.Printf("%+v", t.Post(args, files))

	// From file to base64 work!
	//args["data64"] = readFileToBase64("/Volumes/RamDisk/45070001_1.jpg")
	//log.Printf("%+v", t.Post(args, nil))
}

func TestTumblr_PostPhoto(*testing.T) {
	t := getTumblr()
	t.BaseHost = os.Getenv("TUMBLRUSERBASEHOST")
	t.Token = os.Getenv("TUMBLRUSERTOKEN")
	t.TokenSecret = os.Getenv("TUMBLRUSERSECRET")

	args := make(map[string]string)
	args["caption"] = "<b>Toomore</b> You are the great!!"
	args["source_url"] = "https://www.flickr.com/photos/toomore/23666343391"

	files := []string{"/Volumes/RamDisk/45070001_1.jpg", "/Volumes/RamDisk/45070002_1.jpg"}
	t.PostPhoto(args, files)
}
