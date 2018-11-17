// Mgmt
// Copyright (C) 2013-2018+ James Shubin and the project contributors
// Written by James Shubin <james@shubin.ca> and the project contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// +build !root

package util

import (
	"reflect"
	"sort"
	"testing"
)

func TestNumToAlpha(t *testing.T) {
	var numToAlphaTests = []struct {
		number int
		result string
	}{
		{0, "a"},
		{25, "z"},
		{26, "aa"},
		{27, "ab"},
		{702, "aaa"},
		{703, "aab"},
		{63269, "cool"},
	}

	for _, test := range numToAlphaTests {
		actual := NumToAlpha(test.number)
		if actual != test.result {
			t.Errorf("NumToAlpha(%d): expected %s, actual %s", test.number, test.result, actual)
		}
	}
}

func TestUtilT1(t *testing.T) {

	if Dirname("/foo/bar/baz") != "/foo/bar/" {
		t.Errorf("Result is incorrect.")
	}

	if Dirname("/foo/bar/baz/") != "/foo/bar/" {
		t.Errorf("Result is incorrect.")
	}

	if Dirname("/foo/") != "/" {
		t.Errorf("Result is incorrect.")
	}

	if Dirname("/") != "" { // TODO: should this equal "/" or "" ?
		t.Errorf("Result is incorrect.")
	}

	if Basename("/foo/bar/baz") != "baz" {
		t.Errorf("Result is incorrect.")
	}

	if Basename("/foo/bar/baz/") != "baz/" {
		t.Errorf("Result is incorrect.")
	}

	if Basename("/foo/") != "foo/" {
		t.Errorf("Result is incorrect.")
	}

	if Basename("/") != "/" { // TODO: should this equal "" or "/" ?
		t.Errorf("Result is incorrect.")
	}

	if Basename("") != "" { // TODO: should this equal something different?
		t.Errorf("Result is incorrect.")
	}
}

func TestUtilT2(t *testing.T) {

	// TODO: compare the output with the actual list
	p0 := "/"
	r0 := []string{""} // TODO: is this correct?
	if len(PathSplit(p0)) != len(r0) {
		t.Errorf("Result should be: %q.", r0)
		t.Errorf("Result should have a length of: %v.", len(r0))
	}

	p1 := "/foo/bar/baz"
	r1 := []string{"", "foo", "bar", "baz"}
	if len(PathSplit(p1)) != len(r1) {
		//t.Errorf("Result should be: %q.", r1)
		t.Errorf("Result should have a length of: %v.", len(r1))
	}

	p2 := "/foo/bar/baz/"
	r2 := []string{"", "foo", "bar", "baz"}
	if len(PathSplit(p2)) != len(r2) {
		t.Errorf("Result should have a length of: %v.", len(r2))
	}
}

func TestUtilT3(t *testing.T) {

	if HasPathPrefix("/foo/bar/baz", "/foo/ba") != false {
		t.Errorf("Result should be false.")
	}

	if HasPathPrefix("/foo/bar/baz", "/foo/bar") != true {
		t.Errorf("Result should be true.")
	}

	if HasPathPrefix("/foo/bar/baz", "/foo/bar/") != true {
		t.Errorf("Result should be true.")
	}

	if HasPathPrefix("/foo/bar/baz/", "/foo/bar") != true {
		t.Errorf("Result should be true.")
	}

	if HasPathPrefix("/foo/bar/baz/", "/foo/bar/") != true {
		t.Errorf("Result should be true.")
	}

	if HasPathPrefix("/foo/bar/baz/", "/foo/bar/baz/dude") != false {
		t.Errorf("Result should be false.")
	}

	if HasPathPrefix("/foo/bar/baz/boo/", "/foo/") != true {
		t.Errorf("Result should be true.")
	}
}

func TestUtilT4(t *testing.T) {

	if PathPrefixDelta("/foo/bar/baz", "/foo/ba") != -1 {
		t.Errorf("Result should be -1.")
	}

	if PathPrefixDelta("/foo/bar/baz", "/foo/bar") != 1 {
		t.Errorf("Result should be 1.")
	}

	if PathPrefixDelta("/foo/bar/baz", "/foo/bar/") != 1 {
		t.Errorf("Result should be 1.")
	}

	if PathPrefixDelta("/foo/bar/baz/", "/foo/bar") != 1 {
		t.Errorf("Result should be 1.")
	}

	if PathPrefixDelta("/foo/bar/baz/", "/foo/bar/") != 1 {
		t.Errorf("Result should be 1.")
	}

	if PathPrefixDelta("/foo/bar/baz/", "/foo/bar/baz/dude") != -1 {
		t.Errorf("Result should be -1.")
	}

	if PathPrefixDelta("/foo/bar/baz/a/b/c/", "/foo/bar/baz") != 3 {
		t.Errorf("Result should be 3.")
	}

	if PathPrefixDelta("/foo/bar/baz/", "/foo/bar/baz") != 0 {
		t.Errorf("Result should be 0.")
	}
}

