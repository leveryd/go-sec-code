package unsafe

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path"
	"strings"
)

func DeCompressTar(srcFilePath, destDirPath string) error {
	// Create destination directory
	os.Mkdir(destDirPath, os.ModePerm)
	fr, err := os.Open(srcFilePath)
	if err != nil {
		return err
	}

	defer func(fr *os.File) {
		err := fr.Close()
		if err != nil {
			panic(err)
		}
	}(fr)

	// Gzip reader
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}

	defer func(gr *gzip.Reader) {
		err := gr.Close()
		if err != nil {

		}
	}(gr)

	// Tar reader
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// End of tar archive
			break
		}
		if err != nil {
			return err
		}
		fmt.Println("UnTarGzing file..." + hdr.Name)
		// Check if it is diretory or file
		if hdr.Typeflag != tar.TypeDir {
			// Get files from archive
			// Create diretory before create file
			err := os.MkdirAll(destDirPath+"/"+path.Dir(hdr.Name), os.ModePerm)
			if err != nil {
				return err
			}
			// Write data to file
			fmt.Println(hdr.Name)
			fw, _ := os.Create(destDirPath + "/" + hdr.Name)
			if err != nil {
				return err
			}
			_, err = io.Copy(fw, tr)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func BadTarDecompress(c *gin.Context) {
	// Get file from request
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Unable to get file from request",
		})
		return
	}
	// Get filename from request
	filename := file.Filename

	if strings.Contains(filename, "..") {
		c.JSON(400, gin.H{
			"message": "Invalid filename",
		})
		return
	}
	// Save file to server
	err = c.SaveUploadedFile(file, "./tmp/"+filename)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Unable to save file to server",
		})
	}
	// UnTarGz file
	err = DeCompressTar("./tmp/"+filename, "./tmp/")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Unable to unTarGz file",
		})
		return
	}

	// Delete file
	err = os.Remove("./tmp/" + filename)
	if err != nil {
		c.JSON(200, gin.H{
			"message": "Unable to delete file",
		})
	}
	c.JSON(200, gin.H{
		"message": "UnTarGz file successfully",
	})
}

// doc/untar.md
