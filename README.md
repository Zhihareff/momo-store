# СТРУКТУРА ПРОЕКТА


```bash
|
├── backend                       - исходный код бэкэнда + Dockerfile + gitlab-ci.yml
├── frontend                      - исходный код фронтэнда + Dockerfile + gitlab-ci.yml
├── momo-store-chart              - helm-чарты для развертывания приложения momo-store в k8s кластере
|   └── charts
├── terraform                     - конфигурационные файлы terraform для развертывания кластера
├── .gitlab-ci.yml                - родительский пайплайн для сборки и релиза образов бэкенда и фронтенда в Container Registry
└── .helm-chart.gitlab-ci.yml     - пайплайн для загрузки helm-чартов в репозиторий Nexus
```


# ЗАПУСК ПРИЛОЖЕНИЯ

## 1. Создайте кластер в Яндекс.Облаке

создайте с помощью Terraform необходимую инфраструктуру в Яндекс.Облаке (кластер, группу узлов, сеть, подсеть и т.д.)
```bash
cd terraform
terraform apply
```

перед этим нужно добавить переменные окружения:
ключ к хранилищу S3 в Яндекс.Облаке, в котором будет хранится состояние Terraform
```bash
export AWS_ACCESS_KEY_ID="<идентификатор_ключа>"
export AWS_SECRET_ACCESS_KEY="<секретный_ключ>"
```

(само хранилище необходимо также предварительно создать - имя "momo-store-terraformstate")


## 2. Создайте в облаке хранилище с именем "momo-store-pictures"

и загрузите туда картинки пельменей
(в данном проекте это хранилище создается терраформом)

<img width="900" alt="image" src="https://github.com/Zhihareff/momo-store/raw/main/image/s3-bucket.png">


## 3. Настройте /.kube/config

  - получите креды
```bash
yc managed-kubernetes cluster get-credentials --id <id_кластера> --external
```
  - проверка доступности кластера:
```bash
kubectl cluster-info
```
  - сделайте бэкап текущего ./kube/config:
```bash
cp ~/.kube/config ~/.kube/config.bak
```
  - создайте манифест gitlab-admin-service-account.yaml:

```bash
apiVersion: v1
kind: ServiceAccount
metadata:
  name: gitlab-admin
  namespace: kube-system
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: gitlab-admin-role
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: gitlab-admin
  namespace: kube-system
```
  и примените его:
```bash
kubectl apply -f gitlab-admin-service-account.yaml
```
  - получите endpoint: публичный ip адрес находится по пути Managed Service for Kubernetes/Кластеры/ваш_кластер -> обзор -> основное -> Публичный IPv4

  - получите KUBE_TOKEN:
```bash
kubectl -n kube-system get secrets -o json | jq -r '.items[] | select(.metadata.name | startswith("gitlab-admin")) | .data.token' | base64 --decode
```
  - сгенерируйте конфиг:
```bash
export KUBE_URL=https://<см. пункт выше>   # Важно перед IP указать https://
export KUBE_TOKEN=<см.пункт выше>
export KUBE_USERNAME=gitlab-admin
export KUBE_CLUSTER_NAME=<id_кластера> как в пункте_3

kubectl config set-cluster "$KUBE_CLUSTER_NAME" --server="$KUBE_URL" --insecure-skip-tls-verify=true
kubectl config set-credentials "$KUBE_USERNAME" --token="$KUBE_TOKEN"
kubectl config set-context default --cluster="$KUBE_CLUSTER_NAME" --user="$KUBE_USERNAME"
kubectl config use-context default
```


## 4. Установка Ingress-контроллера NGINX с менеджером для сертификатов Let's Encrypt

 Чтобы с помощью Kubernetes создать Ingress-контроллер NGINX и защитить его сертификатом Let's Encrypt®, выполните следующие действия:

  - установите NGINX Ingress Controller

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.5.1/deploy/static/provider/cloud/deploy.yaml
```

  - установите менеджер сертификатов
  
```bash
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.6.1/cert-manager.yaml
```


## 5. Узнайте IP-адрес Ingress-контроллера

значение в колонке EXTERNAL-IP:
```bash
kubectl get svc -n ingress-nginx
```


## 6. На сайте https://freedns.afraid.org/subdomain/ создайте субдомен

(например, my-pelmen.mooo.com), укажите IP-адрес из п.5
сохраните этот субдомен в переменной SHOP_URL в Gitlab в настройках CI/CD


## 7. Создайте манифест acme-issuer.yaml

```bash
apiVersion: cert-manager.io/v1
kind: ClusterIssuer
metadata:
  name: letsencrypt
  namespace: cert-manager
