# R Scripts

For all R scripts, you can execute them in any tool with access to the R language - R Studio is a common tool for this.
If you do not have access to any IDE software, try copy/pasting the code [online tools](https://makemeanalyst.com/run-your-r-code/) which will run the script for you.

Currently, the data in each script is output using simple `print()` commands, so you can extend/format the results however is most suitable.

## List of datasets and their dimensions
Load the [dimensions-by-datasets.r](r/dimensions-by-datasets.r) script in your preferred R tooling and run. 

The console should contain outputs roughly like this: 
```
[1] "Personal well-being estimates by local authority"
                      label                code_list
1                  Estimate       wellbeing-estimate
2                 Geography administrative-geography
3 All measures of wellbeing     measure-of-wellbeing
4                      Time                  yyyy-yy
```
With one such occurence for each dataset available via the API.

## List of latest CSV URLs
Load the [latest_downloads.r](r/latest_downloads.r) script in your preferred R tooling and run. 

The console should contain outputs roughly like this: 
```
                                                                            title
1 Earnings and hours worked, UK region by public and private sector: ASHE Table 25
  frequency version
1    Annual       3
                                                                                                     csv
1 https://download.beta.ons.gov.uk/downloads/datasets/ashe-tables-25/editions/time-series/versions/3.csv
```
With one such occurence for each dataset available via the API.

NB: The latest version and CSV URL indicates the latest published version across all editions. If a particular
edition is required, this will need specifying in the requests.