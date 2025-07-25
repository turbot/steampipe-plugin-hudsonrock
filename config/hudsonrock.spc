connection "hudsonrock" {
  plugin = "hudsonrock"
  # The maximum number of attempts (including the initial call) Steampipe will make for failing API calls.
  # Defaults to 3 and must be greater than or equal to 1.
  # max_retries = 3

  # The minimum delay between API calls in seconds.
  # Defaults to 1 and must be greater than or equal to 0.
  # min_delay = 1
}
