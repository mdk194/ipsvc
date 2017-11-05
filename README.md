sample web server when visited will render html page showing local ip of server
and version number

Contains:
- main.go: server golang source code
- ipsvc.service: systemd service to auto start the webserver at boot
- packer.json: packer template for building ami
- buildspec.yml: codebuild definition to build binary then build ami with packer
- cf/pipeline.yaml: cloudformation template to create aws codepipeline
- cf/deploy.yaml: cloudformation template for pipeline to deploy new built ami

deploy:
- create aws codecommit repo
- aws cloudformation create-stack --stack-name ipsvc-pipeline --capabilities CAPABILITY_NAMED_IAM --template-body file://cf/pipeline.yaml