func TestUtilT8(t *testing.T) {

	r0 := []string{"/"}
	if fullList0 := PathSplitFullReversed("/"); !reflect.DeepEqual(r0, fullList0) {
		t.Errorf("PathSplitFullReversed expected: %v; got: %v.", r0, fullList0)
	}

	r1 := []string{"/foo/bar/baz/file", "/foo/bar/baz/", "/foo/bar/", "/foo/", "/"}
	if fullList1 := PathSplitFullReversed("/foo/bar/baz/file"); !reflect.DeepEqual(r1, fullList1) {
		t.Errorf("PathSplitFullReversed expected: %v; got: %v.", r1, fullList1)
	}

	r2 := []string{"/foo/bar/baz/dir/", "/foo/bar/baz/", "/foo/bar/", "/foo/", "/"}
	if fullList2 := PathSplitFullReversed("/foo/bar/baz/dir/"); !reflect.DeepEqual(r2, fullList2) {
		t.Errorf("PathSplitFullReversed expected: %v; got: %v.", r2, fullList2)
	}

}

func TestUtilT9(t *testing.T) {
	fileListIn := []string{ // list taken from drbd-utils package
		"/etc/drbd.conf",
		"/etc/drbd.d/global_common.conf",
		"/lib/drbd/drbd",
		"/lib/drbd/drbdadm-83",
		"/lib/drbd/drbdadm-84",
		"/lib/drbd/drbdsetup-83",
		"/lib/drbd/drbdsetup-84",
		"/usr/lib/drbd/crm-fence-peer.sh",
		"/usr/lib/drbd/crm-unfence-peer.sh",
		"/usr/lib/drbd/notify-emergency-reboot.sh",
		"/usr/lib/drbd/notify-emergency-shutdown.sh",
		"/usr/lib/drbd/notify-io-error.sh",
		"/usr/lib/drbd/notify-out-of-sync.sh",
		"/usr/lib/drbd/notify-pri-lost-after-sb.sh",
		"/usr/lib/drbd/notify-pri-lost.sh",
		"/usr/lib/drbd/notify-pri-on-incon-degr.sh",
		"/usr/lib/drbd/notify-split-brain.sh",
		"/usr/lib/drbd/notify.sh",
		"/usr/lib/drbd/outdate-peer.sh",
		"/usr/lib/drbd/rhcs_fence",
		"/usr/lib/drbd/snapshot-resync-target-lvm.sh",
		"/usr/lib/drbd/stonith_admin-fence-peer.sh",
		"/usr/lib/drbd/unsnapshot-resync-target-lvm.sh",
		"/usr/lib/systemd/system/drbd.service",
		"/usr/lib/tmpfiles.d/drbd.conf",
		"/usr/sbin/drbd-overview",
		"/usr/sbin/drbdadm",
		"/usr/sbin/drbdmeta",
		"/usr/sbin/drbdsetup",
		"/usr/share/doc/drbd-utils/COPYING",
		"/usr/share/doc/drbd-utils/ChangeLog",
		"/usr/share/doc/drbd-utils/README",
		"/usr/share/doc/drbd-utils/drbd.conf.example",
		"/usr/share/man/man5/drbd.conf-8.3.5.gz",
		"/usr/share/man/man5/drbd.conf-8.4.5.gz",
		"/usr/share/man/man5/drbd.conf-9.0.5.gz",
		"/usr/share/man/man5/drbd.conf.5.gz",
		"/usr/share/man/man8/drbd-8.3.8.gz",
		"/usr/share/man/man8/drbd-8.4.8.gz",
		"/usr/share/man/man8/drbd-9.0.8.gz",
		"/usr/share/man/man8/drbd-overview-9.0.8.gz",
		"/usr/share/man/man8/drbd-overview.8.gz",
		"/usr/share/man/man8/drbd.8.gz",
		"/usr/share/man/man8/drbdadm-8.3.8.gz",
		"/usr/share/man/man8/drbdadm-8.4.8.gz",
		"/usr/share/man/man8/drbdadm-9.0.8.gz",
		"/usr/share/man/man8/drbdadm.8.gz",
		"/usr/share/man/man8/drbddisk-8.3.8.gz",
		"/usr/share/man/man8/drbddisk-8.4.8.gz",
		"/usr/share/man/man8/drbdmeta-8.3.8.gz",
		"/usr/share/man/man8/drbdmeta-8.4.8.gz",
		"/usr/share/man/man8/drbdmeta-9.0.8.gz",
		"/usr/share/man/man8/drbdmeta.8.gz",
		"/usr/share/man/man8/drbdsetup-8.3.8.gz",
		"/usr/share/man/man8/drbdsetup-8.4.8.gz",
		"/usr/share/man/man8/drbdsetup-9.0.8.gz",
		"/usr/share/man/man8/drbdsetup.8.gz",
		"/etc/drbd.d",
		"/usr/share/doc/drbd-utils",
		"/var/lib/drbd",
	}
	sort.Strings(fileListIn)

	fileListOut := []string{ // fixed up manually
		"/etc/drbd.conf",
		"/etc/drbd.d/global_common.conf",
		"/lib/drbd/drbd",
		"/lib/drbd/drbdadm-83",
		"/lib/drbd/drbdadm-84",
		"/lib/drbd/drbdsetup-83",
		"/lib/drbd/drbdsetup-84",
		"/usr/lib/drbd/crm-fence-peer.sh",
		"/usr/lib/drbd/crm-unfence-peer.sh",
		"/usr/lib/drbd/notify-emergency-reboot.sh",
		"/usr/lib/drbd/notify-emergency-shutdown.sh",
		"/usr/lib/drbd/notify-io-error.sh",
		"/usr/lib/drbd/notify-out-of-sync.sh",
		"/usr/lib/drbd/notify-pri-lost-after-sb.sh",
		"/usr/lib/drbd/notify-pri-lost.sh",
		"/usr/lib/drbd/notify-pri-on-incon-degr.sh",
		"/usr/lib/drbd/notify-split-brain.sh",
		"/usr/lib/drbd/notify.sh",
		"/usr/lib/drbd/outdate-peer.sh",
		"/usr/lib/drbd/rhcs_fence",
		"/usr/lib/drbd/snapshot-resync-target-lvm.sh",
		"/usr/lib/drbd/stonith_admin-fence-peer.sh",
		"/usr/lib/drbd/unsnapshot-resync-target-lvm.sh",
		"/usr/lib/systemd/system/drbd.service",
		"/usr/lib/tmpfiles.d/drbd.conf",
		"/usr/sbin/drbd-overview",
		"/usr/sbin/drbdadm",
		"/usr/sbin/drbdmeta",
		"/usr/sbin/drbdsetup",
		"/usr/share/doc/drbd-utils/COPYING",
		"/usr/share/doc/drbd-utils/ChangeLog",
		"/usr/share/doc/drbd-utils/README",
		"/usr/share/doc/drbd-utils/drbd.conf.example",
		"/usr/share/man/man5/drbd.conf-8.3.5.gz",
		"/usr/share/man/man5/drbd.conf-8.4.5.gz",
		"/usr/share/man/man5/drbd.conf-9.0.5.gz",
		"/usr/share/man/man5/drbd.conf.5.gz",
		"/usr/share/man/man8/drbd-8.3.8.gz",
		"/usr/share/man/man8/drbd-8.4.8.gz",
		"/usr/share/man/man8/drbd-9.0.8.gz",
		"/usr/share/man/man8/drbd-overview-9.0.8.gz",
		"/usr/share/man/man8/drbd-overview.8.gz",
		"/usr/share/man/man8/drbd.8.gz",
		"/usr/share/man/man8/drbdadm-8.3.8.gz",
		"/usr/share/man/man8/drbdadm-8.4.8.gz",
		"/usr/share/man/man8/drbdadm-9.0.8.gz",
		"/usr/share/man/man8/drbdadm.8.gz",
		"/usr/share/man/man8/drbddisk-8.3.8.gz",
		"/usr/share/man/man8/drbddisk-8.4.8.gz",
		"/usr/share/man/man8/drbdmeta-8.3.8.gz",
		"/usr/share/man/man8/drbdmeta-8.4.8.gz",
		"/usr/share/man/man8/drbdmeta-9.0.8.gz",
		"/usr/share/man/man8/drbdmeta.8.gz",
		"/usr/share/man/man8/drbdsetup-8.3.8.gz",
		"/usr/share/man/man8/drbdsetup-8.4.8.gz",
		"/usr/share/man/man8/drbdsetup-9.0.8.gz",
		"/usr/share/man/man8/drbdsetup.8.gz",
		"/etc/drbd.d/",               // added trailing slash
		"/usr/share/doc/drbd-utils/", // added trailing slash
		"/var/lib/drbd",              // can't be fixed :(
	}
	sort.Strings(fileListOut)

	dirify := DirifyFileList(fileListIn, false) // TODO: test with true
	sort.Strings(dirify)
	equals := reflect.DeepEqual(fileListOut, dirify)
	if a, b := len(fileListOut), len(dirify); a != b {
		t.Errorf("DirifyFileList counts didn't match: %d != %d", a, b)
	} else if !equals {
		t.Errorf("DirifyFileList did not match expected!")
		for i := 0; i < len(dirify); i++ {
			if fileListOut[i] != dirify[i] {
				t.Errorf("# %d: %v <> %v", i, fileListOut[i], dirify[i])
			}
		}
	}
}

