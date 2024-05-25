# LLM in Go - An Example

![Go version](https://img.shields.io/badge/Go-1.21-blue)


Uses [go-llama.cpp](https://pkg.go.dev/github.com/go-skynet/go-llama.cpp)

## Installation

**Building llama.cpp dependencies**:

```bash
meson build
```

```bash
cd build/ && meson compile && meson compile copy_and_rename_files copy_and_rename_files2 && cd ..
```

**Building the go code**:

```bash
go1.21.10 build
```

## Usage 

```python
./gollm  --model <path to model> --prompt '<Your prompt>' 
```

## Downloading a model

You can download gguf models from Hugging Face. The following script simplifies the downloading process.

```bash
python scripts/download.py --help
usage: download.py [-h] [--hfrepo HFREPO] [--hffile HFFILE] --odir ODIR

options:
  -h, --help       show this help message and exit
  --hfrepo HFREPO  Name of Hugging Face repository
  --hffile HFFILE  Name of the file to download
  --odir ODIR      Directory where to save the gguf file
```
