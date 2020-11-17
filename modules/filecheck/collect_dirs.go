package filecheck

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/netdata/go.d.plugin/agent/module"
)

func (fc *Filecheck) collectDirs(mx map[string]int64) {
	curTime := time.Now()
	for _, dirpath := range fc.Dirs.Include {
		fc.collectDir(mx, dirpath, curTime)
	}
}

func (fc *Filecheck) collectDir(mx map[string]int64, dirpath string, curTime time.Time) {
	if !fc.collectedDirs[dirpath] {
		fc.collectedDirs[dirpath] = true
		fc.addDirToCharts(dirpath)
	}

	info, err := os.Stat(dirpath)
	if err != nil {
		if os.IsNotExist(err) {
			mx[dirDimID(dirpath, "exists")] = 0
		} else {
			mx[dirDimID(dirpath, "exists")] = 1
		}
		fc.Debug(err)
		return
	}

	if !info.IsDir() {
		return
	}

	mx[dirDimID(dirpath, "exists")] = 1
	mx[dirDimID(dirpath, "mtime_ago")] = int64(curTime.Sub(info.ModTime()).Seconds())
	if num, err := calcDirNumOfFiles(dirpath); err == nil {
		mx[dirDimID(dirpath, "num_of_files")] = int64(num)
	}
	if fc.Dirs.CollectDirSize {
		if size, err := calcDirSize(dirpath); err == nil {
			mx[dirDimID(dirpath, "size_bytes")] = size
		}
	}
}

func (fc *Filecheck) addDirToCharts(dirpath string) {
	for _, chart := range *fc.Charts() {
		if !strings.HasPrefix(chart.ID, "dir_") {
			continue
		}

		var id string
		switch chart.ID {
		case dirExistenceChart.ID:
			id = dirDimID(dirpath, "exists")
		case dirModTimeChart.ID:
			id = dirDimID(dirpath, "mtime_ago")
		case dirNumOfFilesChart.ID:
			id = dirDimID(dirpath, "num_of_files")
		case dirSizeChart.ID:
			id = dirDimID(dirpath, "size_bytes")
		default:
			fc.Warningf("add dimension: couldn't dim id for '%s' chart (dir '%s')", chart.ID, dirpath)
			continue
		}

		dim := &module.Dim{ID: id, Name: dirpath}

		if err := chart.AddDim(dim); err != nil {
			fc.Warning(err)
			continue
		}
		chart.MarkNotCreated()
	}
}

func dirDimID(dirpath, metric string) string {
	return fmt.Sprintf("dir_%s_%s", dirpath, metric)
}

func calcDirNumOfFiles(dirpath string) (int, error) {
	f, err := os.Open(dirpath)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	// TODO: include dirs?
	names, err := f.Readdirnames(-1)
	return len(names), err
}

func calcDirSize(dirpath string) (int64, error) {
	var size int64
	err := filepath.Walk(dirpath, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return size, err
}
