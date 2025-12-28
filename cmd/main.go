package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
	"passgen/internal/generator"
)

func expandCombinedFlags(args []string) []string {
	var result []string
	result = append(result, args[0]) // program name

	for _, arg := range args[1:] {
		// Only process combined short flags like -uls
		if len(arg) > 1 && arg[0] == '-' && arg[1] != '-' {
			for _, ch := range arg[1:] {
				result = append(result, "-"+string(ch))
			}
		} else {
			result = append(result, arg)
		}
	}

	return result
}

var version = "dev"

func main() {
	var (
		length        int
		upper         bool
		lower         bool
		digits        bool
		symbols       bool
		count         int
		clipboardFlag bool
		noAmbiguous   bool
		showVersion   bool
	)
	flag.BoolVar(&showVersion, "version", false, "Show version information")
	flag.BoolVar(&showVersion, "v", false, "Show version information (shorthand)")

	flag.IntVar(&length, "length", 12, "Password length")
	flag.IntVar(&length, "L", 12, "Password length (shorthand)")

	flag.BoolVar(&upper, "upper", false, "Include uppercase letters")
	flag.BoolVar(&upper, "u", false, "Include uppercase letters (shorthand)")

	flag.BoolVar(&lower, "lower", false, "Include lowercase letters")
	flag.BoolVar(&lower, "l", false, "Include lowercase letters (shorthand)")

	flag.BoolVar(&digits, "digits", false, "Include digits")
	flag.BoolVar(&digits, "d", false, "Include digits (shorthand)")

	flag.BoolVar(&symbols, "symbols", false, "Include symbols")
	flag.BoolVar(&symbols, "s", false, "Include symbols (shorthand)")

	flag.IntVar(&count, "count", 1, "Number of passwords to generate")
	flag.IntVar(&count, "c", 1, "Number of passwords to generate (shorthand)")

	flag.BoolVar(&clipboardFlag, "clipboard", false, "Copy password to clipboard")
	flag.BoolVar(&clipboardFlag, "C", false, "Copy password to clipboard (shorthand)")

	flag.BoolVar(&noAmbiguous, "no-ambiguous", false, "Exclude ambiguous characters")
	flag.BoolVar(&noAmbiguous, "A", false, "Exclude ambiguous characters (shorthand)")

	flag.Usage = func() {
		fmt.Println("PassGen - Secure Password Generator (Windows CLI)")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  passgen [options]")
		fmt.Println()
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  passgen --length 16 --upper --lower --digits")
		fmt.Println("  passgen --length 20 --upper --lower --digits --symbols")
		fmt.Println("  passgen --length 12 --upper --lower --digits --no-ambiguous")
		fmt.Println("  passgen --length 16 --upper --lower --clipboard")
		fmt.Println("  passgen --length 10 --upper --digits --count 5")
		fmt.Println("  passgen -L 16 -u -l -d -s -C -A")
		fmt.Println("  passgen -L 16 -u -l -d -s -c 10 -C -A")
		fmt.Println("  passgen -L 8 -udC -c 10")
	}

	os.Args = expandCombinedFlags(os.Args)

	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	if showVersion {
		fmt.Printf("passgen version %s\n", version)
		return
	}

	if count <= 0 {
		fmt.Println("Error: count must be greater than 0")
		os.Exit(1)
	}

	fmt.Println("Generated Passwords:")

	var firstPassword string

	for i := 1; i <= count; i++ {
		password, err := generator.Generate(
			length,
			upper,
			lower,
			digits,
			symbols,
			noAmbiguous,
		)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		if i == 1 {
			firstPassword = password
		}
		fmt.Printf("%d) %s\n", i, password)
	}

	if clipboardFlag {
		err := clipboard.WriteAll(firstPassword)
		if err != nil {
			fmt.Println("Failed to copy to clipboard:", err)
			os.Exit(1)
		}
		fmt.Println("\nPassword copied to clipboard ✔")
	}

}
