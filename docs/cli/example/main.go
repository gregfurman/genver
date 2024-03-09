package main

import "github.com/sirupsen/logrus"

//go:generate make genver
func main() {
	logrus.Infof("Currently using logrus %s@%s", SirupsenLogrus_Github_Path, SirupsenLogrus_Github_Version)
}
