
deploy:
	# ubuntu 18.04
	aws ec2 run-instances --image-id ami-0cd744adeca97abb1 --instance-type t2.large --key-name CUSTOM_KEY_NAME  --user-data file://${PWD}/init.sh
