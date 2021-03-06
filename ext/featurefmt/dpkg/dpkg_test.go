// Copyright 2017 clair authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dpkg

import (
	"testing"

	"github.com/quay/clair/v3/database"
	"github.com/quay/clair/v3/ext/featurefmt"
	"github.com/quay/clair/v3/ext/versionfmt/dpkg"
)

func TestListFeatures(t *testing.T) {
	for _, test := range []featurefmt.TestCase{
		{
			"valid status file",
			map[string]string{"var/lib/dpkg/status": "dpkg/testdata/valid"},
			[]database.LayerFeature{
				{Feature: database.Feature{"libapt-pkg5.0", "1.6.3ubuntu0.1", "dpkg", "binary"}},
				{Feature: database.Feature{"perl-base", "5.26.1-6ubuntu0.2", "dpkg", "binary"}},
				{Feature: database.Feature{"libmount1", "2.31.1-0.4ubuntu3.1", "dpkg", "binary"}},
				{Feature: database.Feature{"perl", "5.26.1-6ubuntu0.2", "dpkg", "source"}},
				{Feature: database.Feature{"libgnutls30", "3.5.18-1ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"liblzma5", "5.2.2-1.3", "dpkg", "binary"}},
				{Feature: database.Feature{"ncurses-bin", "6.1-1ubuntu1.18.04", "dpkg", "binary"}},
				{Feature: database.Feature{"lsb", "9.20170808ubuntu1", "dpkg", "source"}},
				{Feature: database.Feature{"sed", "4.4-2", "dpkg", "source"}},
				{Feature: database.Feature{"libsystemd0", "237-3ubuntu10.3", "dpkg", "binary"}},
				{Feature: database.Feature{"procps", "2:3.3.12-3ubuntu1.1", "dpkg", "source"}},
				{Feature: database.Feature{"login", "1:4.5-1ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"libunistring2", "0.9.9-0ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"sed", "4.4-2", "dpkg", "binary"}},
				{Feature: database.Feature{"libselinux", "2.7-2build2", "dpkg", "source"}},
				{Feature: database.Feature{"libseccomp", "2.3.1-2.1ubuntu4", "dpkg", "source"}},
				{Feature: database.Feature{"libss2", "1.44.1-1", "dpkg", "binary"}},
				{Feature: database.Feature{"liblz4-1", "0.0~r131-2ubuntu3", "dpkg", "binary"}},
				{Feature: database.Feature{"libsemanage1", "2.7-2build2", "dpkg", "binary"}},
				{Feature: database.Feature{"libtasn1-6", "4.13-2", "dpkg", "source"}},
				{Feature: database.Feature{"libzstd1", "1.3.3+dfsg-2ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"fdisk", "2.31.1-0.4ubuntu3.1", "dpkg", "binary"}},
				{Feature: database.Feature{"xz-utils", "5.2.2-1.3", "dpkg", "source"}},
				{Feature: database.Feature{"lsb-base", "9.20170808ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"libpam-modules-bin", "1.1.8-3.6ubuntu2", "dpkg", "binary"}},
				{Feature: database.Feature{"dash", "0.5.8-2.10", "dpkg", "binary"}},
				{Feature: database.Feature{"gnupg2", "2.2.4-1ubuntu1.1", "dpkg", "source"}},
				{Feature: database.Feature{"libfdisk1", "2.31.1-0.4ubuntu3.1", "dpkg", "binary"}},
				{Feature: database.Feature{"lz4", "0.0~r131-2ubuntu3", "dpkg", "source"}},
				{Feature: database.Feature{"libpam0g", "1.1.8-3.6ubuntu2", "dpkg", "binary"}},
				{Feature: database.Feature{"libc-bin", "2.27-3ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"libcap-ng", "0.7.7-3.1", "dpkg", "source"}},
				{Feature: database.Feature{"libcom-err2", "1.44.1-1", "dpkg", "binary"}},
				{Feature: database.Feature{"libudev1", "237-3ubuntu10.3", "dpkg", "binary"}},
				{Feature: database.Feature{"debconf", "1.5.66", "dpkg", "binary"}},
				{Feature: database.Feature{"tar", "1.29b-2", "dpkg", "binary"}},
				{Feature: database.Feature{"diffutils", "1:3.6-1", "dpkg", "source"}},
				{Feature: database.Feature{"gcc-8", "8-20180414-1ubuntu2", "dpkg", "source"}},
				{Feature: database.Feature{"e2fsprogs", "1.44.1-1", "dpkg", "source"}},
				{Feature: database.Feature{"bzip2", "1.0.6-8.1", "dpkg", "source"}},
				{Feature: database.Feature{"diffutils", "1:3.6-1", "dpkg", "binary"}},
				{Feature: database.Feature{"grep", "3.1-2", "dpkg", "binary"}},
				{Feature: database.Feature{"libgcc1", "1:8-20180414-1ubuntu2", "dpkg", "binary"}},
				{Feature: database.Feature{"bash", "4.4.18-2ubuntu1", "dpkg", "source"}},
				{Feature: database.Feature{"libtinfo5", "6.1-1ubuntu1.18.04", "dpkg", "binary"}},
				{Feature: database.Feature{"procps", "2:3.3.12-3ubuntu1.1", "dpkg", "binary"}},
				{Feature: database.Feature{"bzip2", "1.0.6-8.1", "dpkg", "binary"}},
				{Feature: database.Feature{"init-system-helpers", "1.51", "dpkg", "binary"}},
				{Feature: database.Feature{"libncursesw5", "6.1-1ubuntu1.18.04", "dpkg", "binary"}},
				{Feature: database.Feature{"init-system-helpers", "1.51", "dpkg", "source"}},
				{Feature: database.Feature{"libpam-modules", "1.1.8-3.6ubuntu2", "dpkg", "binary"}},
				{Feature: database.Feature{"libext2fs2", "1.44.1-1", "dpkg", "binary"}},
				{Feature: database.Feature{"libacl1", "2.2.52-3build1", "dpkg", "binary"}},
				{Feature: database.Feature{"hostname", "3.20", "dpkg", "binary"}},
				{Feature: database.Feature{"libgpg-error", "1.27-6", "dpkg", "source"}},
				{Feature: database.Feature{"acl", "2.2.52-3build1", "dpkg", "source"}},
				{Feature: database.Feature{"apt", "1.6.3ubuntu0.1", "dpkg", "binary"}},
				{Feature: database.Feature{"base-files", "10.1ubuntu2.2", "dpkg", "source"}},
				{Feature: database.Feature{"libgpg-error0", "1.27-6", "dpkg", "binary"}},
				{Feature: database.Feature{"audit", "1:2.8.2-1ubuntu1", "dpkg", "source"}},
				{Feature: database.Feature{"hostname", "3.20", "dpkg", "source"}},
				{Feature: database.Feature{"gzip", "1.6-5ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"libc6", "2.27-3ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"libnettle6", "3.4-1", "dpkg", "binary"}},
				{Feature: database.Feature{"sysvinit-utils", "2.88dsf-59.10ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"debianutils", "4.8.4", "dpkg", "source"}},
				{Feature: database.Feature{"libstdc++6", "8-20180414-1ubuntu2", "dpkg", "binary"}},
				{Feature: database.Feature{"libsepol", "2.7-1", "dpkg", "source"}},
				{Feature: database.Feature{"libpcre3", "2:8.39-9", "dpkg", "binary"}},
				{Feature: database.Feature{"libuuid1", "2.31.1-0.4ubuntu3.1", "dpkg", "binary"}},
				{Feature: database.Feature{"systemd", "237-3ubuntu10.3", "dpkg", "source"}},
				{Feature: database.Feature{"tar", "1.29b-2", "dpkg", "source"}},
				{Feature: database.Feature{"ubuntu-keyring", "2018.02.28", "dpkg", "source"}},
				{Feature: database.Feature{"passwd", "1:4.5-1ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"sysvinit", "2.88dsf-59.10ubuntu1", "dpkg", "source"}},
				{Feature: database.Feature{"libidn2-0", "2.0.4-1.1build2", "dpkg", "binary"}},
				{Feature: database.Feature{"libhogweed4", "3.4-1", "dpkg", "binary"}},
				{Feature: database.Feature{"db5.3", "5.3.28-13.1ubuntu1", "dpkg", "source"}},
				{Feature: database.Feature{"sensible-utils", "0.0.12", "dpkg", "source"}},
				{Feature: database.Feature{"dpkg", "1.19.0.5ubuntu2", "dpkg", "source"}},
				{Feature: database.Feature{"libp11-kit0", "0.23.9-2", "dpkg", "binary"}},
				{Feature: database.Feature{"glibc", "2.27-3ubuntu1", "dpkg", "source"}},
				{Feature: database.Feature{"mount", "2.31.1-0.4ubuntu3.1", "dpkg", "binary"}},
				{Feature: database.Feature{"libsemanage-common", "2.7-2build2", "dpkg", "binary"}},
				{Feature: database.Feature{"libblkid1", "2.31.1-0.4ubuntu3.1", "dpkg", "binary"}},
				{Feature: database.Feature{"libdebconfclient0", "0.213ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"libffi", "3.2.1-8", "dpkg", "source"}},
				{Feature: database.Feature{"pam", "1.1.8-3.6ubuntu2", "dpkg", "source"}},
				{Feature: database.Feature{"bsdutils", "1:2.31.1-0.4ubuntu3.1", "dpkg", "binary"}},
				{Feature: database.Feature{"libtasn1-6", "4.13-2", "dpkg", "binary"}},
				{Feature: database.Feature{"libaudit-common", "1:2.8.2-1ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"gpgv", "2.2.4-1ubuntu1.1", "dpkg", "binary"}},
				{Feature: database.Feature{"libzstd", "1.3.3+dfsg-2ubuntu1", "dpkg", "source"}},
				{Feature: database.Feature{"base-passwd", "3.5.44", "dpkg", "source"}},
				{Feature: database.Feature{"adduser", "3.116ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"libattr1", "1:2.4.47-2build1", "dpkg", "binary"}},
				{Feature: database.Feature{"libncurses5", "6.1-1ubuntu1.18.04", "dpkg", "binary"}},
				{Feature: database.Feature{"coreutils", "8.28-1ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"base-passwd", "3.5.44", "dpkg", "binary"}},
				{Feature: database.Feature{"ubuntu-keyring", "2018.02.28", "dpkg", "binary"}},
				{Feature: database.Feature{"adduser", "3.116ubuntu1", "dpkg", "source"}},
				{Feature: database.Feature{"libsmartcols1", "2.31.1-0.4ubuntu3.1", "dpkg", "binary"}},
				{Feature: database.Feature{"libunistring", "0.9.9-0ubuntu1", "dpkg", "source"}},
				{Feature: database.Feature{"mawk", "1.3.3-17ubuntu3", "dpkg", "source"}},
				{Feature: database.Feature{"coreutils", "8.28-1ubuntu1", "dpkg", "source"}},
				{Feature: database.Feature{"attr", "1:2.4.47-2build1", "dpkg", "source"}},
				{Feature: database.Feature{"gmp", "2:6.1.2+dfsg-2", "dpkg", "source"}},
				{Feature: database.Feature{"libsemanage", "2.7-2build2", "dpkg", "source"}},
				{Feature: database.Feature{"libselinux1", "2.7-2build2", "dpkg", "binary"}},
				{Feature: database.Feature{"libseccomp2", "2.3.1-2.1ubuntu4", "dpkg", "binary"}},
				{Feature: database.Feature{"zlib1g", "1:1.2.11.dfsg-0ubuntu2", "dpkg", "binary"}},
				{Feature: database.Feature{"dash", "0.5.8-2.10", "dpkg", "source"}},
				{Feature: database.Feature{"gnutls28", "3.5.18-1ubuntu1", "dpkg", "source"}},
				{Feature: database.Feature{"libpam-runtime", "1.1.8-3.6ubuntu2", "dpkg", "binary"}},
				{Feature: database.Feature{"libgcrypt20", "1.8.1-4ubuntu1.1", "dpkg", "source"}},
				{Feature: database.Feature{"sensible-utils", "0.0.12", "dpkg", "binary"}},
				{Feature: database.Feature{"p11-kit", "0.23.9-2", "dpkg", "source"}},
				{Feature: database.Feature{"ncurses-base", "6.1-1ubuntu1.18.04", "dpkg", "binary"}},
				{Feature: database.Feature{"e2fsprogs", "1.44.1-1", "dpkg", "binary"}},
				{Feature: database.Feature{"libgcrypt20", "1.8.1-4ubuntu1.1", "dpkg", "binary"}},
				{Feature: database.Feature{"libprocps6", "2:3.3.12-3ubuntu1.1", "dpkg", "binary"}},
				{Feature: database.Feature{"debconf", "1.5.66", "dpkg", "source"}},
				{Feature: database.Feature{"gcc-8-base", "8-20180414-1ubuntu2", "dpkg", "binary"}},
				{Feature: database.Feature{"base-files", "10.1ubuntu2.2", "dpkg", "binary"}},
				{Feature: database.Feature{"libbz2-1.0", "1.0.6-8.1", "dpkg", "binary"}},
				{Feature: database.Feature{"grep", "3.1-2", "dpkg", "source"}},
				{Feature: database.Feature{"bash", "4.4.18-2ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"libgmp10", "2:6.1.2+dfsg-2", "dpkg", "binary"}},
				{Feature: database.Feature{"shadow", "1:4.5-1ubuntu1", "dpkg", "source"}},
				{Feature: database.Feature{"libidn2", "2.0.4-1.1build2", "dpkg", "source"}},
				{Feature: database.Feature{"gzip", "1.6-5ubuntu1", "dpkg", "source"}},
				{Feature: database.Feature{"util-linux", "2.31.1-0.4ubuntu3.1", "dpkg", "binary"}},
				{Feature: database.Feature{"libaudit1", "1:2.8.2-1ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"libsepol1", "2.7-1", "dpkg", "binary"}},
				{Feature: database.Feature{"pcre3", "2:8.39-9", "dpkg", "source"}},
				{Feature: database.Feature{"apt", "1.6.3ubuntu0.1", "dpkg", "source"}},
				{Feature: database.Feature{"nettle", "3.4-1", "dpkg", "source"}},
				{Feature: database.Feature{"util-linux", "2.31.1-0.4ubuntu3.1", "dpkg", "source"}},
				{Feature: database.Feature{"libcap-ng0", "0.7.7-3.1", "dpkg", "binary"}},
				{Feature: database.Feature{"debianutils", "4.8.4", "dpkg", "binary"}},
				{Feature: database.Feature{"ncurses", "6.1-1ubuntu1.18.04", "dpkg", "source"}},
				{Feature: database.Feature{"libffi6", "3.2.1-8", "dpkg", "binary"}},
				{Feature: database.Feature{"cdebconf", "0.213ubuntu1", "dpkg", "source"}},
				{Feature: database.Feature{"findutils", "4.6.0+git+20170828-2", "dpkg", "source"}},
				{Feature: database.Feature{"libdb5.3", "5.3.28-13.1ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"zlib", "1:1.2.11.dfsg-0ubuntu2", "dpkg", "source"}},
				{Feature: database.Feature{"findutils", "4.6.0+git+20170828-2", "dpkg", "binary"}},
				{Feature: database.Feature{"dpkg", "1.19.0.5ubuntu2", "dpkg", "binary"}},
				{Feature: database.Feature{"mawk", "1.3.3-17ubuntu3", "dpkg", "binary"}},
			},
		},
		{
			"corrupted status file",
			map[string]string{"var/lib/dpkg/status": "dpkg/testdata/corrupted"},
			[]database.LayerFeature{
				{Feature: database.Feature{"libpam-modules-bin", "1.1.8-3.1ubuntu3", "dpkg", "binary"}},
				{Feature: database.Feature{"gcc-5", "5.1.1-12ubuntu1", "dpkg", "source"}},
				{Feature: database.Feature{"makedev", "2.3.1-93ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"libgcc1", "1:5.1.1-12ubuntu1", "dpkg", "binary"}},
				{Feature: database.Feature{"pam", "1.1.8-3.1ubuntu3", "dpkg", "source"}},
				{Feature: database.Feature{"makedev", "2.3.1-93ubuntu1", "dpkg", "source"}},
				{Feature: database.Feature{"libpam-runtime", "1.1.8-3.1ubuntu3", "dpkg", "binary"}},
			},
		},
	} {
		featurefmt.RunTest(t, test, &lister{}, dpkg.ParserName)
	}
}
