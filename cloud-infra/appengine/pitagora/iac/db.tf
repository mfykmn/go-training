resource "random_id" "db_name_suffix" {
  byte_length = 4
}

resource "google_sql_database_instance" "master" {
  name = "master-instance-${random_id.db_name_suffix.hex}"
  database_version = "MYSQL_5_7"

  settings {
    tier = "db-f1-micro"
  }
}

resource "google_sql_database" "users" {
  name      = "users-db"
  instance  = "${google_sql_database_instance.master.name}"
  charset   = "utf8mb4"
  collation = "utf8mb4_bin"
}

resource "google_sql_user" "users" {
  name     = "root"
  instance = "${google_sql_database_instance.master.name}"
  host     = "%"
  password = "password"
}