variable "DB_USER" {
    type = string
    default = getenv("DB_USER")
}
variable "DB_PASS" {
    type = string
    default = getenv("DB_PASS")
}
variable "DB_HOST" {
    type = string
    default = getenv("DB_HOST")
}
variable "DB_PORT" {
    type = string
    default = getenv("DB_PORT")
}
variable "DB_NAME" {
    type = string
    default = getenv("DB_NAME")
}

env "local" {
    url = "postgres://${var.DB_USER}:${var.DB_PASS}@${var.DB_HOST}:${var.DB_PORT}/${var.DB_NAME}?sslmode=disable"
    dev = "docker://postgres/15/dev"

    migration {
        dir = "file://db/migrations"
    }
}