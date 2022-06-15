## ONS API Scripts

A place to gather helpful scripts and examples of how to extract data from ONS APIs. The [Developer Hub](developer.ons.gov.uk) contains official documentation of the APIs, these scripts are not a replacement for that.

### List of datasets and their dimensions

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