func TestUtilT10(t *testing.T) {
	fileListIn := []string{ // fake package list
		"/etc/drbd.conf",
		"/usr/share/man/man8/drbdsetup.8.gz",
		"/etc/drbd.d",
		"/etc/drbd.d/foo",
		"/var/lib/drbd",
		"/var/somedir/",
	}
	sort.Strings(fileListIn)

	fileListOut := []string{ // fixed up manually
		"/etc/drbd.conf",
		"/usr/share/man/man8/drbdsetup.8.gz",
		"/etc/drbd.d/", // added trailing slash
		"/etc/drbd.d/foo",
		"/var/lib/drbd", // can't be fixed :(
		"/var/somedir/", // stays the same
	}
	sort.Strings(fileListOut)

	dirify := DirifyFileList(fileListIn, false) // TODO: test with true
	sort.Strings(dirify)
	equals := reflect.DeepEqual(fileListOut, dirify)
	if a, b := len(fileListOut), len(dirify); a != b {
		t.Errorf("DirifyFileList counts didn't match: %d != %d", a, b)
	} else if !equals {
		t.Errorf("DirifyFileList did not match expected!")
		for i := 0; i < len(dirify); i++ {
			if fileListOut[i] != dirify[i] {
				t.Errorf("# %d: %v <> %v", i, fileListOut[i], dirify[i])
			}
		}
	}
}

