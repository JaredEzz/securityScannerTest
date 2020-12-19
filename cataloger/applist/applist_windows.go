//+build windows

package applist

func Files() []os.FileInfo {
	return _files("\\")
}