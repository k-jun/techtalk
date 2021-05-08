
deploy:
	# ubuntu 18.04
	aws ec2 run-instances --image-id ami-0cd744adeca97abb1 --instance-type t2.large --key-name dena-macbook-pro-16 --user-data file://${PWD}/init.sh

load:
	vegeta attack -rate=100 -duration=10s -targets requests.txt | vegeta report
