/*
 * Copyright 2025 CloudWeGo Authors
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

package einoagent

import (
	"context"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
)

func newChatModel(ctx context.Context) (cm model.ToolCallingChatModel, err error) {
	// TODO Modify component configuration here.
	// config := &ark.ChatModelConfig{
	// 	Model:  os.Getenv("ARK_CHAT_MODEL"),
	// 	APIKey: os.Getenv("ARK_API_KEY"),
	// }
	// cm, err = ark.NewChatModel(ctx, config)

	temperature := float32(0.7)
	config := &openai.ChatModelConfig{
		// Model:  os.Getenv("OPENAI_MODEL_NAME"),
		// APIKey: os.Getenv("OPENAI_API_KEY"),
		// BaseURL: os.Getenv("OPENAI_BASE_URL"),
		Model:  "kimi-k2-turbo-preview",
		APIKey: "sk-sv1EiymdFaMiXdDPodLq4OrCVmrJdtuVD94cIDjGlRYQagpe",
		BaseURL: "https://api.moonshot.cn/v1",
		Temperature: &temperature,
	}
	cm, err = openai.NewChatModel(ctx, config)
	if err != nil {
		return nil, err
	}
	return cm, nil
}
