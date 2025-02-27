# Things to do

[X] We should rewrite the perl script using go and make it part of the applicaiton instead of using a cgi.
* Improve the security checking.
* We need pipeline actions to build the container and test it.
  * This will probably mean moving the 'excuses' to a database or seperate files, so testing on the application can make sure that it returns proper values.
* We need to setup deployment.
  * A helm chart in kubernetes?
  * An EC2 image running docker, with an ASG and LB?
  * Use a defined Cloud Service?
* Add static code checking.
* Add FinOps.
