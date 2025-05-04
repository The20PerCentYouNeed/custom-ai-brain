import sys
import json
from transformers import XLMRobertaTokenizer

tokenizer = XLMRobertaTokenizer.from_pretrained("xlm-roberta-base")

text = sys.stdin.read().strip()

tokens = tokenizer.tokenize(text)

decoded = tokenizer.convert_tokens_to_string(tokens)
if text != decoded:
    print(f"Warning: Decoded text differs from input: {decoded[:100]}...", file=sys.stderr)

print(json.dumps(tokens))
