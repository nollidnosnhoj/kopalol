package images

func GetCacheKey(filename string) string {
	return "image:" + filename
}
