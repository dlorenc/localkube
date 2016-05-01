#!/usr/bin/env perl
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

use strict;

my $command = "mksysnum_linux.pl ". join(' ', @ARGV);

print <<EOF;
// $command
// MACHINE GENERATED BY THE ABOVE COMMAND; DO NOT EDIT

package unix

const(
EOF

sub fmt {
	my ($name, $num) = @_;
	$name =~ y/a-z/A-Z/;
	print "	SYS_$name = $num;\n";
}

my $prev;
while(<>){
	if(/^#define __NR_(\w+)\s+([0-9]+)/){
		$prev = $2;
		fmt($1, $2);
	}
	elsif(/^#define __NR_(\w+)\s+\(\w+\+\s*([0-9]+)\)/){
		fmt($1, $prev+$2)
	}
}

print <<EOF;
)
EOF
