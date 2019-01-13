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

func Untar(dst string, r io.Reader) error {
	tr := tar.NewReader(r)
	for {
		header, err := tr.Next()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return nil
		case header == nil:
			continue
		}
		target := fmt.Sprintf("%s/%s", dst, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}
			f.Close()
		}
	}
}
