package cli

import (
	"github.com/leobeosab/sharingan/internal/models"
	"github.com/urfave/cli/v2"
)

func GetGlobalFlags(s *models.ScanSettings) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "target",
			Value:       "",
			Usage:       "Target domain",
			Destination: &s.Target,
		},
		&cli.IntFlag{
			Name:        "threads",
			Value:       20,
			Usage:       "Max number of go routines",
			Destination: &s.Threads,
		},
		&cli.BoolFlag{
			Name:        "no-prompt",
			Value:       false,
			Usage:       "Disable prompts and continue without confirmation",
			Destination: &s.NoPrompt,
		},
	}
}

func GetNMapFlags(s *models.ScanSettings) []cli.Flag {
	return []cli.Flag{}
}

func GetDNSFlags(s *models.ScanSettings) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "dns-wordlist",
			Value:       "",
			Usage:       "Wordlist for DNS bruteforcing",
			Destination: &s.DNSWordlist,
		},
		&cli.StringFlag{
			Name:        "root-domain",
			Value:       "",
			Usage:       "Basis for subdomain scanning",
			Destination: &s.RootDomain,
		},
		&cli.BoolFlag{
			Name:        "skip-probe",
			Usage:       "Skips host-up nmap scan",
			Destination: &s.SkipProbe,
		},
		&cli.BoolFlag{
			Name:        "replace-subs",
			Usage:       "used with add subs to replace all subs for program",
			Destination: &s.ReplaceSubs,
		},
		&cli.BoolFlag{
			Name:        "rescan",
			Usage:       "Scans domain regardless of the existance of previous results",
			Destination: &s.Rescan,
		},
	}

}