func TestUtilT11(t *testing.T) {
	in1 := []string{"/", "/usr/", "/usr/lib/", "/usr/share/"} // input
	ex1 := []string{"/usr/lib/", "/usr/share/"}               // expected
	sort.Strings(ex1)
	out1 := RemoveCommonFilePrefixes(in1)
	sort.Strings(out1)
	if !reflect.DeepEqual(ex1, out1) {
		t.Errorf("RemoveCommonFilePrefixes expected: %v; got: %v.", ex1, out1)
	}

	in2 := []string{"/", "/usr/"}
	ex2 := []string{"/usr/"}
	sort.Strings(ex2)
	out2 := RemoveCommonFilePrefixes(in2)
	sort.Strings(out2)
	if !reflect.DeepEqual(ex2, out2) {
		t.Errorf("RemoveCommonFilePrefixes expected: %v; got: %v.", ex2, out2)
	}

	in3 := []string{"/"}
	ex3 := []string{"/"}
	out3 := RemoveCommonFilePrefixes(in3)
	if !reflect.DeepEqual(ex3, out3) {
		t.Errorf("RemoveCommonFilePrefixes expected: %v; got: %v.", ex3, out3)
	}

	in4 := []string{"/usr/bin/foo", "/usr/bin/bar", "/usr/lib/", "/usr/share/"}
	ex4 := []string{"/usr/bin/foo", "/usr/bin/bar", "/usr/lib/", "/usr/share/"}
	sort.Strings(ex4)
	out4 := RemoveCommonFilePrefixes(in4)
	sort.Strings(out4)
	if !reflect.DeepEqual(ex4, out4) {
		t.Errorf("RemoveCommonFilePrefixes expected: %v; got: %v.", ex4, out4)
	}

	in5 := []string{"/usr/bin/foo", "/usr/bin/bar", "/usr/lib/", "/usr/share/", "/usr/bin"}
	ex5 := []string{"/usr/bin/foo", "/usr/bin/bar", "/usr/lib/", "/usr/share/"}
	sort.Strings(ex5)
	out5 := RemoveCommonFilePrefixes(in5)
	sort.Strings(out5)
	if !reflect.DeepEqual(ex5, out5) {
		t.Errorf("RemoveCommonFilePrefixes expected: %v; got: %v.", ex5, out5)
	}

	in6 := []string{"/etc/drbd.d/", "/lib/drbd/", "/usr/lib/drbd/", "/usr/lib/systemd/system/", "/usr/lib/tmpfiles.d/", "/usr/sbin/", "/usr/share/doc/drbd-utils/", "/usr/share/man/man5/", "/usr/share/man/man8/", "/usr/share/doc/", "/var/lib/"}
	ex6 := []string{"/etc/drbd.d/", "/lib/drbd/", "/usr/lib/drbd/", "/usr/lib/systemd/system/", "/usr/lib/tmpfiles.d/", "/usr/sbin/", "/usr/share/doc/drbd-utils/", "/usr/share/man/man5/", "/usr/share/man/man8/", "/var/lib/"}
	sort.Strings(ex6)
	out6 := RemoveCommonFilePrefixes(in6)
	sort.Strings(out6)
	if !reflect.DeepEqual(ex6, out6) {
		t.Errorf("RemoveCommonFilePrefixes expected: %v; got: %v.", ex6, out6)
	}

	in7 := []string{"/etc/", "/lib/", "/usr/lib/", "/usr/lib/systemd/", "/usr/", "/usr/share/doc/", "/usr/share/man/", "/var/"}
	ex7 := []string{"/etc/", "/lib/", "/usr/lib/systemd/", "/usr/share/doc/", "/usr/share/man/", "/var/"}
	sort.Strings(ex7)
	out7 := RemoveCommonFilePrefixes(in7)
	sort.Strings(out7)
	if !reflect.DeepEqual(ex7, out7) {
		t.Errorf("RemoveCommonFilePrefixes expected: %v; got: %v.", ex7, out7)
	}

	in8 := []string{
		"/etc/drbd.conf",
		"/etc/drbd.d/global_common.conf",
		"/lib/drbd/drbd",
		"/lib/drbd/drbdadm-83",
		"/lib/drbd/drbdadm-84",
		"/lib/drbd/drbdsetup-83",
		"/lib/drbd/drbdsetup-84",
		"/usr/lib/drbd/crm-fence-peer.sh",
		"/usr/lib/drbd/crm-unfence-peer.sh",
		"/usr/lib/drbd/notify-emergency-reboot.sh",
		"/usr/lib/drbd/notify-emergency-shutdown.sh",
		"/usr/lib/drbd/notify-io-error.sh",
		"/usr/lib/drbd/notify-out-of-sync.sh",
		"/usr/lib/drbd/notify-pri-lost-after-sb.sh",
		"/usr/lib/drbd/notify-pri-lost.sh",
		"/usr/lib/drbd/notify-pri-on-incon-degr.sh",
		"/usr/lib/drbd/notify-split-brain.sh",
		"/usr/lib/drbd/notify.sh",
		"/usr/lib/drbd/outdate-peer.sh",
		"/usr/lib/drbd/rhcs_fence",
		"/usr/lib/drbd/snapshot-resync-target-lvm.sh",
		"/usr/lib/drbd/stonith_admin-fence-peer.sh",
		"/usr/lib/drbd/unsnapshot-resync-target-lvm.sh",
		"/usr/lib/systemd/system/drbd.service",
		"/usr/lib/tmpfiles.d/drbd.conf",
		"/usr/sbin/drbd-overview",
		"/usr/sbin/drbdadm",
		"/usr/sbin/drbdmeta",
		"/usr/sbin/drbdsetup",
		"/usr/share/doc/drbd-utils/COPYING",
		"/usr/share/doc/drbd-utils/ChangeLog",
		"/usr/share/doc/drbd-utils/README",
		"/usr/share/doc/drbd-utils/drbd.conf.example",
		"/usr/share/man/man5/drbd.conf-8.3.5.gz",
		"/usr/share/man/man5/drbd.conf-8.4.5.gz",
		"/usr/share/man/man5/drbd.conf-9.0.5.gz",
		"/usr/share/man/man5/drbd.conf.5.gz",
		"/usr/share/man/man8/drbd-8.3.8.gz",
		"/usr/share/man/man8/drbd-8.4.8.gz",
		"/usr/share/man/man8/drbd-9.0.8.gz",
		"/usr/share/man/man8/drbd-overview-9.0.8.gz",
		"/usr/share/man/man8/drbd-overview.8.gz",
		"/usr/share/man/man8/drbd.8.gz",
		"/usr/share/man/man8/drbdadm-8.3.8.gz",
		"/usr/share/man/man8/drbdadm-8.4.8.gz",
		"/usr/share/man/man8/drbdadm-9.0.8.gz",
		"/usr/share/man/man8/drbdadm.8.gz",
		"/usr/share/man/man8/drbddisk-8.3.8.gz",
		"/usr/share/man/man8/drbddisk-8.4.8.gz",
		"/usr/share/man/man8/drbdmeta-8.3.8.gz",
		"/usr/share/man/man8/drbdmeta-8.4.8.gz",
		"/usr/share/man/man8/drbdmeta-9.0.8.gz",
		"/usr/share/man/man8/drbdmeta.8.gz",
		"/usr/share/man/man8/drbdsetup-8.3.8.gz",
		"/usr/share/man/man8/drbdsetup-8.4.8.gz",
		"/usr/share/man/man8/drbdsetup-9.0.8.gz",
		"/usr/share/man/man8/drbdsetup.8.gz",
		"/etc/drbd.d/",
		"/usr/share/doc/drbd-utils/",
		"/var/lib/drbd",
	}
	ex8 := []string{
		"/etc/drbd.conf",
		"/etc/drbd.d/global_common.conf",
		"/lib/drbd/drbd",
		"/lib/drbd/drbdadm-83",
		"/lib/drbd/drbdadm-84",
		"/lib/drbd/drbdsetup-83",
		"/lib/drbd/drbdsetup-84",
		"/usr/lib/drbd/crm-fence-peer.sh",
		"/usr/lib/drbd/crm-unfence-peer.sh",
		"/usr/lib/drbd/notify-emergency-reboot.sh",
		"/usr/lib/drbd/notify-emergency-shutdown.sh",
		"/usr/lib/drbd/notify-io-error.sh",
		"/usr/lib/drbd/notify-out-of-sync.sh",
		"/usr/lib/drbd/notify-pri-lost-after-sb.sh",
		"/usr/lib/drbd/notify-pri-lost.sh",
		"/usr/lib/drbd/notify-pri-on-incon-degr.sh",
		"/usr/lib/drbd/notify-split-brain.sh",
		"/usr/lib/drbd/notify.sh",
		"/usr/lib/drbd/outdate-peer.sh",
		"/usr/lib/drbd/rhcs_fence",
		"/usr/lib/drbd/snapshot-resync-target-lvm.sh",
		"/usr/lib/drbd/stonith_admin-fence-peer.sh",
		"/usr/lib/drbd/unsnapshot-resync-target-lvm.sh",
		"/usr/lib/systemd/system/drbd.service",
		"/usr/lib/tmpfiles.d/drbd.conf",
		"/usr/sbin/drbd-overview",
		"/usr/sbin/drbdadm",
		"/usr/sbin/drbdmeta",
		"/usr/sbin/drbdsetup",
		"/usr/share/doc/drbd-utils/COPYING",
		"/usr/share/doc/drbd-utils/ChangeLog",
		"/usr/share/doc/drbd-utils/README",
		"/usr/share/doc/drbd-utils/drbd.conf.example",
		"/usr/share/man/man5/drbd.conf-8.3.5.gz",
		"/usr/share/man/man5/drbd.conf-8.4.5.gz",
		"/usr/share/man/man5/drbd.conf-9.0.5.gz",
		"/usr/share/man/man5/drbd.conf.5.gz",
		"/usr/share/man/man8/drbd-8.3.8.gz",
		"/usr/share/man/man8/drbd-8.4.8.gz",
		"/usr/share/man/man8/drbd-9.0.8.gz",
		"/usr/share/man/man8/drbd-overview-9.0.8.gz",
		"/usr/share/man/man8/drbd-overview.8.gz",
		"/usr/share/man/man8/drbd.8.gz",
		"/usr/share/man/man8/drbdadm-8.3.8.gz",
		"/usr/share/man/man8/drbdadm-8.4.8.gz",
		"/usr/share/man/man8/drbdadm-9.0.8.gz",
		"/usr/share/man/man8/drbdadm.8.gz",
		"/usr/share/man/man8/drbddisk-8.3.8.gz",
		"/usr/share/man/man8/drbddisk-8.4.8.gz",
		"/usr/share/man/man8/drbdmeta-8.3.8.gz",
		"/usr/share/man/man8/drbdmeta-8.4.8.gz",
		"/usr/share/man/man8/drbdmeta-9.0.8.gz",
		"/usr/share/man/man8/drbdmeta.8.gz",
		"/usr/share/man/man8/drbdsetup-8.3.8.gz",
		"/usr/share/man/man8/drbdsetup-8.4.8.gz",
		"/usr/share/man/man8/drbdsetup-9.0.8.gz",
		"/usr/share/man/man8/drbdsetup.8.gz",
		"/var/lib/drbd",
	}
	sort.Strings(ex8)
	out8 := RemoveCommonFilePrefixes(in8)
	sort.Strings(out8)
	if !reflect.DeepEqual(ex8, out8) {
		t.Errorf("RemoveCommonFilePrefixes expected: %v; got: %v.", ex8, out8)
	}

	in9 := []string{
		"/etc/drbd.conf",
		"/etc/drbd.d/",
		"/lib/drbd/drbd",
		"/lib/drbd/",
		"/lib/drbd/",
		"/lib/drbd/",
		"/usr/lib/drbd/",
		"/usr/lib/drbd/",
		"/usr/lib/drbd/",
		"/usr/lib/drbd/",
		"/usr/lib/drbd/",
		"/usr/lib/systemd/system/",
		"/usr/lib/tmpfiles.d/",
		"/usr/sbin/",
		"/usr/sbin/",
		"/usr/share/doc/drbd-utils/",
		"/usr/share/doc/drbd-utils/",
		"/usr/share/man/man5/",
		"/usr/share/man/man5/",
		"/usr/share/man/man8/",
		"/usr/share/man/man8/",
		"/usr/share/man/man8/",
		"/etc/drbd.d/",
		"/usr/share/doc/drbd-utils/",
		"/var/lib/drbd",
	}
	ex9 := []string{
		"/etc/drbd.conf",
		"/etc/drbd.d/",
		"/lib/drbd/drbd",
		"/usr/lib/drbd/",
		"/usr/lib/systemd/system/",
		"/usr/lib/tmpfiles.d/",
		"/usr/sbin/",
		"/usr/share/doc/drbd-utils/",
		"/usr/share/man/man5/",
		"/usr/share/man/man8/",
		"/var/lib/drbd",
	}
	sort.Strings(ex9)
	out9 := RemoveCommonFilePrefixes(in9)
	sort.Strings(out9)
	if !reflect.DeepEqual(ex9, out9) {
		t.Errorf("RemoveCommonFilePrefixes expected: %v; got: %v.", ex9, out9)
	}

	in10 := []string{
		"/etc/drbd.conf",
		"/etc/drbd.d/",                   // watch me, i'm a dir
		"/etc/drbd.d/global_common.conf", // and watch me i'm a file!
		"/lib/drbd/drbd",
		"/lib/drbd/drbdadm-83",
		"/lib/drbd/drbdadm-84",
		"/lib/drbd/drbdsetup-83",
		"/lib/drbd/drbdsetup-84",
		"/usr/lib/drbd/crm-fence-peer.sh",
		"/usr/lib/drbd/crm-unfence-peer.sh",
		"/usr/lib/drbd/notify-emergency-reboot.sh",
		"/usr/lib/drbd/notify-emergency-shutdown.sh",
		"/usr/lib/drbd/notify-io-error.sh",
		"/usr/lib/drbd/notify-out-of-sync.sh",
		"/usr/lib/drbd/notify-pri-lost-after-sb.sh",
		"/usr/lib/drbd/notify-pri-lost.sh",
		"/usr/lib/drbd/notify-pri-on-incon-degr.sh",
		"/usr/lib/drbd/notify-split-brain.sh",
		"/usr/lib/drbd/notify.sh",
		"/usr/lib/drbd/outdate-peer.sh",
		"/usr/lib/drbd/rhcs_fence",
		"/usr/lib/drbd/snapshot-resync-target-lvm.sh",
		"/usr/lib/drbd/stonith_admin-fence-peer.sh",
		"/usr/lib/drbd/unsnapshot-resync-target-lvm.sh",
		"/usr/lib/systemd/system/drbd.service",
		"/usr/lib/tmpfiles.d/drbd.conf",
		"/usr/sbin/drbd-overview",
		"/usr/sbin/drbdadm",
		"/usr/sbin/drbdmeta",
		"/usr/sbin/drbdsetup",
		"/usr/share/doc/drbd-utils/", // watch me, i'm a dir too
		"/usr/share/doc/drbd-utils/COPYING",
		"/usr/share/doc/drbd-utils/ChangeLog",
		"/usr/share/doc/drbd-utils/README",
		"/usr/share/doc/drbd-utils/drbd.conf.example",
		"/usr/share/man/man5/drbd.conf-8.3.5.gz",
		"/usr/share/man/man5/drbd.conf-8.4.5.gz",
		"/usr/share/man/man5/drbd.conf-9.0.5.gz",
		"/usr/share/man/man5/drbd.conf.5.gz",
		"/usr/share/man/man8/drbd-8.3.8.gz",
		"/usr/share/man/man8/drbd-8.4.8.gz",
		"/usr/share/man/man8/drbd-9.0.8.gz",
		"/usr/share/man/man8/drbd-overview-9.0.8.gz",
		"/usr/share/man/man8/drbd-overview.8.gz",
		"/usr/share/man/man8/drbd.8.gz",
		"/usr/share/man/man8/drbdadm-8.3.8.gz",
		"/usr/share/man/man8/drbdadm-8.4.8.gz",
		"/usr/share/man/man8/drbdadm-9.0.8.gz",
		"/usr/share/man/man8/drbdadm.8.gz",
		"/usr/share/man/man8/drbddisk-8.3.8.gz",
		"/usr/share/man/man8/drbddisk-8.4.8.gz",
		"/usr/share/man/man8/drbdmeta-8.3.8.gz",
		"/usr/share/man/man8/drbdmeta-8.4.8.gz",
		"/usr/share/man/man8/drbdmeta-9.0.8.gz",
		"/usr/share/man/man8/drbdmeta.8.gz",
		"/usr/share/man/man8/drbdsetup-8.3.8.gz",
		"/usr/share/man/man8/drbdsetup-8.4.8.gz",
		"/usr/share/man/man8/drbdsetup-9.0.8.gz",
		"/usr/share/man/man8/drbdsetup.8.gz",
		"/var/lib/drbd",
	}
	ex10 := []string{
		"/etc/drbd.conf",
		"/etc/drbd.d/global_common.conf",
		"/lib/drbd/drbd",
		"/lib/drbd/drbdadm-83",
		"/lib/drbd/drbdadm-84",
		"/lib/drbd/drbdsetup-83",
		"/lib/drbd/drbdsetup-84",
		"/usr/lib/drbd/crm-fence-peer.sh",
		"/usr/lib/drbd/crm-unfence-peer.sh",
		"/usr/lib/drbd/notify-emergency-reboot.sh",
		"/usr/lib/drbd/notify-emergency-shutdown.sh",
		"/usr/lib/drbd/notify-io-error.sh",
		"/usr/lib/drbd/notify-out-of-sync.sh",
		"/usr/lib/drbd/notify-pri-lost-after-sb.sh",
		"/usr/lib/drbd/notify-pri-lost.sh",
		"/usr/lib/drbd/notify-pri-on-incon-degr.sh",
		"/usr/lib/drbd/notify-split-brain.sh",
		"/usr/lib/drbd/notify.sh",
		"/usr/lib/drbd/outdate-peer.sh",
		"/usr/lib/drbd/rhcs_fence",
		"/usr/lib/drbd/snapshot-resync-target-lvm.sh",
		"/usr/lib/drbd/stonith_admin-fence-peer.sh",
		"/usr/lib/drbd/unsnapshot-resync-target-lvm.sh",
		"/usr/lib/systemd/system/drbd.service",
		"/usr/lib/tmpfiles.d/drbd.conf",
		"/usr/sbin/drbd-overview",
		"/usr/sbin/drbdadm",
		"/usr/sbin/drbdmeta",
		"/usr/sbin/drbdsetup",
		"/usr/share/doc/drbd-utils/COPYING",
		"/usr/share/doc/drbd-utils/ChangeLog",
		"/usr/share/doc/drbd-utils/README",
		"/usr/share/doc/drbd-utils/drbd.conf.example",
		"/usr/share/man/man5/drbd.conf-8.3.5.gz",
		"/usr/share/man/man5/drbd.conf-8.4.5.gz",
		"/usr/share/man/man5/drbd.conf-9.0.5.gz",
		"/usr/share/man/man5/drbd.conf.5.gz",
		"/usr/share/man/man8/drbd-8.3.8.gz",
		"/usr/share/man/man8/drbd-8.4.8.gz",
		"/usr/share/man/man8/drbd-9.0.8.gz",
		"/usr/share/man/man8/drbd-overview-9.0.8.gz",
		"/usr/share/man/man8/drbd-overview.8.gz",
		"/usr/share/man/man8/drbd.8.gz",
		"/usr/share/man/man8/drbdadm-8.3.8.gz",
		"/usr/share/man/man8/drbdadm-8.4.8.gz",
		"/usr/share/man/man8/drbdadm-9.0.8.gz",
		"/usr/share/man/man8/drbdadm.8.gz",
		"/usr/share/man/man8/drbddisk-8.3.8.gz",
		"/usr/share/man/man8/drbddisk-8.4.8.gz",
		"/usr/share/man/man8/drbdmeta-8.3.8.gz",
		"/usr/share/man/man8/drbdmeta-8.4.8.gz",
		"/usr/share/man/man8/drbdmeta-9.0.8.gz",
		"/usr/share/man/man8/drbdmeta.8.gz",
		"/usr/share/man/man8/drbdsetup-8.3.8.gz",
		"/usr/share/man/man8/drbdsetup-8.4.8.gz",
		"/usr/share/man/man8/drbdsetup-9.0.8.gz",
		"/usr/share/man/man8/drbdsetup.8.gz",
		"/var/lib/drbd",
	}
	sort.Strings(ex10)
	out10 := RemoveCommonFilePrefixes(in10)
	sort.Strings(out10)
	if !reflect.DeepEqual(ex10, out10) {
		t.Errorf("RemoveCommonFilePrefixes expected: %v; got: %v.", ex10, out10)
		for i := 0; i < len(ex10); i++ {
			if ex10[i] != out10[i] {
				t.Errorf("# %d: %v <> %v", i, ex10[i], out10[i])
			}
		}
	}
}

