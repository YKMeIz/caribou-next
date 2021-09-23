package handler

import (
	"encoding/json"
	"github.com/YKMeIz/KV"
	"github.com/YKMeIz/Pill"
	"github.com/YKMeIz/logc"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

var (
	standardTTL = time.Duration(24 * time.Hour)
	cache       = KV.NewNameSpace(KV.Config{
		DiskLess: true,
	})
)

func getMetadata(id string, norm bool) []byte {
	// handle with cache
	if os.Getenv("CARIBOU_CACHE") == "1" {
		cacheKey := id
		if norm {
			cacheKey += ":NORM"
		}
		b, e := cache.Get([]byte(cacheKey))
		if e == nil {
			return b
		}

		logc.Default("cahce ", cacheKey, " not found in redis")
	}

	i, e := pill.Pixiv(id)
	if e != nil {
		logc.Warning("error happened when grab pixiv illust metadata with id: ", id, "\n", e.Error())
		return nil
	}

	if norm {
		// Normalize URLs
		i.Author.Avatar = strings.ReplaceAll(i.Author.Avatar, "https://i.pximg.net", domain)
		for index, v := range i.Sources {
			i.Sources[index] = strings.ReplaceAll(v, "https://i.pximg.net", domain)
		}

		// Normalize URLs in description
		jumps := regexp.MustCompile(`href=\".+?"`).FindAllString(i.Description, -1)
		for _, jump := range jumps {
			oURL, e := url.PathUnescape(jump[6:])
			if e != nil {
				logc.Warning("error happened when parse url after /jump.php: ", e.Error())
				continue
			}
			i.Description = strings.ReplaceAll(i.Description, jump[6:], oURL[10:])
		}
	}

	b, e := json.Marshal(i)
	if e != nil {
		logc.Warning("error happened when marshal json for metadata of id: ", id, "\n", e.Error())
		return nil
	}

	// handle with cache
	if os.Getenv("CARIBOU_CACHE") == "1" {
		cacheKey := id
		if norm {
			cacheKey += ":NORM"
		}
		e := cache.Put([]byte(cacheKey), b, standardTTL)
		if e == nil {
			return b
		}

		logc.Default("cahce ", cacheKey, " cannot be stored: ", e.Error())
	}

	return b
}

func getSourceURLByName(name string) string {
	id := strings.Split(name, "_")[0]
	i, e := pill.Pixiv(id)
	if e != nil {
		logc.Warning("error happened when grab pixiv illust metadata with id: ", id, "\n", e.Error())
		return ""
	}

	for _, v := range i.Sources {
		if strings.Contains(v, name) {
			return v
		}
	}

	return ""
}
