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
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/openai"
	"github.com/cloudwego/eino/components/embedding"
)

func newEmbedding(ctx context.Context) (eb embedding.Embedder, err error) {
	// TODO Modify component configuration here.
	// config := &ark.EmbeddingConfig{
	// 	Model:  os.Getenv("ARK_EMBEDDING_MODEL"),
	// 	APIKey: os.Getenv("ARK_API_KEY"),
	// }
	// eb, err = ark.NewEmbedder(ctx, config)

	dimensions := 1024
	config := &openai.EmbeddingConfig{
		Model:      os.Getenv("OPENAI_EMBEDDING_MODEL_NAME"),
		APIKey:     os.Getenv("OPENAI_API_KEY"),
		BaseURL:    os.Getenv("OPENAI_BASE_URL"),
		Dimensions: &dimensions,
	}
	eb, err = openai.NewEmbedder(ctx, config)
	if err != nil {
		return nil, err
	}
	return eb, nil
}
