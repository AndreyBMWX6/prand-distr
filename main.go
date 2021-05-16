package main

import (
	"bufio"
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	//@note: all flags are pointers to values
	filename := flag.String("file", "", "Name of file with student names")
	numbilets := flag.Uint("numbilets", 0, "Amount of bilets for distribution")
	parameter := flag.Int("parameter", 0, "Parameter changing distribution")
	flag.Parse()

	// Command arguments errors processing
	if *filename == "" || *filename == "--numbilets" {
		log.Fatal("Command arguments error: No student names file")
	}
	if *numbilets <= 0 {
		log.Fatal("Value error: `numbilets` argument should be positive (> 0)")
	}
	// using abs to avoid type conversion errors
	if *parameter < 0 {
		*parameter = int(math.Abs(float64(*parameter)))
	}

	// slice used to mark, which numbers are already used
	// default values are false
	bilets := make([]bool, *numbilets)

	// converting distribution parameter to []byte
	// we will append it to student name []byte slices later
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, uint16(*parameter))
	if err != nil {
		fmt.Println(err)
	}

	// opening file
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// reading line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := scanner.Text()
		data := append([]byte(name), buff.Bytes()...)
		hash := sha1.Sum(data)

		/* indexes for cases of big repetitions of used bilet numbers amount
		   `i` and `j` indexes are used for nested loop in which we search for bytes combination for hash
		   to get unused bilet number
		   `k` index is used in case we haven't found any combination.
		    In that case we change position of number byte slice in data slice used for generating hash
		*/
		i := 0
		j := 1
		k := 0
		for {
			// 2 bytes is enough for 2^16 bilets
			hashNum := uint16(hash[i]) << 8 | uint16(hash[j])
			num := hashNum % uint16(*numbilets)
			if bilets[num] == false {
				bilets[num] = true
				fmt.Println(name, num + 1)
				break
			} else {
				j++
			}

			// changing indexes to avoid out of range exception
			if j == len(hash) {
				if i == len(hash) {
					k++
					byteName := []byte(name)
					data := append(byteName[:len(byteName) - k], buff.Bytes()...)
					data = append(byteName, byteName[len(byteName) - k:]...)
					hash = sha1.Sum(data)
				} else {
					i++
					j = i + 1
				}
			}
		}
	}
}