func TestUtilFlattenListWithSplit1(t *testing.T) {
	{
		in := []string{} // input
		ex := []string{} // expected
		out := FlattenListWithSplit(in, []string{",", ";", " "})
		sort.Strings(out)
		sort.Strings(ex)
		if !reflect.DeepEqual(ex, out) {
			t.Errorf("FlattenListWithSplit expected: %v; got: %v.", ex, out)
		}
	}

	{
		in := []string{"hey"} // input
		ex := []string{"hey"} // expected
		out := FlattenListWithSplit(in, []string{",", ";", " "})
		sort.Strings(out)
		sort.Strings(ex)
		if !reflect.DeepEqual(ex, out) {
			t.Errorf("FlattenListWithSplit expected: %v; got: %v.", ex, out)
		}
	}

	{
		in := []string{"a", "b", "c", "d"} // input
		ex := []string{"a", "b", "c", "d"} // expected
		out := FlattenListWithSplit(in, []string{",", ";", " "})
		sort.Strings(out)
		sort.Strings(ex)
		if !reflect.DeepEqual(ex, out) {
			t.Errorf("FlattenListWithSplit expected: %v; got: %v.", ex, out)
		}
	}

	{
		in := []string{"a,b,c,d"}          // input
		ex := []string{"a", "b", "c", "d"} // expected
		out := FlattenListWithSplit(in, []string{",", ";", " "})
		sort.Strings(out)
		sort.Strings(ex)
		if !reflect.DeepEqual(ex, out) {
			t.Errorf("FlattenListWithSplit expected: %v; got: %v.", ex, out)
		}
	}

	{
		in := []string{"a,b;c d"}          // input (mixed)
		ex := []string{"a", "b", "c", "d"} // expected
		out := FlattenListWithSplit(in, []string{",", ";", " "})
		sort.Strings(out)
		sort.Strings(ex)
		if !reflect.DeepEqual(ex, out) {
			t.Errorf("FlattenListWithSplit expected: %v; got: %v.", ex, out)
		}
	}

	{
		in := []string{"a,b,c,d;e,f,g,h;i,j,k,l;m,n,o,p q,r,s,t;u,v,w,x y z"}                                                                            // input (mixed)
		ex := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"} // expected
		out := FlattenListWithSplit(in, []string{",", ";", " "})
		sort.Strings(out)
		sort.Strings(ex)
		if !reflect.DeepEqual(ex, out) {
			t.Errorf("FlattenListWithSplit expected: %v; got: %v.", ex, out)
		}
	}
}

