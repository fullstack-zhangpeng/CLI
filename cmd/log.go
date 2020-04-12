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
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
)

var log = &cobra.Command{
	Use:   "log",
	Short: "获取日志",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("请输入手机号：")
		var phone string
		_, _ = fmt.Scanln(&phone)
		fmt.Println("请输入日期，例：20190807：")
		var date string
		_, _ = fmt.Scanln(&date)
		fmt.Println("请输入平台，例：ios/android/windows：")
		var platform string
		_, _ = fmt.Scanln(&platform)
		fmt.Println("-----------------------------")
		resp, err := http.Get("https://hx.zuoyebang.com/flipped-classin/loginfo/getlivelog?phone=15900242911&date=20190724&os=ios")
		if err != nil {
			// handle error
		}
		//fmt.Println(resp.Body)
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(body)
	},
}

func init() {
	//rootCmd.AddCommand(log)
}
