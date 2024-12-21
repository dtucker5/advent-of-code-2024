package main

import (
	"os"
	"strconv"
)

type space struct {
	index int
	size  int
}

func main() {
	input, err := os.ReadFile("Day 9: Disk Fragmenter/Part 1/input.txt")
	if err != nil {
		panic(err)
	}
	fs := newFilesystem(string(input))
	checksum := fs.compact()
	println(checksum) // 7854700201 not right
}

type filesystem struct {
	emptyBlockIndices []int64
	spaces            []space
	fileMap           map[int64]int64
	fileSizes         map[int64]int64
	totalUsage        int64
	disk              blocks
}

func newFilesystem(diskMap string) *filesystem {
	disk := make(blocks, 0)
	isFile := true
	fileId := int64(0)
	spaces := make([]space, 0)         // List of free space block indices and sizes
	fileMap := make(map[int64]int64)   // Map of file blocks (file ID -> index)
	fileSizes := make(map[int64]int64) // Map of file sizes (file ID -> size)
	totalUsage := int64(0)
	emptyBlockIndices := make([]int64, 0)
	totalBlockCount := int64(0)
	for _, char := range diskMap {
		if isFile {
			fileSize, _ := strconv.ParseInt(string(char), 10, 64)
			fileMap[fileId] = int64(len(disk))
			fileSizes[fileId] = fileSize
			for i := 0; i < int(fileSize); i++ {
				disk = append(disk, block{fileId: fileId})
				totalBlockCount++
			}
			fileId++
			totalUsage += fileSize
		} else {
			freeSpaceSize, _ := strconv.Atoi(string(char))
			spaces = append(spaces, space{index: len(disk), size: freeSpaceSize})
			for i := 0; i < freeSpaceSize; i++ {
				disk = append(disk, block{fileId: -1})
				totalBlockCount++
				emptyBlockIndices = append(emptyBlockIndices, totalBlockCount-1)
			}
		}
		isFile = !isFile
	}
	//println(disk.sprint())
	return &filesystem{
		fileMap:           fileMap,
		spaces:            spaces,
		fileSizes:         fileSizes,
		totalUsage:        totalUsage,
		disk:              disk,
		emptyBlockIndices: emptyBlockIndices,
	}
}

func (fs *filesystem) compact() int64 {
	isComplete := false
	for {
		fileId := int64(len(fs.fileMap))
		for fileId := fileId - 1; fileId >= 0; fileId-- {
			//println(fileId)
			fileIndex := fs.fileMap[fileId]
			fileSize := fs.fileSizes[fileId]
			for i := int64(0); i < fileSize; i++ {
				nextFreeSpaceIndex := fs.getNextFreeSpaceIndex()
				if nextFreeSpaceIndex == -1 || nextFreeSpaceIndex == int64(len(fs.disk)) {
					isComplete = true
					break
				}
				fs.disk[nextFreeSpaceIndex] = fs.disk[fileIndex+fileSize-1-i]
				fs.disk[fileIndex+fileSize-1-i] = block{fileId: -1}
				//println(string(fs.disk.sprint()), nextFreeSpaceIndex, fileIndex+fileSize-1-i)
				if fs.getCompactedLen() == fs.totalUsage {
					isComplete = true
					break
				}
			}
			if isComplete {
				break
			}
		}
		if isComplete {
			break
		}
	}
	checksum := int64(0)
	for i, v := range fs.disk {
		if v.fileId == -1 {
			break
		}
		checksum += int64(i) * v.fileId
	}
	return checksum
}

func (fs *filesystem) getNextFreeSpaceIndex() int64 {
	//fmt.Println(fs.emptyBlockIndices)
	if len(fs.emptyBlockIndices) == 0 {
		return -1
	}
	emptyBlockIndex := fs.emptyBlockIndices[0]
	fs.emptyBlockIndices = fs.emptyBlockIndices[1:]
	return emptyBlockIndex

	//nextFreeSpaceIndex := fs.spaces[0].index
	//fs.spaces[0].index++
	//fs.spaces[0].size--
	//if fs.spaces[0].size == 0 {
	//	fs.spaces = fs.spaces[1:]
	//}
	//return nextFreeSpaceIndex
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
