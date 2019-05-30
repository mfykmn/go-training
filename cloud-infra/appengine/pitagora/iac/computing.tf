resource "google_app_engine_application" "default" {
  location_id = "asia-northeast1"

}

resource "google_app_engine_application" "worker" {
  location_id = "asia-northeast1"
}