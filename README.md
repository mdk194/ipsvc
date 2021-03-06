sample web server when visited will render html page showing local ip of server
and version number

Contains:
```
.
├── buildspec.yml       # AWS Codebuild steps
├── cf                  # folder contains Cloudformation (CF) templates
│   ├── config.json     # configuration parameters for deploy.yaml CF template
│   ├── deploy.yaml     # EC2 instance + Elastic IP that hosting the web server
│   └── pipeline.yaml   # codepipeline + codebuild + iam roles ....
├── k8s                 # Kubernetes resources definitions
│   ├── ipsvc-de.yaml   # deployment object for ipsvc
│   └── ipsvc-svc.yaml  # service object for ipsvc
├── ipsvc.service       # systemd service to run the web server at boot
├── main.go             # web server source code in Golang
├── Makefile            # instructions to compile main.go
├── packer.json         # packer template to build AMI
└── README.md
```

AMI deploy:
- create aws codecommit repo
- `aws cloudformation create-stack --stack-name ipsvc-pipeline --capabilities CAPABILITY_NAMED_IAM --template-body file://cf/pipeline.yaml`

K8s deploy:
```
kubectl apply -f k8s/
```

![pipeline](./images/pipeline.png)
