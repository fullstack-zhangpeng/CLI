// Copyright © 2019 zhangpeng zhangpeng.0304@aliyun.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "display local ip address",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var ip []byte
		var err error
		t := exec.Command("/bin/bash", "-c", "ifconfig | grep 'inet' | grep -v '127.0.0.1' | grep -v 'inet6' | awk '{print $2}'")
		if ip, err = t.Output(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(strings.Trim(string(ip), "\n"))
	},
}

func test(args ...string) {
	fmt.Println(args)
}

func init() {
	rootCmd.AddCommand(ipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