spec:
  acme:
    server: https://acme-v02.api.letsencrypt.org/directory
    email: <емейл-адрес>
    privateKeySecretRef:
      name: letsencrypt
    solvers:
    - http01:
        ingress:
          class: nginx
```

примените его
```bash
kubectl apply -f acme-issuer.yaml
```


## 8. Для скачивания образов из Container Registry в k8s необходимо создать:
  - секрет:
```bash
kubectl create secret docker-registry docker-config-secret --docker-server=gitlab.praktikum-services.ru:5050   --docker-username=<указать_свой_логин>   --docker-password=<указать_свой_пароль>
```

  - сервис-аккаунт:
```bash
kubectl create serviceaccount my-serviceaccount
kubectl patch serviceaccount my-serviceaccount -p '{"imagePullSecrets": [{"name": "docker-config-secret"}]}' -n default 
```


## 9. Создайте в Nexus репозиторий для хранения helm-чартов

после чего в Gitlab в настройках CI/CD добавьте переменные со следующими именами: NEXUS_REPO_USER (укажите свой логин в Nexus), NEXUS_REPO_PASS (пароль), NEXUS_REPO_URL (ссылка на репозиторий)


## 10. Установите ArgoCD

(см. https://cloud.yandex.ru/docs/managed-kubernetes/operations/applications/argo-cd)

```bash
helm pull oci://cr.yandex/yc-marketplace/yandex-cloud/argo/chart/argo-cd \
  --version 5.4.3-7 \
  --untar && \
helm install \
  --namespace argocd \
  --create-namespace \
  argo-cd ./argo-cd/
```


## 11. Создайте приложение в ArgoCD
  - выполните команду:
```bash
kubectl port-forward service/argocd-server -n argocd 8080:443
```

  - зайдите браузером на http://localhost:8080
    логин: admin;
    пароль получите с помощью команды:
```bash
kubectl -n kube-system get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d
```

  - подключите HELM репозиторий: Settings / Repositories - Connect repo using HTTPS - веберите тип helm, укажите путь до репозитория (в нашем случае Nexus, см. п.9)

  - создайте новое приложение со следующими параметрами:
	Application Name: momo-store
	Project: momo-store
	SYNC POLLICY: Automatic
	Repository URL: указываем url HELM-репозитория (подключенного в предыдущем пункте)
	Chart: momo-store
	Version: x

## Готово!

Приложение доступно по адресу: https://my-pelmen.mooo.com

<img width="900" alt="image" src="https://github.com/Zhihareff/momo-store/raw/main/image/site.png">


<img width="900" alt="image" src="https://github.com/Zhihareff/momo-store/raw/main/image/argocd.png">


<img width="900" alt="image" src="https://github.com/Zhihareff/momo-store/raw/main/image/nexus.png">



# ARGOCD

Для того, чтобы в web-интерфейс ArgoCD можно было заходить по своему url необходимо сделать следующее:
  - в репозитории argo в values.yaml найти следующий блок
```bash
ingress:
    annotations: {}
    enabled: false
```
и заменить "enabled: false" на "enabled: true"
после этого в этом же каталоге выполнить команду:
```bash
helm upgrade --install --namespace=argocd argocd .
```
  - после этого в неймспейсе argocd появится Ingress, нужно будет его отредактировать:
```bash
kubectl edit service argocd-server -n argocd  
```
манифест должен выглядить подобным образом (обязательно разрешить ssl-passthrough):
```bash
# Please edit the object below. Lines beginning with a '#' will be ignored,
# and an empty file will abort the edit. If an error occurs while saving this file will be
# reopened with the relevant failures.
#
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    acme.cert-manager.io/http01-edit-in-place: "true"
    cert-manager.io/cluster-issuer: letsencrypt
    kubernetes.io/ingress.class: nginx
    kubernetes.io/tls-acme: "true"
    meta.helm.sh/release-name: argocd
    meta.helm.sh/release-namespace: argocd
    nginx.ingress.kubernetes.io/backend-protocol: HTTPS
    nginx.ingress.kubernetes.io/ssl-passthrough: "true"
  creationTimestamp: "2023-07-16T05:02:21Z"
  generation: 6
  labels:
    app.kubernetes.io/component: server
    app.kubernetes.io/instance: argocd
    app.kubernetes.io/managed-by: Helm
    app.kubernetes.io/name: argocd-server
    app.kubernetes.io/part-of: argocd
    helm.sh/chart: argo-cd-4.5.3-1
  name: argocd-server
  namespace: argocd
  resourceVersion: "1113383"
  uid: 84a9bf18-80f6-4bf5-a227-d916cdd12b1e