func TestSortedStrSliceCompare0(t *testing.T) {
	slice0 := []string{"foo", "bar", "baz"}
	slice1 := []string{"bar", "foo", "baz"}

	if err := SortedStrSliceCompare(slice0, slice1); err != nil {
		t.Errorf("slices were not evaluated as equivalent: %v, %v", slice0, slice1)
	}
}

func TestSortedStrSliceCompare1(t *testing.T) {
	slice0 := []string{"foo", "bar", "baz"}
	slice1 := []string{"fi", "fi", "fo"}

	if err := SortedStrSliceCompare(slice0, slice1); err == nil {
		t.Errorf("slices were evaluated as equivalent: %v, %v", slice0, slice1)
	}
}

func TestSortedStrSliceCompare2(t *testing.T) {
	slice0 := []string{"foo", "bar", "baz"}
	slice1 := []string{"foo", "bar"}

	if err := SortedStrSliceCompare(slice0, slice1); err == nil {
		t.Errorf("slices were evaluated as equivalent: %v, %v", slice0, slice1)
	}
}

func TestSortedStrSliceCompare3(t *testing.T) {
	slice0 := []string{"foo", "bar", "baz"}
	slice1 := []string{"zip", "zap", "zop"}

	_ = SortedStrSliceCompare(slice0, slice1)

	if slice0[0] != "foo" || slice0[1] != "bar" || slice0[2] != "baz" {
		t.Errorf("input slice reordered to: %v", slice0)
	}

	if slice1[0] != "zip" || slice1[1] != "zap" || slice1[2] != "zop" {
		t.Errorf("input slice reordered to: %v", slice1)
	}
}

