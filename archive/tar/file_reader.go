package tar

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
)

var (
	// ErrFileNotFound indicates that we reached the end of the file without finding the file
	ErrFileNotFound = errors.New("File not found in the archive")

	// ErrNotAFile indicates that there is a match for the path, but it's not a file (eg. it's a directory)
	ErrNotAFile = errors.New("Not a file")
)

// tarFileReader reads a certain file from a TAR archive
// which can also be optionally decompressed during the process
type tarFileReader struct {
	decompressor io.Closer // When there is decompression involved, the decompressor might have to be closed
	archive      *tar.Reader
	found        bool   // Whether the file has been found or not
	file         string // The path to the file inside the archive
}

// NewTarGzFileReader returns a new Reader which reads a specific file from a .tar.gz archive
func NewTarGzFileReader(r io.Reader, f string) (io.ReadCloser, error) {
	gz, err := gzip.NewReader(r)
	if err != nil {
		return nil, err
	}

	return &tarFileReader{
		decompressor: gz,
		archive:      tar.NewReader(gz),

		file: f,
	}, nil
}

func (r *tarFileReader) Read(p []byte) (int, error) {
	// The file is already found, but hasn't been read entirely
	if r.found {
		return r.archive.Read(p)
	}

	// Find the file
	for {
		header, err := r.archive.Next()
		if err == io.EOF {
			return 0, ErrFileNotFound
		} else if err != nil {
			return 0, err
		}

		if header.Name == r.file {
			switch header.Typeflag {
			case tar.TypeReg, tar.TypeRegA:
				r.found = true
				return r.archive.Read(p)
			default: // Avoid scanning the rest of the archive
				return 0, ErrNotAFile
			}
		}
	}
}

func (r *tarFileReader) Close() error {
	if r.decompressor != nil {
		return r.decompressor.Close()
	}

	return nil
}
