package limiter_test

import (
	"errors"
	"os/exec"

	limiterpkg "code.cloudfoundry.org/grootfs/store/volume_driver/drax/limiter"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagertest"

	"github.com/cloudfoundry/gunk/command_runner/fake_command_runner"
	. "github.com/cloudfoundry/gunk/command_runner/fake_command_runner/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Btrfs", func() {
	var (
		fakeCommandRunner *fake_command_runner.FakeCommandRunner
		limiter           *limiterpkg.BtrfsLimiter
		logger            lager.Logger
	)

	BeforeEach(func() {
		fakeCommandRunner = fake_command_runner.New()

		limiter = limiterpkg.NewBtrfsLimiter(fakeCommandRunner)

		logger = lagertest.NewTestLogger("drax-limiter")
	})

	Describe("ApplyDiskLimit", func() {

		It("limits the provided volume", func() {
			Expect(limiter.ApplyDiskLimit(logger, "/full/path/to/volume", 1024*1024)).To(Succeed())

			Expect(fakeCommandRunner).Should(HaveExecutedSerially(fake_command_runner.CommandSpec{
				Path: "btrfs",
				Args: []string{"qgroup", "limit", "1048576", "/full/path/to/volume"},
			}))
		})

		Context("when setting the limit fails", func() {
			BeforeEach(func() {
				fakeCommandRunner.WhenRunning(fake_command_runner.CommandSpec{
					Path: "btrfs",
				}, func(cmd *exec.Cmd) error {
					cmd.Stdout.Write([]byte("failed to set btrfs limit"))
					cmd.Stderr.Write([]byte("some stderr text"))

					return errors.New("exit status 1")
				})
			})

			It("forwards the stdout and stderr", func() {
				err := limiter.ApplyDiskLimit(logger, "/full/path/to/volume", 1024*1024)

				Expect(err).To(MatchError(ContainSubstring("failed to set btrfs limit")))
				Expect(err).To(MatchError(ContainSubstring("some stderr text")))
			})
		})
	})

	Describe("DestroyQuotaGroup", func() {
		It("destroys the qgroup for the path", func() {
			Expect(limiter.DestroyQuotaGroup(logger, "/full/path/to/volume")).To(Succeed())

			Expect(fakeCommandRunner).Should(HaveExecutedSerially(fake_command_runner.CommandSpec{
				Path: "btrfs",
				Args: []string{"qgroup", "destroy", "/full/path/to/volume", "/full/path/to/volume"},
			}))
		})

		Context("when destroying the qgroup fails", func() {
			BeforeEach(func() {
				fakeCommandRunner.WhenRunning(fake_command_runner.CommandSpec{
					Path: "btrfs",
				}, func(cmd *exec.Cmd) error {
					cmd.Stdout.Write([]byte("failed to destroy qgroup"))
					cmd.Stderr.Write([]byte("some stderr text"))

					return errors.New("exit status 1")
				})
			})

			It("forwards the stdout and stderr", func() {
				err := limiter.DestroyQuotaGroup(logger, "/full/path/to/volume")

				Expect(err).To(MatchError(ContainSubstring("failed to destroy qgroup")))
				Expect(err).To(MatchError(ContainSubstring("some stderr text")))
			})
		})
	})
})
