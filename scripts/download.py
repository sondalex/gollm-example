from argparse import ArgumentParser

from huggingface_hub import hf_hub_download


def cli():
    parser = ArgumentParser()
    parser.add_argument(
        "--hfrepo",
        dest="hfrepo",
        default="TheBloke/Mistral-7B-Instruct-v0.1-GGUF",
        required=False,
        help="Name of Hugging Face repository"
    )
    parser.add_argument(
        "--hffile",
        dest="hffile",
        default="mistral-7b-instruct-v0.1.Q4_K_M.gguf",
        required=False,
        help="Name of the file to download"
    )
    parser.add_argument("--odir", 
                        dest="odir", 
                        required=True,
                        help="Directory where to save the gguf file"
                        )
    return parser


def download(repo, filename, local_dir):
    hf_hub_download(repo_id=repo, filename=filename, local_dir=local_dir)


def main():
    parser = cli()
    args = parser.parse_args()
    download(args.hfrepo, args.hffile, args.odir)


if __name__ == "__main__":
    import sys

    sys.exit(main())
