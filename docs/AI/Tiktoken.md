# Tiktoken

`tiktoken` is a fast BPE (Byte Pair Encoding) tokenizer created by OpenAI. It is optimized for use with OpenAI's models (such as GPT-3, GPT-4, and GPT-4o) and is written in Rust, which allows it to be exceptionally fast.

## What is a Tokenizer?

Before large language models can process text, the text needs to be broken down into smaller pieces called **tokens**. A tokenizer is responsible for splitting text into tokens (encoding) and converting tokens back into text (decoding).
- 1 token is approximately 4 characters in English.
- 100 tokens are roughly 75 words.

## Why use Tiktoken?

- **Speed:** `tiktoken` is significantly faster than comparable tokenizers (e.g., Hugging Face tokenizers or standard Python implementations) because of its core Rust implementation.
- **Accuracy:** It's the official tokenizer for OpenAI models, ensuring that you accurately calculate the number of tokens in your prompts, helping to track and predict API costs.
- **Offline:** It runs locally without making any API calls to OpenAI.

## Installation

You can install `tiktoken` using `pip`:

```bash
pip install tiktoken
```

## Basic Usage

### 1. Getting an Encoding
You can get the encoding directly by its name, or you can have `tiktoken` determine the correct encoding for a specific model.

```python
import tiktoken

# Get encoding by name
encoding = tiktoken.get_encoding("cl100k_base")

# Get encoding for a specific model
encoding = tiktoken.encoding_for_model("gpt-3.5-turbo")
```

OpenAI uses different encodings for different models:
- `o200k_base`: `gpt-4o`, `gpt-4o-mini`
- `cl100k_base`: `gpt-4`, `gpt-3.5-turbo`, text embedding models
- `p50k_base`: Codex models, `text-davinci-002`, `text-davinci-003`
- `r50k_base` (or `gpt2`): GPT-3 models like `davinci`

### 2. Encoding Text
To convert text into a list of token integers, use the `.encode()` method.

```python
text = "Hello, how are you today?"
tokens = encoding.encode(text)

print(tokens)
# Output: [9906, 11, 1268, 527, 499, 3351, 30]

print(f"Number of tokens: {len(tokens)}")
# Output: Number of tokens: 7
```

### 3. Decoding Tokens
To convert a list of token integers back into text, use the `.decode()` method.

```python
decoded_text = encoding.decode(tokens)

print(decoded_text)
# Output: Hello, how are you today?
```

### 4. Decoding Single Tokens
You can decode individual tokens to see exactly how the text was split. Note that tokens are byte sequences, which is why `.decode_single_token_bytes()` is used.

```python
for token in tokens:
    print(f"Token: {token} -> {encoding.decode_single_token_bytes(token)}")

# Output:
# Token: 9906 -> b'Hello'
# Token: 11 -> b','
# Token: 1268 -> b' how'
# Token: 527 -> b' are'
# Token: 499 -> b' you'
# Token: 3351 -> b' today'
# Token: 30 -> b'?'
```

## Special Tokens

Some models use special tokens for formatting, such as `<|endoftext|>`. By default, `encode()` throws an error if it encounters a special token in the text to prevent injection attacks. You can allow them explicitly:

```python
encoding.encode(
    "Hello <|endoftext|>",
    allowed_special={"<|endoftext|>"}
)
# Or allow all special tokens
encoding.encode(
    "Hello <|endoftext|>",
    allowed_special="all"
)
```

## References
- [Tiktoken GitHub Repository](https://github.com/openai/tiktoken)
- [OpenAI Tokenizer Web App](https://platform.openai.com/tokenizer)
