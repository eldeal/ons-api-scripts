# load packages
library(tidyverse)

# get JSON data from datasets API
api <- jsonlite::fromJSON("https://api.beta.ons.gov.uk/v1/datasets?limit=200")

# see what's available
datasets <- data.frame(
  title = api$items$title,
  latest_version_link = api$items$links$latest_version$href,
  latest_version = api$items$links$latest_version$id,
  frequency = api$items$release_frequency,
  stringsAsFactors = FALSE
)

# loop through datasets
for (i in 1:nrow(datasets)) {
  # use the latest_version link to get to the list of dimensions
  latest_version <- jsonlite::fromJSON(datasets$latest_version_link[i])

  # if any of the fields we care about are empty, replace with NA
  freq <- datasets$frequency[i]
  if (length(freq) == 0) {
    freq <- NA
  }

  t <- datasets$title[i]
  if (length(t) == 0) {
    t <- NA
  }

  v <- datasets$latest_version[i]
  if (length(v) == 0) {
    v <- NA
  }

  c <- latest_version$downloads$csv$href
  if (length(c) == 0) {
    c <- NA
  }

  csv <- data.frame(
    title = t,
    frequency = freq,
    version = v,
    csv = c
  )

  # print the title, frequency, version number and CSV URL for each dataset
  print(csv)
}
