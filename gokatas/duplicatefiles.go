package gokatas

import (
	"strings"
)

func findDuplicates(paths []string) [][]string {
	fmap := make(map[string][]string)
	duplicatedMap := make(map[string]bool)

	for i := 0; i < len(paths); i++ {
		// split path to get the directory and its files
		dc := strings.Split(paths[i], " ")

		// ignore the directory by starting iterate at index 1
		for i := 1; i < len(dc); i++ {

			//get the file content and the file itself
			fc, f := extractFileAndContent(dc[i])

			// build the file path without the content
			fp := dc[0] + "/" + f

			// check whenever the file content exists in the file map
			if kv, ok := fmap[fc]; ok {

				// when it does, append the new file that has this specific content
				kv = append(kv, fp)
				fmap[fc] = kv

				// store all duplicated contents
				if _, y := duplicatedMap[fc]; !y {
					duplicatedMap[fc] = true
				}

			} else {
				fmap[fc] = []string{fp}
			}

		}
	}

	// fetch all the files that has duplicated content
	duplicates := make([][]string, len(duplicatedMap))
	i := 0
	for k := range duplicatedMap {
		v, _ := fmap[k]
		duplicates[i] = v
		i++
	}

	return duplicates
}

func extractFileAndContent(filePath string) (string, string) {
	start := strings.Index(filePath, "(")
	end := strings.Index(filePath[start:], ")")

	end += start

	return filePath[start+1 : end], filePath[:start]
}
