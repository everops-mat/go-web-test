#!/usr/bin/env perl
use strict;
use warnings;
use Test::More;
use LWP::UserAgent;
use HTTP::Request::Common;

# Define the CGI script location
my $cgi_script = "./eo.pl";

# Check if the script exists and is executable
ok(-e $cgi_script, "CGI script exists");
ok(-x $cgi_script, "CGI script is executable");

# Simulate CGI execution using `open3` to capture output
use IPC::Open3;
use Symbol qw(gensym);

sub run_cgi {
    my $cmd = shift;
    my $stderr = gensym;
    my ($in, $out);
    my $pid = open3($in, $out, $stderr, $cmd);
    waitpid($pid, 0);
    my $output = do { local $/; <$out> };
    return $output;
}

# Run the CGI script and check output
my $output = run_cgi($cgi_script);

# Test that the output contains valid HTML
like($output, qr/Content-type: text\/html/, "Output contains correct Content-Type header");
like($output, qr/The cause of the problem is:/, "Output contains expected BOFH message");

done_testing();
