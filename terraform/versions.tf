terraform {
  required_version = ">= 1.3"
  required_providers {
    yandex = {
      version = ">= 0.85"
      source  = "yandex-cloud/yandex"
    }
  }  
}