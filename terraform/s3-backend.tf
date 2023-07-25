# Описание бэкенда хранения состояния
terraform {
    backend "s3" {
        endpoint   = "storage.yandexcloud.net"
        bucket     = "momo-store-terraformstate"
        region     = "ru-central1"
        key        = "terraform.tfstate"
    
        skip_region_validation      = true
        skip_credentials_validation = true
   }
}