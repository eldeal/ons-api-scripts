#load packages
library(tidyverse)

#get JSON data from datasets API
api <- jsonlite::fromJSON('https://api.beta.ons.gov.uk/v1/datasets?limit=200')

#see what's available
datasets <- data.frame(title = api$items$title, description = api$items$description, editions = api$items$links$editions$href, latest_version = api$items$links$latest_version$href, stringsAsFactors = FALSE)

# loop through datasets
for(i in 1:length(datasets))
{
  title <- datasets$title[i]

  #use the latest_version link to get to the list of dimensions
  latest_version_dimensions <- paste0(datasets$latest_version[i],'/dimensions?limit=50')
  dimensions_json <- jsonlite::fromJSON(latest_version_dimensions)
  dims <- data.frame(label = dimensions_json$items$label, code_list = dimensions_json$items$links$code_list$id)

  #print the title and list of dimensions for each dataset
  print(title)
  print(dims)
}
