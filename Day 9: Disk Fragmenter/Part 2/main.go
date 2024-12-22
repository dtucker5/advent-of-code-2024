package main

import (
	"os"
	"strconv"
)

func main() {
	input, err := os.ReadFile("Day 9: Disk Fragmenter/Part 2/input.txt")
	if err != nil {
		panic(err)
	}
	fs := newFilesystem(string(input))
	checksum := fs.compact()
	println(checksum) // 8597959775130 (too high)
}

func newFilesystem(diskMap string) *filesystem {
	disk := make(blocks, 0)
	isFile := true
	fileId := int64(0)
	emptyBlocks := make([]chunk, 0)
	files := make([]chunk, 0)

	for _, char := range diskMap {
		size, _ := strconv.ParseInt(string(char), 10, 64)
		id := int64(-1)
		if isFile {
			files = append(files, chunk{start: int64(len(disk)), size: size})
			id = fileId
			fileId++
		} else {
			emptyBlocks = append(emptyBlocks, chunk{start: int64(len(disk)), size: size})
		}
		for i := int64(0); i < size; i++ {
			disk = append(disk, block{fileId: id})
		}
		isFile = !isFile
	}

	return &filesystem{
		disk:        disk,
		emptyBlocks: emptyBlocks,
		files:       files,
	}
}

func (fs *filesystem) compact() int64 {
	for fileId := int64(len(fs.files)) - 1; fileId >= 0; fileId-- {
		//println("Attempting to move file", fileId)
		file := fs.files[fileId]
		nextFreeSpaceIndex := fs.getNextContiguousFreeSpaceToLeft(file)
		if nextFreeSpaceIndex == -1 {
			//println("    No contiguous free space found for file", fileId)
			continue
		}
		//println("    Moving file", fileId, "to index", nextFreeSpaceIndex)
		for i := int64(0); i < file.size; i++ {
			fs.disk[nextFreeSpaceIndex] = fs.disk[file.start+file.size-1-i]
			fs.disk[file.start+file.size-1-i] = block{fileId: -1}
			nextFreeSpaceIndex++
		}
		//println(fs.disk.sprint())
		// Update the file's start position
		fs.files[fileId].start = nextFreeSpaceIndex - file.size
	}
	checksum := int64(0)
	for fileId, file := range fs.files {
		for i := int64(0); i < file.size; i++ {
			checksum += (file.start + i) * int64(fileId)
		}
	}
	return checksum
}

func (fs *filesystem) getNextContiguousFreeSpaceToLeft(file chunk) int64 {
	//fmt.Println(fs.emptyBlocks, n)
	for i, v := range fs.emptyBlocks {
		if v.start >= file.start {
			return -1
		}
		if v.size >= file.size {
			if v.size == file.size {
				fs.emptyBlocks = append(fs.emptyBlocks[:i], fs.emptyBlocks[i+1:]...)
			} else {
				fs.emptyBlocks[i].start += file.size
				fs.emptyBlocks[i].size -= file.size
			}
			return v.start
		}
	}
	return -1
}

func (fs *filesystem) getCompactedLen() int64 {
	compactedLen := int64(0)
	for _, v := range fs.disk {
		if v.fileId == -1 {
			break
		}
		compactedLen++
	}
	return compactedLen
}

type block struct {
	fileId int64 // -1 if empty
}

type blocks []block

func (b blocks) sprint() string {
	s := ""
	for _, v := range b {
		if v.fileId == -1 {
			s += "."
		} else {
			s += strconv.FormatInt(v.fileId, 10)
		}
	}
	return s
}

type chunk struct {
	start int64
	size  int64
}

type filesystem struct {
	disk        blocks
	emptyBlocks []chunk
	files       []chunk
}