func TestSortUInt64Slice(t *testing.T) {
	slice0 := []uint64{42, 13, 0}
	sort.Sort(UInt64Slice(slice0))
	if slice0[0] != 0 || slice0[1] != 13 || slice0[2] != 42 {
		t.Errorf("input slice reordered to: %v", slice0)
	}

	slice1 := []uint64{99, 12, 13}
	sort.Sort(UInt64Slice(slice1))
	if slice1[0] != 12 || slice1[1] != 13 || slice1[2] != 99 {
		t.Errorf("input slice reordered to: %v", slice1)
	}
}

func TestSortMapStringValuesByUInt64Keys(t *testing.T) {
	if x := len(SortMapStringValuesByUInt64Keys(nil)); x != 0 {
		t.Errorf("input map of nil caused a: %d", x)
	}

	map0 := map[uint64]string{
		42: "world",
		34: "there",
		13: "hello",
	}
	slice0 := SortMapStringValuesByUInt64Keys(map0)
	if slice0[0] != "hello" || slice0[1] != "there" || slice0[2] != "world" {
		t.Errorf("input slice reordered to: %v", slice0)
	}

	map1 := map[uint64]string{
		99: "a",
		12: "c",
		13: "b",
	}
	slice1 := SortMapStringValuesByUInt64Keys(map1)
	if slice1[0] != "c" || slice1[1] != "b" || slice1[2] != "a" {
		t.Errorf("input slice reordered to: %v", slice1)
	}

	map2 := map[uint64]string{
		12:    "c",
		0:     "d",
		44442: "b",
	}
	slice2 := SortMapStringValuesByUInt64Keys(map2)
	if slice2[0] != "d" || slice2[1] != "c" || slice2[2] != "b" {
		t.Errorf("input slice reordered to: %v", slice2)
	}
}
