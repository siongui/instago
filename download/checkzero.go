package igdl

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/siongui/instago"
)

func (m *IGDownloadManager) ReDownload(path string, info os.FileInfo) (err error) {
	if strings.Contains(info.Name(), "-post-") {
		fmt.Println("Post: ", path)
		code := instago.ExtractPostCodeFromFilename(info.Name())
		if code != "" {
			fmt.Println(code)
			//_, err = DownloadPostNoLogin(code)
			_, err = m.DownloadPost(code)
			return
		}
	}
	if strings.Contains(info.Name(), "-story-") {
		fmt.Println("Story", path)
		_, id := instago.ExtractUsernameIdFromFilename(filepath.Base(path))
		return m.downloadUserReelMedia(id)
	}
	return
}

func moveFile(oldpath, dir string) (err error) {
	filename := filepath.Base(oldpath)
	newpath := filepath.Join(dir, filename)
	return os.Rename(oldpath, newpath)
}

// For some reasons, wget will leave files with 0 size because of download
// failure. This function try to find zero files and re-download them.
// dir1 is the directory which contains downloaded files, and the zero files
// will be moved to dir2.
func (m *IGDownloadManager) CheckZero(dir1, dir2 string) (err error) {
	err = filepath.Walk(dir1, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.Mode().IsRegular() && info.Size() == 0 {
			moveFile(path, dir2)
			m.ReDownload(path, info)
		}

		return nil
	})
	return
}
