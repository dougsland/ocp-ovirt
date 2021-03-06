//
// Copyright (c) 2020 huihui <huihui.fu@cs2c.com.cn>.
// Copyright (c) 2020 Douglas Landgraf <dougsland@redhat.com>
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
	inputRawURL := "https://engine.mydomain.home/ovirt-engine/api"

	conn, err := ovirtsdk4.NewConnectionBuilder().
		URL(inputRawURL).
		Username("admin@internal").
		Password("mysuperpassword").
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

	tagsService := conn.SystemService().TagsService()
	resp, err := tagsService.List().Send()
	if err != nil {
		fmt.Printf("Failed to get tag list, reason: %v\n", err)
		return
	}

	if tagSlice, ok := resp.Tags(); ok {
		for _, tag := range tagSlice.Slice() {
			fmt.Printf("Tag: (")
			if tagName, ok := tag.Name(); ok {
				fmt.Printf(" name: %v", tagName)
			}
			if tagDesc, ok := tag.Description(); ok {
				fmt.Printf(" desc: %v", tagDesc)
			}
			fmt.Println(")")
		}
	}

}
