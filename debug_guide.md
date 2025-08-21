# Redis Debugging Guide for Eino Assistant

## Issue Analysis
You're experiencing issues where queries aren't returning results. Based on the embedding data, I can see:

1. **Vector Dimension Mismatch**: The new embedding shows 1536 dimensions, but the system was configured for 512
2. **Possible Data Not Indexed**: Documents may not have been properly indexed into Redis

## Debugging Steps

### 1. Check Redis Connection
```bash
# Check if Redis is running
redis-cli ping

# Check vector index info
redis-cli FT.INFO eino:doc:vector_index

# List all keys with eino prefix
redis-cli KEYS "eino:doc:*"
```

### 2. Verify Vector Dimensions
```bash
# Check a sample document's vector dimension
redis-cli HGET eino:doc:<document_id> content_vector | wc -w
```

### 3. Check Index Schema
```bash
# Get detailed index information
redis-cli FT.INFO eino:doc:vector_index
```

### 4. Test Vector Search
```bash
# Test with a simple vector search
redis-cli FT.SEARCH eino:doc:vector_index "*=>[KNN 3 @content_vector $vec]" PARAMS 2 vec <your_vector_here> DIALECT 2
```

## Common Issues & Solutions

### Issue 1: Dimension Mismatch
**Problem**: Embedding model outputs 1536 dims, but Redis expects 512
**Solution**: Update dimension in `pkg/redis/redis.go`

### Issue 2: No Documents Indexed
**Problem**: Redis shows no documents
**Solution**: Run indexing process
```bash
cd quickstart/eino_assistant/cmd/knowledgeindexing
go run main.go
```

### Issue 3: Index Not Created
**Problem**: Vector index doesn't exist
**Solution**: Restart the indexing process to create index

## Quick Verification Script

Create a simple Go script to check Redis status:

```go
package main

import (
    "context"
    "fmt"
    "github.com/redis/go-redis/v9"
)

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr: "127.0.0.1:6379",
    })
    
    ctx := context.Background()
    
    // Check connection
    pong, err := rdb.Ping(ctx).Result()
    fmt.Println("Connection:", pong, err)
    
    // Check index
    info, err := rdb.Do(ctx, "FT.INFO", "eino:doc:vector_index").Result()
    fmt.Println("Index info:", info, err)
    
    // Count documents
    keys, err := rdb.Keys(ctx, "eino:doc:*").Result()
    fmt.Println("Document count:", len(keys))
}
```

## Next Steps
1. Run the debugging commands above
2. Check the actual vector dimensions in stored documents
3. Verify the indexing process completed successfully
4. Test with a simple query to isolate the issue