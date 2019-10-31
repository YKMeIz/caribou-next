package handler

import (
	"encoding/json"
	"github.com/YKMeIz/Pill"
	"github.com/YKMeIz/logc"
	"net/url"
	"regexp"
	"strings"
)

func getMetadata(id string) []byte {
	i, e := pill.Pixiv(id)
	if e != nil {
		logc.Warning("error happened when grab pixiv illust metadata with id:", id+"\n"+e.Error())
		return nil
	}

	i.Author.Avatar = strings.ReplaceAll(i.Author.Avatar, "https://i.pximg.net", domain)
	for index, v := range i.Sources {
		i.Sources[index] = strings.ReplaceAll(v, "https://i.pximg.net", domain)
	}

	jumps := regexp.MustCompile(`href=\".+?"`).FindAllString(i.Description, -1)
	for _, jump := range jumps {
		oURL, e := url.PathUnescape(jump[6:])
		if e != nil {
			logc.Warning("error happened when parse url after /jump.php:", e.Error())
			continue
		}
		i.Description = strings.ReplaceAll(i.Description, jump[6:], oURL[10:])
	}

	b, e := json.Marshal(i)
	if e != nil {
		logc.Warning("error happened when marshal json for metadata of id:", id+"\n"+e.Error())
		return nil
	}

	return b
}

func getSourceURLByName(name string) string {
	id := strings.Split(name, "_")[0]
	i, e := pill.Pixiv(id)
	if e != nil {
		logc.Warning("error happened when grab pixiv illust metadata with id:", id+"\n"+e.Error())
		return ""
	}

	for _, v := range i.Sources {
		if strings.Contains(v, name) {
			return v
		}
	}

	return ""
}
