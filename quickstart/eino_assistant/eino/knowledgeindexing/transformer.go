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

package knowledgeindexing

import (
	"context"

	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/markdown"
	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/recursive"
	"github.com/cloudwego/eino-ext/components/document/transformer/splitter/semantic"
	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/schema"
	"github.com/wangle201210/go-rag/server/core/common"
)

// newDocumentTransformer component initialization function of node 'MarkdownSplitter' in graph 'KnowledgeIndexing'
func newDocumentTransformer(ctx context.Context) (tfr document.Transformer, err error) {
	// TODO Modify component configuration here.
	config := &markdown.HeaderConfig{
		Headers: map[string]string{
			"#":  "title",
			"##": "h2",
		},
		TrimHeaders: false}
	tfr, err = markdown.NewHeaderSplitter(ctx, config)
	if err != nil {
		return nil, err
	}
	return tfr, nil
}

// newDocumentTransformer component initialization function of node 'MarkdownSplitter' in graph 'KnowledgeIndexing'
func newSemanticTransformer(ctx context.Context) (tfr document.Transformer, err error) {
	// 初始化分割器
	embeddingModel, err := newEmbedding(ctx)
	if err != nil {
		return nil, err
	}

	splitter, err := semantic.NewSplitter(ctx, &semantic.Config{
		Embedding:    embeddingModel,
		BufferSize:   2,
		MinChunkSize: 100,
		Separators:   []string{"\n", ".", "?", "!"},
		Percentile:   0.9,
	})

	if err != nil {
		return nil, err
	}
	return splitter, nil
}

// newDocumentTransformer component initialization function of node 'DocumentTransformer3' in graph 'rag'
func newDocumentTransformerV2(ctx context.Context) (tfr document.Transformer, err error) {
	trans := &transformer{}
	// 递归分割
	config := &recursive.Config{
		ChunkSize:   1000, // 每段内容1000字
		OverlapSize: 100,  // 有10%的重叠
		Separators:  []string{"\n", "。", "?", "？", "!", "！"},
		LenFunc: func(s string) int {
			// eg: 使用 unicode 字符数而不是字节数
			return len([]rune(s))
		},
	}
	recTrans, err := recursive.NewSplitter(ctx, config)
	if err != nil {
		return nil, err
	}
	// md 文档特殊处理
	mdTrans, err := markdown.NewHeaderSplitter(ctx, &markdown.HeaderConfig{
		Headers:     map[string]string{"#": common.Title1, "##": common.Title2, "###": common.Title3},
		TrimHeaders: false,
	})
	if err != nil {
		return nil, err
	}
	trans.recursive = recTrans
	trans.markdown = mdTrans
	return trans, nil
}

type transformer struct {
	markdown  document.Transformer
	recursive document.Transformer
}

func (x *transformer) Transform(ctx context.Context, docs []*schema.Document, opts ...document.TransformerOption) ([]*schema.Document, error) {
	isMd := false
	for _, doc := range docs {
		// 只需要判断第一个是不是.md
		if doc.MetaData["_extension"] == ".md" {
			isMd = true
			break
		}
	}
	if isMd {
		return x.markdown.Transform(ctx, docs, opts...)
	}
	return x.recursive.Transform(ctx, docs, opts...)
}
