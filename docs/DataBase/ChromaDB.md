# ChromaDB

`ChromaDB` (or simply Chroma) is an open-source, AI-native vector database designed to store, retrieve, and search vector embeddings. It is heavily used in Large Language Model (LLM) applications, particularly for implementing Retrieval-Augmented Generation (RAG) and semantic search.

## Key Features

- **Developer-first:** Designed to be simple, lightweight, and easy to set up.
- **In-memory & Persistent:** Runs in-memory for quick testing/prototyping, persists to disk for local development, or runs in a client-server mode for production.
- **Built-in Embeddings:** Automatically handles text embedding generation out of the box using default or custom embedding functions (like OpenAI, Hugging Face, Cohere, etc.).
- **Metadata Filtering:** Supports rich filtering on metadata alongside semantic similarity searches.

## Installation

Install `chromadb` using `pip`:

```bash
pip install chromadb
```

For JS/TS applications, you can install the NPM package:

```bash
npm install chromadb
```

## Running Modes

### 1. In-Memory (Ephemeral)
Suitable for unit tests or transient workloads. Data is lost when the process terminates.

```python
import chromadb
client = chromadb.EphemeralClient()
```

### 2. Persistent Local Storage
Stores the SQLite database and index files on disk in the specified directory.

```python
import chromadb
client = chromadb.PersistentClient(path="./chroma_db")
```

### 3. Client-Server Mode
Runs Chroma as a standalone backend service (e.g., via Docker) and connects using a client.

```bash
# Start Chroma server
chroma run --path ./chroma_db
```

```python
import chromadb
client = chromadb.HttpClient(host="localhost", port=8000)
```

## Core Workflow

### 1. Collections
Chroma organizes data into **Collections** (similar to tables in SQL databases or collections in MongoDB).

```python
# Create or retrieve a collection
collection = client.get_or_create_collection(name="my_collection")
```

### 2. Adding Data
You can add raw text documents to a collection. By default, Chroma will automatically generate embeddings using the `all-MiniLM-L6-v2` Sentence Transformer model.

```python
collection.add(
    documents=[
        "Chroma is a vector database designed for developer productivity.",
        "SQL databases organize data in tables with rows and columns.",
        "Neural networks use layers of nodes to process and learn from data."
    ],
    metadatas=[
        {"topic": "databases"},
        {"topic": "databases"},
        {"topic": "machine_learning"}
    ],
    ids=["doc1", "doc2", "doc3"]
)
```

> [!NOTE]
> If you already have pre-computed embeddings, you can pass them using the `embeddings` parameter:
> ```python
> collection.add(
>     embeddings=[[0.1, 0.2, ...], [0.3, 0.4, ...]],
>     documents=["Doc 1", "Doc 2"],
>     ids=["id1", "id2"]
> )
> ```

### 3. Querying
You can perform similarity searches by passing query texts. Chroma will embed the query text and retrieve the most similar documents based on cosine distance (default).

```python
results = collection.query(
    query_texts=["What database should I use for AI application embeddings?"],
    n_results=2
)

# Print results
for i in range(len(results['documents'][0])):
    print(f"ID: {results['ids'][0][i]}")
    print(f"Document: {results['documents'][0][i]}")
    print(f"Distance: {results['distances'][0][i]}")
    print(f"Metadata: {results['metadatas'][0][i]}\n")
```

### 4. Updating and Deleting
- **Update:** Modify documents, metadata, or embeddings.
- **Upsert:** Insert if the ID doesn't exist, otherwise update.
- **Delete:** Remove entries by ID or metadata filter.

```python
# Update existing entry
collection.update(
    ids=["doc1"],
    documents=["ChromaDB is a super fast open-source vector store."]
)

# Delete entries
collection.delete(ids=["doc2"])
```

## Metadata & Document Filtering

Chroma allows querying with filters using MongoDB-like operator syntax:

```python
results = collection.query(
    query_texts=["AI databases"],
    n_results=5,
    # Filter by metadata attributes
    where={"topic": "databases"},
    # Filter by document contents
    where_document={"$contains": "vector"}
)
```

Supported operators: `$eq`, `$ne`, `$gt`, `$gte`, `$lt`, `$lte`, `$in`, `$nin`, `$and`, `$or`.

## Custom Embedding Functions

You can configure Chroma collections to use alternative embedding models:

### Using OpenAI Embeddings

```python
from chromadb.utils import embedding_functions

openai_ef = embedding_functions.OpenAIEmbeddingFunction(
    api_key="your-api-key",
    model_name="text-embedding-3-small"
)

collection = client.get_or_create_collection(
    name="openai_collection",
    embedding_function=openai_ef
)
```

## References
- [ChromaDB Official Documentation](https://docs.trychroma.com/)
- [ChromaDB GitHub Repository](https://github.com/chroma-core/chroma)
