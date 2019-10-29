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
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "s2o",
	Short: "string -> json, output",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("param error, please check and try again")
			return
		}
		// "{\"workNumber\":\"0000\",\"name\":\"zhangpeng\",\"email\":\"zhangpeng.0304@aliyun.com\",\"personalInfo\":{\"age\":18,\"phone\":\"00000000000\",\"maritalStatus\":true,\"address\":\"Beijing\"},\"workInfo\":{\"jobs\":\"RD\",\"position\":\"iOS Developer\",\"department\":\"HuanXiong\"}}"
		jsonString := strings.TrimSpace(string(args[0]))

		if !json.Valid([]byte(jsonString)) {
			fmt.Println("invalid json")
		}

		fmt.Println("---------------------------")

		var m map[string]interface{}
		/// 反序列化
		err := json.Unmarshal([]byte(jsonString), &m)
		if err != nil {
			fmt.Println(err)
			return
		}

		/// 序列化
		data, err := json.MarshalIndent(m, "", "    ")
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)
}
