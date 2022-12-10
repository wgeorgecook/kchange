# kchange
KChange gets the change cause of the provided deployments. 

### Install
Clone this repo and then `go install` from it. 

### Usage
Deployments can be provided as a list of strings using the `-d` 
shorthand, or `--deployments`. 

### Example
```
╰─ kchange -d deployment1 -d deployment2
kubectl set image deployment deployment1 pod=repo/imagename:123456 --record=true
kubectl set image deployment deployment2 pod=repo/imagename:123456 --record=true
```