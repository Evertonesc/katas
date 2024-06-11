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

		// ignore the directory to build the full working path
		fullPaths := make([]string, len(dc)-1)

		// build the full working path
		for i := 0; i < len(fullPaths); i++ {
			fullPaths[i] = dc[0] + "/" + dc[i+1]
		}

		// for each working path, extract the content of the file and the file name
		for _, v := range fullPaths {
			fc, sv := extractFileAndContent(v)
			// fc := extractContent(v)
			// sv := removeContent(v)

			// check whenever the file content is present in the map
			if kv, ok := fmap[fc]; ok {

				// when it does, append the new file that has this specific content
				kv = append(kv, sv)
				fmap[fc] = kv

				// store all duplicated contents
				if _, y := duplicatedMap[fc]; !y {
					duplicatedMap[fc] = true
				}

			} else {
				fmap[fc] = []string{sv}
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
