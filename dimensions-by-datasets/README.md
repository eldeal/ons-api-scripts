## List of datasets and their dimensions

### How to run
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

### Limitations
This currently assumes less than 200 datasets exist in the API, and does not paginate through the full list. If more than 200 datasets are present, the code will need to be updated.

### What it's doing

The code all lives in the main.go file. It is performing the following steps
- Starting with the `main()` function as an entry point
    - Get a list of all datasets via the `getDatasets` function
        - This function calls `https://api.beta.ons.gov.uk/v1/datasets?limit=200` and converts the response into an object of type `DatasetList`. This object type (defined in the structs at the top of the file) only maps the specific fields needed to satisfy this script. In order to reuse this for other scripts, other fields may need to be added
    - Check that the list isn't empty
    - Create an empty map, where we will map titles to lists of dimension names
    - Loop over the `Items` in the `DatasetList` response: for each title, `getDimensions` of the provided version
        - `getDimensions` takes the latestVersion link from the dataset and attaches `/dimensions` to the URL, returning all the metadata about all dimensions for that dataset
        - This data is mapped to a `DimensionList` object, with the same limited fields as described for the previous function (only what is needed is mapped)
    - Once we have a map with the information we're looking for, we can start to format a result by looping over the results
        - This loop just prints to stdout, rather than writing to a file. This is why the execution command contains ` > results.txt` to pass the result to a file. Without this, the output will just be shown directly on screen and not captured.