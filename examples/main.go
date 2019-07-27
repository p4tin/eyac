package main

import (
	"fmt"

	"github.com/p4tin/eyac"
)

type Settings struct {
	AppName     string `yaml:"app_name"`
	Environment string `yaml:"environment"`
	Mongo       struct {
		Host              string `yaml:"host"`
		UseSsl            string `yaml:"use_ssl"`
		CertPath          string `yaml:"cert_path"`
		Rootca            string `yaml:"rootca"`
		ClientKey         string `yaml:"client_key"`
		ClientCertificate string `yaml:"client_certificate"`
		Database          string `yaml:"database"`
		ProductCollection string `yaml:"product_collection"`
		CartCollection    string `yaml:"cart_collection"`
		OrdersCollection  string `yaml:"orders_collection"`
	} `yaml:"mongo"`
	Aws struct {
		URL             string `yaml:"url"`
		ClientID        string `yaml:"client_id"`
		RegionID        string `yaml:"region_id"`
		AccessKeyID     string `yaml:"access_key_id"`
		SecretAccessKey string `yaml:"secret_access_key"`
		InventoryQueue  string `yaml:"inventory_queue"`
	} `yaml:"aws"`
}

func main() {
	settings := new(Settings)
	eyac.LoadConfig("settings.yaml", &settings)
	fmt.Printf("%+v\n", settings)
}
