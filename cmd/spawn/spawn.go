package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spectreman/tokyo/args"
	"github.com/spectreman/tokyo/lax"
	"github.com/spectreman/tokyo/spawn"
)

func main() {
	args := args.New(os.Args)

	if args.Size() != 2 {
		fmt.Fprintf(os.Stderr, "Needs MaxSpawn (0 => Spawn All, Reads commands from Stdin)\n")
		os.Exit(1)
	}

	max := int(lax.ParseUint32(args.Get(1)))

	spawner := spawn.NewSpawner(max)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		cmd := scanner.Text()
		spawner.Add(cmd)
	}

	spawner.Run()
}
