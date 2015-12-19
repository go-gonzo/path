package path

import (
	"strings"

	"github.com/go-gonzo/util"
	"github.com/omeid/gonzo"
)

func TrimPrefix(prefix string) gonzo.Stage {
	return Rename(func(name string) string {
		return strings.TrimPrefix(name, prefix)
	})
}

func Rename(rename func(string) string) gonzo.Stage {
	return util.Do(func(file gonzo.File) gonzo.File {
		file.FileInfo().SetName(rename(file.FileInfo().Name()))
		return file
	})
}
