variable "token" {
  description = "user token"
  sensitive = true
  nullable  = false
}

variable "cloud_id" {
  type        = string
  description = "virtual cloud id"
  default     = "b1g69nmv0ba0cfoo7a1h"
  nullable    = false
}

variable "folder_id" {
  type        = string
  description = "id of the folder in cloud"
  default     = "b1g9bkqoqo3havemb2nf"
  nullable    = false
}

variable "zone" {
  type        = string
  description = "geo zone id"
  default     = "ru-central1-a"
  nullable    = false
}