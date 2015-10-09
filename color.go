package main

import (
        "github.com/fatih/color"
)

func applyColor(b bool) {
        if ! b {
                color.NoColor = true
        }
}

func bBlue(s string) string {
        return color.New(color.FgBlue).SprintFunc()(s)
}

func bGreen(s string) string {
        return color.New(color.FgGreen, color.Bold).SprintFunc()(s)
}

func bMagenta(s string) string {
        return color.New(color.FgMagenta).SprintFunc()(s)
}

func bCyan(s string) string {
        return color.New(color.FgCyan, color.Bold).SprintFunc()(s)
}

func bRed(s string) string {
        return color.New(color.FgRed, color.Bold).SprintFunc()(s)
}

func bYellow(s string) string {
        return color.New(color.FgYellow, color.Bold).SprintFunc()(s)
}