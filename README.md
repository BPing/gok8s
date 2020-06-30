# gok8s

k8s发布go web 项目【学习】



## # 启动

```shell script
# 编译web镜像
docker build -t cbping/gok8s:1.0 .

# 创建configmap
# kubectl create configmap gok8s-config --from-file=./conf/config.json
kubectl create  -f configmap.yaml

# 启动ingress-controller
# - minikube: minikube addons enable ingress
# 

# k8s部署
kubectl apply -f gok8s-demo.yaml

# 查找ingress地址
kubectl get ingress 

# 访问
http://{$ingress_address}/(ping|version)
```