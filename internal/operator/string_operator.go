/*
 * Copyright 2023 ByteDance Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package operator

import (
	"fmt"
	"github.com/bytedance/arishem/internal/tools"
	"regexp"
	"strings"
)

type StringOperator struct {
	opr func(left, right string) bool
}

func StringStartWithOperate(left, right string) bool {
	return strings.HasPrefix(left, right)
}

func StringNotStartWithOperate(left, right string) bool {
	return !strings.HasPrefix(left, right)
}

func StringEndWithOperate(left, right string) bool {
	return strings.HasSuffix(left, right)
}

func StringNotEndWithOperate(left, right string) bool {
	return !strings.HasSuffix(left, right)
}

func StringContainsOperate(left, right string) bool {
	return strings.Contains(left, right)
}

func StringNotContainsOperate(left, right string) bool {
	return !strings.Contains(left, right)
}

func ContainRegexpOperate(left, right string) bool {
	compile, err := regexp.Compile(right)
	if err != nil {
		return false
	}
	return compile.MatchString(left)
}

func NotContainRegexpOperate(left, right string) bool {
	return !ContainRegexpOperate(left, right)
}

func (b *StringOperator) Operate(left, right interface{}) (bool, error) {
	lj, err := tools.ConvToJSONValidStr(left)
	if err != nil {
		return false, fmt.Errorf("left(%T,%v) conver to string err=>%s", left, left, err.Error())
	}
	rj, err := tools.ConvToJSONValidStr(right)
	if err != nil {
		return false, fmt.Errorf("right(%T,%v) convert to string err=>%s", right, right, err.Error())
	}
	return b.opr(lj, rj), nil
}
