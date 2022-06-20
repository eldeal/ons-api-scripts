## ONS API Scripts

A place to gather helpful scripts and examples of how to extract data from ONS APIs. The [Developer Hub](developer.ons.gov.uk) contains official documentation of the APIs, these scripts are not a replacement for that.

### List of datasets and their dimensions

This script is intended to give a high level overview of the data available through the ONS API.
It lists datasets, and the dimensions available within them so multiple datasets with similar 
dimensions or characteristics can easily be identified.

#### R Script

Load the [dimensions-by-datasets.r](r/dimensions-by-datasets.r) script in your preferred R tooling (RStudio etc) and run. 

If you do not have access to any IDE software, try copy/pasting the code [online tools](https://makemeanalyst.com/run-your-r-code/) which will run the script for you.

Currently, the data is output using simple `print()` commands, so you can extend it/format the results however is most suitable.

#### Executable Go Script

Execute the code in the [dimensions-by-datasets](dimensions-by-datasets) directory, and pass the results to a file.

> ./dimensions-by-datasets > results.txt

This will give you results in the format:

```
Life Expectancy by Local Authority
  Contains:
  - Label: "Age group" CodeList: "age-groups"
  - Label: "Geography" CodeList: "administrative-geography"
  - Label: "Sex" CodeList: "sex"
  - Label: "Time" CodeList: "two-year-intervals"
  ```
For more details see the [README](dimensions-by-datasets/README.md).

