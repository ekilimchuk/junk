package util

import (
	"archive/tar"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func addFileToTar(tarWriter *tar.Writer, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if stat, err := file.Stat(); err == nil {
		header := new(tar.Header)
		header.Name = filePath
		header.Size = stat.Size()
		header.Mode = int64(stat.Mode())
		header.ModTime = stat.ModTime()
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}
		if _, err := io.Copy(tarWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func Tar(src string, dst string) error {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	tarFile, err := os.Create(fmt.Sprintf("./%s", dst))
	if err != nil {
		return err
	}
	tarWriter := tar.NewWriter(tarFile)
	defer tarWriter.Close()
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if err := addFileToTar(tarWriter, fmt.Sprintf("%s/%s", src, file.Name())); err != nil {
			return err
		}
	}
	return nil
}