spec:
  rules:
  - host: argocd-pelmen.mooo.com
    http:
      paths:
      - backend:
          service:
            name: argocd-server
            port:
              name: https
        path: /
        pathType: Prefix
  tls:
  - hosts:
    - argocd-pelmen.mooo.com
    secretName: letsencrypt
```
в качестве хоста прописать необходимый домен


# МОНИТОРИНГ

  - скопируйте репозиторий отсюда: https://gitlab.praktikum-services.ru/root/monitoring-tools
  - удостоверьтесь, что в описание приложения бэкенда пельменной (spec.template.annotations) добавлены аннотации:
  ```bash
prometheus.io/path: /metrics
prometheus.io/port: "8081"
prometheus.io/scrape: "true"
```

## 1. Установка Prometheus

  - в скопированном репозитории в prometheus/templates создайте файл clusterRole.yaml:
```bash
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: prometheus
rules:
- apiGroups: [""]
  resources:
    - nodes
    - nodes/proxy
    - nodes/metrics
    - services
    - endpoints
    - pods
  verbs: ["get", "list", "watch"]
- apiGroups:
  - extensions
  resources:
  - ingresses
  verbs: ["get", "list", "watch"]
- nonResourceURLs: ["/metrics"]
  verbs: ["get"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: prometheus
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: prometheus
subjects:
- kind: ServiceAccount
  name: default
  namespace: default
```
  - выполните команду:
```bash
helm upgrade --install prometheus prometheus
```
  - при необходимости настройте Ingress или заходите в веб-интерфейс с помощью port-forward:
```bash  
kubectl port-forward <имя_пода_prometheus> -n default 9090
```

Пример того, что должно получиться:

<img width="900" alt="image" src="https://github.com/Zhihareff/momo-store/raw/main/image/prometheus01.png">
<img width="900" alt="image" src="https://github.com/Zhihareff/momo-store/raw/main/image/prometheus02.png">
<img width="900" alt="image" src="https://github.com/Zhihareff/momo-store/raw/main/image/prometheus03.png">


## 2. Установка Grafana

  - выполните команду:
```bash
helm upgrade --install grafana  grafana
```

  - при необходимости настройте Ingress или заходите в веб-интерфейс с помощью port-forward:
```bash  
kubectl port-forward <имя_пода_grafana> -n default 3000
```

  - импортируйте или настройте самостояльно нужный дашборд

Пример того, что должно получиться:

<img width="900" alt="image" src="https://github.com/Zhihareff/momo-store/raw/main/image/grafana.png">

# БОНУС. POLARIS

https://github.com/FairwindsOps/polaris

Polaris — проект с открытым исходным кодом для Kubernetes, который проверяет и исправляет конфигурацию ресурсов.
Иными словами, помогает поддерживать «здоровье» кластера.
Он включает в себя более 30 встроенных политик конфигурации, а также возможность создавать собственные политики с помощью JSON.
При запуске из командной строки или в качестве изменяемого вебхука Polaris может автоматически устранять проблемы на основе критериев политики.

<img width="900" alt="image" src="https://camo.githubusercontent.com/cf8dc8d0a68de78de27ef0e156c7192e41269388509d89c847589367048774e7/68747470733a2f2f706f6c617269732e646f63732e6661697277696e64732e636f6d2f696d672f6172636869746563747572652e737667">

Polaris может работать в трех разных режимах:

  - admission controller - автоматически отклоняйте или изменяйте деплои, которые не соответствуют политикам
  - command-line tool - включите политику-как-код (policy-as-code) в процесс CI/CD для тестирования локальных файлов YAML (реализовано в данном проекте, см. пайплайн helm-chart.gitlab-ci.yml)
  - дашборд мониторинга - проверьте ресурсы Kubernetes на соответствие политикам:
```bash
helm repo add fairwinds-stable https://charts.fairwinds.com/stable
helm upgrade --install polaris fairwinds-stable/polaris --namespace polaris --create-namespace
kubectl port-forward --namespace polaris svc/polaris-dashboard 8080:80
```
пример дашборда:
<img width="900" alt="image" src="https://github.com/Zhihareff/momo-store/raw/main/image/polaris01.png">
<img width="900" alt="image" src="https://github.com/Zhihareff/momo-store/raw/main/image/polaris02.png">