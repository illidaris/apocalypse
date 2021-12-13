package file

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"

	"github.com/illidaris/apocalypse/log/logger"
)

// MockTempFiles
/**
 * @Description:
 * @param exec
 */
var location = "./testTemp"

func deferFuc(path string) {
	e := PathDelete(path)
	if e != nil {
		logger.Info(e.Error())
	}
}
func MockTempFiles(exec func(paths ...string)) {
	defer func() {
		err := PathDelete(location)
		if err != nil {
			panic(err)
		}
	}()
	tmpPath := path.Join(location, "tmp")
	err := MkdirIfNotExist(tmpPath)
	if err != nil {
		panic(err)
	}
	defer deferFuc(tmpPath)
	file1 := path.Join(location, "1.txt")
	err = ioutil.WriteFile(file1, []byte("123456A"), fs.ModePerm)
	if err != nil {
		panic(err)
	}
	defer func(path string) {
		e := PathDelete(path)
		if e != nil {
			logger.Info(e.Error())
		}
	}(file1)
	file2 := path.Join(tmpPath, "2.txt")
	err = ioutil.WriteFile(file2, []byte("123456B"), fs.ModePerm)
	if err != nil {
		panic(err)
	}
	defer deferFuc(file2)
	exec(tmpPath, file1)
}

// TestCompress
/**
 * @Description:
 * @param t
 */
func TestZipCompress(t *testing.T) {
	compressor := NewCompressor(CompressZip)
	MockTempFiles(func(paths ...string) {
		output := path.Join(location, "tmp.zip")
		defer deferFuc(output)
		files := make([]*os.File, 0)
		for _, p := range paths {
			f, err := os.Open(p)
			if err != nil {
				t.Fatal(err)
			}
			files = append(files, f)
		}
		defer func() {
			for _, f := range files {
				err := f.Close()
				if err != nil {
					t.Log(err.Error())
				}
			}
		}()
		beg := time.Now()
		err := compressor.Compress(output, files...)
		t.Logf("compress cost %dms", time.Since(beg).Milliseconds())
		if err != nil {
			t.Fatal(err)
		}
		beg2 := time.Now()
		err = compressor.UnCompress(output, "./testTemp/tmp")
		t.Logf("uncompress cost %dms", time.Since(beg2).Milliseconds())
		if err != nil {
			t.Fatal(err)
		}
	})
}
