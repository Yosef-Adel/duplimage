# DuplImage

DuplImage is a fast and efficient command-line tool for finding duplicate images in directories using SHA-256 hash comparison. It recursively scans directories and identifies identical images regardless of their filenames.

## Features

- Fast SHA-256 hash comparison
- Recursive directory scanning
- Support for common image formats (JPG, JPEG, PNG, GIF, BMP)
- Memory-efficient processing
- Simple command-line interface

## Installation

### Prerequisites

- Go 1.16 or higher

### Building from source

```bash
git clone https://github.com/Yosef-Adel/duplimage
cd duplimage
go build -o duplimage
```

## Usage

Basic usage:

```bash
./duplimage -d "/path/to/your/images"
```

### Command Line Options

| Flag | Description                                            |
| ---- | ------------------------------------------------------ |
| -d   | Directory path to scan for duplicate images (required) |

### Example Output

```
Found 2 groups of duplicate images:

Duplicate group 1:
- /photos/vacation/beach.jpg
- /photos/backup/beach_copy.jpg

Duplicate group 2:
- /photos/events/party.png
- /photos/events/party_duplicate.png
```

## Supported Image Formats

- JPEG (.jpg, .jpeg)
- PNG (.png)
- GIF (.gif)
- BMP (.bmp)

## How It Works

1. DuplImage recursively walks through the specified directory
2. For each image file:
   - Calculates SHA-256 hash of the file content
   - Groups files with identical hashes
3. Reports groups of files that share the same hash

## Technical Details

- Uses Go's standard crypto/sha256 package for hashing
- Processes files in chunks to maintain low memory usage
- Ignores non-image files based on file extensions

## Performance

- Memory usage is constant regardless of image size
- Processing time is linear with respect to the total size of all images
- Hash comparisons are performed in-memory for quick results

## Limitations

- Only finds exactly identical files
- Does not detect visually similar images that are not bit-for-bit identical
- Does not account for image metadata differences
- No support for comparing images across different formats

## Future Improvements

- Add support for more image formats
- Implement perceptual hashing for finding similar (not identical) images
- Add options for different hash algorithms
- Add parallel processing for faster scanning
- Add option to automatically handle duplicates (delete/move)

## License

MIT License
