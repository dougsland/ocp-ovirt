//
// Copyright (c) 2017 Joey <majunjiev@gmail.com>.
// Copyright (c) 2020 Douglas S Landgraf <dougsland@redhat.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package main

import (
	"fmt"
	"time"

	ovirtsdk4 "github.com/ovirt/go-ovirt"
)

func main() {
	inputRawURL := "https://engine.medogz.home/ovirt-engine/api"

	conn, err := ovirtsdk4.NewConnectionBuilder().
		URL(inputRawURL).
		Username("admin@internal").
		Password("superpass").
		Insecure(true).
		Compress(true).
		Timeout(time.Second * 10).
		Build()
	if err != nil {
		fmt.Printf("Make connection failed, reason: %v\n", err)
		return
	}
	defer conn.Close()

	// To use `Must` methods, you should recover it if panics
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panics occurs, try the non-Must methods to find the reason")
		}
	}()

	// Get the reference to the "templatesservice" service:
	templateService := conn.SystemService().TemplatesService()

	// Use the "list" method of the "template" service to list all the templates of the system:
	templateResponse, err := templateService.List().Search("name=ocp-jrs4q-*").Send()
	if err != nil {
		fmt.Printf("Failed to get template list, reason: %v\n", err)
		return
	}

	if templates, ok := templateResponse.Templates(); ok {
		// Print the datacenter names and identifiers:
		for _, template := range templates.Slice() {
			fmt.Printf("Template")
			service := conn.SystemService().TemplatesService().TemplateService(template.MustId())
			_, err := service.Remove().Send()
			if err != nil {
				fmt.Println("error...")
			}
		}
	}
}
