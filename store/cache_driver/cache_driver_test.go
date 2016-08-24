package cache_driver_test

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"code.cloudfoundry.org/grootfs/fetcher"
	"code.cloudfoundry.org/grootfs/store"
	"code.cloudfoundry.org/grootfs/store/cache_driver"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("CacheDriver", func() {
	var (
		cacheDriver *cache_driver.CacheDriver
		storePath   string

		logger              *lagertest.TestLogger
		streamBlobCallCount int
		streamBlob          fetcher.StreamBlob
	)

	BeforeEach(func() {
		var err error
		storePath, err = ioutil.TempDir("", "store")
		Expect(err).ToNot(HaveOccurred())
		Expect(os.MkdirAll(filepath.Join(storePath, "cache", "blobs"), 0755)).To(Succeed())

		logger = lagertest.NewTestLogger("cacheDriver")
		cacheDriver = cache_driver.NewCacheDriver(storePath)

		streamBlobCallCount = 0
		streamBlob = func(logger lager.Logger) (io.ReadCloser, error) {
			streamBlobCallCount += 1

			buffer := gbytes.NewBuffer()
			buffer.Write([]byte("hello world"))
			return buffer, nil
		}
	})

	AfterEach(func() {
		Expect(os.RemoveAll(storePath)).To(Succeed())
	})

	Context("when the blob is not cached", func() {
		It("calls the streamBlob function", func() {
			_, err := cacheDriver.Blob(logger, "my-blob", streamBlob)
			Expect(err).ToNot(HaveOccurred())
			Expect(streamBlobCallCount).To(Equal(1))
		})

		It("returns the stream returned by streamBlob", func() {
			stream, err := cacheDriver.Blob(logger, "my-blob", streamBlob)
			Expect(err).ToNot(HaveOccurred())

			contents, err := ioutil.ReadAll(stream)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(contents)).To(Equal("hello world"))
		})

		Context("when the store does not exist", func() {
			BeforeEach(func() {
				cacheDriver = cache_driver.NewCacheDriver("/non/existing/store")
			})

			It("returns an error", func() {
				_, err := cacheDriver.Blob(logger, "my-blob", streamBlob)
				Expect(err).To(MatchError(ContainSubstring("creating cached blob file")))
			})
		})

		It("stores the stream returned by streamBlob in the cache", func() {
			stream, err := cacheDriver.Blob(logger, "my-blob", streamBlob)
			Expect(err).ToNot(HaveOccurred())

			theBlobPath := blobPath(storePath, "my-blob")
			Expect(theBlobPath).To(BeARegularFile())

			// consume the stream first
			_, err = ioutil.ReadAll(stream)
			Expect(err).NotTo(HaveOccurred())

			cachedBlobContents, err := ioutil.ReadFile(theBlobPath)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(cachedBlobContents)).To(Equal("hello world"))
		})

		Context("when streamBlob fails", func() {
			BeforeEach(func() {
				streamBlob = func(logger lager.Logger) (io.ReadCloser, error) {
					return nil, errors.New("failed getting remote stream")
				}
			})

			It("returns the error", func() {
				_, err := cacheDriver.Blob(logger, "my-blob", streamBlob)
				Expect(err).To(MatchError(ContainSubstring("failed getting remote stream")))
			})
		})
	})

	Context("when the blob is cached", func() {
		BeforeEach(func() {
			_, err := cacheDriver.Blob(logger, "my-blob", streamBlob)
			Expect(err).ToNot(HaveOccurred())

			// reset the test counter
			streamBlobCallCount = 0
		})

		It("does not call streamBlob", func() {
			_, err := cacheDriver.Blob(logger, "my-blob", streamBlob)
			Expect(err).ToNot(HaveOccurred())

			Expect(streamBlobCallCount).To(Equal(0))
		})

		Context("but the cached file is not a file", func() {
			BeforeEach(func() {
				theBlobPath := blobPath(storePath, "my-blob")
				Expect(os.Remove(theBlobPath)).To(Succeed())
				Expect(os.MkdirAll(theBlobPath, 0755)).To(Succeed())
			})

			It("returns an error", func() {
				_, err := cacheDriver.Blob(logger, "my-blob", streamBlob)
				Expect(err).To(MatchError(ContainSubstring("exists but it's not a regular file")))
			})
		})

		Context("but it does not have access to the cache", func() {
			BeforeEach(func() {
				Expect(os.RemoveAll(filepath.Join(storePath, "cache", "blobs"))).To(Succeed())
				Expect(os.MkdirAll(filepath.Join(storePath, "cache", "blobs"), 0000)).To(Succeed())
			})

			It("returns an error", func() {
				_, err := cacheDriver.Blob(logger, "my-blob", streamBlob)
				Expect(err).To(MatchError(ContainSubstring("checking if the blob exists")))
			})
		})

		Context("but it does not have access to the cached blob", func() {
			BeforeEach(func() {
				theBlobPath := blobPath(storePath, "my-blob")
				Expect(os.RemoveAll(theBlobPath)).To(Succeed())
				Expect(ioutil.WriteFile(theBlobPath, []byte("hello world"), 000)).To(Succeed())
			})

			It("returns an error", func() {
				_, err := cacheDriver.Blob(logger, "my-blob", streamBlob)
				Expect(err).To(MatchError(ContainSubstring("accessing the cached blob")))
			})
		})
	})
})

func blobPath(storePath, id string) string {
	return filepath.Join(storePath, store.CACHE_DIR_NAME, "blobs", id)
}