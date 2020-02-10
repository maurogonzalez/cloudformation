# Ansible + AWSCloudformation Examples

This project tries to cover some AWS management through
[AWSCloudformation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html) using [Ansible](https://docs.ansible.com/ansible/latest/index.html) as automation tool.

# Table of contents
1. [Requirements](#requirements)
1. [Getting ready](#getting-ready)
1. [Roles](#roles)
    1. [eks](#eks)
    1. [iam](#iam)
    1. [network](#iam)
1. [Author](#author)

## Requirements <a name="requirements" />

- [ansible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)

- [boto3](https://pypi.org/project/boto3/):
  Follow the AWS Credentials setup to use a particular profile in the playbooks.

## Getting ready <a name="getting-ready" />

Clone this repo:
```
# Move to your working path
$ git clone https://github.com/maurogonzalez/cloudformation.git
$ cd cloudformation
```

## Roles <a name="roles" />

The general `playbook` is `playbook.yml`. It receives the role variable
to set which role is going to be played. Each role needs an `environment`
variable which is already set in `inventory/dev.yml` but it can be 
overridden in the command (I suggest to use inventory as it is easier to
set variables/values corresponding to a given scope).

The command to run each implemented role is as follows:
```
$  ansible-playbook -vvv -i inventory/dev.yml -e "role=$ROLE profile=default" playbook.yml
```


### EKS <a name="eks" />

Manages an [EKS](https://aws.amazon.com/eks/) Cluster.

Requires:
- `iam` role
- `network` role

### IAM <a name="iam" />

Manages [IAM](https://aws.amazon.com/iam/) resources.

### Network <a name="network" />

Manages [VPC](https://aws.amazon.com/ec2/) resources such.

## Author <a name="author" />

If you have any questions regarding to this project contact:  
Mauro González <jmajma8@gmail.com>